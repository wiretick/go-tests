package sum

func Sum(n []int) int {
	sum := 0
	for _, nr := range n {
		sum += nr
	}
	return sum
}

// O(n^2)
func SumAll(numbers ...[]int) []int {
	var results []int
	for _, s := range numbers {
		results = append(results, Sum(s))
	}
	return results
}
