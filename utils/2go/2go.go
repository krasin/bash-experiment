package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// #define LALA 13 -> const lala = 13
var constDefine = regexp.MustCompile("^([\t ]*)#[\t ]*define[\t ]+([A-Za-z_][A-Za-z0-9_]*)[\t ]+(.*)$")

// int lala; -> lala int
var varDefine = regexp.MustCompile("^([\t ]*)(static[\t ]+|)([A-Za-z_][A-Za-z0-9_]*)[\t ]+([A-Za-z_][A-Za-z0-9_]*)[\t ]*;[\n\t ]*$")

// int *lala; -> lala *int
var pointerDefine = regexp.MustCompile("^([\t ]*)(static[\t ]+|)([A-Za-z_][A-Za-z0-9_]*)[\t ]*[*][\t ]*([A-Za-z_][A-Za-z0-9_]*)[\t ]*;[\n\t ]*$")


func enhance(line string) {
	if (constDefine.MatchString(line)) {
		groups := constDefine.FindStringSubmatch(line)
		indent := groups[1]
		name := groups[2]
		value := strings.TrimSpace(groups[3])
		fmt.Printf("%sconst %s = %s\n", indent, name, value)
		return
	}
	if (varDefine.MatchString(line)) {
		groups := varDefine.FindStringSubmatch(line)
		indent := groups[1]
		typ := groups[3]
		name := groups[4]
		fmt.Printf("%s%s %s\n", indent, name, typ)
		return
	}
	if (pointerDefine.MatchString(line)) {
		groups := pointerDefine.FindStringSubmatch(line)
		indent := groups[1]
		typ := groups[3]
		name := groups[4]
		fmt.Printf("%s%s *%s\n", indent, name, typ)
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
