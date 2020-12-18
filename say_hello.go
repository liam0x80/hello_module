package hello_module

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("hello, %s!", name)
}