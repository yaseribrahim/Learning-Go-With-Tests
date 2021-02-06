package sum

//Sum takes an array of numbers and returns their sum
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

//SumAll takes in arrays of numbers and returns an array of sums
func SumAll(numbersToSum ...[]int) []int {
		lengthOfNumbers := len(numbersToSum)
		sums := make([]int, lengthOfNumbers)

		for i, numbers := range numbersToSum {
			sums[i] = Sum(numbers)
		}

	return sums
}
