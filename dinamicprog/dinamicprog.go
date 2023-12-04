package dinamicprog

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func makeTable(size int, capacity int) [][]int {
	table := make([][]int, size+1)
	for i := 0; i <= size; i++ {
		table[i] = make([]int, capacity+1)
		for j := 0; j < len(table[i]); j++ {
			table[i][j] = 0
		}
	}
	return table
}

func printTable(table [][]int) {
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			print(table[i][j], " ")
		}
		println()
	}
}

func Knapsack(weights []int, values []int, capacity int) (int, []int) {
	maxWeightCarried := 0
	itemsCarried := []int{}

	size := len(weights)

	// Cria e inicializa a tabela de memória
	table := makeTable(size, capacity)

	// printTable(table)

	for i := 1; i <= size; i++ {
		for w := 1; w <= capacity; w++ {
			if weights[i-1] > w {
				table[i][w] = table[i-1][w]
			} else {
				table[i][w] = max(table[i-1][w], values[i-1]+table[i-1][w-weights[i-1]])
			}
		}
		// fmt.Println("\n\nIteração", i)
		// printTable(table)
	}

	maxWeightCarried = table[size][capacity]

	for i := size; i > 0; i-- {
		if table[i][capacity] != table[i-1][capacity] {
			itemsCarried = append(itemsCarried, i-1)
			capacity -= weights[i-1]
		}
	}

	return maxWeightCarried, itemsCarried
}
