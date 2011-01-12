package main

import (
	"fmt"
	"gobash"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: gobash <path_to_script>")
		os.Exit(1)
	}
	os.Exit(gobash.ExecuteScript(os.Args[1]))
}
