CREATE DATABASE twitter;
use twitter;
CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT,
    username varchar(255) NOT NULL,
    password varchar(255),
    PRIMARY KEY (id)
);
