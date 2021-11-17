CREATE DATABASE twitter;
use twitter;
CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT,
    username varchar(255) NOT NULL,
    password varchar(255),
    PRIMARY KEY (id)
);
CREATE TABLE tweets (
    id int NOT NULL AUTO_INCREMENT,
    tweet varchar(255) NOT NULL,
    user_id int NOT NULL,
    posted_at datetime NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);