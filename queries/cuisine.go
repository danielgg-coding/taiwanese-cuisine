package queries

import (
	"context"
	"database/sql"
	"log"

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
func GetCuisineScore(client *firestore.Client, id string) (*models.FirestoreCuisine, error) {
	dsnap, err := client.Collection("scores").Doc(id).Get(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// Extract the docuemt's data into a vault of type FirestoreCuisine
	var cuisine models.FirestoreCuisine
	if err := dsnap.DataTo(&cuisine); err != nil {
		return nil, err
	}

	return &cuisine, nil
}

// UpdateCuisineScore update an entry to firestore
func UpdateCuisineScore(client *firestore.Client, cuisine *models.FirestoreCuisine) error {
	result, err := client.Collection("scores").Doc(cuisine.ID).Set(context.Background(), cuisine)
	if err != nil {
		return err
	}
	log.Print(result)
	return nil
}
