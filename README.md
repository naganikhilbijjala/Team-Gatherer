# Team Gatherer

Instructions to run the code

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

