package anomalous

import "testing"

func TestBuildTreeBinaryFeatures(t *testing.T) {
	X := [][]float64{
		{1.0, 0.0, 0.0},
		{1.0, 1.0, 0.0},
	}
	tree := makeTree(X, 3)

	pos1 := tree.pathLength([]float64{1.0, 0.0, 0.0})
	pos2 := tree.pathLength([]float64{1.0, 1.0, 0.0})

	var neg float64

	neg = tree.pathLength([]float64{0.0, 0.0, 1.0})
	if pos1 < neg || pos2 < neg {
		t.Error("we expect the clearly different x sample to have a shorter path length")
	}

	neg = tree.pathLength([]float64{0.0, 1.0, 1.0})
	if pos1 < neg || pos2 < neg {
		t.Error("we expect the clearly different x sample to have a shorter path length")
	}
}

func TestBuildTreeRealFeatures(t *testing.T) {
	X := [][]float64{
		{1.0, 0.0},
		{1.0, 1.0},
	}
	tree := makeTree(X, 3)

	pos1 := tree.pathLength([]float64{1.0, 0.0})
	pos2 := tree.pathLength([]float64{1.0, 1.0})

	var neg float64

	neg = tree.pathLength([]float64{0.0, 1000.0})
	if pos1 < neg || pos2 < neg {
		t.Error("we expect the clearly different x sample to have a shorter path length")
	}

	neg = tree.pathLength([]float64{0.1, 0.1})
	if pos1 < neg || pos2 < neg {
		t.Error("we expect the clearly different x sample to have a shorter path length")
	}
}
