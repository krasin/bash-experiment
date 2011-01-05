package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// #define LALA 13 -> const lala = 13
var constDefine = regexp.MustCompile("^([\t ]*)#[\t ]*define[\t ]+([A-Za-z_]+[A-Za-z0-9_]*)[\t ]+(.*)$")

// int lala; -> lala int
var intVarDefine = regexp.MustCompile("^([\t ]*)(static[\t ]+|)int[\t ]+([A-Za-z_]+[A-Za-z0-9_]*)[\t ]*;[\n\t ]*$")

func enhance(line string) {
	if (constDefine.MatchString(line)) {
		groups := constDefine.FindStringSubmatch(line)
		indent := groups[1]
		name := groups[2]
		value := strings.TrimSpace(groups[3])
		fmt.Printf("%sconst %s = %s\n", indent, name, value)
		return
	}
	if (intVarDefine.MatchString(line)) {
		groups := intVarDefine.FindStringSubmatch(line)
		indent := groups[1]
		name := groups[3]
		fmt.Printf("%s%s int\n", indent, name)
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
