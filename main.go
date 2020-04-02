package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/danielgg-coding/taiwanese-cuisine/controllers"
	"github.com/gin-gonic/gin"

	firebase "firebase.google.com/go"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

func main() {

	router := gin.Default()

	// db, err := sql.Open("mysql", "root:thedaniel@tcp(localhost:3306)/taiwancuisine")

	// if err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }

	// Initiate firebase app
	sa := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, sa)

	if err != nil {
		panic(err)
	}

	// Initiate firestore client
	client, err := app.Firestore(context.Background())

	dynamodbSess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-2")},
	)

	// Create DynamoDB client
	dyanmodbClient := dynamodb.New(dynamodbSess)

	defer client.Close()

	router.GET("/", controllers.Index())
	router.GET("/cuisinef/", controllers.GetCuisineFirestore(client))
	router.GET("/votef", controllers.VoteCuisineFirestore(client))
	router.GET("/cuisinec", controllers.GetCuisineDynamo(dyanmodbClient))
	router.GET("/votec", controllers.VoteCuisineDynamo(dyanmodbClient))
	port := os.ExpandEnv(":$PORT")
	router.Run(port) // listen and serve on 0.0.0.0:8081
}
