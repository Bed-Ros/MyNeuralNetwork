package main

import (
	"errors"
	"math/rand"
)

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
				toNeuron.In = append(toNeuron.In, newConn)
			}
		}
	}

	return &result
}

func (n *NeuralNetwork) Calculate(inputs []float64) ([]float64, error) {
	//Проверяем количество вводных данных
	if len(inputs) != len(n.Layers[0]) {
		return nil, errors.New("количество вводных значений не равно количеству нейронов входного уровня")
	}
	//Вводим значения для входного слоя
	for i, inputNeuron := range n.Layers[0] {
		inputNeuron.Value = &inputs[i]
	}
	//Считаем знаяения для всех последующих слоев
	for _, layer := range n.Layers {
		for _, neuron := range layer {
			err := neuron.CalcValue()
			if err != nil {
				return nil, err
			}
		}
	}
	//Возвращаем значения нейронов выходного слоя
	var result []float64
	for _, neuron := range n.Layers[len(n.Layers)-1] {
		result = append(result, *neuron.Value)
	}
	return result, nil
}
