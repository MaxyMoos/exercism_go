package hello

const TestVersion = 1

func HelloWorld(input string) string {
    if len(input) == 0 { input = "World" }
    return "Hello, " + input + "!"
}
