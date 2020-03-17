package queries

import (
	"database/sql"
	"taiwanese-cuisine/models"
)

func GetCuisine(db *sql.DB, id int) (*models.Cuisine, error) {

	var cuisine models.Cuisine

	err := db.QueryRow("SELECT * FROM cuisine WHERE id=?", id).Scan(&cuisine.ID, &cuisine.Name, &cuisine.Score)

	if err != nil {
		return nil, err
	}

	return &cuisine, nil
}

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