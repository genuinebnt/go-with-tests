package hello_world

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World!"
	}

	prefix := englishHelloPrefix

	switch language {
	case "Spanish":
		prefix = "Hola, "
	case "French":
		prefix = "Bonjour, "
	}

	return prefix + name
}
