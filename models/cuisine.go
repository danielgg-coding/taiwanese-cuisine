package models

// Cuisine ...
type Cuisine struct {
	ID    int64
	Name  string
	Score int64
}

// FirestoreCuisine ...
type FirestoreCuisine struct {
	ID     string `firestore:"id"`
	Score  int64  `firestore:"score"`
	Played int64  `firestore:"played"`
}
