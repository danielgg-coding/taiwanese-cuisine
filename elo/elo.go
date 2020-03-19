package elo

import (
	"math"
	"taiwanese-cuisine/models"
)

const (
	// K is the default K-factor
	K = 32
	// D is the default devision
	D float64 = 400.0
)

func prob(r1, r2 int64) float64 {
	return 1 / (1 + math.Pow(10, float64(r1-r2)/D))
}

// Elorating is the function calculate new rating score
func Elorating(winner, loser *models.FirestoreCuisine) (*models.FirestoreCuisine, *models.FirestoreCuisine) {
	prob := prob(loser.Score, winner.Score)
	delta := int64(K * (1 - prob))

	winner.Score += delta
	winner.Played++

	loser.Score -= delta
	loser.Played++

	return winner, loser
}
