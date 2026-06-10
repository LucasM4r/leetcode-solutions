package integertoroman

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected string
	}{
		{"Basic single digit", 3, "III"},
		{"Subtractive case 4", 4, "IV"},
		{"Subtractive case 9", 9, "IX"},
		{"Standard combination 58", 58, "LVIII"},
		{"Complex case 1994", 1994, "MCMXCIV"},
		{"Boundary minimum", 1, "I"},
		{"Boundary maximum", 3999, "MMMCMXCIX"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.value); got != tt.expected {
				t.Errorf("intToRoman(%d) = %s; want %s", tt.value, got, tt.expected)
			}
		})
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	benchmarks := []struct {
		name  string
		value int
	}{
		{"Small", 3},
		{"Medium", 58},
		{"Large", 1994},
		{"Maximum", 3999},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = intToRoman(bm.value)
			}
		})
	}
}
