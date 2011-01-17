echo ff{c,b,a}
echo f{d,e,f}g
echo {l,n,m}xyz
echo {abc\,def}
echo {abc}

echo \{a,b,c,d,e}
echo {x,y,\{a,b,c}}
echo {x\,y,\{abc\},trie}

echo /usr/{ucb/{ex,edit},lib/{ex,how_ex}}

echo XXXX\{`echo a b c | tr ' ' ','`\}
eval echo XXXX\{`echo a b c | tr ' ' ','`\}

echo {}
echo { }
echo }
echo {
echo abcd{efgh

echo foo {1,2} bar
echo `echo foo {1,2} bar`
echo $(echo foo {1,2} bar)

var=baz
varx=vx
vary=vy

echo foo{bar,${var}.}
echo foo{bar,${var}}

echo "${var}"{x,y}
echo $var{x,y}
echo ${var}{x,y}

unset var varx vary

# new sequence brace operators
echo {1..10}

# this doesn't work yet
echo {0..10,braces}
# but this does
echo {{0..10},braces}
echo x{{0..10},braces}y

echo {3..3}
echo x{3..3}y
echo {10..1}
echo {10..1}y
echo x{10..1}y

echo {a..f}
echo {f..a}

echo {a..A}
echo {A..a}

echo {f..f}

# mixes are incorrectly-formed brace expansions
echo {1..f}
echo {f..1}

echo 0{1..9} {10..20}

# do negative numbers work?
echo {-1..-10}
echo {-20..0}

# weirdly-formed brace expansions -- fixed in post-bash-3.1
echo a-{b{d,e}}-c

echo a-{bdef-{g,i}-c

echo {"klklkl"}{1,2,3}
echo {"x,x"}

echo {1..10..2}
echo {-1..-10..2}
echo {-1..-10..-2}

echo {10..1..-2}
echo {10..1..2}

echo {1..20..2}
echo {1..20..20}

echo {100..0..5}
echo {100..0..-5}

echo {a..z}
echo {a..z..2}
echo {z..a..-2}

# unwanted zero-padding -- fixed post-bash-4.0
echo {10..0..2}
echo {10..0..-2}
echo {-50..-0..5}

# bad
echo {1..10.f}
echo {1..ff}
echo {1..10..ff}
echo {1.20..2}
echo {1..20..f2}
echo {1..20..2f}
echo {1..2f..2}
echo {1..ff..2}
echo {1..ff}
echo {1..f}
echo {1..0f}
echo {1..10f}
echo {1..10.f}
echo {1..10.f}
