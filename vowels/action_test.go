package vowels

import (
	"fmt"
	"testing"
)

func TestNewAction(t *testing.T) {
	a := NewAction("sleep")
	if a.Display != "sleep" {
		t.Error("expected sleep got", a.Display)
	}
}

func ExamplePrintAction() {
	a := NewAction("sleep")
	fmt.Println(a.PrintAction())
	// Output: <Action> Display: sleep
}
