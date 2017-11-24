// This package implements an "isolation forest" [Liu et al. 2008] for outlier detection.
package anomalous

type IsolationForest struct {
	trees []*IsolationTree
}

// Return a float indicating how anomalous the given sample is. Larger values indicate that the
// sample is more likely to be anomalous.
func (forest *IsolationForest) IsoScore(x []float64) float64 {
	pathLength := 0.0
	for _, tree := range forest.trees {
		pathLength += tree.pathLength(x)
	}
	pathLength = pathLength / float64(len(forest.trees))
	return 1.0 / (1.0 + pathLength)
}

// Train an isolation forest model for the given samples which contains `nbTrees` trees, each of
// which has a maximum depth of `maxDepth`. The training data `X` should be a slice which contains
// slices of samples.
func BuildForest(X [][]float64, maxDepth int, nbTrees int) *IsolationForest {
	forest := &IsolationForest{
		trees: make([]*IsolationTree, nbTrees),
	}
	for i := 0; i < nbTrees; i++ {
		forest.trees[i] = makeTree(X, maxDepth)
	}
	return forest
}
