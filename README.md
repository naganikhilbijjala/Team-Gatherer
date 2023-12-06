# Team Gatherer


## Overview:
Finding teammates for specific games and coordinating playtimes can often be a challenging task. Enter Team Gatherer, a solution designed to streamline this process. With Team Gatherer, you have the ability to create a game, specifying important details such as its name, scheduled time, required number of players, and the designated location where the game is going to be played. Once listed on the platform, the game becomes visible to other users, complete with a convenient "join" button. Interested users can easily join games that align with their preferences, as long as the player limit has not been reached. This approach simplifies the difficult task of finding suitable teammates for gaming sessions.

## Tech Stack:
We used React for the frontend, GoLang for the backend, and MySQL to store all the information.

Instructions to setup the project

## Frontend Setup
1. Go to Frontend folder and run the following commands
   1. `npm i` Installs all the required dependencies
   2. `npm start` Starts running on port number 3000 by default

## Backend Setup
2. Go to Backend folder
   1. Download the scripts.sql file and import the database into the mysql db inorder to start with the inital data
   2. To import mysql data first login to your mysql and create a database then use the command `mysql -u username -p your_database_name < dumpfilename.sql` to import the data into db
   3. Change the user name and password and the database name as defined in your system
   4. To run the backend, go to backend folder and run the command `go run .\main.go`