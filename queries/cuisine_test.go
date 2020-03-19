package queries

import (
	"testing"

	"taiwanese-cuisine/models"

	_ "github.com/DATA-DOG/go-sqlmock"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestQueryALLCuisine(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	// testing with db containing correct rows
	rows := sqlmock.NewRows([]string{"id", "name", "score"}).AddRow(1, "水餃", 0).AddRow(2, "韭菜", 0)

	mock.ExpectQuery("SELECT (.+) FROM cuisine").
		WillReturnRows(rows)

	cuisines, err := GetAllCuisine(db)

	var expectedCuisines []models.Cuisine

	cuisine1 := models.Cuisine{
		ID:    1,
		Name:  "水餃",
		Score: 0,
	}

	cuisine2 := models.Cuisine{
		ID:    2,
		Name:  "韭菜",
		Score: 0,
	}

	expectedCuisines = append(expectedCuisines, cuisine1)
	expectedCuisines = append(expectedCuisines, cuisine2)

	assert.Equal(t, expectedCuisines, cuisines)

	// testing with db containing empty rows
	rows = sqlmock.NewRows([]string{"id", "name", "score"})

	mock.ExpectQuery("SELECT (.+) FROM cuisine*").WillReturnRows(rows)

	cuisines, err = GetAllCuisine(db)

	assert.Equal(t, []models.Cuisine{}, cuisines)

	// testing with db containing wrong schema
	rows = sqlmock.NewRows([]string{"id", "name", "score", "bad_column"}).AddRow(1, "水餃", 0, true)

	mock.ExpectQuery("SELECT (.+) FROM cuisine").WillReturnRows(rows)

	cuisines, err = GetAllCuisine(db)

	assert.Nil(t, cuisines)

}

func TestQueryOneCuisine(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	testingID := 1

	// testing with db containing correct rows
	rows := sqlmock.NewRows([]string{"id", "name", "score"}).AddRow(1, "水餃", 0)

	mock.ExpectQuery("SELECT (.+) FROM cuisine*").
		WithArgs(testingID).
		WillReturnRows(rows)

	cuisines, err := GetCuisine(db, testingID)

	expectedCuisine := models.Cuisine{
		ID:    1,
		Name:  "水餃",
		Score: 0,
	}

	assert.Equal(t, &expectedCuisine, cuisines)

	// testing with db containing empty rows
	rows = sqlmock.NewRows([]string{"id", "name", "score"})

	mock.ExpectQuery("SELECT (.+) FROM cuisine*").
		WithArgs(testingID).
		WillReturnRows(rows)

	cuisines, err = GetCuisine(db, testingID)

	expectedCuisine = models.Cuisine{}
	assert.Nil(t, cuisines)

	// testing with db containing wrong schema
	rows = sqlmock.NewRows([]string{"id", "name", "score", "bad_column"}).AddRow(1, "水餃", 0, true)

	mock.ExpectQuery("SELECT (.+) FROM cuisine").
		WithArgs(testingID).
		WillReturnRows(rows)

	cuisines, err = GetCuisine(db, testingID)

	assert.Nil(t, cuisines)
}
