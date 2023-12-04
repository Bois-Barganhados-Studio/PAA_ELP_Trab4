package greedy

func Knapsack(weights []int, values []int, capacity int) (int, []int) {
	maxValueCarried := 0
	itemsCarried := []int{}

	for i := 0; i < len(weights); i++ {
		if weights[i] <= capacity {
			maxValueCarried += values[i]
			itemsCarried = append(itemsCarried, i)
			capacity -= weights[i]
		}
	}

	return maxValueCarried, itemsCarried
}
