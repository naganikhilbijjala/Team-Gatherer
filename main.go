package main

import (
	"database/sql"
	_ "fmt"
	"log"
	"net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:gaurav11596@tcp(127.0.0.1:3306)/TEAMPROJECT")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the Gin router
	r := gin.Default()

	// Define the routes
	r.POST("/teams", createTeam)
	r.GET("/teams", getTeams)

	// Start the server
	r.Run(":8080")
}

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Time string `json:"Time"`
}

func createTeam(c *gin.Context) {
	var team Team
	if err := c.ShouldBindJSON(&team); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	result, err := db.Exec("INSERT INTO teams (name) VALUES (?)", team.Name)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	team.ID = int(id)
	c.JSON(http.StatusCreated, team)
}

func getTeams(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM teams")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer rows.Close()

	teams := []Team{}
	for rows.Next() {
		var team Team
		err := rows.Scan(&team.ID, &team.Name, &team.Time)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		teams = append(teams, team)
	}
	c.JSON(http.StatusOK, teams)
}
