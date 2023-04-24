package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user struct {
			Email    string `json:"email"`
			Passcode string `json:"passcode"`
			Name     string `json:"name"`
		}
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Check if the user with the given email already exists in the database
		var count int
		if err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", user.Email).Scan(&count); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - User Exists"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User with the given email already exists"})
			return
		}

		// Insert the user into the database
		result, err := db.Exec("INSERT INTO users (email, passcode, name) VALUES (?, ?, ?)", user.Email, user.Passcode, user.Name)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error - Insert Error"})
			return
		}

		// Get the ID of the newly inserted user
		id, err := result.LastInsertId()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Return the ID of the newly inserted user
		c.JSON(http.StatusCreated, gin.H{"id": id})

	}
}

func GetUserInfo(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get email from the request query parameters
		email := c.Query("email")
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email parameter is missing"})
			return
		}

		// Query the database to retrieve the ID and name of the user with the given email
		var id int
		var name string
		err := db.QueryRow("SELECT id, name FROM users WHERE email = ?", email).Scan(&id, &name)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
			return
		}

		// Return the ID and name of the user
		c.JSON(http.StatusOK, gin.H{"id": id, "name": name})
	}
}

func CheckUser(db *sql.DB) gin.HandlerFunc {
	// Get email and passcode from the request body
	return func(c *gin.Context) {
		var user struct {
			Email    string `json:"email"`
			Passcode string `json:"passcode"`
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Query the database to check if the user exists
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ? AND passcode = ?", user.Email, user.Passcode).Scan(&count)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Return the result
		if count > 0 {
			c.JSON(http.StatusOK, gin.H{"message": "User exists"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "User does not exist"})
		}
	}
}
