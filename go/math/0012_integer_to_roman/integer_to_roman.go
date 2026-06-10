package integertoroman

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var romanValue string

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			romanValue += symbols[i]
			num -= values[i]
		}
	}

	return romanValue
}
