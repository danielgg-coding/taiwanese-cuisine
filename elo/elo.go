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
func Elorating(winnerScore, loserScore int) (int, int) {
	prob := prob(winnerScore, loserScore)
	delta := int(K * (1 - prob))

	winnerScore += delta
	loserScore -= delta

	return winnerScore, loserScore
}
