package togo

import (
	"fmt"
	"regexp"
	"strings"
)

// #define LALA 13 -> const lala = 13
var constDefine = regexp.MustCompile("^([\t ]*)#[\t ]*define[\t ]+([A-Za-z_][A-Za-z0-9_]*)[\t ]+(.*)$")

// int lala; -> lala int
var varDefine = regexp.MustCompile("^([\t ]*)(static[\t ]+|)([A-Za-z_][A-Za-z0-9_]*)[\t ]+([A-Za-z_][A-Za-z0-9_]*)[\t ]*;[\n\t ]*(/[*][^*]*[*]/[\n\t ]*|)$")

// int *lala; -> lala *int
var pointerDefine = regexp.MustCompile("^([\t ]*)(static[\t ]+|)([A-Za-z_][A-Za-z0-9_]*)[\t ]*[*][\t ]*([A-Za-z_][A-Za-z0-9_]*)[\t ]*;[\n\t ]*(/[*][^*]*[*]/[\n\t ]*|)$")


func EnhanceLine(line string) string {
	if (constDefine.MatchString(line)) {
		groups := constDefine.FindStringSubmatch(line)
		indent := groups[1]
		name := groups[2]
		value := strings.TrimSpace(groups[3])
		return fmt.Sprintf("%sconst %s = %s", indent, name, value)
	}
	if (varDefine.MatchString(line)) {
		groups := varDefine.FindStringSubmatch(line)
		indent := groups[1]
		typ := groups[3]
		name := groups[4]
		comment := strings.TrimSpace(groups[5])
		if comment != "" {
			comment = " " + comment
		}
		return fmt.Sprintf("%s%s %s%s", indent, name, typ, comment)
	}
	if (pointerDefine.MatchString(line)) {
		groups := pointerDefine.FindStringSubmatch(line)
		indent := groups[1]
		typ := groups[3]
		name := groups[4]
		comment := strings.TrimSpace(groups[5])
		if comment != "" {
			comment = " " + comment
		}
		return fmt.Sprintf("%s%s *%s%s", indent, name, typ, comment)
	}
	return line
}

