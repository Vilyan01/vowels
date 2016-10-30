package version

import (
	"fmt"
)

const VERSION = "0.1"

func PrintVersion() {
	fmt.Println(fmt.Sprintf("Vowels v%s", VERSION))
}
