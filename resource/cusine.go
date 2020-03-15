package resource

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetCuisine(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		cuisineId := c.Param("cuisineId")

		var cuisineName string

		row := db.QueryRow("SELECT cuisine_name FROM cuisine WHERE id = ?", cuisineId)

		err := row.Scan(&cuisineName)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		c.String(200, cuisineName)
	}

	return gin.HandlerFunc(fn)
}
