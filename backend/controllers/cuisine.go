package controllers

import (
	"fmt"

	"github.com/danielgg-coding/taiwanese-cuisine/backend/elo"

	"github.com/danielgg-coding/taiwanese-cuisine/backend/queries"
	"github.com/gin-gonic/gin"

	"cloud.google.com/go/firestore"
)

// GetCuisineFirestore gets cuisine data from firestore
func GetCuisineFirestore(client *firestore.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		cuisines, err := queries.GetAllFromFire(client)
		if err != nil {
			panic(err)
		}
		c.JSON(200, cuisines)
	}

	return gin.HandlerFunc(fn)
}

// VoteCuisineFirestore update cuisine data to firestore
func VoteCuisineFirestore(client *firestore.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		winnerID := c.Query("winner")
		loserID := c.Query("loser")

		players, err := queries.GetCuisinesFromFire(client, []string{winnerID, loserID})
		if err != nil {
			c.String(404, fmt.Sprint("Id not found"))
			panic(err)
			// panic(err)
		}
		winner, loser := players[0], players[1]
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
