package main

import (
	"fmt"
	"os"

	libs "libs"
)

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	err := libs.Tree(args[0], "")

	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
