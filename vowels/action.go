package vowels

import (
	"fmt"
)

type Action struct {
	Display string
}

func NewAction(display string) Action {
	return Action{Display: display}
}

func (a Action) PrintAction() string {
	return fmt.Sprintf("<Action> Display: %s", a.Display)
}
