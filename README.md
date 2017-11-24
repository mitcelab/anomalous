# anomalous
[![Build Status](https://travis-ci.org/mitcelab/anomalous.svg?branch=master)](https://travis-ci.org/mitcelab/anomalous)
[![GoDoc](https://godoc.org/github.com/mitcelab/anomalous?status.svg)](https://godoc.org/github.com/mitcelab/anomalous)

Anomaly detection in Go with isolation forests.

### install
> go install github.com/mitcelab/anomalous

### usage
```go
X := [][]float64{
	{1.0, 0.0, 0.0},
	{1.0, 1.0, 0.0},
	{1.0, 2.0, 0.0},
	{1.0, 3.0, 0.0},
	{1.0, 4.0, 0.0},
	{1.0, 5.0, 0.0},
}
detector := NewDetector(X)

positive := detector.Predict(X[0])
maybe := detector.Predict([]float64{1.0, 4.5, 0.0})
no := detector.Predict([]float64{1.0, 4.5, 1.0})

if positive > maybe {
	log.Fatal("the `maybe` sample should have higher anomaly probability")
}

if maybe > no {
	log.Fatal("the `no` sample should have a *much* higher anomaly probability")
}
```
