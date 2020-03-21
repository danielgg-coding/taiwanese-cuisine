package queries

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"taiwanese-cuisine/models"

	"cloud.google.com/go/firestore"
)

// GetCuisine query cuisine by id from DB
func GetCuisine(db *sql.DB, id int) (*models.Cuisine, error) {

	var cuisine models.Cuisine

	err := db.QueryRow("SELECT * FROM cuisine WHERE id=?", id).Scan(&cuisine.ID, &cuisine.Name, &cuisine.Score)

	if err != nil {
		return nil, err
	}

	return &cuisine, nil
}

// GetAllCuisine query all cuisine from DB
func GetAllCuisine(db *sql.DB) ([]models.Cuisine, error) {

	var rows *sql.Rows
	var err error

	rows, err = db.Query("SELECT * FROM cuisine")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cuisines := []models.Cuisine{}

	for rows.Next() {

		var cuisine models.Cuisine

		err := rows.Scan(&cuisine.ID, &cuisine.Name, &cuisine.Score)

		if err != nil {
			return nil, err
		}

		cuisines = append(cuisines, cuisine)
	}

	return cuisines, nil
}

// GetCuisineScore get current score of a cuisine
func GetCuisineScore(client *firestore.Client, id int) (int64, error) {
	dsnap, err := client.Collection("scores").Doc(strconv.Itoa(id)).Get(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// Extract the docuemt's data into a vault of type FirestoreCuisine
	var nyData models.FirestoreCuisine
	if err := dsnap.DataTo(&nyData); err != nil {
		return 0, err
	}

	return nyData.Score, nil
}

// UpdateCuisineScore update an entry to firestore
func UpdateCuisineScore(client *firestore.Client, id int, score int64, played int64) error {
	nyData := models.FirestoreCuisine{
		Played: played,
		Score:  score,
	}
	result, err := client.Collection("scores").Doc(strconv.Itoa(id)).Set(context.Background(), nyData)
	if err != nil {
		return err
	}
	log.Print(result)
	return nil
}
