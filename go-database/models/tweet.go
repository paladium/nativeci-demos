package models

type Tweet struct {
	ID       string
	Tweet    string
	UserID   int
	PostedAt string
}

type ResponseTweet struct {
	Tweet
	Username string
}
