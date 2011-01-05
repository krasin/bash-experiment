package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// #define LALA 13 -> const lala = 13
var constDefine = regexp.MustCompile("^[\t ]*#[\t ]*define[\t ]+([A-Za-z_]+[A-Za-z0-9_]*)[\t ]+(.*)$")

func enhance(line string) {
	if (constDefine.MatchString(line)) {
		groups := constDefine.FindStringSubmatch(line)
		name := groups[1]
		value := strings.TrimSpace(groups[2])
		fmt.Printf("const %s = %s\n", name, value)
		return
	}
	fmt.Printf(line)
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	for {
		line, err := buf.ReadString('\n')
		if line != "" {
			enhance(line)
		}
		if err != nil {
			break
		}
	}
}
