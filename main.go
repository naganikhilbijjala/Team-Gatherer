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

type Player struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TeamID int    `json:"team_id"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:gaurav11596@tcp(127.0.0.1:3306)/TEAMPROJECT")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/players", getPlayers)
	r.POST("/players", createPlayer)
	r.Run(":8080")
}

func getPlayers(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM players")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer rows.Close()

	players := []Player{}
	for rows.Next() {
		var player Player
		err := rows.Scan(&player.ID, &player.Name, &player.TeamID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		players = append(players, player)
	}
	c.JSON(http.StatusOK, players)
}

func createPlayer(c *gin.Context) {
	var player Player
	if err := c.ShouldBindJSON(&player); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Check if the team exists
	var ID int
	err := db.QueryRow("SELECT id FROM teams WHERE id = ?", player.TeamID).Scan(&ID)
	log.Println(player.TeamID)
	log.Println(ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Team not found"})
			return
		}
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error, Team Not Found"})
		return
	}

	result, err := db.Exec("INSERT INTO players (name, team_id) VALUES (?, ?)", player.Name, player.TeamID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error, Player not inserted"})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	player.ID = int(id)
	c.JSON(http.StatusCreated, player)
}
