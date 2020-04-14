package queries

import (
	"context"
	"log"

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

// GetCuisinesFromFire get cuisines by id list from Firestore
func GetCuisinesFromFire(client *firestore.Client, ids []string) ([]*models.FirestoreCuisine, error) {
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
