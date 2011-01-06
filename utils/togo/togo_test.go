// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package togo

import (
	"testing"
)

type enhanceLineTest struct {
	from string
	to   string
}

var enhanceLineTests = []*enhanceLineTest{
	&enhanceLineTest{"int a;", "a int"},
	&enhanceLineTest{"Lala b;", "b Lala"},
	&enhanceLineTest{"int b;\n", "b int"},
	&enhanceLineTest{"  int c;\r\n", "  c int"},
	&enhanceLineTest{" /* good comment */", " /* good comment */"},
	&enhanceLineTest{" /* some comment */\n", " /* some comment */"},
	&enhanceLineTest{"int a = 0;", "var a int = 0"},
	&enhanceLineTest{"int a = b;", "var a int = b"},
	&enhanceLineTest{"int a = 3.4;", "var a int = 3.4"},
	&enhanceLineTest{"int a = -5;", "var a int = -5"},
	&enhanceLineTest{"int a = 5 5;", "int a = 5 5;"},
	&enhanceLineTest{"#define REST -3  /* comment. */\n", "const REST = -3 /* comment. */"},
}

func TestEnhanceLine(t *testing.T) {
	for _, test := range enhanceLineTests {
		enhanced := EnhanceLine(test.from)
		if test.to != enhanced {
			t.Errorf("togo.EnhanceLine('%s') failed. Expected: '%s', has: '%s'", test.from, test.to, enhanced)
		}
	}
}
