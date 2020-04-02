package queries

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/danielgg-coding/taiwanese-cuisine/backend/models"

	"cloud.google.com/go/firestore"
)

// GetCuisineFromFire get current info of a cuisine
func GetCuisineFromFire(client *firestore.Client, id string) (*models.FirestoreCuisine, error) {
	doc, err := client.Collection("cuisine").Doc(id).Get(context.Background())
	if err != nil {
		panic(err)
	}
	// Extract the docuemt's data into a vault of type FirestoreCuisine
	var cuisine models.FirestoreCuisine
	if err := doc.DataTo(&cuisine); err != nil {
		return nil, err
	}
	return &cuisine, nil
}

// GetAllFromFire get all cuisine from firestore
func GetAllFromFire(client *firestore.Client) ([]*models.FirestoreCuisine, error) {
	docrefs, err := client.Collection("cuisine").DocumentRefs(context.Background()).GetAll()
	if err != nil {
		panic(err)
	}

	docs, err := client.GetAll(context.Background(), docrefs)
	if err != nil {
		panic(err)
	}

	var cuisines []*models.FirestoreCuisine
	for _, doc := range docs {
		var cuisine models.FirestoreCuisine
		if err := doc.DataTo(&cuisine); err != nil {
			return nil, err
		}
		cuisine.ID = doc.Ref.ID
		cuisines = append(cuisines, &cuisine)
	}
	return cuisines, nil
}

// GetCuisines get cuisines by id list from Firestore
func GetCuisines(client *firestore.Client, ids []string) ([]*models.FirestoreCuisine, error) {
	var docrefs []*firestore.DocumentRef
	for _, id := range ids {
		docrefs = append(docrefs, client.Collection("cuisine").Doc(id))
	}

	docs, err := client.GetAll(context.Background(), docrefs)

	if err != nil {
		panic(err)
	}
	var cuisines []*models.FirestoreCuisine
	for _, doc := range docs {
		var cuisine models.FirestoreCuisine
		if err := doc.DataTo(&cuisine); err != nil {
			return nil, err
		}
		cuisines = append(cuisines, &cuisine)
	}
	return cuisines, nil
}

// UpdateCuisineToFire update an entry to firestore
func UpdateCuisineToFire(client *firestore.Client, cuisine *models.FirestoreCuisine, id string) error {
	result, err := client.Collection("cuisine").Doc(id).Set(context.Background(), cuisine)
	if err != nil {
		return err
	}
	log.Print(result)
	return nil
}

// Testing query one from DynamoDB
func GetOneCuisinesFromDynamodb(dyanmodbClient *dynamodb.DynamoDB, id string) (models.DynamodbCuisine, error) {

	tableName := "food"

	result, err := dyanmodbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(id),
			},
		},
	})

	item := models.DynamodbCuisine{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return item, nil
}

// Testing query all from DynamoDB
func GetCuisinesFromDynamodb(dyanmodbClient *dynamodb.DynamoDB) ([]models.DynamodbCuisine, error) {

	tableName := "food"

	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	queryResult, _ := dyanmodbClient.Scan(params)
	items := []models.DynamodbCuisine{}

	for _, i := range queryResult.Items {
		item := models.DynamodbCuisine{}
		err := dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Testing update to DynamoDB
func UpdateDynamodb(dyanmodbClient *dynamodb.DynamoDB, item models.DynamodbCuisine) error {

	tableName := "food"

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":newScore": {
				N: aws.String(strconv.Itoa(item.Score)),
			},
			":newPlayed": {
				N: aws.String(strconv.Itoa(item.Played)),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(item.ID)),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Score = :newScore, Played = :newPlayed"),
	}

	_, err := dyanmodbClient.UpdateItem(input)

	if err != nil {
		return err
	}

	return nil
	// fmt.Println("Successfully updated '" + movieName + "' (" + movieYear + ") rating to " + movieRating)

}

// GetCuisine query cuisine by id from DB
// func GetCuisine(db *sql.DB, id int) (*models.Cuisine, error) {

// 	var cuisine models.Cuisine

// 	err := db.QueryRow("SELECT * FROM cuisine WHERE id=?", id).Scan(&cuisine.ID, &cuisine.Name, &cuisine.Score)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &cuisine, nil
// }

// GetAllCuisine query all cuisine from DB
// func GetAllCuisine(db *sql.DB) ([]models.Cuisine, error) {

// 	var rows *sql.Rows
// 	var err error

// 	rows, err = db.Query("SELECT * FROM cuisine")

// 	if err != nil {
// 		return nil, err
// 	}

// 	defer rows.Close()

// 	cuisines := []models.Cuisine{}

// 	for rows.Next() {

// 		var cuisine models.Cuisine

// 		err := rows.Scan(&cuisine.ID, &cuisine.Name, &cuisine.Score)

// 		if err != nil {
// 			return nil, err
// 		}

// 		cuisines = append(cuisines, cuisine)
// 	}

// 	return cuisines, nil
// }
