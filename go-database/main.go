package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"go-database/db"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func createJwtToken(userID string) (*string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func tryDecodeToken(c *gin.Context) (*string, error) {
	tokenCookie, err := c.Cookie("token")
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(tokenCookie, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, err
	}
	userID, ok := token.Claims.(jwt.MapClaims)["user_id"].(string)
	if !ok {
		return nil, errors.New("Cannot get the user id")
	}
	return &userID, nil
}

func main() {
	r := gin.Default()
	dbRepository := db.NewDatabaseRepository(os.Getenv("DB_URL"))
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		_, err := tryDecodeToken(c)
		if err == nil {
			c.Redirect(301, "/chat")
		} else {
			c.HTML(http.StatusOK, "index.html", nil)
		}
	})
	r.GET("/chat", func(c *gin.Context) {
		_, err := tryDecodeToken(c)
		if err == nil {
			tweets, err := dbRepository.GetAllTweets()
			if err != nil {
				log.Println(err.Error())
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Cannot get the tweets"})
				return
			}
			c.HTML(http.StatusOK, "chat.html", gin.H{
				"tweets": tweets,
			})
		} else {
			c.Redirect(301, "/")
		}
	})
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "" || password == "" {
			c.Redirect(301, "")
			return
		}
		//Try to find the user in database or register him
		userID, err := dbRepository.FindOrRegisterUser(username, password)
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid username or password"})
			return
		}
		fmt.Printf("New user_id %d\n", *userID)
		//Create a jwt cookie and set is response
		token, err := createJwtToken(strconv.Itoa(int(*userID)))
		if err != nil {
			log.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid username or password"})
			return
		}
		fmt.Println("Generated jwt token")
		c.SetCookie("token", *token, 1000, "/", ".", false, false)
		c.Redirect(301, "/chat")
	})
	r.POST("/tweet", func(c *gin.Context) {
		userID, err := tryDecodeToken(c)
		if err == nil {
			tweet := c.PostForm("tweet")
			err := dbRepository.PostTweet(*userID, tweet)
			if err != nil {
				log.Println(err.Error())
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Cannot post the tweet"})
				return
			}
			c.Redirect(301, "/chat")
		} else {
			c.HTML(http.StatusOK, "index.html", nil)
		}
	})
	r.Run("0.0.0.0:8000")
}
