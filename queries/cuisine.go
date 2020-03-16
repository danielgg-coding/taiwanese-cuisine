package queries

import (
	"database/sql"
	"taiwanese-cuisine/models"
)

func GetCuisine(db *sql.DB, id string) ([]models.Cuisine, error) {

	var rows *sql.Rows
	var err error

	if id == "all" {
		rows, err = db.Query("SELECT * FROM cuisine")
	} else {
		rows, err = db.Query("SELECT * FROM cuisine WHERE id = ?", id)
	}

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
