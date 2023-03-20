CREATE DATABASE IF NOT EXISTS TEAMPROJECT;
DROP TABLE IF EXISTS players;
DROP TABLE IF EXISTS teams;
CREATE TABLE teams (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE players (
                         id INT AUTO_INCREMENT PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         team_id INT NOT NULL,
                         FOREIGN KEY (team_id) REFERENCES teams(id)
);
insert into teams (name) values ('Arsenal');
insert into teams (name) values ('Real Madrid');
insert into players (name, team_id) values ('Cristiano Ronaldo', 1);