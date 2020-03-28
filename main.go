package main

import (
	"database/sql"

	"github.com/danielgg-coding/taiwanese-cuisine/controllers"

	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func main() {

	router := gin.Default()

	db, err := sql.Open("mysql", "root:thedaniel@tcp(localhost:3306)/taiwancuisine")

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Initiate firebase app
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	if err != nil {
		panic(err)
	}

	// Initiate firestore client
	client, err := app.Firestore(context.Background())

	defer client.Close()
	defer db.Close()
	router.GET("/", controllers.Index())
	router.GET("/cuisinef/", controllers.GetCuisineFirestore(client))
	router.GET("/votef", controllers.VoteCuisineFirestore(client))
	router.GET("/cuisine/:cuisineId", controllers.GetCuisine(db))
	router.GET("/cuisines", controllers.GetAllCuisine(db))
	router.GET("/vote", controllers.VoteCuisine(db))
	router.Run(":8081") // listen and serve on 0.0.0.0:8081
}
