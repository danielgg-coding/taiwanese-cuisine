package main

import (
	"os"

	firebase "firebase.google.com/go"
	"github.com/danielgg-coding/taiwanese-cuisine/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func main() {

	seed := os.Getenv("SEED")
	key := DecryptFile("serviceAccountKey", seed)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// Initiate firebase app
	sa := option.WithCredentialsJSON(key)
	app, err := firebase.NewApp(context.Background(), nil, sa)

	if err != nil {
		panic(err)
	}

	// Initiate firestore client
	client, err := app.Firestore(context.Background())

	defer client.Close()

	router.Use(static.Serve("/", static.LocalFile("./build/", true)))

	router.GET("/cuisinef/", controllers.GetCuisineFirestore(client))
	router.GET("/votef", controllers.VoteCuisineFirestore(client))
	router.Run() // listen and serve on 0.0.0.0:8080
}
