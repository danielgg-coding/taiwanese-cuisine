package queries

import (
	"context"
	"database/sql"
	"log"

	"github.com/danielgg-coding/taiwanese-cuisine/models"

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
