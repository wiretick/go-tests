package arrays

func Sum(n [4]int) int {
	sum := 0
	for _, nr := range n {
		sum += nr
	}
	return sum
}
