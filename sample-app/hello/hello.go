package main

import (
	"fmt"

	"sample-app/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
