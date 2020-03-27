package elo

import (
	"math"
)

const (
	// K is the default K-factor
	K = 32
	// D is the default devision
	D float64 = 400.0
)

func prob(r1, r2 int) float64 {
	return 1 / (1 + math.Pow(10, float64(r1-r2)/D))
}

// Elorating is the function calculate new rating score
<<<<<<< HEAD
func Elorating(winner, loser *models.FirestoreCuisine) (*models.FirestoreCuisine, *models.FirestoreCuisine) {
	prob := prob(loser.Score, winner.Score)
	delta := int(K * (1 - prob))

	winner.Score += delta
	winner.Played++

	loser.Score -= delta
	loser.Played++

	return winner, loser
=======
func Elorating(winnerScore, loserScore int) (int, int) {
	prob := prob(winnerScore, loserScore)
	delta := int(K * (1 - prob))

	winnerScore += delta
	loserScore -= delta

	return winnerScore, loserScore
>>>>>>> 9a272c18d9f1d1faac822150595c3603eac9e58f
}
