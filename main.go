package main

import (
	"database/sql"

	"github.com/danielgg-coding/taiwanese-cuisine/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	db, err := sql.Open("mysql", "root:thedaniel@tcp(localhost:3306)/taiwancuisine")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer db.Close()

	router.GET("/cuisine/:cuisineId", controllers.GetCuisine(db))
	router.GET("/cuisines", controllers.GetAllCuisine(db))
	router.GET("/vote", controllers.VoteCuisine(db))
	router.Run() // listen and serve on 0.0.0.0:8080
}
