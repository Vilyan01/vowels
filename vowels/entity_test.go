package vowels

import (
	"testing"
)

var e Entity

func init() {
	// actions
	a := NewAction("talk")
	a2 := NewAction("steal")
	a3 := NewAction("beg")
	aCombined := []Action{a, a2, a3}

	// interests
	i := NewInterest("gold", 0.65)
	i2 := NewInterest("socialization", 0.3)
	iCombined := []Interest{i, i2}

	e = NewEntity("orc", aCombined, iCombined)
}

func TestNewEntity(t *testing.T) {
	// entity was already created
	if e.Display != "orc" {
		t.Error("Expected orc got", e.Display)
	}
	if len(e.Actions) != 3 {
		t.Error("Expected actions length of 3, got", len(e.Actions))
	}
	if len(e.Interests) != 2 {
		t.Error("Expected interests length of 2 got", len(e.Interests))
	}
}
