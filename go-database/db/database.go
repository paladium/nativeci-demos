package db

import (
	"database/sql"
	"errors"
	"go-database/models"
	"log"
	"net/url"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseRepository struct {
	client *sql.DB
}

func NewDatabaseRepository(
	connectionURL string,
) *DatabaseRepository {
	repository := &DatabaseRepository{}
	err := repository.init(connectionURL)
	if err != nil {
		panic(err)
	}
	return repository
}

func (db *DatabaseRepository) getDatabaseConnectionConfig(connectionURL string) (*string, error) {
	u, err := url.Parse(connectionURL)
	if err != nil {
		return nil, err
	}
	password, _ := u.User.Password()
	config := mysql.Config{
		User:                 u.User.Username(),
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 u.Host,
		DBName:               "twitter",
		AllowNativePasswords: true,
	}
	connection := config.FormatDSN()
	return &connection, nil
}

func (db *DatabaseRepository) init(connectionURL string) (err error) {
	log.Println("Connecting to database")
	connection, err := db.getDatabaseConnectionConfig(connectionURL)
	if err != nil {
		return
	}
	db.client, err = sql.Open("mysql", *connection)
	if err != nil {
		return
	}
	err = db.client.Ping()
	if err != nil {
		return
	}
	log.Println("Connected to database")
	return
}

func (db *DatabaseRepository) FindOrRegisterUser(username string, password string) (userID *int64, err error) {
	//Find the user in users table by the username
	//If the hashed passwords match return no error
	//Return error otherwise
	//If the user does not exist, register him and return no error
	row := db.client.QueryRow("SELECT id,username,password FROM users WHERE username=?", username)
	var existingUser models.User
	err = row.Scan(&existingUser.ID, &existingUser.Username, &existingUser.Password)
	if err != nil {
		//Insert the user
		var hashedPassword string
		hashedPassword, err = HashPassword(password)
		if err != nil {
			return
		}
		var result sql.Result
		result, err = db.client.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			return
		}
		var insertedId int64
		insertedId, err = result.LastInsertId()
		if err != nil {
			return
		}
		userID = &insertedId
	} else {
		//Check the hashed password
		if !CheckPasswordHash(password, existingUser.Password) {
			err = errors.New("passwords do not match")
			return
		}
		userID = &existingUser.ID
	}
	return
}

func (db *DatabaseRepository) GetAllTweets() ([]models.ResponseTweet, error) {
	rows, err := db.client.Query("SELECT t.id, tweet, user_id, posted_at, u.username FROM tweets t INNER JOIN users u ON t.user_id = u.id order by posted_at desc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := []models.ResponseTweet{}
	for rows.Next() {
		var result models.ResponseTweet
		err := rows.Scan(&result.ID, &result.Tweet.Tweet, &result.UserID, &result.PostedAt, &result.Username)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (db *DatabaseRepository) PostTweet(userID string, tweet string) error {
	_, err := db.client.Exec("INSERT INTO tweets(tweet, user_id, posted_at) VALUES(?, ?, ?)", tweet, userID, time.Now().Format("2006-01-02 03:04:05"))
	if err != nil {
		return err
	}
	return nil
}
