package togo

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// #define LALA 13 -> const lala = 13
var constDefine = regexp.MustCompile("^#[\t ]*define[\t ]+([A-Za-z_][A-Za-z0-9_]*)[\t ]+([A-Za-z0-9_.\\-+]+)$")

// int lala; -> lala int
var varDefine = regexp.MustCompile("^([A-Za-z_][A-Za-z0-9_]*)[\t ]+([A-Za-z_][A-Za-z0-9_]*)[\t ]*;$")

// int *lala; -> lala *int
var pointerDefine = regexp.MustCompile("^([A-Za-z_][A-Za-z0-9_]*)[\t ]*[*][\t ]*([A-Za-z_][A-Za-z0-9_]*)[\t ]*;$")

// type_name lala = 70.0; -> var lala type_name = 70.0
var initializedDefine = regexp.MustCompile("^([A-Za-z_][A-Za-z0-9_]*)[\t ]+([A-Za-z_][A-Za-z0-9_]*)[\t ]*=[\t ]*([A-Za-z0-9_.\\-+]+)[\t ]*;$")

// typedef struct for_com { -> type for_com struct {
var structDefine = regexp.MustCompile("^typedef[\t ]+struct[\t ]+([A-Za-z_][A-Za-z0-9_]*)[\t ]*[{]$")

var trailingComment = regexp.MustCompile("^([^/]*)(/[*][^*]*[*]/)$")
var trailingSpace = regexp.MustCompile("^(.*)([\t\r\n ]*)$")
var indent = regexp.MustCompile("^([\t ]*)([^\t ].*)$")
var static = regexp.MustCompile("static[\t ](.*)$")

type filter func(line string) (cur string, finalizer filterFinalizer)
type filterFinalizer func(cur string) string

type filterDef struct {
	r *regexp.Regexp
	f filter
}

func Ident(cur string) string {
	return cur
}

var filters = []*filterDef{
	&filterDef{trailingSpace, filterTrailingSpace},
	&filterDef{indent, filterIndent},
	&filterDef{static, filterStatic},
	&filterDef{trailingComment, filterTrailingComment},
	&filterDef{trailingSpace, filterTrailingSpace},
	&filterDef{varDefine, filterVarDefine},
	&filterDef{pointerDefine, filterPointerDefine},
	&filterDef{constDefine, filterConstDefine},
	&filterDef{initializedDefine, filterInitializedDefine},
	&filterDef{structDefine, filterStructDefine},
}

func EnhanceLine(line string) (res string) {
	for _, def := range filters {
		//		fmt.Printf("line: '%s'\n", line)
		var fin filterFinalizer
		if def.r.MatchString(line) {
			line, fin = def.f(line)
			if fin == nil {
				return line
			}
			defer func() {
				res = fin(res)
				//				fmt.Printf("res: '%s'\n", res)
			}()
		}
	}
	return line
}

func filterTrailingSpace(line string) (string, filterFinalizer) {
	line = strings.TrimRightFunc(line, unicode.IsSpace)
	return line, Ident
}

func filterIndent(line string) (string, filterFinalizer) {
	groups := indent.FindStringSubmatch(line)
	indent := groups[1]
	line = groups[2]
	return line, func(cur string) string { return fmt.Sprintf("%s%s", indent, cur) }
}

func filterStatic(line string) (string, filterFinalizer) {
	groups := static.FindStringSubmatch(line)
	line = groups[1]
	return line, Ident
}

func filterTrailingComment(line string) (string, filterFinalizer) {
	comment := ""
	groups := trailingComment.FindStringSubmatch(line)
	line = groups[1]
	comment = strings.TrimSpace(groups[2])
	return line, func(cur string) string {
		if cur != "" {
			comment = " " + comment
		}
		return fmt.Sprintf("%s%s", cur, comment)
	}
}

func filterVarDefine(line string) (string, filterFinalizer) {
	groups := varDefine.FindStringSubmatch(line)
	typ := groups[1]
	name := groups[2]
	return fmt.Sprintf("%s %s", name, typ), nil
}

func filterConstDefine(line string) (string, filterFinalizer) {
	groups := constDefine.FindStringSubmatch(line)
	name := groups[1]
	value := strings.TrimSpace(groups[2])
	return fmt.Sprintf("const %s = %s", name, value), nil
}

func filterPointerDefine(line string) (string, filterFinalizer) {
	groups := pointerDefine.FindStringSubmatch(line)
	typ := groups[1]
	name := groups[2]
	return fmt.Sprintf("%s *%s", name, typ), nil
}

func filterInitializedDefine(line string) (string, filterFinalizer) {
	groups := initializedDefine.FindStringSubmatch(line)
	typ := groups[1]
	name := groups[2]
	value := groups[3]
	return fmt.Sprintf("var %s %s = %s", name, typ, value), nil
}

func filterStructDefine(line string) (string, filterFinalizer) {
	groups := structDefine.FindStringSubmatch(line)
	typ := groups[1]
	return fmt.Sprintf("type %s struct {", typ), nil
}
