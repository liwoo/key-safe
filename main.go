package main

import (
	"fmt"
	"passwordGen/pkg/passgen"
)

func main() {
	config := passgen.DefaultConfig()
	password := passgen.GeneratePassword(config)
	fmt.Println("Your password is: ", password)
}
