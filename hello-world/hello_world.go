package greeting

const testVersion = 3

func HelloWorld(name string) string {
	if len(name) == 0 {
		name = "World"
	}
	return "Hello, " + name + "!"
}