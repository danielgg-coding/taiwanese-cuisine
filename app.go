package main

import (
	"database/sql"

	"taiwanese-cuisine/resource"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	db, err := sql.Open("mysql", "root:thedaniel@tcp(localhost:3306)/taiwancuisine")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer db.Close()

	router.GET("/cuisine/:cuisineId", resource.GetCuisine(db))
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
