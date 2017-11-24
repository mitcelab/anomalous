package anomalous

import "testing"

func TestBuildForestBinaryFeatures(t *testing.T) {
	X := [][]float64{
		{1.0, 0.0, 0.0},
		{1.0, 1.0, 0.0},
	}
	forest := BuildForest(X, 10, 100)

	pos1 := forest.IsoScore([]float64{1.0, 0.0, 0.0})
	pos2 := forest.IsoScore([]float64{1.0, 1.0, 0.0})

	var neg float64

	neg = forest.IsoScore([]float64{0.0, 0.0, 1.0})
	if pos1 > neg || pos2 > neg {
		t.Error("we expect the clearly different x sample to have a shorter path length")
	}

	neg = forest.IsoScore([]float64{0.0, 1.0, 1.0})
	if pos1 > neg || pos2 > neg {
		t.Error("we expect the clearly different x sample to have a shorter path length")
	}
}

func TestBuildForestRealFeatures(t *testing.T) {
	X := [][]float64{
		{1.0, 0.0},
		{1.0, 1.0},
	}
	forest := BuildForest(X, 10, 100)

	pos1 := forest.IsoScore([]float64{1.0, 0.0})
	pos2 := forest.IsoScore([]float64{1.0, 1.0})

	var neg float64

	neg = forest.IsoScore([]float64{0.0, 1000.0})
	if pos1 > neg || pos2 > neg {
		t.Error("we expect the clearly different x sample to have a shorter path length")
	}

	neg = forest.IsoScore([]float64{0.1, 0.1})
	if pos1 > neg || pos2 > neg {
		t.Error("we expect the clearly different x sample to have a shorter path length")
	}
}
