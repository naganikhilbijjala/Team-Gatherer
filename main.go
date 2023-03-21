package main

import (
	"database/sql"
	_ "fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	_ "strconv"
)

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Time string `json:"Time"`
}

type Player struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TeamID int    `json:"team_id"`
}

var db *sql.DB

func main() {
	// Connect to the database
	var err error
	db, err = sql.Open("mysql", "root:gaurav11596@tcp(127.0.0.1:3306)/TEAMPROJECT")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the Gin router
	r := gin.Default()

	// Define the routes
	r.GET("/teams", getTeams)
	r.POST("/teams", createTeam)
	r.DELETE("/teams/:id", deleteTeam)
	r.PUT("/teams/:id", updateTeam)
	//r.GET("search/:name", searchTeam)
	r.GET("/players", getPlayers)
	r.POST("/players", createPlayer)
	r.DELETE("/players/:id", deletePlayer)
	//r.PUT("/players/:id", updatePlayer)
	r.GET("getPlayers/:id", getPlayersByTeamID)
	// Start the server
	r.Run(":8080")
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

//func searchTeam(c *gin.Context) {
//	name := c.Param("name")
//
//	rows, err := db.Query("SELECT * FROM teams WHERE name = ?", name)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
//		return
//	}
//	defer rows.Close()
//
//	teams := []Team{}
//	for rows.Next() {
//		var team Team
//		err := rows.Scan(&team.ID, &team.Name, &team.Time)
//		if err != nil {
//			log.Println(err)
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
//			return
//		}
//		teams = append(teams, team)
//	}
//	c.JSON(http.StatusOK, teams)
//}

// Have to handle the error for foreign key constraint for players
func deleteTeam(c *gin.Context) {
	id := c.Param("id")
	result, err := db.Exec("DELETE FROM teams WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Team deleted"})
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

func updateTeam(c *gin.Context) {
	id := c.Param("id")

	var team Team
	if err := c.ShouldBindJSON(&team); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	result, err := db.Exec("UPDATE teams SET name = ? WHERE id = ?", team.Name, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team not found"})
		return
	}

	c.JSON(http.StatusOK, team)
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

func deletePlayer(c *gin.Context) {
	id := c.Param("id")
	result, err := db.Exec("DELETE FROM players WHERE id = ?", id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player deleted"})
}

//func updatePlayer(c *gin.Context) {
//	id := c.Param("id")
//
//	var player Player
//	if err := c.ShouldBindJSON(&player); err != nil {
//		log.Println(err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
//		return
//	}
//
//	result, err := db.Exec("UPDATE players SET name = ?, team_id = ? WHERE id = ?", player.Name, player.TeamID, id)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
//		return
//	}
//
//	rowsAffected, err := result.RowsAffected()
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
//		return
//	}
//
//	if rowsAffected == 0 {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
//		return
//	}
//
//	c.JSON(http.StatusOK, player)
//}

func getPlayersByTeamID(c *gin.Context) {
	teamID := c.Param("id")

	rows, err := db.Query("SELECT * FROM players WHERE team_id = ?", teamID)
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
