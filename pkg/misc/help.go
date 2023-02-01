package misc

import (
	"fmt"
	"os"
)

// NoImplemented is a silly function that shows a message and terminates the execution of the program.
func NoImplemented(message string) {
	fmt.Printf("%s", message)
	os.Exit(-5)
}
