package main

import (
	"github.com/danielgg-coding/taiwanese-cuisine/backend/controllers"
	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func main() {

	router := gin.Default()

	// Initiate firebase app
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	if err != nil {
		panic(err)
	}

	// Initiate firestore client
	client, err := app.Firestore(context.Background())

	defer client.Close()

	router.GET("/cuisinef/", controllers.GetCuisineFirestore(client))
	router.GET("/votef", controllers.VoteCuisineFirestore(client))
	router.Run() // listen and serve on 0.0.0.0:8080
}
