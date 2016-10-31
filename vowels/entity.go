package vowels

import (
	"fmt"
)

type Entity struct {
	Display   string
	Actions   []Action
	Interests []Interest
}

func NewEntity(display string, actions []Action, interests []Interest) Entity {
	return Entity{Display: display, Actions: actions, Interests: interests}
}

func (e Entity) PrintEntity() {
	fmt.Println("placeholder")
}
