package app

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGameWithPlayers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("getGameWithPlayers")
		fmt.Println(c.Query("id"))
		var team Team
		gameId := c.Query("id")
		err := db.QueryRow("SELECT id, name, time, game_time, game_date, game_period, owner FROM teams WHERE id = ?", gameId).Scan(&team.ID, &team.Name, &team.Time, &team.GameTime, &team.GameDate, &team.GamePeriod, &team.Owner)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Game not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error-1"})
			}
			return
		}

		players := make([]Player, 0)
		rows, err := db.Query("SELECT id, name, team_id, user_id FROM players WHERE team_id = ?", gameId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error-2"})
			return
		}
		defer rows.Close()
		for rows.Next() {
			var player Player
			if err := rows.Scan(&player.ID, &player.Name, &player.TeamID, &player.UserID); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error-3"})
				return
			}
			players = append(players, player)
		}
		c.JSON(http.StatusOK, gin.H{"game": team, "players": players})
	}
}

func LeaveGame(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			TeamID int `json:"team_id"`
			UserID int `json:"user_id"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if the player exists in the team
		var playerID int
		err := db.QueryRow("SELECT id FROM players WHERE team_id = ? AND user_id = ?", req.TeamID, req.UserID).Scan(&playerID)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Player not found in team"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error-1"})
			}
			return
		}

		// Remove the player from the team
		_, err = db.Exec("DELETE FROM players WHERE id = ?", playerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error-2"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Player left the team"})
	}
}
