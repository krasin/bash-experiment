package main

import (
	"bufio"
	"fmt"
	"os"
	"togo"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	for {
		line, err := buf.ReadString('\n')
		if line != "" {
			fmt.Printf("%s\n", togo.EnhanceLine(line))
		}
		if err != nil {
			break
		}
	}
}
