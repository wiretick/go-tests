package iteration

const repeatCount = 5

func Repeat(symbol string) string {
	var result string

	for i := 0; i < repeatCount; i++ {
		result += symbol
	}

	return result
}
