package main

import (
	"chepra/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")

	repl.Start(os.Stdin, os.Stdout)
}
