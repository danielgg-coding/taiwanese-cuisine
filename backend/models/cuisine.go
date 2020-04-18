package models

// FirestoreCuisine struct
type FirestoreCuisine struct {
	Name   string `firestore:"string" json:"name"`
	Image  string `firestore:"image" json:"image"`
	Score  int    `firestore:"score" json:"score"`
	Played int    `firestore:"played" json:"played"`
	ID     string `json:"id"`
}
