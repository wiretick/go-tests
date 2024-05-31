package sum

func Sum(n []int) int {
	sum := 0
	for _, nr := range n {
		sum += nr
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var results []int
	for _, s := range numbers {
		results = append(results, Sum(s))
	}
	return results
}

func SumAllTails(numbers ...[]int) []int {
	results := []int{}

	for _, slic := range numbers {
		if len(slic) <= 0 {
			results = append(results, 0)
		} else {
			results = append(results, Sum(slic[1:]))
		}
	}

	return results
}
