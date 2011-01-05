package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	for {
		line, err := buf.ReadString('\n')
		if line != "" {
			fmt.Printf("// %s", line)
		}
		if err != nil {
			break
		}
	}
}
