package dog

import (
	"fmt"
	"testing"
)

func TestYears(T *testing.T) {
	type test struct {
		data   int
		answer int
	}

	tests := []test{
		test{7, 49},
		test{10, 70},
		test{4, 28},
	}

	for _, t := range tests {
		y := Years(t.data)
		if y != t.answer {
			T.Error("got", y, "want", t.answer)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(7)

	if y != 49 {
		t.Error("Expected 49 got", y)
	}
}

func ExampleYears() {
	fmt.Println(Years(7))
	// Output:
	// 49
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(7))
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(7)
	}
}
