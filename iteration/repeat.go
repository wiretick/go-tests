package iteration

func Repeat(symbol string, nr int) string {
	var result string

	for i := 0; i < nr; i++ {
		result += symbol
	}

	return result
}
