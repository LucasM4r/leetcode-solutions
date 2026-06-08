package printfoobaralternately

import (
	"bytes"
	"strings"
	"sync"
	"testing"
)

func TestPrintFoobarAlternately(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{"print 1", 1},
		{"print 2", 2},
		{"print 0", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foobar := NewFooBar(5)
			foobar.n = tt.n
			var buf bytes.Buffer
			var mu sync.Mutex

			printFoo := func() {
				mu.Lock()
				buf.WriteString("Foo")
				mu.Unlock()
			}

			printBar := func() {
				mu.Lock()
				buf.WriteString("Bar")
				mu.Unlock()
			}

			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				defer wg.Done()
				foobar.Foo(printFoo)
			}()

			go func() {
				defer wg.Done()
				foobar.Bar(printBar)
			}()

			wg.Wait()

			expected := strings.Repeat("FooBar", tt.n)
			if got := buf.String(); got != expected {
				t.Errorf("[ERROR] Expected %q, but got %q", expected, got)
			}

		})
	}
}

func BenchmarkPrintFooBar(b *testing.B) {
	n := 100

	printFoo := func() {}
	printBar := func() {}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		foobar := NewFooBar(n)

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			foobar.Foo(printFoo)
		}()

		go func() {
			defer wg.Done()
			foobar.Bar(printBar)
		}()

		wg.Wait()
	}
}
