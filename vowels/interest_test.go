package vowels

import (
	"fmt"
	"testing"
)

func TestNewInterest(t *testing.T) {
	interest := NewInterest("test", 0.5)
	if interest.Display != "test" {
		t.Error("Expected display of test, got", interest.Display)
	}
	if interest.Weight != 0.5 {
		t.Error("Expected weight of 0.5 got", interest.Weight)
	}
}

func ExamplePrintInterest() {
	i := NewInterest("test", 0.5)
	fmt.Println(i.PrintInterest())
	// Output: <Interest> Display: test, Weight: 0.50
}
