package anomalous

import "testing"

func TestNewDetector(t *testing.T) {
	X := [][]float64{
		{1.0, 0.0, 0.0},
		{1.0, 1.0, 0.0},
		{1.0, 2.0, 0.0},
		{1.0, 3.0, 0.0},
		{1.0, 4.0, 0.0},
		{1.0, 5.0, 0.0},
	}
	detector := NewDetector(X)

	if detector.Predict(X[0]) > detector.Predict([]float64{1.0, 4.5, 0.0}) {
		t.Error("the new sample should have higher probability of being an anomaly")
	}

	if detector.Predict([]float64{1.0, 4.5, 0.0}) > detector.Predict([]float64{1.0, 4.5, 1.0}) {
		t.Error("the new sample should have a much higher probability of being an anomaly")
	}
}
