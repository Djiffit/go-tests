package main

// Sum array of numbers and return it
func Sum(values []int) (sum int) {
	for _, value := range values {
		sum += value
	}
	return
}

func main() {
	Sum([]int{11, 2, 3, 4, 5})
}

// SumAll returns the sum of all arrays passed as a parameter
func SumAll(arrsToSum ...[]int) (sums []int) {
	for _, arr := range arrsToSum {
		sums = append(sums, Sum(arr))
	}
	return
}

// SumAllTails returns the sum of all tails
func SumAllTails(tailsToSum ...[]int) (sums []int) {
	for _, arr := range tailsToSum {
		if len(arr) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(arr[1:]))
		}
	}
	return
}
