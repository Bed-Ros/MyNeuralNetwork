package main

import "fmt"

func main() {
	qwe := NewNeuralNetwork()
	r, err := qwe.Calculate([]float64{1, 2, 3})
	fmt.Printf("%v %v", r, err)
}
