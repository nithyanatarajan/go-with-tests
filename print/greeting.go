package print

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	// revive:disable-next-line
	fmt.Fprintf(writer, "Hello, %s", name)
}
