package anomalous

import "math/rand"

type IsolationTree struct {
	N     int
	dim   int
	val   float64
	left  *IsolationTree
	right *IsolationTree
}

// Recursively traverse the tree until no more splits are possible.
func (tree *IsolationTree) pathLength(x []float64) float64 {
	if tree.N == 0 {
		return 0.0
	}
	if x[tree.dim] <= tree.val {
		return 1.0 + tree.left.pathLength(x)
	} else {
		return 1.0 + tree.right.pathLength(x)
	}
}

// Build an isolation tree by randomly selecting dimensions and splits until either maxDepth is
// reached or all samples have been split into their own leaves.
func makeTree(X [][]float64, maxDepth int) *IsolationTree {
	nbSamples := len(X)
	if nbSamples == 0 || maxDepth == 0 {
		return &IsolationTree{N: 0}
	}

	root := &IsolationTree{}
	root.N = nbSamples
	root.dim = rand.Intn(len(X[0]))

	minValue := 1e10
	maxValue := -1e10
	for _, x := range X {
		if x[root.dim] < minValue {
			minValue = x[root.dim]
		}
		if x[root.dim] > maxValue {
			maxValue = x[root.dim]
		}
	}
	root.val = rand.Float64()*(maxValue-minValue) + minValue

	var leftX [][]float64
	var rightX [][]float64
	for _, x := range X {
		if x[root.dim] <= root.val {
			leftX = append(leftX, x)
		} else {
			rightX = append(rightX, x)
		}
	}
	root.left = makeTree(leftX, maxDepth-1)
	root.right = makeTree(rightX, maxDepth-1)

	return root
}
