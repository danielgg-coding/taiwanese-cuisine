package controllers

import (
	"database/sql"
	"fmt"
	"taiwanese-cuisine/queries"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetCuisine(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		cuisineId := c.Param("cuisineId")

		cuisines, err := queries.GetCuisine(db, cuisineId)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		c.JSON(200, cuisines)
	}

	return gin.HandlerFunc(fn)
}

func GetAllCuisine(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		cuisines, err := queries.GetCuisine(db, "all")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		c.JSON(200, cuisines)
	}
	return gin.HandlerFunc(fn)
}

func VoteCuisine(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		winner := c.Query("winner")
		losser := c.Query("losser")

		c.String(200, fmt.Sprintf("id %s beats id %s", winner, losser))
	}

	return gin.HandlerFunc(fn)
}
