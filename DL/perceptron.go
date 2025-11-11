// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )


// type Perceptron struct {
// 	weights      []float64
// 	learningRate float64
// 	epochs       int
// }


// func NewPerceptron(numInputs int, learningRate float64, epochs int) *Perceptron {

// 	rand.Seed(time.Now().UnixNano())

// 	weights := make([]float64, numInputs+1)
// 	for i := range weights {

// 		weights[i] = rand.Float64()
// 	}

// 	return &Perceptron{
// 		weights:      weights,
// 		learningRate: learningRate,
// 		epochs:       epochs,
// 	}
// }

// // activationFunction implements the Heaviside step function.
// func (p *Perceptron) activationFunction(z float64) int {
// 	if z >= 0 {
// 		return 1
// 	}
// 	return 0
// }

// func (p *Perceptron) Predict(inputs []float64) int {
// 	// The weighted sum starts with the bias term (p.weights[0])
// 	weightedSum := p.weights[0]

// 	// Calculate the dot product of inputs and feature weights (p.weights[1:])
// 	for i, input := range inputs {
// 		// i+1 because p.weights[0] is the bias
// 		weightedSum += input * p.weights[i+1]
// 	}

// 	return p.activationFunction(weightedSum)
// }

// func (p *Perceptron) Train(X [][]float64, y []int) {
// 	numSamples := len(X)

// 	for epoch := 1; epoch <= p.epochs; epoch++ {
// 		totalError := 0
// 		for i := 0; i < numSamples; i++ {
// 			inputs := X[i]
// 			target := y[i]

// 			prediction := p.Predict(inputs)

// 			// Update weights and bias only if the prediction is incorrect
// 			if prediction != target {
// 				error := float64(target - prediction)
// 				totalError += 1

// 				// Update bias (p.weights[0])
// 				p.weights[0] += p.learningRate * error

// 				// Update feature weights (p.weights[1:])
// 				for j, input := range inputs {
// 					// j+1 because p.weights[0] is the bias
// 					p.weights[j+1] += p.learningRate * error * input
// 				}
// 			}
// 		}
// 		// If totalError is 0, the model has converged, and we can stop training early.
// 		if totalError == 0 {
// 			// fmt.Printf("Converged after epoch %d\n", epoch)
// 			break
// 		}
// 	}
// }

// func main() {

// 	XAnd := [][]float64{
// 		{0.0, 0.0},
// 		{0.0, 1.0},
// 		{1.0, 0.0},
// 		{1.0, 1.0},
// 	}
// 	yAnd := []int{0, 0, 0, 1}

	
// 	perceptronAnd := NewPerceptron(2, 0.1, 10)
// 	perceptronAnd.Train(XAnd, yAnd)

// 	fmt.Println("--- AND Gate Predictions (after training) ---")
// 	for i, inputs := range XAnd {
// 		prediction := perceptronAnd.Predict(inputs)
// 		fmt.Printf("Input: %v, Target: %d, Prediction: %d\n", inputs, yAnd[i], prediction)
// 	}


// 	XOr := [][]float64{
// 		{0.0, 0.0},
// 		{0.0, 1.0},
// 		{1.0, 0.0},
// 		{1.0, 1.0},
// 	}
// 	yOr := []int{0, 1, 1, 1}

// 	perceptronOr := NewPerceptron(2, 0.1, 10)
// 	perceptronOr.Train(XOr, yOr)

// 	fmt.Println("\n--- OR Gate Predictions (after training) ---")
// 	for i, inputs := range XOr {
// 		prediction := perceptronOr.Predict(inputs)
// 		fmt.Printf("Input: %v, Target: %d, Prediction: %d\n", inputs, yOr[i], prediction)
// 	}
// }