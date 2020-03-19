package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"taiwanese-cuisine/elo"
	"taiwanese-cuisine/queries"

	"github.com/gin-gonic/gin"

	"cloud.google.com/go/firestore"
)

// GetCuisine get cuisine by id
func GetCuisine(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		cuisineIDString := c.Param("cuisineId")
		cuisineID, err := strconv.Atoi(cuisineIDString)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		cuisines, err := queries.GetCuisine(db, cuisineID)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		c.JSON(200, cuisines)
	}

	return gin.HandlerFunc(fn)

}

// GetAllCuisine get all cuisine
func GetAllCuisine(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		cuisines, err := queries.GetAllCuisine(db)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		c.JSON(200, cuisines)
	}

	return gin.HandlerFunc(fn)

}

// VoteCuisine write comparison entry
func VoteCuisine(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		winner := c.Query("winner")
		losser := c.Query("losser")

		c.String(200, fmt.Sprintf("id %s beats id %s", winner, losser))
	}

	return gin.HandlerFunc(fn)
}

// GetCuisineFirestore gets cuisine data from firestore
func GetCuisineFirestore(client *firestore.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		dsnap, err := client.Collection("cuisine").Doc("food").Get(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		m := dsnap.Data()
		c.JSON(200, m)
	}

	return gin.HandlerFunc(fn)
}

// VoteCuisineFirestore update cuisine data to firestore
func VoteCuisineFirestore(client *firestore.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		winnerID := c.Query("winner")
		loserID := c.Query("loser")

		winner, err := queries.GetCuisineScore(client, winnerID)
		if err != nil {
			log.Fatalln(err)
		}

		loser, err := queries.GetCuisineScore(client, loserID)
		if err != nil {
			log.Fatalln(err)
		}

		winner, loser = elo.Elorating(winner, loser)

		err = queries.UpdateCuisineScore(client, winner)
		if err != nil {
			log.Fatalln(err)
		}

		err = queries.UpdateCuisineScore(client, loser)
		if err != nil {
			log.Fatalln(err)
		}

		c.String(200, fmt.Sprintf("id %s beats id %s", winnerID, loserID))
	}

	return gin.HandlerFunc(fn)
}
