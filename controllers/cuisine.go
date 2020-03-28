package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/danielgg-coding/taiwanese-cuisine/elo"

	"github.com/danielgg-coding/taiwanese-cuisine/queries"

	"github.com/gin-gonic/gin"

	"cloud.google.com/go/firestore"
)

func Index() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.String(200, "Everthing is fine, YAY !!!")
	}
	return gin.HandlerFunc(fn)
}

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
			panic(err)
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

		start := time.Now()
		players, err := queries.GetCuisines(client, []string{winnerID, loserID})
		if err != nil {
			c.String(404, fmt.Sprint("Id not found"))
			panic(err)
			// panic(err)
		}
		winner, loser := players[0], players[1]
		fmt.Println(time.Now().Sub(start))
		// winner, err := queries.GetCuisineFromFire(client, winnerID)
		// if err != nil {
		// 	panic(err)
		// }

		// loser, err := queries.GetCuisineFromFire(client, loserID)
		// if err != nil {
		// 	panic(err)
		// }

		winner.Score, loser.Score = elo.Elorating(winner.Score, loser.Score)
		winner.Played++
		loser.Played++

		err = queries.UpdateCuisineToFire(client, winner, winnerID)
		if err != nil {
			panic(err)
		}

		err = queries.UpdateCuisineToFire(client, loser, loserID)
		if err != nil {
			panic(err)
		}

		c.String(200, fmt.Sprintf("id %s beats id %s", winnerID, loserID))
	}

	return gin.HandlerFunc(fn)
}
