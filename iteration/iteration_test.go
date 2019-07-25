package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("Repeats 3 times", func(t *testing.T) {
		repeated := Repeat("x", 3)
		expected := "xxx"

		if repeated != expected {
			t.Errorf("Expected %q got %q", expected, repeated)
		}
	})

	t.Run("Repeats 5 times", func(t *testing.T) {
		repeated := Repeat("s", 3)
		expected := "sss"

		if repeated != expected {
			t.Errorf("Expected %q got %q", expected, repeated)
		}
	})

	t.Run("Repeats 0 times", func(t *testing.T) {
		repeated := Repeat("s", 0)
		expected := ""

		if repeated != expected {
			t.Errorf("Expected %q got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("x", 5)
	}
}

func ExampleRepeat() {
	output := Repeat("x", 5)
	fmt.Println(output)
	// Output: xxxxx
}
