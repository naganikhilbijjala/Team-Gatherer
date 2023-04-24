package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Team struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Time       string `json:"Time"`
	GameDate   string `json:"gameDate"`
	GameTime   string `json:"gameTime"`
	GamePeriod string `json:"gamePeriod"`
	Owner      int    `json:"Owner"`
	Min        int    `json:"min"`
	Max        int    `json:"max"`
	Current    int    `json:"current"`
	Location   string `json:"location"`
}

func CreateTeam(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var team Team
		if err := c.ShouldBindJSON(&team); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
			return
		}

		result, err := db.Exec("INSERT INTO teams (name, game_date, game_time, game_period, Owner, min, max, current, location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", team.Name, team.GameDate, team.GameTime, team.GamePeriod, team.Owner, team.Min, team.Max, team.Current, team.Location)
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

		var ownerName string
		err = db.QueryRow("SELECT name FROM users WHERE id = ?", team.Owner).Scan(&ownerName)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		_, err = db.Exec("INSERT INTO players (name, team_id, user_id) VALUES (?, ?, ?)", ownerName, team.ID, team.Owner)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error for inserting players"})
		}

		c.JSON(http.StatusCreated, team)
	}
}
func GetTeams(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
			err := rows.Scan(&team.ID, &team.Name, &team.Time, &team.Owner, &team.GameTime, &team.GameDate, &team.GamePeriod, &team.Max, &team.Min, &team.Current, &team.Location)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
				return
			}
			teams = append(teams, team)
		}
		c.JSON(http.StatusOK, teams)
	}
}

func GetTeamsByUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var teams []int

		userID := c.Query("user_id")

		rows, err := db.Query("SELECT team_id FROM players WHERE user_id = ?", userID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		defer rows.Close()

		for rows.Next() {
			var teamID int
			err := rows.Scan(&teamID)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
				return
			}

			teams = append(teams, teamID)
		}

		if err := rows.Err(); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"teams": teams})
	}
}

func UpdateTeam(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}

func DeleteTeam(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
}
