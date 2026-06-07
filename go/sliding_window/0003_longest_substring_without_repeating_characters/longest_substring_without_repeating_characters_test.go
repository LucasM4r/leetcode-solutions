package longestsubstringwithoutrepeatingcharacters

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"Exemplo 1 (abcabcbb)", "abcabcbb", 3},
		{"Exemplo 2 (bbbbb)", "bbbbb", 1},
		{"Exemplo 3 (pwwkew)", "pwwkew", 3},

		{"String Vazia", "", 0},
		{"Apenas um espaço em branco", " ", 1},
		{"Dois espaços em branco", "  ", 1},
		{"Letras Maiúsculas e Minúsculas", "aA", 2},
		{"Todos caracteres únicos", "abcdef", 6},
		{"Caracteres Especiais", "!@#$%^&*()", 10},
		{"Padrão repetitivo rápido", "abba", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LengthOfLongestSubstringHashMap(tt.s); got != tt.want {
				t.Errorf("LengthOfLongestSubstringHashMap(%q) = %v, want %v", tt.s, got, tt.want)
			}

			if got := lengthOfLongestSubstringSlidingWindow(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringSlidingWindow(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

var benchStr = "abcabcbbxyz1234567890qwertyuiop!@#abcabcabcbbxyz1234567890qwertyuiop"

func BenchmarkLengthOfLongestSubstringHashMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LengthOfLongestSubstringHashMap(benchStr)
	}
}

func BenchmarkLengthOfLongestSubstringSlidingWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstringSlidingWindow(benchStr)
	}
}
