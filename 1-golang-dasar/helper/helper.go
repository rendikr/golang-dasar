package helper

var version = "1.0.0"      // non accessible to other package
var Application = "golang" // accessible to other package

func sayGoodbye(name string) string {
	return "Goodbye, " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
