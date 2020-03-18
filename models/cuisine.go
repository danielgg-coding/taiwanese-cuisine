package models

type Cuisine struct {
	ID    int64
	Name  string
	Score int64
}

type FirestoreCuisine struct {
	Score  int64 `firestore:"score"`
	Played int64 `firestore:"played"`
}
