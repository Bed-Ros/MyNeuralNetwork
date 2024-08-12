package main

import "math/rand"

type Layer []*Neuron

type NeuralNetwork struct {
	Layers []Layer
}

func NewNeuralNetwork() *NeuralNetwork {
	var result NeuralNetwork

	//Тестовое количество слоев и нейронов в них
	numberOfNeuronsPerLayer := []int{3, 2, 1}

	//Создание слоев
	for _, n := range numberOfNeuronsPerLayer {
		var newLayer Layer
		for i := 0; i < n; i++ {
			newLayer = append(newLayer, &Neuron{})
		}
		result.Layers = append(result.Layers, newLayer)
	}

	//Назначение связей нейронов между слоями со случайными значениями
	for i := 0; i < len(result.Layers)-1; i++ {
		fromLayer := result.Layers[i]
		toLayer := result.Layers[i+1]

		for _, fromNeuron := range fromLayer {
			for _, toNeuron := range toLayer {
				newConn := &Connection{
					Weight:     rand.Float64(),
					FromNeuron: fromNeuron,
					ToNeuron:   toNeuron,
				}
				fromNeuron.Out = append(fromNeuron.Out, newConn)
				toNeuron.In = append(fromNeuron.Out, newConn)
			}
		}
	}

	return &result
}
