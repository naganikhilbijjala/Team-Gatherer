package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Player struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TeamID int    `json:"team_id"`
	UserID int    `json:"user_id"`
}

func GetPlayers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func CreatePlayer(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		tx, err := db.Begin()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Insert the player into the players table
		result, err := tx.Exec("INSERT INTO players (name, team_id, user_id) VALUES (?, ?, ?)", player.Name, player.TeamID, player.UserID)
		if err != nil {
			log.Println(err)
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error, Player not inserted"})
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Println(err)
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		player.ID = int(id)

		// Update the current count in the teams table
		_, err = tx.Exec("UPDATE teams SET current = current + 1 WHERE id = ?", player.TeamID)
		if err != nil {
			log.Println(err)
			err := tx.Rollback()
			if err != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error, Team Current Status not updated"})
			return
		}

		if err := tx.Commit(); err != nil {
			log.Println(err)
			err := tx.Rollback()
			if err != nil {
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusCreated, player)
	}
}

func DeletePlayer(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func GetPlayersByTeamID(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}
