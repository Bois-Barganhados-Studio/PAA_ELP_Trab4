package main

import (
	"fmt"
	dp "main/dinamicprog"
	gd "main/greedy"
	"os"
	"strconv"
	"time"
)

func ErrorCheck(err error, message string) {
	if err != nil {
		fmt.Println(message)
		os.Exit(1)
	}
}

func main() {
	n_attempts := 1000
	n_size, err := strconv.Atoi(os.Args[1])
	ErrorCheck(err, "Error parsing input")

	method := os.Args[2]

	filename := "inputs/input" + fmt.Sprintf("%d", n_size) + ".txt"
	file, err := os.Open(filename)
	ErrorCheck(err, "Error opening file")
	defer file.Close()

	// Leitura dos inputs, o arquivo esta estruturado da seguinte forma:
	// Primeira linha: capacidade da mochila
	// Demais linhas: peso e valor de cada item
	maxWeight, weights, values := readInput(file)

	fmt.Println("Max Weight:", maxWeight)
	// if len(weights) > 25 {
	// 	fmt.Println("Weights:", weights[:25], "...")
	// 	fmt.Println("Values:", values[:25], "...")
	// } else {
	// 	fmt.Println("Weights:", weights)
	// 	fmt.Println("Values:", values)
	// }

	// Executa o algoritmo n_attempts vezes
	maxWeightCarried := 0
	itemsCarried := []int{}
	var initTime, endTime time.Time
	if method == "dp" {
		initTime = time.Now()
		for i := 0; i < n_attempts; i++ {
			maxWeightCarried, itemsCarried = dp.Knapsack(weights, values, maxWeight)
		}
		endTime = time.Now()
	} else if method == "gd" {
		initTime = time.Now()
		for i := 0; i < n_attempts; i++ {
			maxWeightCarried, itemsCarried = gd.Knapsack(weights, values, maxWeight)
		}
		endTime = time.Now()
	}

	// Soma os pesos
	totalWeight := 0
	for _, item := range itemsCarried {
		totalWeight += weights[item]
	}
	// Imprime o resultado
	fmt.Println("Max value Carried:", maxWeightCarried)
	fmt.Println("Total Weight:", totalWeight)
	fmt.Println("Items Carried:", itemsCarried)
	fmt.Println("Time:", endTime.Sub(initTime), " - ", endTime.Nanosecond()-initTime.Nanosecond())

}

func readInput(file *os.File) (int, []int, []int) {
	var maxWeight int
	var weights []int
	var values []int

	fmt.Fscanf(file, "%d\n", &maxWeight)

	for {
		var weight, value int
		_, err := fmt.Fscanf(file, "%d %d\n", &weight, &value)
		if err != nil {
			break
		}
		weights = append(weights, weight)
		values = append(values, value)
	}

	return maxWeight, weights, values
}
