package main

import (
	app "SPL-Spring2023/Backend/app"
	"database/sql"
	_ "fmt"
	"log"
	_ "strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Connect to the database
	var err error
	db, err = sql.Open("mysql", "root:Nikhil@22@tcp(127.0.0.1:3306)/TEAMPROJECT")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the Gin router
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // replace with your frontend URL
	r.Use(cors.New(config))

	// Define the routes
	r.GET("/teams", app.GetTeams(db))
	r.POST("/teams", app.CreateTeam(db))
	r.DELETE("/teams/:id", app.DeleteTeam(db))
	r.PUT("/teams/:id", app.UpdateTeam(db))
	//r.GET("search/:name", searchTeam)
	r.GET("/players", app.GetPlayers(db))
	r.POST("/players", app.CreatePlayer(db))
	r.DELETE("/players/:id", app.DeletePlayer(db))
	r.GET("getPlayers/:id", app.GetPlayersByTeamID(db))
	r.POST("/register", app.RegisterUser(db))
	r.POST("/check", app.CheckUser(db))
	r.GET("/getUserInfo", app.GetUserInfo(db))
	r.GET("/getTeamsByUser", app.GetTeamsByUser(db))
	r.GET("/getGameWithPlayers", app.GetGameWithPlayers(db))
	r.POST("/leave-game", app.LeaveGame(db))
	// Start the server
	r.Run(":8080")
}
