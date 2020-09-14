package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(7)

	if y != 49 {
		t.Error("Expected 49 got", y)
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
