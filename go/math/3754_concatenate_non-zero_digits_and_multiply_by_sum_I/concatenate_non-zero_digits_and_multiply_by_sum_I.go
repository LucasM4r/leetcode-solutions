package concatenatenonzerodigitsandmultiplybysumi

import "strconv"

func sumAndMultiply(n int) int64 {
	nStr := strconv.Itoa(n)
	var sum int64 = 0

	var xStr string

	for _, char := range nStr {
		if char == '0' {
			continue
		}
		xStr += string(char)
		digit := int64(char - '0')
		sum += digit
	}
	if xStr == "" {
		return 0
	}

	x, _ := strconv.ParseInt(xStr, 10, 64)
	return x * sum
}
