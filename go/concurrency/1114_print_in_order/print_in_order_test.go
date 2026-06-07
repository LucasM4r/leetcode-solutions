package printinorder

import (
	"bytes"
	"sync"
	"testing"
)

func TestPrintInOrder(t *testing.T) {
	tests := []struct {
		name  string
		order []int
	}{
		{"Order 1-2-3", []int{1, 2, 3}},
		{"Order 1-3-2", []int{1, 3, 2}},
		{"Order 2-1-3", []int{2, 1, 3}},
		{"Order 2-3-1", []int{2, 3, 1}},
		{"Order 3-1-2", []int{3, 1, 2}},
		{"Order 3-2-1", []int{3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			foo := NewFoo()

			var buf bytes.Buffer
			var mu sync.Mutex

			printFirst := func() {
				mu.Lock()
				buf.WriteString("first")
				mu.Unlock()
			}
			printSecond := func() {
				mu.Lock()
				buf.WriteString("second")
				mu.Unlock()
			}
			printThird := func() {
				mu.Lock()
				buf.WriteString("third")
				mu.Unlock()
			}

			var wg sync.WaitGroup

			for _, val := range tt.order {
				wg.Add(1)
				switch val {
				case 1:
					go func() {
						defer wg.Done()
						foo.First(printFirst)
					}()
				case 2:
					go func() {
						defer wg.Done()
						foo.Second(printSecond)
					}()
				case 3:
					go func() {
						defer wg.Done()
						foo.Third(printThird)
					}()
				}
			}

			wg.Wait()

			expected := "firstsecondthird"
			if buf.String() != expected {
				t.Errorf("got %q, want %q", buf.String(), expected)
			}
		})
	}
}

func BenchmarkPrintInOrder(b *testing.B) {
	printFirst := func() {}
	printSecond := func() {}
	printThird := func() {}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		foo := NewFoo()
		var wg sync.WaitGroup

		wg.Add(3)

		go func() {
			defer wg.Done()
			foo.Third(printThird)
		}()

		go func() {
			defer wg.Done()
			foo.Second(printSecond)
		}()

		go func() {
			defer wg.Done()
			foo.First(printFirst)
		}()

		wg.Wait()
	}
}
