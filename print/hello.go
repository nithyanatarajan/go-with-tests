package print

import "fmt"

func Hello(name string) string {
	if name == "" {
		return fmt.Sprintf("Hello, world!")
	}
	return fmt.Sprintf("Hello, %s!", name)
}
