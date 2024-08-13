package main

import (
	"errors"
	"math"
)

type Connection struct {
	Weight     float64
	FromNeuron *Neuron
	ToNeuron   *Neuron
}

type Neuron struct {
	Value *float64
	In    []*Connection
	Out   []*Connection
}

func (n *Neuron) CalcValue() error {
	//Считаем сумму входных сигналов
	var inputSum float64
	for _, connection := range n.In {
		//Проверяем есть ли у предыдущего нейрона значение
		if connection.FromNeuron.Value == nil {
			return errors.New("невозможно расчитать значение нейрона, " +
				"потому что у одного из предыдущих нейронов отсутствует значение")
		}
		inputSum += connection.Weight * *connection.FromNeuron.Value
	}
	//Применяем активационную функцию
	//	в данном случает это Сигмоида y=1/(1+e^(-x))
	v := 1 / (1 + math.Exp(-inputSum))
	n.Value = &v
	return nil
}
