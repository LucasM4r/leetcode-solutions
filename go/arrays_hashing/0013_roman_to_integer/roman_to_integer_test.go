package romantointeger

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name        string
		romanNumber string
		want        int
	}{
		{"Example 1", "XIX", 19},
		{"Example 2", "XVII", 17},
		{"Example 3", "IV", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.romanNumber); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

var result int

func BenchmarkRomanToInt(b *testing.B) {
	romanStr := "MCMXCIV"
	var r int

	for i := 0; i < b.N; i++ {
		r = romanToInt(romanStr)
	}

	result = r
}
