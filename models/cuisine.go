package models

// Cuisine ...
type Cuisine struct {
	ID    int64
	Name  string
	Score int64
}

// FirestoreCuisine ...
type FirestoreCuisine struct {
	Name   string `firestore:"string"`
	Image  string `firestore:"image"`
	Score  int    `firestore:"score"`
	Played int    `firestore:"played"`
	ID     string
}

type DynamodbCuisine struct {
	ID     int `dynamodbav:"Id"`
	Name   string
	Image  string
	Score  int
	Played int
}
