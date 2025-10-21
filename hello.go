// file: hello.go
package main

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func SayBye(name string) string {
	return fmt.Sprintf("Goodbye, %s!", name)
}
