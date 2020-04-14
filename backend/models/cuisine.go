package models

// FirestoreCuisine struct
type FirestoreCuisine struct {
	Name   string `firestore:"string"`
	Image  string `firestore:"image"`
	Score  int    `firestore:"score"`
	Played int    `firestore:"played"`
	ID     string
}
