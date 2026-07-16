package sequentialdigits

func sequentialDigits(low int, high int) []int {
	var result []int

	for length := 2; length <= 9; length++ {
		start := 0
		adder := 0

		for i := 1; i <= length; i++ {
			start = start*10 + i
			adder = adder*10 + 1
		}

		for start <= high {
			if start >= low {
				result = append(result, start)
			}

			if start%10 == 9 {
				break
			}

			start += adder
		}
	}

	return result
}
