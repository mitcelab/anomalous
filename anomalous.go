package anomalous

import (
	"math"
)

type Detector struct {
	w, b float64
	forest *IsolationForest
}

// Return the probability that this sample is an anomaly.
func (detector *Detector) Predict(x []float64) (float64) {
	y := math.Exp(detector.forest.IsoScore(x) * detector.w + detector.b)
	return y / (y + 1.0)
}

// Create a new anomaly detector.
func NewDetector(X [][]float64) *Detector {
	detector := &Detector{
		forest: BuildForest(X, 7, 13),
	}

	minScore := 1e10
	maxScore := -1e10
	for _, x := range X {
		score := detector.forest.IsoScore(x)
		if score < minScore { minScore = score }
		if score > maxScore { maxScore = score }
	}

	// Using Mathematica, we calibrated these parameters to produce some probabilities:
	//
	//   sigmoid[x_] = E^x/(E^x + 1.0);
	//   Solve[{
	//     sigmoid[xmax*w + b] == 0.5,
	//     sigmoid[xmin*w + b] == 0.1
	//   }, {w, b}]
	//
	// Note that these probabilities are not well-calibrated.

	detector.w = 2.19722 / (1e-1 + maxScore - minScore)
	detector.b = - (2.19722 * maxScore) / (1e-1 + maxScore - minScore)
	return detector
}
