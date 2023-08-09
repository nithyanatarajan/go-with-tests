package print

import "fmt"

// Hello prints given name in given language.
// Defaults to 'Hello, World!' if no name or language is passed
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)

	return fmt.Sprintf("%s, %s!", prefix, name)
}

func greetingPrefix(language string) string {
	languages := map[string]string{
		"Spanish": "Hola",
		"French":  "Bonjour",
	}
	prefix, found := languages[language]
	if !found {
		prefix = "Hello"
	}
	return prefix
}
