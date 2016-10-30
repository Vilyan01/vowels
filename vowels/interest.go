package vowels

import (
	"fmt"
)

type Interest struct {
	Display string
	Weight  float64
}

func NewInterest(name string, weight float64) Interest {
	return Interest{Display: name, Weight: weight}
}

func (i Interest) PrintInterest() string {
	return fmt.Sprintf("<Interest> Display: %s, Weight: %.2f", i.Display, i.Weight)
}
