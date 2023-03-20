CREATE DATABASE IF NOT EXISTS TEAMPROJECT;
DROP TABLE IF EXISTS teams;
CREATE TABLE teams (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


insert into teams (name) values ('Arsenal');
insert into teams (name) values ('Real Madrid');
