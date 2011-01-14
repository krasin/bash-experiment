package gobash

/* A Bison parser, made by GNU Bison 2.3.  */

/* Skeleton implementation for Bison's Yacc-like parsers in C

   Copyright (C) 1984, 1989, 1990, 2000, 2001, 2002, 2003, 2004, 2005, 2006
   Free Software Foundation, Inc.

   This program is free software; you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation; either version 2, or (at your option)
   any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program; if not, write to the Free Software
   Foundation, Inc., 51 Franklin Street, Fifth Floor,
   Boston, MA 02110-1301, USA.  */

/* As a special exception, you may create a larger work that contains
   part or all of the Bison parser skeleton and distribute that work
   under terms of your choice, so long as that work isn't itself a
   parser generator using the skeleton or a modified version thereof
   as a parser skeleton.  Alternatively, if you modify or redistribute
   the parser skeleton itself, you may (at your option) remove this
   special exception, which will cause the skeleton and the resulting
   Bison output files to be licensed under the GNU General Public
   License without this special exception.

   This special exception was added by the Free Software Foundation in
   version 2.2 of Bison.  */

/* C LALR(1) parser skeleton written by Richard Stallman, by
   simplifying the original so-called "semantic" parser.  */

/* All symbols defined below should begin with yy or YY, to avoid
   infringing on user name space.  This should be done even for local
   variables, as they might otherwise be expanded by user macros.
   There are some unavoidable exceptions within include files to
   define necessary library symbols; they are noted "INFRINGES ON
   USER NAME SPACE" below.  */

import (
	"fmt"
	"math"
	"os"
    "strconv"
    "unicode"
)

const EOF = -1

/* Identify Bison output.  */
const YYBISON = 1

/* Bison version.  */
const YYBISON_VERSION = "2.3"

/* Skeleton name.  */
const YYSKELETON_NAME = "yacc.c"

/* Pure parsers.  */
const YYPURE = 0

/* Using locations.  */
const YYLSP_NEEDED = 0



/* Tokens.  */
const NO_TOKEN = 0
const IF = 258
const THEN = 259
const ELSE = 260
const ELIF = 261
const FI = 262
const CASE = 263
const ESAC = 264
const FOR = 265
const SELECT = 266
const WHILE = 267
const UNTIL = 268
const DO = 269
const DONE = 270
const FUNCTION = 271
const COPROC = 272
const COND_START = 273
const COND_END = 274
const COND_ERROR = 275
const IN = 276
const BANG = 277
const TIME = 278
const TIMEOPT = 279
const WORD = 280
const ASSIGNMENT_WORD = 281
const REDIR_WORD = 282
const NUMBER = 283
const ARITH_CMD = 284
const ARITH_FOR_EXPRS = 285
const COND_CMD = 286
const AND_AND = 287
const OR_OR = 288
const GREATER_GREATER = 289
const LESS_LESS = 290
const LESS_AND = 291
const LESS_LESS_LESS = 292
const GREATER_AND = 293
const SEMI_SEMI = 294
const SEMI_AND = 295
const SEMI_SEMI_AND = 296
const LESS_LESS_MINUS = 297
const AND_GREATER = 298
const AND_GREATER_GREATER = 299
const LESS_GREATER = 300
const GREATER_BAR = 301
const BAR_AND = 302
const yacc_EOF = 303


const RE_READ_TOKEN = -99
const NO_EXPANSION = -100

/* The line number in a script where the word in a `case WORD', `select WORD'
   or `for WORD' begins.  This is a nested command maximum, since the array
   index is decremented after a case, select, or for command is parsed. */
const MAX_CASE_NEST = 128

func enlargeBuffer(buf []int, newsize int) (res []int) {
	if newsize <= len(buf) {
		return buf
	}
	res = make([]int, newsize)
	for i, v := range buf {
		res[i] = v
	}
	return
}

func resizeBuffer(buf []int, cind int, room int, csize int, sincr int) (res []int) {
    if cind + room >= csize {
	for cind + room >= csize {
	  csize += sincr
	}
	return enlargeBuffer(buf, csize)
    }
    return buf
}


// 
// 
// extern int current_command_number;
// extern int sourcelevel, parse_and_execute_level;
// extern int posixly_correct;
// extern pid_t last_command_subst_pid;
// extern char *shell_name, *current_host_name;
// extern char *dist_version;
// extern int patch_level;
// extern int dump_translatable_strings, dump_po_strings;
// extern sh_builtin_func_t *last_shell_builtin, *this_shell_builtin;
// extern int bash_input_fd_changed;
// 

type StringSaver struct {
  next *StringSaver
  expand_alias bool /* Value to set expand_alias to when string is popped. */
  saved_line []int
  expander *alias_t /* alias that caused this line to be pushed. */
  saved_line_size int
  saved_line_index int
  saved_line_terminator int
}

type ParserState struct {

// TODO(krasin): consider moving this to options
posixly_correct bool

bashInput BashInput

pushed_string_list *StringSaver

global_command *Command

last_command_exit_value int

/* Non-zero means we expand aliases in commands. */
expand_aliases bool

extended_glob bool

/* If non-zero, $'...' and $"..." are expanded when they appear within
   a ${...} expansion, even when the expansion appears within double
   quotes. */
extended_quote int

/* When non-zero, an open-brace used to create a group is awaiting a close
   brace partner. */
open_brace_count int

/* The number of lines read from input while creating the current command. */
current_command_line_count int

/* The token that currently denotes the end of parse. */
shell_eof_token int

/* The token currently being read. */
current_token int

/* The current parser state. */
parser_state int

/* Variables to manage the task of reading here documents, because we need to
   defer the reading until after a complete command has been collected. */
redir_stack [10]*Redirect
need_here_doc int

/* Where shell input comes from.  History expansion is performed on each
   line when the shell is interactive. */
shell_input_line []int
shell_input_line_index int
shell_input_line_size int	/* Amount allocated for shell_input_line. */
shell_input_line_len int	/* strlen (shell_input_line) */

/* Either zero or EOF. */
shell_input_line_terminator int

/* The line number in a script on which a function definition starts. */
function_dstart int

/* The line number in a script on which a function body starts. */
function_bstart int

/* The line number in a script at which an arithmetic for command starts. */
arith_for_lineno int

/* The last read token, or NULL.  read_token () uses this for context
   checking. */
last_read_token int

/* The token read prior to gps.last_read_token. */
token_before_that int

/* The token read prior to gps.token_before_that. */
two_tokens_ago int

global_extglob bool

word_lineno [MAX_CASE_NEST]int
word_top int

/* If non-zero, it is the token that we want read_token to return
   regardless of what text is (or isn't) present to be read.  This
   is reset by read_token.  If token_to_read == WORD or
   ASSIGNMENT_WORD, yylval.word should be set to word_desc_to_read. */
token_to_read int
word_desc_to_read *word_desc

source Redirectee
redir Redirectee

/* The look-ahead symbol.  */
yychar int

/* The semantic value of the look-ahead symbol.  */
yylval YYSTYPE

/* Number of syntax errors so far.  */
yynerrs int

/* Global var is non-zero when end of file has been reached. */
EOF_Reached bool

/* Variable containing the current get and unget functions.
   See ./input.h for a clearer description. */
bash_input BashInput

/* The globally known line number. */
line_number int

cond_lineno int
cond_token int

dstack *intStack
builtins *BuiltinsManager

/* This implements one-character lookahead/lookbehind across physical input
   lines, to avoid something being lost because it's pushed back with
   shell_ungetc when we're at the start of a line. */
eol_ungetc_lookahead int

// TODO(krasin): this should go to the options.
echo_input_at_read bool

/* When non-zero, we have read the required tokens
   which allow ESAC to be the next one read. */
esacs_needed_count int


} // ParserState

func newParserState(bashInput BashInput) *ParserState {
	state := new(ParserState)
	state.bashInput = bashInput
	state.extended_quote = 1
	state.word_top = -1
	state.dstack = newIntStack()
    state.builtins = NewBuiltinsManager()
	return state
}

type YYSTYPE struct {
// #line 320 "/Users/chet/src/bash/src/parse.y"
  word *word_desc /* the word that we read. */
  number int /* the number that we read. */
  word_list *word_list
  command *Command
  redirect *Redirect
  element ELEMENT
  pattern *PatternList
}


/* Copy the second part of user declarations.  */


/* Line 216 of yacc.c.  */
// #line 514 "y.tab.c"

const YYSIZE_MAXIMUM = math.MaxUint64

/* YYFINAL -- State number of the termination state.  */
const YYFINAL = 113
/* YYLAST -- Last index in YYTABLE.  */
const YYLAST = 760

/* YYNTOKENS -- Number of terminals.  */
const YYNTOKENS = 60
/* YYNNTS -- Number of nonterminals.  */
const YYNNTS = 38
/* YYNRULES -- Number of rules.  */
const YYNRULES = 167
/* YYNRULES -- Number of states.  */
const YYNSTATES = 344

/* YYTRANSLATE(YYLEX) -- Bison symbol number corresponding to YYLEX.  */
const YYUNDEFTOK = 2
const YYMAXUTOK = 303

func YYTRANSLATE(yyx int) int {
	if (yyx <= YYMAXUTOK) {
		return yytranslate[yyx]
	}
	return YYUNDEFTOK
}

/* YYTRANSLATE[YYLEX] -- Bison symbol number corresponding to YYLEX.  */
var yytranslate = []int {
       0,     2,     2,     2,     2,     2,     2,     2,     2,     2,
      50,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,    48,     2,
      58,    59,     2,     2,     2,    55,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,    49,
      54,     2,    53,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,    56,    52,    57,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     2,     2,     2,     2,
       2,     2,     2,     2,     2,     2,     1,     2,     3,     4,
       5,     6,     7,     8,     9,    10,    11,    12,    13,    14,
      15,    16,    17,    18,    19,    20,    21,    22,    23,    24,
      25,    26,    27,    28,    29,    30,    31,    32,    33,    34,
      35,    36,    37,    38,    39,    40,    41,    42,    43,    44,
      45,    46,    47,    51,
}

/* YYR1[YYN] -- Symbol number of symbol that rule YYN derives.  */
var yyr1 = []int {
       0,    60,    61,    61,    61,    61,    62,    62,    63,    63,
      63,    63,    63,    63,    63,    63,    63,    63,    63,    63,
      63,    63,    63,    63,    63,    63,    63,    63,    63,    63,
      63,    63,    63,    63,    63,    63,    63,    63,    63,    63,
      63,    63,    63,    63,    63,    63,    63,    63,    63,    63,
      63,    63,    64,    64,    64,    65,    65,    66,    66,    67,
      67,    67,    67,    67,    68,    68,    68,    68,    68,    68,
      68,    68,    68,    68,    68,    69,    69,    69,    69,    69,
      69,    69,    69,    70,    70,    70,    70,    71,    71,    71,
      71,    71,    71,    72,    72,    72,    73,    73,    73,    74,
      74,    75,    76,    76,    76,    76,    76,    77,    77,    77,
      78,    79,    80,    81,    81,    81,    82,    82,    83,    83,
      83,    83,    84,    84,    84,    84,    84,    84,    85,    85,
      86,    87,    87,    88,    88,    88,    89,    89,    89,    89,
      89,    89,    90,    90,    91,    91,    91,    92,    92,    93,
      93,    93,    94,    94,    94,    94,    94,    95,    95,    95,
      95,    95,    95,    96,    96,    96,    97,    97,
}

/* YYR2[YYN] -- Number of symbols composing right hand side of rule YYN.  */
var yyr2 = []int {
       0,     2,     2,     1,     2,     1,     1,     2,     2,     2,
       3,     3,     3,     3,     2,     3,     3,     2,     3,     3,
       2,     3,     3,     2,     3,     3,     2,     3,     3,     2,
       3,     3,     2,     3,     3,     2,     3,     3,     2,     3,
       3,     2,     3,     3,     2,     3,     3,     2,     3,     3,
       2,     2,     1,     1,     1,     1,     2,     1,     2,     1,
       1,     2,     1,     1,     1,     1,     5,     5,     1,     1,
       1,     1,     1,     1,     1,     6,     6,     7,     7,    10,
      10,     9,     9,     7,     7,     5,     5,     6,     6,     7,
       7,    10,    10,     6,     7,     6,     5,     6,     4,     1,
       2,     3,     2,     3,     3,     4,     2,     5,     7,     6,
       3,     1,     3,     4,     6,     5,     1,     2,     4,     4,
       5,     5,     2,     3,     2,     3,     2,     3,     1,     3,
       2,     1,     2,     3,     3,     3,     4,     4,     4,     4,
       4,     1,     1,     1,     1,     1,     1,     0,     2,     1,
       2,     2,     4,     4,     3,     3,     1,     1,     2,     2,
       3,     3,     2,     4,     4,     1,     1,     2,
}

/* YYDEFACT[STATE-NAME] -- Default rule to reduce with in state
   STATE-NUM when YYTABLE doesn't specify something else to do.  Zero
   means the default is an error.  */
var yydefact = []int {
       0,     0,   147,     0,     0,     0,   147,   147,     0,     0,
       0,     0,   166,    52,    53,     0,     0,   111,     0,     0,
       0,     0,     0,     0,     0,     0,     0,     0,     3,     5,
       0,     0,   147,   147,     0,    54,    57,    59,   165,    60,
      64,    74,    68,    65,    62,    70,    63,    69,    71,    72,
      73,     0,   149,   156,   157,     0,     4,   131,     0,     0,
     147,   147,     0,   147,     0,     0,   147,    52,   106,   102,
       0,   158,     0,   167,     0,     0,     0,     0,     0,     0,
       0,     0,     0,     0,     0,     0,     0,     0,     0,     0,
       0,     0,     0,     0,     0,    14,    23,    38,    32,    47,
      29,    41,    35,    44,    26,    50,    51,    20,    17,     8,
       9,     0,     0,     1,    52,    58,    55,    61,   142,   143,
       2,   147,   147,   150,   151,   147,   147,     0,   145,   144,
     146,   162,   159,   147,   148,   130,   132,   141,     0,   147,
       0,   147,   147,   147,   147,     0,   147,   147,     0,     0,
     104,   103,   112,   161,   147,    16,    25,    40,    34,    49,
      31,    43,    37,    46,    28,    22,    19,    12,    13,    15,
      24,    39,    33,    48,    30,    42,    36,    45,    27,    21,
      18,    10,    11,   110,   101,    56,     0,     0,   154,   155,
       0,     0,   160,     0,   147,   147,   147,   147,   147,   147,
       0,   147,     0,   147,     0,     0,     0,     0,   147,     0,
     147,     0,     0,   147,    99,    98,   105,     0,   152,   153,
       0,     0,   164,   163,   147,   147,   107,     0,     0,     0,
     134,   135,   133,     0,   116,   147,     0,   147,   147,     0,
       6,     0,   147,     0,    85,    86,   147,   147,   147,   147,
       0,     0,     0,     0,    66,    67,     0,   100,    96,     0,
       0,   109,   136,   137,   138,   139,   140,    95,   122,   124,
     126,   117,     0,    93,   128,     0,     0,     0,     0,    75,
       7,   147,     0,    76,     0,     0,     0,     0,    87,     0,
     147,    88,    97,   108,   147,   147,   147,   147,   123,   125,
     127,    94,     0,     0,   147,    77,    78,     0,   147,   147,
      83,    84,    89,    90,     0,   113,     0,     0,     0,   147,
     129,   118,   119,   147,   147,     0,     0,   147,   147,   147,
     115,   120,   121,     0,     0,    81,    82,     0,     0,   114,
      79,    80,    91,    92,
}

/* YYDEFGOTO[NTERM-NUM].  */
var yydefgoto = []int {
      -1,    34,   241,    35,    36,   117,    37,    38,    39,    40,
      41,    42,    43,    44,   215,    45,    46,    47,    48,    49,
      50,   227,   233,   234,   235,   276,    57,    58,   135,   136,
     120,   131,    59,    51,   188,   137,    54,    55,
}

/* YYPACT[STATE-NUM] -- Index in YYTABLE of the portion describing
   STATE-NUM.  */
const YYPACT_NINF = -212
var yypact = []int {
     318,   -40,  -212,    11,    10,    32,  -212,  -212,    34,   661,
      39,   514,    52,    21,  -212,   255,   706,  -212,    65,    79,
      47,    98,    55,   133,   134,   139,   151,   152,  -212,  -212,
     156,   157,  -212,  -212,   142,  -212,  -212,   232,  -212,   693,
    -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,
    -212,    49,   213,  -212,   -28,   367,  -212,  -212,   150,   416,
    -212,   135,    -3,   137,   184,   194,   158,    31,   232,   693,
     191,   -28,   612,  -212,   166,   199,   202,    69,   209,   128,
     210,   214,   218,   219,   225,   227,   247,   165,   248,   181,
     256,   257,   258,   259,   262,  -212,  -212,  -212,  -212,  -212,
    -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,
    -212,   237,   236,  -212,  -212,  -212,  -212,   693,  -212,  -212,
    -212,  -212,  -212,   465,   465,  -212,  -212,   612,  -212,  -212,
    -212,  -212,   -28,  -212,  -212,  -212,   205,  -212,   -13,  -212,
     116,  -212,  -212,  -212,  -212,   117,  -212,  -212,   240,    53,
     693,   693,  -212,   -28,  -212,  -212,  -212,  -212,  -212,  -212,
    -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,
    -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,
    -212,  -212,  -212,  -212,  -212,  -212,   416,   416,    75,    75,
     563,   563,   -28,    15,  -212,  -212,  -212,  -212,  -212,  -212,
      72,  -212,   120,  -212,   281,   249,   105,   118,  -212,   279,
    -212,   290,   292,  -212,   693,  -212,   693,    53,  -212,  -212,
     465,   465,   -28,   -28,  -212,  -212,  -212,   303,   416,   416,
     416,   416,   416,   302,   174,  -212,     0,  -212,  -212,   297,
    -212,   179,  -212,   263,  -212,  -212,  -212,  -212,  -212,  -212,
     299,   416,   179,   265,  -212,  -212,    53,   693,  -212,   308,
     312,  -212,  -212,  -212,    80,    80,    80,  -212,  -212,  -212,
    -212,   224,    43,  -212,  -212,   300,    33,   309,   270,  -212,
    -212,  -212,   129,  -212,   317,   276,   322,   282,  -212,   205,
    -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,  -212,
    -212,  -212,    46,   313,  -212,  -212,  -212,   149,  -212,  -212,
    -212,  -212,  -212,  -212,   161,   115,   416,   416,   416,  -212,
    -212,  -212,   416,  -212,  -212,   327,   291,  -212,  -212,  -212,
    -212,  -212,   416,   334,   293,  -212,  -212,   336,   301,  -212,
    -212,  -212,  -212,  -212,
}

/* YYPGOTO[NTERM-NUM].  */
var yypgoto = []int {
    -212,  -212,   148,   -36,     1,   -62,   350,  -212,    -5,  -212,
    -212,  -212,  -212,  -212,  -211,  -212,  -212,  -212,  -212,  -212,
    -212,    50,  -212,   131,  -212,    92,  -194,    -6,  -212,  -200,
    -212,   -45,   -48,  -212,     5,     2,    12,   362,
}

/* YYTABLE[YYPACT[STATE-NUM]].  What to do in state STATE-NUM.  If
   positive, shift that token.  If negative, reduce the rule which
   number is the opposite.  If zero, do what YYDEFACT says.
   If YYTABLE_NINF, syntax error.  */
const YYTABLE_NINF = -1
var yytable = []int {
      64,    65,    53,   116,    69,    52,   258,   151,   199,   273,
      56,   141,   138,   140,   250,   145,   253,   143,   149,   125,
     224,   225,   226,    71,   126,   274,   111,   112,   262,   263,
     264,   265,   266,   116,     2,    61,    60,   134,   115,     3,
      62,     4,     5,     6,     7,   292,   128,   129,   130,    10,
     134,   289,   301,   142,   286,   287,     2,    63,   275,    66,
      17,     3,   150,     4,     5,     6,     7,   132,   274,   115,
      70,    10,    97,   186,   187,    98,    73,   190,   191,    74,
     101,   185,    17,   102,   153,   303,   237,    32,   216,    33,
      95,   200,   304,   134,   157,   206,   207,   158,   303,   118,
     119,   275,    99,   134,    96,   319,   217,   121,   122,    32,
     103,    33,   194,   195,   116,   185,   264,   265,   266,   246,
     329,   225,   134,   100,   159,    53,    53,   193,   238,   189,
     201,   208,   248,   337,   338,   204,   205,   202,   209,   192,
     211,   212,   113,   308,   214,   240,   228,   229,   230,   231,
     232,   236,   257,   161,   133,   134,   162,   242,   104,   105,
     251,   247,   251,   323,   106,   256,   134,   134,   134,   128,
     129,   130,   203,   210,   249,   327,   107,   108,   116,   134,
     185,   109,   110,   163,   139,   309,   144,   272,    53,    53,
     171,   218,   219,   172,   282,   239,   281,   243,   146,   134,
     251,   251,   222,   223,   280,   324,   175,   290,   147,   176,
     152,   134,   214,   268,   269,   270,   148,   328,   259,   260,
     173,   185,    53,    53,   155,   154,   189,   156,   128,   129,
     130,   277,   278,   307,   160,   164,   177,   194,   195,   165,
     284,   285,   314,   166,   167,   121,   122,   316,   317,   318,
     168,   214,   169,   196,   197,   198,   322,   114,    14,    15,
      16,   123,   124,   298,   299,   300,    18,    19,    20,    21,
      22,   332,   170,   174,    23,    24,    25,    26,    27,   251,
     251,   178,   179,   180,   181,    30,    31,   182,   315,    75,
      76,    77,    78,    79,   183,   184,   244,    80,   321,   213,
      81,    82,   325,   326,   240,   254,   245,   255,    83,    84,
     261,   267,   279,   331,   288,   293,   294,   333,   334,     1,
     283,     2,   291,   339,   305,   274,     3,   306,     4,     5,
       6,     7,   310,   311,     8,     9,    10,   312,   320,   313,
      11,    12,   335,    13,    14,    15,    16,    17,   336,   340,
     341,   342,    18,    19,    20,    21,    22,   252,   343,    68,
      23,    24,    25,    26,    27,   330,   271,   302,    28,    29,
       2,    30,    31,    72,    32,     3,    33,     4,     5,     6,
       7,     0,     0,     8,     9,    10,     0,     0,     0,   127,
       0,     0,    13,    14,    15,    16,    17,     0,     0,     0,
       0,    18,    19,    20,    21,    22,     0,     0,     0,    23,
      24,    25,    26,    27,     0,     0,   128,   129,   130,     2,
      30,    31,     0,    32,     3,    33,     4,     5,     6,     7,
       0,     0,     8,     9,    10,     0,     0,     0,    11,    12,
       0,    13,    14,    15,    16,    17,     0,     0,     0,     0,
      18,    19,    20,    21,    22,     0,     0,     0,    23,    24,
      25,    26,    27,     0,     0,     0,   134,     0,     2,    30,
      31,     0,    32,     3,    33,     4,     5,     6,     7,     0,
       0,     8,     9,    10,     0,     0,     0,    11,    12,     0,
      13,    14,    15,    16,    17,     0,     0,     0,     0,    18,
      19,    20,    21,    22,     0,     0,     0,    23,    24,    25,
      26,    27,     0,     0,     0,     0,     0,     2,    30,    31,
       0,    32,     3,    33,     4,     5,     6,     7,     0,     0,
       8,     9,    10,     0,     0,     0,     0,    12,     0,    13,
      14,    15,    16,    17,     0,     0,     0,     0,    18,    19,
      20,    21,    22,     0,     0,     0,    23,    24,    25,    26,
      27,     0,     0,     0,     0,     0,     2,    30,    31,     0,
      32,     3,    33,     4,     5,     6,     7,     0,     0,     8,
       9,    10,     0,     0,     0,     0,     0,     0,    13,    14,
      15,    16,    17,     0,     0,     0,     0,    18,    19,    20,
      21,    22,     0,     0,     0,    23,    24,    25,    26,    27,
       0,     0,     0,   134,     0,     2,    30,    31,     0,    32,
       3,    33,     4,     5,     6,     7,     0,     0,     8,     9,
      10,     0,     0,     0,     0,     0,     0,    13,    14,    15,
      16,    17,     0,     0,     0,     0,    18,    19,    20,    21,
      22,     0,     0,     0,    23,    24,    25,    26,    27,     0,
       0,     0,     0,     0,     2,    30,    31,     0,    32,     3,
      33,     4,     5,     6,     7,     0,     0,     0,     0,    10,
       0,     0,     0,     0,     0,     0,    67,    14,    15,    16,
      17,     0,     0,     0,     0,    18,    19,    20,    21,    22,
       0,     0,     0,    23,    24,    25,    26,    27,     0,     0,
       0,     0,     0,     0,    30,    31,     0,    32,     0,    33,
      15,    16,     0,     0,     0,     0,     0,    18,    19,    20,
      21,    22,     0,     0,     0,    23,    24,    25,    26,    27,
      85,    86,    87,    88,    89,     0,    30,    31,    90,     0,
       0,    91,    92,     0,     0,     0,     0,     0,     0,    93,
      94,
}

var yycheck = []int {
       6,     7,     0,    39,     9,     0,   217,    69,    21,     9,
      50,    14,    60,    61,   208,    63,   210,    62,    66,    47,
       5,     6,     7,    11,    52,    25,    32,    33,   228,   229,
     230,   231,   232,    69,     3,    25,    25,    50,    37,     8,
      30,    10,    11,    12,    13,   256,    49,    50,    51,    18,
      50,   251,     9,    56,   248,   249,     3,    25,    58,    25,
      29,     8,    67,    10,    11,    12,    13,    55,    25,    68,
      31,    18,    25,   121,   122,    28,    24,   125,   126,    58,
      25,   117,    29,    28,    72,    52,    14,    56,   150,    58,
      25,   139,    59,    50,    25,   143,   144,    28,    52,    50,
      51,    58,    55,    50,    25,    59,   154,    32,    33,    56,
      55,    58,    32,    33,   150,   151,   316,   317,   318,    14,
       5,     6,    50,    25,    55,   123,   124,   133,    56,   124,
      14,    14,    14,   327,   328,   141,   142,    21,    21,   127,
     146,   147,     0,    14,   149,    25,   194,   195,   196,   197,
     198,   199,   214,    25,     4,    50,    28,   202,    25,    25,
     208,    56,   210,    14,    25,   213,    50,    50,    50,    49,
      50,    51,    56,    56,    56,    14,    25,    25,   214,    50,
     216,    25,    25,    55,    49,    56,    49,   235,   186,   187,
      25,   186,   187,    28,   242,   201,   241,   203,    14,    50,
     248,   249,   190,   191,    25,    56,    25,   252,    14,    28,
      19,    50,   217,    39,    40,    41,    58,    56,   224,   225,
      55,   257,   220,   221,    25,    59,   221,    25,    49,    50,
      51,   237,   238,   281,    25,    25,    55,    32,    33,    25,
     246,   247,   290,    25,    25,    32,    33,   295,   296,   297,
      25,   256,    25,    48,    49,    50,   304,    25,    26,    27,
      28,    48,    49,    39,    40,    41,    34,    35,    36,    37,
      38,   319,    25,    25,    42,    43,    44,    45,    46,   327,
     328,    25,    25,    25,    25,    53,    54,    25,   294,    34,
      35,    36,    37,    38,    57,    59,    15,    42,   304,    59,
      45,    46,   308,   309,    25,    15,    57,    15,    53,    54,
       7,     9,    15,   319,    15,     7,     4,   323,   324,     1,
      57,     3,    57,   329,    15,    25,     8,    57,    10,    11,
      12,    13,    15,    57,    16,    17,    18,    15,    25,    57,
      22,    23,    15,    25,    26,    27,    28,    29,    57,    15,
      57,    15,    34,    35,    36,    37,    38,   209,    57,     9,
      42,    43,    44,    45,    46,   315,   235,   275,    50,    51,
       3,    53,    54,    11,    56,     8,    58,    10,    11,    12,
      13,    -1,    -1,    16,    17,    18,    -1,    -1,    -1,    22,
      -1,    -1,    25,    26,    27,    28,    29,    -1,    -1,    -1,
      -1,    34,    35,    36,    37,    38,    -1,    -1,    -1,    42,
      43,    44,    45,    46,    -1,    -1,    49,    50,    51,     3,
      53,    54,    -1,    56,     8,    58,    10,    11,    12,    13,
      -1,    -1,    16,    17,    18,    -1,    -1,    -1,    22,    23,
      -1,    25,    26,    27,    28,    29,    -1,    -1,    -1,    -1,
      34,    35,    36,    37,    38,    -1,    -1,    -1,    42,    43,
      44,    45,    46,    -1,    -1,    -1,    50,    -1,     3,    53,
      54,    -1,    56,     8,    58,    10,    11,    12,    13,    -1,
      -1,    16,    17,    18,    -1,    -1,    -1,    22,    23,    -1,
      25,    26,    27,    28,    29,    -1,    -1,    -1,    -1,    34,
      35,    36,    37,    38,    -1,    -1,    -1,    42,    43,    44,
      45,    46,    -1,    -1,    -1,    -1,    -1,     3,    53,    54,
      -1,    56,     8,    58,    10,    11,    12,    13,    -1,    -1,
      16,    17,    18,    -1,    -1,    -1,    -1,    23,    -1,    25,
      26,    27,    28,    29,    -1,    -1,    -1,    -1,    34,    35,
      36,    37,    38,    -1,    -1,    -1,    42,    43,    44,    45,
      46,    -1,    -1,    -1,    -1,    -1,     3,    53,    54,    -1,
      56,     8,    58,    10,    11,    12,    13,    -1,    -1,    16,
      17,    18,    -1,    -1,    -1,    -1,    -1,    -1,    25,    26,
      27,    28,    29,    -1,    -1,    -1,    -1,    34,    35,    36,
      37,    38,    -1,    -1,    -1,    42,    43,    44,    45,    46,
      -1,    -1,    -1,    50,    -1,     3,    53,    54,    -1,    56,
       8,    58,    10,    11,    12,    13,    -1,    -1,    16,    17,
      18,    -1,    -1,    -1,    -1,    -1,    -1,    25,    26,    27,
      28,    29,    -1,    -1,    -1,    -1,    34,    35,    36,    37,
      38,    -1,    -1,    -1,    42,    43,    44,    45,    46,    -1,
      -1,    -1,    -1,    -1,     3,    53,    54,    -1,    56,     8,
      58,    10,    11,    12,    13,    -1,    -1,    -1,    -1,    18,
      -1,    -1,    -1,    -1,    -1,    -1,    25,    26,    27,    28,
      29,    -1,    -1,    -1,    -1,    34,    35,    36,    37,    38,
      -1,    -1,    -1,    42,    43,    44,    45,    46,    -1,    -1,
      -1,    -1,    -1,    -1,    53,    54,    -1,    56,    -1,    58,
      27,    28,    -1,    -1,    -1,    -1,    -1,    34,    35,    36,
      37,    38,    -1,    -1,    -1,    42,    43,    44,    45,    46,
      34,    35,    36,    37,    38,    -1,    53,    54,    42,    -1,
      -1,    45,    46,    -1,    -1,    -1,    -1,    -1,    -1,    53,
      54,
}

/* YYSTOS[STATE-NUM] -- The (internal number of the) accessing
   symbol of state STATE-NUM.  */
var yystos = []int {
       0,     1,     3,     8,    10,    11,    12,    13,    16,    17,
      18,    22,    23,    25,    26,    27,    28,    29,    34,    35,
      36,    37,    38,    42,    43,    44,    45,    46,    50,    51,
      53,    54,    56,    58,    61,    63,    64,    66,    67,    68,
      69,    70,    71,    72,    73,    75,    76,    77,    78,    79,
      80,    93,    94,    95,    96,    97,    50,    86,    87,    92,
      25,    25,    30,    25,    87,    87,    25,    25,    66,    68,
      31,    96,    97,    24,    58,    34,    35,    36,    37,    38,
      42,    45,    46,    53,    54,    34,    35,    36,    37,    38,
      42,    45,    46,    53,    54,    25,    25,    25,    28,    55,
      25,    25,    28,    55,    25,    25,    25,    25,    25,    25,
      25,    87,    87,     0,    25,    64,    63,    65,    50,    51,
      90,    32,    33,    48,    49,    47,    52,    22,    49,    50,
      51,    91,    96,     4,    50,    88,    89,    95,    92,    49,
      92,    14,    56,    91,    49,    92,    14,    14,    58,    92,
      68,    65,    19,    96,    59,    25,    25,    25,    28,    55,
      25,    25,    28,    55,    25,    25,    25,    25,    25,    25,
      25,    25,    28,    55,    25,    25,    28,    55,    25,    25,
      25,    25,    25,    57,    59,    63,    92,    92,    94,    94,
      92,    92,    96,    87,    32,    33,    48,    49,    50,    21,
      92,    14,    21,    56,    87,    87,    92,    92,    14,    21,
      56,    87,    87,    59,    68,    74,    65,    92,    94,    94,
      48,    49,    96,    96,     5,     6,     7,    81,    92,    92,
      92,    92,    92,    82,    83,    84,    92,    14,    56,    87,
      25,    62,    91,    87,    15,    57,    14,    56,    14,    56,
      86,    92,    62,    86,    15,    15,    92,    65,    74,    87,
      87,     7,    89,    89,    89,    89,    89,     9,    39,    40,
      41,    83,    92,     9,    25,    58,    85,    87,    87,    15,
      25,    91,    92,    57,    87,    87,    86,    86,    15,    89,
      91,    57,    74,     7,     4,    48,    49,    50,    39,    40,
      41,     9,    85,    52,    59,    15,    57,    92,    14,    56,
      15,    57,    15,    57,    92,    87,    92,    92,    92,    59,
      25,    87,    92,    14,    56,    87,    87,    14,    56,     5,
      81,    87,    92,    87,    87,    15,    57,    86,    86,    87,
      15,    57,    15,    57,
}

const YYEMPTY = (-2)
const YYEOF = 0

const YYTERROR = 1
const YYERRCODE = 256

/* YYINITDEPTH -- initial size of the parser's stacks.  */
const YYINITDEPTH = 200

/* YYMAXDEPTH -- maximum size the stacks can grow to (effective only
   if the built-in stack extension method is used).

   Do not make this value too large; the results are undefined if
   YYSTACK_ALLOC_MAXIMUM < YYSTACK_BYTES (YYMAXDEPTH)
   evaluated with infinite-precision integer arithmetic.  */

const YYMAXDEPTH = 10000

type yyparseState int

const (
	yysetstate yyparseState = 0
	yybackup yyparseState = iota
	yydefault yyparseState = iota
	yyreduce yyparseState = iota
	yyerrorlab yyparseState = iota
	yyerrlab yyparseState = iota
	yyerrlab1 yyparseState = iota
	yyacceptlab yyparseState = iota
	yyabortlab yyparseState = iota
	yyreturn yyparseState = iota
)

/*----------.
| yyparse.  |
`----------*/

func (gps *ParserState) Yyparse () (yyresult int) {
  var yystate int
  var yyn int
  /* Number of tokens to shift before error messages enabled.  */
  var yyerrstatus int
  /* Look-ahead token as an internal (translated) token number.  */
  var yytoken int = 0
  /* Two stacks and their tools:
     `yyss': related to states,
     `yyvs': related to semantic values,

     Refer to the stacks thru separate pointers, to allow yyoverflow
     to reallocate them elsewhere.  */

  /* The state stack.  */
	yyss := newIntStack()

  /* The semantic value stack.  */
	yyvs := newYYSTYPEStack()

	popStack := func(n int) {
		yyss.PopMany(n)
		yyvs.PopMany(n)
	}


  /* The variables used to return semantic value and location from the
     action routines.  */
  var yyval YYSTYPE


  /* The number of symbols on the RHS of the reduced rule.
     Keep to zero when no symbol should be popped. */
  var yylen int = 0

  yystate = 0;
  yyerrstatus = 0;
  gps.yynerrs = 0;
  gps.yychar = YYEMPTY;		/* Cause a token to be read.  */

	yyparseState := yysetstate

for {
	switch yyparseState {

/*------------------------------------------------------------.
| yysetstate -- Push a new state, which is found in yystate.  |
`------------------------------------------------------------*/
case yysetstate:
	yyss.Push(yystate)
	yyparseState = yybackup

/*-----------.
| yybackup.  |
`-----------*/
case yybackup:

  /* Do appropriate processing given the current state.  Read a
     look-ahead token if we need one and don't already have one.  */

  /* First try to decide what to do without reference to look-ahead token.  */
  yyn = yypact[yystate];
  if (yyn == YYPACT_NINF) {
    yyparseState = yydefault;
	continue
  }

  /* Not known => get a look-ahead token if don't already have one.  */

  /* YYCHAR is either YYEMPTY or YYEOF or a valid look-ahead symbol.  */
  if (gps.yychar == YYEMPTY) {
      gps.yychar = gps.yylex();
  }

  if (gps.yychar <= YYEOF) {
      yytoken = YYEOF;
	gps.yychar = YYEOF
    } else {
      yytoken = YYTRANSLATE (gps.yychar);
    }

  /* If the proper action on seeing token YYTOKEN is to reduce or to
     detect an error, take that action.  */
  yyn += yytoken;
	if (yyn < 0 || YYLAST < yyn || yycheck[yyn] != int(yytoken)) {
		yyparseState = yydefault
		continue
	}
  yyn = yytable[yyn];
  if (yyn <= 0)    {
      if (yyn == 0 || yyn == YYTABLE_NINF) {
		yyparseState = yyerrlab
		continue
	}
      yyn = -yyn;
      yyparseState = yyreduce;
	continue
    }

  if (yyn == YYFINAL) {
    yyparseState = yyacceptlab; continue;
}

  /* Count tokens shifted since error; after three, turn off error
     status.  */
  if (yyerrstatus > 0) {
    yyerrstatus--;
  }

  /* Discard the shifted token unless it is eof.  */
  if (gps.yychar != YYEOF) {
    gps.yychar = YYEMPTY;
  }

  yystate = yyn;
  yyvs.Push(gps.yylval);

  yyparseState = yysetstate;
	continue


/*-----------------------------------------------------------.
| yydefault -- do the default action for the current state.  |
`-----------------------------------------------------------*/
case yydefault:
  yyn = yydefact[yystate];
  if (yyn == 0) {
    yyparseState = yyerrlab; continue
  }
  yyparseState = yyreduce; continue;


/*-----------------------------.
| yyreduce -- Do a reduction.  |
`-----------------------------*/
case yyreduce:
  /* yyn is the number of a rule to reduce with.  */
  yylen = yyr2[yyn];

  /* If YYLEN is nonzero, implement the default value of the action:
     `$$ = $1'.

     Otherwise, the following line sets YYVAL to garbage.
     This behavior is undocumented and Bison
     users should not rely upon it.  Assigning to YYVAL
     unconditionally makes the parser a bit smaller, and it avoids a
     GCC warning that YYVAL may be used uninitialized.  */
  yyval = yyvs.PeekN(1-yylen)


  switch (yyn)  {
        case 2:
// #line 374 "/Users/chet/src/bash/src/parse.y"
    {
			  /* Case of regular command.  Discard the error
			     safety net,and return the command just parsed. */
			  gps.global_command = (yyvs.PeekN((1) - (2)).command);
			  /* discard_parser_constructs (0); */
			  if (gps.parser_state & PST_CMDSUBST != 0) {
			    gps.parser_state |= PST_EOFTOKEN;
			  }
			  yyparseState = yyacceptlab; continue;
			}
    break;

  case 3:
// #line 385 "/Users/chet/src/bash/src/parse.y"
    {
			  /* Case of regular command, but not a very
			     interesting one.  Return a NULL command. */
			  gps.global_command = nil;
			  if (gps.parser_state & PST_CMDSUBST != 0) {
			    gps.parser_state |= PST_EOFTOKEN;
			  }
			  yyparseState = yyacceptlab; continue;
			}
    break;

  case 4:
// #line 394 "/Users/chet/src/bash/src/parse.y"
    {
			  /* Error during parsing.  Return NULL command. */
			  gps.global_command = nil
			  /* discard_parser_constructs (1); */
			  yyparseState = yyabortlab; continue;
			}
    break;

  case 5:
// #line 409 "/Users/chet/src/bash/src/parse.y"
    {
			  /* Case of EOF seen by itself.  Do ignoreeof or
			     not. */
			  gps.global_command = nil
			  gps.handle_eof_input_unit ();
			  yyparseState = yyacceptlab; continue;
			}
    break;

  case 6:
// #line 419 "/Users/chet/src/bash/src/parse.y"
    { (yyval.word_list) = makeWordList ((yyvs.PeekN((1) - (1)).word), nil); }
    break;

  case 7:
// #line 421 "/Users/chet/src/bash/src/parse.y"
    { (yyval.word_list) = makeWordList ((yyvs.PeekN((2) - (2)).word), (yyvs.PeekN((1) - (2)).word_list)); }
    break;

  case 8:
// #line 425 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 1;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_output_direction, gps.redir, 0);
			}
    break;

  case 9:
// #line 431 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 0;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_input_direction, gps.redir, 0);
			}
    break;

  case 10:
// #line 437 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_output_direction, gps.redir, 0);
			}
    break;

  case 11:
// #line 443 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_input_direction, gps.redir, 0);
			}
    break;

  case 12:
// #line 449 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_output_direction, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 13:
// #line 455 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_input_direction, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 14:
// #line 461 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 1;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_appending_to, gps.redir, 0);
			}
    break;

  case 15:
// #line 467 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_appending_to, gps.redir, 0);
			}
    break;

  case 16:
// #line 473 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_appending_to, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 17:
// #line 479 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 1;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_output_force, gps.redir, 0);
			}
    break;

  case 18:
// #line 485 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_output_force, gps.redir, 0);
			}
    break;

  case 19:
// #line 491 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_output_force, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 20:
// #line 497 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 0;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_input_output, gps.redir, 0);
			}
    break;

  case 21:
// #line 503 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_input_output, gps.redir, 0);
			}
    break;

  case 22:
// #line 509 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_input_output, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 23:
// #line 515 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 0;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_reading_until, gps.redir, 0);
			  gps.redir_stack[gps.need_here_doc] = (yyval.redirect);
				gps.need_here_doc++
			}
    break;

  case 24:
// #line 522 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_reading_until, gps.redir, 0);
			  gps.redir_stack[gps.need_here_doc] = (yyval.redirect);
				gps.need_here_doc++
			}
    break;

  case 25:
// #line 529 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_reading_until, gps.redir, REDIR_VARASSIGN);
			  gps.redir_stack[gps.need_here_doc] = (yyval.redirect);
				gps.need_here_doc++
			}
    break;

  case 26:
// #line 536 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 0;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_deblank_reading_until, gps.redir, 0);
			  gps.redir_stack[gps.need_here_doc] = (yyval.redirect);
				gps.need_here_doc++
			}
    break;

  case 27:
// #line 543 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_deblank_reading_until, gps.redir, 0);
			  gps.redir_stack[gps.need_here_doc] = (yyval.redirect);
				gps.need_here_doc++
			}
    break;

  case 28:
// #line 550 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_deblank_reading_until, gps.redir, REDIR_VARASSIGN);
			  gps.redir_stack[gps.need_here_doc] = (yyval.redirect);
				gps.need_here_doc++
			}
    break;

  case 29:
// #line 557 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 0;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_reading_string, gps.redir, 0);
			}
    break;

  case 30:
// #line 563 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_reading_string, gps.redir, 0);
			}
    break;

  case 31:
// #line 569 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_reading_string, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 32:
// #line 575 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 0;
			  gps.redir.dest = (yyvs.PeekN((2) - (2)).number);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_input, gps.redir, 0);
			}
    break;

  case 33:
// #line 581 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.dest = (yyvs.PeekN((3) - (3)).number);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_input, gps.redir, 0);
			}
    break;

  case 34:
// #line 587 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.dest = (yyvs.PeekN((3) - (3)).number);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_input, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 35:
// #line 593 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 1;
			  gps.redir.dest = (yyvs.PeekN((2) - (2)).number);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_output, gps.redir, 0);
			}
    break;

  case 36:
// #line 599 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.dest = (yyvs.PeekN((3) - (3)).number);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_output, gps.redir, 0);
			}
    break;

  case 37:
// #line 605 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.dest = (yyvs.PeekN((3) - (3)).number);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_output, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 38:
// #line 611 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 0;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_input_word, gps.redir, 0);
			}
    break;

  case 39:
// #line 617 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_input_word, gps.redir, 0);
			}
    break;

  case 40:
// #line 623 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_input_word, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 41:
// #line 629 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 1;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_output_word, gps.redir, 0);
			}
    break;

  case 42:
// #line 635 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_output_word, gps.redir, 0);
			}
    break;

  case 43:
// #line 641 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.filename = (yyvs.PeekN((3) - (3)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_duplicating_output_word, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 44:
// #line 647 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 1;
			  gps.redir.dest = 0;
			  (yyval.redirect) = makeRedirection (gps.source, r_close_this, gps.redir, 0);
			}
    break;

  case 45:
// #line 653 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.dest = 0;
			  (yyval.redirect) = makeRedirection (gps.source, r_close_this, gps.redir, 0);
			}
    break;

  case 46:
// #line 659 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.dest = 0;
			  (yyval.redirect) = makeRedirection (gps.source, r_close_this, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 47:
// #line 665 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 0;
			  gps.redir.dest = 0;
			  (yyval.redirect) = makeRedirection (gps.source, r_close_this, gps.redir, 0);
			}
    break;

  case 48:
// #line 671 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = (yyvs.PeekN((1) - (3)).number);
			  gps.redir.dest = 0;
			  (yyval.redirect) = makeRedirection (gps.source, r_close_this, gps.redir, 0);
			}
    break;

  case 49:
// #line 677 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.filename = (yyvs.PeekN((1) - (3)).word);
			  gps.redir.dest = 0;
			  (yyval.redirect) = makeRedirection (gps.source, r_close_this, gps.redir, REDIR_VARASSIGN);
			}
    break;

  case 50:
// #line 683 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 1;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_err_and_out, gps.redir, 0);
			}
    break;

  case 51:
// #line 689 "/Users/chet/src/bash/src/parse.y"
    {
			  gps.source.dest = 1;
			  gps.redir.filename = (yyvs.PeekN((2) - (2)).word);
			  (yyval.redirect) = makeRedirection (gps.source, r_append_err_and_out, gps.redir, 0);
			}
    break;

  case 52:
// #line 697 "/Users/chet/src/bash/src/parse.y"
    { (yyval.element).word = (yyvs.PeekN((1) - (1)).word); (yyval.element).redirect = nil; }
    break;

  case 53:
// #line 699 "/Users/chet/src/bash/src/parse.y"
    { (yyval.element).word = (yyvs.PeekN((1) - (1)).word); (yyval.element).redirect = nil; }
    break;

  case 54:
// #line 701 "/Users/chet/src/bash/src/parse.y"
    { (yyval.element).redirect = (yyvs.PeekN((1) - (1)).redirect); (yyval.element).word = nil; }
    break;

  case 55:
// #line 705 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.redirect) = (yyvs.PeekN((1) - (1)).redirect);
			}
    break;

  case 56:
// #line 709 "/Users/chet/src/bash/src/parse.y"
    {
			  var t *Redirect

			  for t = (yyvs.PeekN((1) - (2)).redirect); t.next != nil; t = t.next {
			  }
			  t.next = (yyvs.PeekN((2) - (2)).redirect);
			  (yyval.redirect) = (yyvs.PeekN((1) - (2)).redirect);
			}
    break;

  case 57:
// #line 720 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_simple_command ((yyvs.PeekN((1) - (1)).element), nil); }
    break;

  case 58:
// #line 722 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_simple_command ((yyvs.PeekN((2) - (2)).element), (yyvs.PeekN((1) - (2)).command)); }
    break;

  case 59:
// #line 726 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.clean_simple_command ((yyvs.PeekN((1) - (1)).command)); }
    break;

  case 60:
// #line 728 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 61:
// #line 730 "/Users/chet/src/bash/src/parse.y"
    {
			  var tc *Command

			  tc = (yyvs.PeekN((1) - (2)).command);
			  if (tc.redirects != nil)			    {
			      var t *Redirect;
			      for t = tc.redirects; t.next != nil; t = t.next {
				}
			      t.next = (yyvs.PeekN((2) - (2)).redirect);
			    } else {
			    tc.redirects = (yyvs.PeekN((2) - (2)).redirect);
			}
			  (yyval.command) = (yyvs.PeekN((1) - (2)).command);
			}
    break;

  case 62:
// #line 746 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 63:
// #line 748 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 64:
// #line 752 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 65:
// #line 754 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 66:
// #line 756 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_while_command ((yyvs.PeekN((2) - (5)).command), (yyvs.PeekN((4) - (5)).command)); }
    break;

  case 67:
// #line 758 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_until_command ((yyvs.PeekN((2) - (5)).command), (yyvs.PeekN((4) - (5)).command)); }
    break;

  case 68:
// #line 760 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 69:
// #line 762 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 70:
// #line 764 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 71:
// #line 766 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 72:
// #line 768 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 73:
// #line 770 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 74:
// #line 772 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 75:
// #line 776 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_for_command ((yyvs.PeekN((2) - (6)).word), add_string_to_list ("\"$@\"", nil), (yyvs.PeekN((5) - (6)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top--; }
			}
    break;

  case 76:
// #line 781 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_for_command ((yyvs.PeekN((2) - (6)).word), add_string_to_list ("\"$@\"", nil), (yyvs.PeekN((5) - (6)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top--; }
			}
    break;

  case 77:
// #line 786 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_for_command ((yyvs.PeekN((2) - (7)).word), add_string_to_list ("\"$@\"", nil), (yyvs.PeekN((6) - (7)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top--; }
			}
    break;

  case 78:
// #line 791 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_for_command ((yyvs.PeekN((2) - (7)).word), add_string_to_list ("\"$@\"", nil), (yyvs.PeekN((6) - (7)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top--; }
			}
    break;

  case 79:
// #line 796 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_for_command ((yyvs.PeekN((2) - (10)).word), reverseWordList(yyvs.PeekN((5) - (10)).word_list), (yyvs.PeekN((9) - (10)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top--; }
			}
    break;

  case 80:
// #line 801 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_for_command ((yyvs.PeekN((2) - (10)).word), reverseWordList(yyvs.PeekN((5) - (10)).word_list), (yyvs.PeekN((9) - (10)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top--; }
			}
    break;

  case 81:
// #line 806 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_for_command ((yyvs.PeekN((2) - (9)).word), nil, (yyvs.PeekN((8) - (9)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top--; }
			}
    break;

  case 82:
// #line 811 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_for_command ((yyvs.PeekN((2) - (9)).word), nil, (yyvs.PeekN((8) - (9)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top--; }
			}
    break;

  case 83:
// #line 818 "/Users/chet/src/bash/src/parse.y"
    {
				  (yyval.command) = gps.make_arith_for_command ((yyvs.PeekN((2) - (7)).word_list), (yyvs.PeekN((6) - (7)).command), gps.arith_for_lineno);
				  if (gps.word_top > 0) { gps.word_top--; }
				}
    break;

  case 84:
// #line 823 "/Users/chet/src/bash/src/parse.y"
    {
				  (yyval.command) = gps.make_arith_for_command ((yyvs.PeekN((2) - (7)).word_list), (yyvs.PeekN((6) - (7)).command), gps.arith_for_lineno);
				  if (gps.word_top > 0) { gps.word_top--; }
				}
    break;

  case 85:
// #line 828 "/Users/chet/src/bash/src/parse.y"
    {
				  (yyval.command) = gps.make_arith_for_command ((yyvs.PeekN((2) - (5)).word_list), (yyvs.PeekN((4) - (5)).command), gps.arith_for_lineno);
				  if (gps.word_top > 0) { gps.word_top--; }
				}
    break;

  case 86:
// #line 833 "/Users/chet/src/bash/src/parse.y"
    {
				  (yyval.command) = gps.make_arith_for_command ((yyvs.PeekN((2) - (5)).word_list), (yyvs.PeekN((4) - (5)).command), gps.arith_for_lineno);
				  if (gps.word_top > 0) { gps.word_top--; }
				}
    break;

  case 87:
// #line 840 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_select_command ((yyvs.PeekN((2) - (6)).word), add_string_to_list ("\"$@\"", nil), (yyvs.PeekN((5) - (6)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top-- }
			}
    break;

  case 88:
// #line 845 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_select_command ((yyvs.PeekN((2) - (6)).word), add_string_to_list ("\"$@\"", nil), (yyvs.PeekN((5) - (6)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top-- }
			}
    break;

  case 89:
// #line 850 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_select_command ((yyvs.PeekN((2) - (7)).word), add_string_to_list ("\"$@\"", nil), (yyvs.PeekN((6) - (7)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top-- }
			}
    break;

  case 90:
// #line 855 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_select_command ((yyvs.PeekN((2) - (7)).word), add_string_to_list ("\"$@\"", nil), (yyvs.PeekN((6) - (7)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top-- }
			}
    break;

  case 91:
// #line 860 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_select_command ((yyvs.PeekN((2) - (10)).word), reverseWordList(yyvs.PeekN((5) - (10)).word_list), (yyvs.PeekN((9) - (10)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top-- }
			}
    break;

  case 92:
// #line 865 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_select_command ((yyvs.PeekN((2) - (10)).word), reverseWordList(yyvs.PeekN((5) - (10)).word_list), (yyvs.PeekN((9) - (10)).command), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top-- }
			}
    break;

  case 93:
// #line 872 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_case_command ((yyvs.PeekN((2) - (6)).word), nil, gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top-- }
			}
    break;

  case 94:
// #line 877 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_case_command ((yyvs.PeekN((2) - (7)).word), (yyvs.PeekN((5) - (7)).pattern), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top-- }
			}
    break;

  case 95:
// #line 882 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_case_command ((yyvs.PeekN((2) - (6)).word), (yyvs.PeekN((5) - (6)).pattern), gps.word_lineno[gps.word_top]);
			  if (gps.word_top > 0) { gps.word_top-- }
			}
    break;

  case 96:
// #line 889 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_function_def ((yyvs.PeekN((1) - (5)).word), (yyvs.PeekN((5) - (5)).command), gps.function_dstart, gps.function_bstart); }
    break;

  case 97:
// #line 892 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_function_def ((yyvs.PeekN((2) - (6)).word), (yyvs.PeekN((6) - (6)).command), gps.function_dstart, gps.function_bstart); }
    break;

  case 98:
// #line 895 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_function_def ((yyvs.PeekN((2) - (4)).word), (yyvs.PeekN((4) - (4)).command), gps.function_dstart, gps.function_bstart); }
    break;

  case 99:
// #line 899 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 100:
// #line 901 "/Users/chet/src/bash/src/parse.y"
    {
			  var tc *Command

			  tc = (yyvs.PeekN((1) - (2)).command);
			  /* According to Posix.2 3.9.5, redirections
			     specified after the body of a function should
			     be attached to the function and performed when
			     the function is executed, not as part of the
			     function definition command. */
			  /* XXX - I don't think it matters, but we might
			     want to change this in the future to avoid
			     problems differentiating between a function
			     definition with a redirection and a function
			     definition containing a single command with a
			     redirection.  The two are semantically equivalent,
			     though -- the only difference is in how the
			     command printing code displays the redirections. */
			  if (tc.redirects != nil) {
			      var t *Redirect
			      for t = tc.redirects; t.next != nil; t = t.next {
			      }
			      t.next = (yyvs.PeekN((2) - (2)).redirect);
			    } else {
			    tc.redirects = (yyvs.PeekN((2) - (2)).redirect);
				}
			  (yyval.command) = (yyvs.PeekN((1) - (2)).command);
			}
    break;

  case 101:
// #line 932 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_subshell_command ((yyvs.PeekN((2) - (3)).command));
			  (yyval.command).flags |= CMD_WANT_SUBSHELL;
			}
    break;

  case 102:
// #line 939 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_coproc_command ("COPROC", (yyvs.PeekN((2) - (2)).command));
			  (yyval.command).flags |= CMD_WANT_SUBSHELL|CMD_COPROC_SUBSHELL;
			}
    break;

  case 103:
// #line 944 "/Users/chet/src/bash/src/parse.y"
    {
			  var tc *Command

			  tc = (yyvs.PeekN((2) - (3)).command);
			  if tc.redirects != nil {
			      var t *Redirect
			      for t = tc.redirects; t.next != nil; t = t.next {}

			      t.next = (yyvs.PeekN((3) - (3)).redirect);
			    }  else {
			    tc.redirects = (yyvs.PeekN((3) - (3)).redirect);
			  }
			  (yyval.command) = gps.make_coproc_command ("COPROC", (yyvs.PeekN((2) - (3)).command));
			  (yyval.command).flags |= CMD_WANT_SUBSHELL|CMD_COPROC_SUBSHELL;
			}
    break;

  case 104:
// #line 961 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_coproc_command ((yyvs.PeekN((2) - (3)).word).word, (yyvs.PeekN((3) - (3)).command));
			  (yyval.command).flags |= CMD_WANT_SUBSHELL|CMD_COPROC_SUBSHELL;
			}
    break;

  case 105:
// #line 966 "/Users/chet/src/bash/src/parse.y"
    {
			  var tc *Command

			  tc = (yyvs.PeekN((3) - (4)).command);
			  if tc.redirects != nil {
			      var t *Redirect
			      for t = tc.redirects; t.next != nil; t = t.next {}

			      t.next = (yyvs.PeekN((4) - (4)).redirect);
			    }			  else {
			    tc.redirects = (yyvs.PeekN((4) - (4)).redirect);
				}
			  (yyval.command) = gps.make_coproc_command ((yyvs.PeekN((2) - (4)).word).word, (yyvs.PeekN((3) - (4)).command));
			  (yyval.command).flags |= CMD_WANT_SUBSHELL|CMD_COPROC_SUBSHELL;
			}
    break;

  case 106:
// #line 983 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = gps.make_coproc_command ("COPROC", gps.clean_simple_command ((yyvs.PeekN((2) - (2)).command)));
			  (yyval.command).flags |= CMD_WANT_SUBSHELL|CMD_COPROC_SUBSHELL;
			}
    break;

  case 107:
// #line 990 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_if_command ((yyvs.PeekN((2) - (5)).command), (yyvs.PeekN((4) - (5)).command), nil); }
    break;

  case 108:
// #line 992 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_if_command ((yyvs.PeekN((2) - (7)).command), (yyvs.PeekN((4) - (7)).command), (yyvs.PeekN((6) - (7)).command)); }
    break;

  case 109:
// #line 994 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_if_command ((yyvs.PeekN((2) - (6)).command), (yyvs.PeekN((4) - (6)).command), (yyvs.PeekN((5) - (6)).command)); }
    break;

  case 110:
// #line 999 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_group_command ((yyvs.PeekN((2) - (3)).command)); }
    break;

  case 111:
// #line 1003 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_arith_command ((yyvs.PeekN((1) - (1)).word_list)); }
    break;

  case 112:
// #line 1007 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((2) - (3)).command); }
    break;

  case 113:
// #line 1011 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_if_command ((yyvs.PeekN((2) - (4)).command), (yyvs.PeekN((4) - (4)).command), nil); }
    break;

  case 114:
// #line 1013 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_if_command ((yyvs.PeekN((2) - (6)).command), (yyvs.PeekN((4) - (6)).command), (yyvs.PeekN((6) - (6)).command)); }
    break;

  case 115:
// #line 1015 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.make_if_command ((yyvs.PeekN((2) - (5)).command), (yyvs.PeekN((4) - (5)).command), (yyvs.PeekN((5) - (5)).command)); }
    break;

  case 117:
// #line 1020 "/Users/chet/src/bash/src/parse.y"
    { (yyvs.PeekN((2) - (2)).pattern).next = (yyvs.PeekN((1) - (2)).pattern); (yyval.pattern) = (yyvs.PeekN((2) - (2)).pattern); }
    break;

  case 118:
// #line 1024 "/Users/chet/src/bash/src/parse.y"
    { (yyval.pattern) = gps.make_pattern_list ((yyvs.PeekN((2) - (4)).word_list), (yyvs.PeekN((4) - (4)).command)); }
    break;

  case 119:
// #line 1026 "/Users/chet/src/bash/src/parse.y"
    { (yyval.pattern) = gps.make_pattern_list ((yyvs.PeekN((2) - (4)).word_list), nil); }
    break;

  case 120:
// #line 1028 "/Users/chet/src/bash/src/parse.y"
    { (yyval.pattern) = gps.make_pattern_list ((yyvs.PeekN((3) - (5)).word_list), (yyvs.PeekN((5) - (5)).command)); }
    break;

  case 121:
// #line 1030 "/Users/chet/src/bash/src/parse.y"
    { (yyval.pattern) = gps.make_pattern_list ((yyvs.PeekN((3) - (5)).word_list), nil); }
    break;

  case 122:
// #line 1034 "/Users/chet/src/bash/src/parse.y"
    { (yyval.pattern) = (yyvs.PeekN((1) - (2)).pattern); }
    break;

  case 123:
// #line 1036 "/Users/chet/src/bash/src/parse.y"
    { (yyvs.PeekN((2) - (3)).pattern).next = (yyvs.PeekN((1) - (3)).pattern); (yyval.pattern) = (yyvs.PeekN((2) - (3)).pattern); }
    break;

  case 124:
// #line 1038 "/Users/chet/src/bash/src/parse.y"
    { (yyvs.PeekN((1) - (2)).pattern).flags |= CASEPAT_FALLTHROUGH; (yyval.pattern) = (yyvs.PeekN((1) - (2)).pattern); }
    break;

  case 125:
// #line 1040 "/Users/chet/src/bash/src/parse.y"
    { (yyvs.PeekN((2) - (3)).pattern).flags |= CASEPAT_FALLTHROUGH; (yyvs.PeekN((2) - (3)).pattern).next = (yyvs.PeekN((1) - (3)).pattern); (yyval.pattern) = (yyvs.PeekN((2) - (3)).pattern); }
    break;

  case 126:
// #line 1042 "/Users/chet/src/bash/src/parse.y"
    { (yyvs.PeekN((1) - (2)).pattern).flags |= CASEPAT_TESTNEXT; (yyval.pattern) = (yyvs.PeekN((1) - (2)).pattern); }
    break;

  case 127:
// #line 1044 "/Users/chet/src/bash/src/parse.y"
    { (yyvs.PeekN((2) - (3)).pattern).flags |= CASEPAT_TESTNEXT; (yyvs.PeekN((2) - (3)).pattern).next = (yyvs.PeekN((1) - (3)).pattern); (yyval.pattern) = (yyvs.PeekN((2) - (3)).pattern); }
    break;

  case 128:
// #line 1048 "/Users/chet/src/bash/src/parse.y"
    { (yyval.word_list) = makeWordList ((yyvs.PeekN((1) - (1)).word), nil); }
    break;

  case 129:
// #line 1050 "/Users/chet/src/bash/src/parse.y"
    { (yyval.word_list) = makeWordList ((yyvs.PeekN((3) - (3)).word), (yyvs.PeekN((1) - (3)).word_list)); }
    break;

  case 130:
// #line 1059 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = (yyvs.PeekN((2) - (2)).command);
			  if (gps.need_here_doc != 0) {
			    gps.gather_here_documents ();
			}
    }
    break;

  case 132:
// #line 1068 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = (yyvs.PeekN((2) - (2)).command);
			}
    break;

  case 134:
// #line 1075 "/Users/chet/src/bash/src/parse.y"
    {
			  if ((yyvs.PeekN((1) - (3)).command).typ == cm_connection) {
			    (yyval.command) = gps.connect_async_list ((yyvs.PeekN((1) - (3)).command), nil, '&');
			  } else {
			    (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (3)).command), nil, '&');
			  }
    }
    break;

  case 136:
// #line 1086 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (4)).command), (yyvs.PeekN((4) - (4)).command), AND_AND); }
    break;

  case 137:
// #line 1088 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (4)).command), (yyvs.PeekN((4) - (4)).command), OR_OR); }
    break;

  case 138:
// #line 1090 "/Users/chet/src/bash/src/parse.y"
    {
			  if ((yyvs.PeekN((1) - (4)).command).typ == cm_connection) {
			    (yyval.command) = gps.connect_async_list ((yyvs.PeekN((1) - (4)).command), (yyvs.PeekN((4) - (4)).command), '&');
			  } else {
			    (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (4)).command), (yyvs.PeekN((4) - (4)).command), '&');
			  }
			}
    break;

  case 139:
// #line 1097 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (4)).command), (yyvs.PeekN((4) - (4)).command), ';'); }
    break;

  case 140:
// #line 1099 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (4)).command), (yyvs.PeekN((4) - (4)).command), ';'); }
    break;

  case 141:
// #line 1101 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 144:
// #line 1109 "/Users/chet/src/bash/src/parse.y"
    { (yyval.number) = '\n'; }
    break;

  case 145:
// #line 1111 "/Users/chet/src/bash/src/parse.y"
    { (yyval.number) = ';'; }
    break;

  case 146:
// #line 1113 "/Users/chet/src/bash/src/parse.y"
    { (yyval.number) = yacc_EOF; }
    break;

  case 149:
// #line 1127 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = (yyvs.PeekN((1) - (1)).command);
			  if (gps.need_here_doc != 0) {
			    gps.gather_here_documents ();
			  }
			  if ((gps.parser_state & PST_CMDSUBST != 0) && gps.current_token == gps.shell_eof_token) {
			      gps.global_command = (yyvs.PeekN((1) - (1)).command);
			      gps.rewind_input_string ();
			      yyparseState = yyacceptlab; continue;
			    }
    }
    break;

  case 150:
// #line 1140 "/Users/chet/src/bash/src/parse.y"
    {
			  if ((yyvs.PeekN((1) - (2)).command).typ == cm_connection) {
			    (yyval.command) = gps.connect_async_list ((yyvs.PeekN((1) - (2)).command), nil, '&');
			  } else {
			    (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (2)).command), nil, '&');
			  }
			  if (gps.need_here_doc != 0) {
			    gps.gather_here_documents ();
			  }
			  if ((gps.parser_state & PST_CMDSUBST != 0) && gps.current_token == gps.shell_eof_token)   {
			      gps.global_command = (yyvs.PeekN((1) - (2)).command);
			      gps.rewind_input_string ();
			      yyparseState = yyacceptlab; continue;
			    }
			}
    break;

  case 151:
// #line 1156 "/Users/chet/src/bash/src/parse.y"
    {
			  (yyval.command) = (yyvs.PeekN((1) - (2)).command);
			  if (gps.need_here_doc != 0) {
			    gps.gather_here_documents ();
			  }
			  if ((gps.parser_state & PST_CMDSUBST != 0) && gps.current_token == gps.shell_eof_token)  {
			      gps.global_command = (yyvs.PeekN((1) - (2)).command);
			      gps.rewind_input_string ();
			      yyparseState = yyacceptlab; continue;
			    }
			}
    break;

  case 152:
// #line 1171 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (4)).command), (yyvs.PeekN((4) - (4)).command), AND_AND); }
    break;

  case 153:
// #line 1173 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (4)).command), (yyvs.PeekN((4) - (4)).command), OR_OR); }
    break;

  case 154:
// #line 1175 "/Users/chet/src/bash/src/parse.y"
    {
			  if ((yyvs.PeekN((1) - (3)).command).typ == cm_connection) {
			    (yyval.command) = gps.connect_async_list ((yyvs.PeekN((1) - (3)).command), (yyvs.PeekN((3) - (3)).command), '&');
			  } else {
			    (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (3)).command), (yyvs.PeekN((3) - (3)).command), '&');
			  }
    }
    break;

  case 155:
// #line 1182 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (3)).command), (yyvs.PeekN((3) - (3)).command), ';'); }
    break;

  case 156:
// #line 1185 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 157:
// #line 1189 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 158:
// #line 1191 "/Users/chet/src/bash/src/parse.y"
    {
			  if ((yyvs.PeekN((2) - (2)).command) != nil) {
			    (yyvs.PeekN((2) - (2)).command).flags |= CMD_INVERT_RETURN;
			  }
			  (yyval.command) = (yyvs.PeekN((2) - (2)).command);
			}
    break;

  case 159:
// #line 1197 "/Users/chet/src/bash/src/parse.y"
    {
			  if ((yyvs.PeekN((2) - (2)).command) != nil) {
			    (yyvs.PeekN((2) - (2)).command).flags |= (yyvs.PeekN((1) - (2)).number);
			  }
			  (yyval.command) = (yyvs.PeekN((2) - (2)).command);
			}
    break;

  case 160:
// #line 1203 "/Users/chet/src/bash/src/parse.y"
    {
			  if ((yyvs.PeekN((3) - (3)).command) != nil) {
			    (yyvs.PeekN((3) - (3)).command).flags |= (yyvs.PeekN((1) - (3)).number)|CMD_INVERT_RETURN;
			  }
			  (yyval.command) = (yyvs.PeekN((3) - (3)).command);
			}
    break;

  case 161:
// #line 1209 "/Users/chet/src/bash/src/parse.y"
    {
			  if ((yyvs.PeekN((3) - (3)).command) != nil) {
			    (yyvs.PeekN((3) - (3)).command).flags |= (yyvs.PeekN((2) - (3)).number)|CMD_INVERT_RETURN;
			  }
			  (yyval.command) = (yyvs.PeekN((3) - (3)).command);
			}
    break;

  case 162:
// #line 1215 "/Users/chet/src/bash/src/parse.y"
    {
			  var x ELEMENT

			  /* Boy, this is unclean.  `time' by itself can
			     time a null command.  We cheat and push a
			     newline back if the list_terminator was a newline
			     to avoid the double-newline problem (one to
			     terminate this, one to terminate the command) */
			  (yyval.command) = gps.make_simple_command (x, nil);
			  (yyval.command).flags |= (yyvs.PeekN((1) - (2)).number);
			  /* XXX - let's cheat and push a newline back */
			  if ((yyvs.PeekN((2) - (2)).number) == '\n') {
			    gps.token_to_read = '\n';
			  }
			}
    break;

  case 163:
// #line 1235 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (4)).command), (yyvs.PeekN((4) - (4)).command), '|'); }
    break;

  case 164:
// #line 1237 "/Users/chet/src/bash/src/parse.y"
    {
			  /* Make cmd1 |& cmd2 equivalent to cmd1 2>&1 | cmd2 */
			  var tc Redirectable
			  var rd Redirectee
			  var sd Redirectee
			  var r *Redirect

			  if (yyvs.PeekN((1) - (4)).command).typ == cm_simple {
				tc = (yyvs.PeekN((1) - (4)).command).value.Simple
			  } else {
				tc = (yyvs.PeekN((1) - (4)).command);
			  }
			  sd.dest = 2;
			  rd.dest = 1;
			  r = makeRedirection (sd, r_duplicating_output, rd, 0);
			  tc.AddRedirect(r)

			  (yyval.command) = gps.command_connect ((yyvs.PeekN((1) - (4)).command), (yyvs.PeekN((4) - (4)).command), '|');
			}
    break;

  case 165:
// #line 1260 "/Users/chet/src/bash/src/parse.y"
    { (yyval.command) = (yyvs.PeekN((1) - (1)).command); }
    break;

  case 166:
// #line 1264 "/Users/chet/src/bash/src/parse.y"
    { (yyval.number) = CMD_TIME_PIPELINE; }
    break;

  case 167:
// #line 1266 "/Users/chet/src/bash/src/parse.y"
    { (yyval.number) = CMD_TIME_PIPELINE|CMD_TIME_POSIX; }
    break;


/* Line 1267 of yacc.c.  */
// #line 3383 "y.tab.c"
      default: break;
    }

  popStack(yylen)
  yylen = 0;

  yyvs.Push(yyval)


  /* Now `shift' the result of the reduction.  Determine what state
     that goes to, based on the state we popped back to and the rule
     number reduced by.  */

  yyn = yyr1[yyn];

  yystate = yypgoto[yyn - YYNTOKENS] + yyss.Peek()
  if (0 <= yystate && yystate <= YYLAST && yycheck[yystate] == yyss.Peek()) {
    yystate = yytable[yystate];
  } else {
    yystate = yydefgoto[yyn - YYNTOKENS];
  }
  yyparseState = yysetstate; continue;


/*------------------------------------.
| yyerrlab -- here on detecting error |
`------------------------------------*/
case yyerrlab:
  /* If not already recovering from an error, report this error.  */
  if (yyerrstatus == 0)    {
      gps.yynerrs++
      //yyerror (("syntax error"));
      fmt.Fprintf(os.Stderr, "syntax error\n")
    }



  if (yyerrstatus == 3)  {
      /* If just tried and failed to reuse look-ahead token after an
	 error, discard it.  */

      if (gps.yychar <= YYEOF)	{
	  /* Return failure if at end of input.  */
	  if (gps.yychar == YYEOF) {
	    yyparseState = yyabortlab; continue;
	  }
	} else {
	  gps.yychar = YYEMPTY;
	}
    }

  /* Else will try to reuse look-ahead token after shifting the error
     token.  */
  yyparseState = yyerrlab1; continue;


/*---------------------------------------------------.
| yyerrorlab -- error raised explicitly by YYERROR.  |
`---------------------------------------------------*/
case yyerrorlab:

  /* Do not reclaim the symbols of the rule which action triggered
     this YYERROR.  */
  popStack(yylen);
  yylen = 0;
  yystate = yyss.Peek();
  yyparseState = yyerrlab1; continue;


/*-------------------------------------------------------------.
| yyerrlab1 -- common code for both syntax error and YYERROR.  |
`-------------------------------------------------------------*/
case yyerrlab1:
  yyerrstatus = 3;	/* Each real token shifted decrements this.  */

  for {
      yyn = yypact[yystate];
      if (yyn != YYPACT_NINF) {
	  yyn += YYTERROR;
	  if (0 <= yyn && yyn <= YYLAST && yycheck[yyn] == YYTERROR) {
	      yyn = yytable[yyn];
	      if (0 < yyn) {
		break;
	      }
	    }
	}

      /* Pop the current state because it cannot handle the error token.  */
      if (yyss.IsEmpty()) {
	yyparseState = yyabortlab; continue;
      }


      popStack(1);
      yystate = yyss.Peek()
    }

  if (yyn == YYFINAL) {
    yyparseState = yyacceptlab; continue;
  }

  yyvs.Push(gps.yylval)


  yystate = yyn;
  yyparseState = yysetstate; continue;


/*-------------------------------------.
| yyacceptlab -- YYACCEPT comes here.  |
`-------------------------------------*/
case yyacceptlab:
  yyresult = 0;
  yyparseState = yyreturn; continue;

/*-----------------------------------.
| yyabortlab -- YYABORT comes here.  |
`-----------------------------------*/
case yyabortlab:
  yyresult = 1;
  yyparseState = yyreturn; continue;

case yyreturn:
 return;
}
}
  return;
}


// #line 1268 "/Users/chet/src/bash/src/parse.y"


/* Initial size to allocate for tokens, and the
   amount to grow them by. */
const TOKEN_DEFAULT_INITIAL_SIZE = 496
const TOKEN_DEFAULT_GROW_SIZE = 512


/* gps.yy_getc () returns the next available character from input or EOF.
   yy_ungetc (c) makes `c' the next character to read.
   init_yy_io (get, unget, typ, location) makes the function GET the
   installed function for getting the next character, makes UNGET the
   installed function for un-getting a character, sets the typ of stream
   (either string or file) from TYPE, and makes LOCATION point to where
   the input is coming from. */

///* Set all of the fields in BASH_INPUT to NULL.  Free bash_input.name if it
//   is non-null, avoiding a memory leak. */
//void
//initialize_bash_input ()
//{
//  bash_input.typ = st_none;
//  bash_input.name = nil;
//  bash_input.location.file = nil;
//  bash_input.location.string = nil;
//  bash_input.getter = nil;
//  bash_input.ungetter = nil;
//}
//
///* Set the contents of the current bash input stream from
//   GET, UNGET, TYPE, NAME, and LOCATION. */
//void
//init_yy_io (get, unget, typ, name, location)
//     sh_cget_func_t *get;
//     sh_cunget_func_t *unget;
//     enum stream_type typ;
//     const char *name;
//     INPUT_STREAM location;
//{
//  bash_input.typ = typ;
//  bash_input.name = name ? savestring (name) : nil;
//
//  /* XXX */
//  bash_input.location = location;
//  bash_input.getter = get;
//  bash_input.ungetter = unget;
//}
//
//char *
//yy_input_name ()
//{
//  return (bash_input.name ? bash_input.name : "stdin");
//}

//* Call this to get the next character of input. */
func (gps *ParserState) yy_getc () int {
  return gps.bashInput.Getc()
}
//
///* Call this to unget C.  That is, to make C the next character
//   to be read. */
//static int
//yy_ungetc (c)
//     int c;
//{
//  return (*(bash_input.ungetter)) (c);
//}
//
//void
//with_input_from_stdin ()
//{
//  with_input_from_stream (stdin, "stdin");
//}
//
///* **************************************************************** */
///*								    */
///*   Let input come from STRING.  STRING is zero terminated.	    */
///*								    */
///* **************************************************************** */
//
//static int
//yy_string_get ()
//{
//  char *string;
//  unsigned char c;
//
//  string = bash_input.location.string;
//
//  /* If the string doesn't exist, or is empty, EOF found. */
//  if (string && *string) {
//      c = *string++;
//      bash_input.location.string = string;
//      return (c);
//    }
//  else
//    return (EOF);
//}
//
//static int
//yy_string_unget (c)
//     int c;
//{
//  *(--bash_input.location.string) = c;
//  return (c);
//}
//
//void
//with_input_from_string (string, name)
//     char *string;
//     const char *name;
//{
//  INPUT_STREAM location;
//
//  location.string = string;
//  init_yy_io (yy_string_get, yy_string_unget, st_string, name, location);
//}
//
///* Count the number of characters we've consumed from bash_input.location.string
//   and read into gps.shell_input_line, but have not returned from shell_getc.
//   That is the true input location.  Rewind bash_input.location.string by
//   that number of characters, so it points to the last character actually
//   consumed by the parser. */
//static void
//rewind_input_string ()
//{
//  int xchars;
//
//  /* number of unconsumed characters in the input -- XXX need to take newlines
//     into account, e.g., $(...\n) */
//  xchars = gps.shell_input_line_len - gps.shell_input_line_index;
//  if (bash_input.location.string[-1] == '\n') {
//    xchars++;
//  }
//  /* XXX - how to reflect bash_input.location.string back to string passed to
//     parse_and_execute or xparse_dolparen?  xparse_dolparen needs to know how
//     far into the string we parsed.  parse_and_execute knows where bash_input.
//     location.string is, and how far from orig_string that is -- that's the
//     number of characters the command consumed. */
//
//  /* bash_input.location.string - xchars should be where we parsed to */
//  /* need to do more validation on xchars value for sanity -- test cases. */
//  bash_input.location.string -= xchars;
//}

func (gps *ParserState) rewind_input_string() {
  // TODO(krasin): implement this
  panic("rewind_input_string is not implemented")
}

//
///* **************************************************************** */
///*								    */
///*		     Let input come from STREAM.		    */
///*								    */
///* **************************************************************** */
//
///* These two functions used to test the value of the HAVE_RESTARTABLE_SYSCALLS
//   define, and just use getc/ungetc if it was defined, but since bash
//   installs its signal handlers without the SA_RESTART flag, some signals
//   (like SIGCHLD, SIGWINCH, etc.) received during a read(2) will not cause
//   the read to be restarted.  We need to restart it ourselves. */
//
//static int
//yy_stream_get ()
//{
//  int result;
//
//  result = EOF;
//  if (bash_input.location.file) {
//      result = getc_with_restart (bash_input.location.file);
//    }
//  return (result);
//}
//
//static int
//yy_stream_unget (c)
//     int c;
//{
//  return (ungetc_with_restart (c, bash_input.location.file));
//}
//
//void
//with_input_from_stream (stream, name)
//     FILE *stream;
//     const char *name;
//{
//  INPUT_STREAM location;
//
//  location.file = stream;
//  init_yy_io (yy_stream_get, yy_stream_unget, st_stream, name, location);
//}
//
//typedef struct stream_saver {
//  struct stream_saver *next;
//  BASH_INPUT bash_input;
//  int line;
//  BUFFERED_STREAM *bstream;
//} STREAM_SAVER;
//

//STREAM_SAVER *stream_list = nil;
//
//void
//push_stream (reset_lineno)
//     int reset_lineno;
//{
//  STREAM_SAVER *saver = (STREAM_SAVER *)xmalloc (sizeof (STREAM_SAVER));
//
//  xbcopy ((char *)&bash_input, (char *)&(saver.bash_input), sizeof (BASH_INPUT));
//
//  saver.bstream = nil;
//  /* If we have a buffered stream, clear out buffers[fd]. */
//  if (bash_input.typ == st_bstream && bash_input.location.buffered_fd >= 0) {
//    saver.bstream = set_buffered_stream (bash_input.location.buffered_fd,
//    					  nil);
//  }
//
//  saver.line = gps.line_number;
//  bash_input.name = nil;
//  saver.next = stream_list;
//  stream_list = saver;
//  gps.EOF_Reached = false;
//  if (reset_lineno) {
//    gps.line_number = 0;
//  }
//}
//
//void
//pop_stream ()
//{
//  if (!stream_list) {
//    gps.EOF_Reached = true;
//  } else {
//      STREAM_SAVER *saver = stream_list;
//
//      gps.EOF_Reached = false;
//      stream_list = stream_list.next;
//
//      init_yy_io (saver.bash_input.getter,
//		  saver.bash_input.ungetter,
//		  saver.bash_input.typ,
//		  saver.bash_input.name,
//		  saver.bash_input.location);
//
//      /* If we have a buffered stream, restore buffers[fd]. */
//      /* If the input file descriptor was changed while this was on the
//	 save stack, update the buffered fd to the new file descriptor and
//	 re-establish the buffer <-> bash_input fd correspondence. */
//      if (bash_input.typ == st_bstream && bash_input.location.buffered_fd >= 0) {
//	  if (bash_input_fd_changed) {
//	      bash_input_fd_changed = 0;
//	      if (default_buffered_input >= 0) {
//		  bash_input.location.buffered_fd = default_buffered_input;
//		  saver.bstream.b_fd = default_buffered_input;
//		  SET_CLOSE_ON_EXEC (default_buffered_input);
//		}
//	    }
//	  /* XXX could free buffered stream returned as result here. */
//	  set_buffered_stream (bash_input.location.buffered_fd, saver.bstream);
//	}
//
//      gps.line_number = saver.line;
//
//    }
//}
//
///* Return 1 if a stream of typ TYPE is saved on the stack. */
//int
//stream_on_stack (typ)
//     enum stream_type typ;
//{
//  STREAM_SAVER *s;
//
//  for (s = stream_list; s; s = s.next) {
//    if (s.bash_input.typ == typ) {
//      return 1;
//    }
//  }
//  return 0;
//}
//
///* Save the current token state and return it in a malloced array. */
//int *
//save_token_state ()
//{
//  int *ret;
//
//  ret = (int *)xmalloc (4 * sizeof (int));
//  ret[0] = gps.last_read_token;
//  ret[1] = gps.token_before_that;
//  ret[2] = gps.two_tokens_ago;
//  ret[3] = gps.current_token;
//  return ret;
//}
//
//void
//restore_token_state (ts)
//     int *ts;
//{
//  if (ts == 0) {
//    return;
//  }
//  gps.last_read_token = ts[0];
//  gps.token_before_that = ts[1];
//  gps.two_tokens_ago = ts[2];
//  gps.current_token = ts[3];
//}
//
///*
// * This is used to inhibit alias expansion and reserved word recognition
// * inside case statement pattern lists.  A `case statement pattern list' is:
// *
// *	everything between the `in' in a `case word in' and the next ')'
// *	or `esac'
// *	everything between a `;;' and the next `)' or `esac'
// */
//
//
//#define END_OF_ALIAS 0
//
///*
// * Pseudo-global variables used in implementing token-wise alias expansion.
// */
//
///*
// * Pushing and popping strings.  This works together with shell_getc to
// * implement alias expansion on a per-token basis.
// */
//
/*
 * Push the current gps.shell_input_line onto a stack of such lines and make S
 * the current input.  Used when expanding aliases.  EXPAND is used to set
 * the value of expand_next_token when the string is popped, so that the
 * word after the alias in the original line is handled correctly when the
 * alias expands to multiple words.  TOKEN is the token that was expanded
 * into S; it is saved and used to prevent infinite recursive expansion.
 */
func (gps *ParserState) push_string(s string, expand bool, ap *alias_t) {
  temp := new(StringSaver)

  temp.expand_alias = expand;
  temp.saved_line = gps.shell_input_line;
  temp.saved_line_size = gps.shell_input_line_size;
  temp.saved_line_index = gps.shell_input_line_index;
  temp.saved_line_terminator = gps.shell_input_line_terminator;
  temp.expander = ap;
  temp.next = gps.pushed_string_list;
  gps.pushed_string_list = temp;

  if ap != nil {
    ap.flags |= AL_BEINGEXPANDED;
  }

  gps.shell_input_line = stringToRunes(s)
  gps.shell_input_line_size = len(s);
  gps.shell_input_line_index = 0;
  gps.shell_input_line_terminator = 0
}

/*
 * Make the top of the pushed_string stack be the current shell input.
 * Only called when there is something on the stack.  Called from shell_getc
 * when it thinks it has consumed the string generated by an alias expansion
 * and needs to return to the original input line.
 */
func (gps *ParserState) pop_string () {
  gps.shell_input_line = gps.pushed_string_list.saved_line;
  gps.shell_input_line_index = gps.pushed_string_list.saved_line_index;
  gps.shell_input_line_size = gps.pushed_string_list.saved_line_size;
  gps.shell_input_line_terminator = gps.pushed_string_list.saved_line_terminator;

  if (gps.pushed_string_list.expand_alias) {
    gps.parser_state |= PST_ALEXPNEXT;
  } else {
    gps.parser_state &= ^PST_ALEXPNEXT;
  }

  t := gps.pushed_string_list;
  gps.pushed_string_list = gps.pushed_string_list.next;

  if t.expander != nil {
    t.expander.flags &= ^AL_BEINGEXPANDED;
  }

}

func (gps *ParserState) free_string_list () {
  for t := gps.pushed_string_list; t != nil; {
      t1 := t.next;
      if (t.expander != nil) {
	  t.expander.flags &= ^AL_BEINGEXPANDED;
      }
      t = t1;
  }
  gps.pushed_string_list = nil;
}

//void
//free_pushed_string_input ()
//{
//  gps.free_string_list ();
//}
//
///* Return a line of text, taken from wherever yylex () reads input.
//   If there is no more input, then we return NULL.  If REMOVE_QUOTED_NEWLINE
//   is non-zero, we remove unquoted \<newline> pairs.  This is used by
//   read_secondary_line to read here documents. */
//static char *
//read_a_line (remove_quoted_newline)
//     int remove_quoted_newline;
//{
//  static char *line_buffer = nil;
//  static int buffer_size = 0;
//  int indx, c, peekc, pass_next;
//
//  pass_next = indx = 0;
//  while (1)
//    {
//      /* Allow immediate exit if interrupted during input. */
//      QUIT;
//
//      c = gps.yy_getc ();
//
//      /* Ignore null bytes in input. */
//      if (c == 0) {
//	  continue;
//	}
//
//      /* If there is no more input, then we return NULL. */
//      if (c == EOF) {
//	  if (indx == 0) {
//	    return (nil);
//        }
//	  c = '\n';
//	}
//
//      /* `+2' in case the final character in the buffer is a newline. */
//      RESIZE_MALLOCED_BUFFER (line_buffer, indx, 2, buffer_size, 128);
//
//      /* IF REMOVE_QUOTED_NEWLINES is non-zero, we are reading a
//	 here document with an unquoted delimiter.  In this case,
//	 the line will be expanded as if it were in double quotes.
//	 We allow a backslash to escape the next character, but we
//	 need to treat the backslash specially only if a backslash
//	 quoting a backslash-newline pair appears in the line. */
//      if (pass_next) {
//	  line_buffer[indx++] = c;
//	  pass_next = 0;
//	} else {
//        if (c == '\\' && remove_quoted_newline) {
//	    peekc = gps.yy_getc ();
//	    if (peekc == '\n') {
//	        gps.line_number++;
//	        continue;	/* Make the unquoted \<newline> pair disappear. */
//	    } else {
//	      yy_ungetc (peekc);
//	      pass_next = 1;
//	      line_buffer[indx++] = c;		/* Preserve the backslash. */
//	    }
//	  } else {
//	    line_buffer[indx++] = c;
//        }
//      }
//
//      if (c == '\n') {
//	  line_buffer[indx] = '\0';
//	  return (line_buffer);
//	}
//    }
//}
//
///* Return a line as in read_a_line ().
//   This is used to read the lines of a here
//   document.  REMOVE_QUOTED_NEWLINE is non-zero if we should remove
//   newlines quoted with backslashes while reading the line.  It is
//   non-zero unless the delimiter of the here document was quoted. */
//char *
//read_secondary_line (remove_quoted_newline)
//     int remove_quoted_newline;
//{
//  char *ret;
//  ret = read_a_line (remove_quoted_newline);
//  return ret;
//}
//
///* **************************************************************** */
///*								    */
///*				YYLEX ()			    */
///*								    */
///* **************************************************************** */
//
/* Reserved words.  These are only recognized as the first word of a
   command. */
var word_token_alist = map[string] int {
  "if": IF,
   "then": THEN,
   "else": ELSE,
   "elif": ELIF,
   "fi": FI,
   "case": CASE,
   "esac": ESAC,
   "for": FOR,
   "select": SELECT,
   "while": WHILE,
   "until": UNTIL,
   "do": DO,
   "done": DONE,
   "in": IN,
   "function": FUNCTION,
   "time": TIME,
   "{": '{',
   "}": ',',
   "!": BANG,
   "[[": COND_START,
   "]]": COND_END,
   "": 0,
}

///* other tokens that can be returned by read_token() */
//STRING_INT_ALIST other_token_alist[] = {
//  /* Multiple-character tokens with special values */
//  { "-p", TIMEOPT },
//  { "&&", AND_AND },
//  { "||", OR_OR },
//  { ">>", GREATER_GREATER },
//  { "<<", LESS_LESS },
//  { "<&", LESS_AND },
//  { ">&", GREATER_AND },
//  { ";;", SEMI_SEMI },
//  { ";&", SEMI_AND },
//  { ";;&", SEMI_SEMI_AND },
//  { "<<-", LESS_LESS_MINUS },
//  { "<<<", LESS_LESS_LESS },
//  { "&>", AND_GREATER },
//  { "&>>", AND_GREATER_GREATER },
//  { "<>", LESS_GREATER },
//  { ">|", GREATER_BAR },
//  { "|&", BAR_AND },
//  { "EOF", yacc_EOF },
//  /* Tokens whose value is the character itself */
//  { ">", '>' },
//  { "<", '<' },
//  { "-", '-' },
//  { "{", '{' },
//  { "}", '}' },
//  { ";", ';' },
//  { "(", '(' },
//  { ")", ')' },
//  { "|", '|' },
//  { "&", '&' },
//  { "newline", '\n' },
//  { (char *)NULL, 0}
//};
//
///* others not listed here:
//	WORD			look at yylval.word
//	ASSIGNMENT_WORD		look at yylval.word
//	NUMBER			look at yylval.number
//	ARITH_CMD		look at yylval.word_list
//	ARITH_FOR_EXPRS		look at yylval.word_list
//	COND_CMD		look at yylval.command
//*/

func (gps *ParserState) current_delimiter() int {
  if gps.dstack.IsEmpty() {
    return 0
  }
  return gps.dstack.Peek()
}

func (gps *ParserState) push_delimiter(character int) {
  gps.dstack.Push(character)
}

func (gps *ParserState) pop_delimiter() {
  gps.dstack.Pop()
}

/* Return the next shell input character.  This always reads characters
   from gps.shell_input_line; when that line is exhausted, it is time to
   read the next line.  This is called by read_token when the shell is
   processing normal command input. */
func (gps *ParserState) shell_getc (remove_quoted_newline bool) int {
  var i int
  var c int
  var uc int

  if (gps.eol_ungetc_lookahead != 0) {
      c = gps.eol_ungetc_lookahead;
      gps.eol_ungetc_lookahead = 0;
      return c
  }

  /* If shell_input_line[gps.shell_input_line_index] == 0, but there is
     something on the pushed list of strings, then we don't want to go
     off and get another line.  We let the code down below handle it. */

  if gps.shell_input_line ==nil || ((gps.shell_input_line[gps.shell_input_line_index] == 0) &&
			    (gps.pushed_string_list == nil)) {
      gps.line_number++;

    restart_read:

      /* Allow immediate exit if interrupted during input. */

      i = 0;
      gps.shell_input_line_terminator = 0;

      // TODO(krasin): enable it
      //cleanup_dead_jobs ();

      // TODO(krasin): decide what to do with bash_input
      //if (bash_input.typ == st_stream) {
//	clearerr (stdin);
//      }

      for {
	  c = gps.yy_getc ();

	  if c == 0 {
	      continue;
	  }

	  gps.shell_input_line = resizeBuffer(gps.shell_input_line, i, 2, gps.shell_input_line_size, 256)

	  if c == EOF {
//  TODO(krasin): decide what to do with bash_input
//	      if bash_input.typ == st_stream {
//		clearerr (stdin);
//              }

	      if i == 0 {
		gps.shell_input_line_terminator = EOF;
              }

	      gps.shell_input_line[i] = 0
	      break;
	  }

	  gps.shell_input_line[i] = c;
	  i++

	  if (c == '\n') {
	      i--
	      gps.shell_input_line[i] = 0
	      gps.current_command_line_count++;
	      break;
	  }
      }

      gps.shell_input_line_index = 0;
      gps.shell_input_line_len = i;		/* == strlen (gps.shell_input_line) */

      if gps.shell_input_line != nil {
	  /* Lines that signify the end of the shell's input should not be
	     echoed. */
	  if (gps.echo_input_at_read && (gps.shell_input_line[0] != 0 ||
				     gps.shell_input_line_terminator != EOF)) {
	    fmt.Fprintf(os.Stderr, "%s\n", runesToString(gps.shell_input_line));
          }
	} else {
	  gps.shell_input_line_size = 0;
	  goto restart_read;
	}

      /* Add the newline to the end of this string, iff the string does
	 not already end in an EOF character.  */
      if (gps.shell_input_line_terminator != EOF) {
	  if (gps.shell_input_line_len + 3 > gps.shell_input_line_size) {
            gps.shell_input_line_size += 2
	    gps.shell_input_line = enlargeBuffer(gps.shell_input_line, 1+gps.shell_input_line_size)
	  }

	  gps.shell_input_line[gps.shell_input_line_len] = '\n';
	  gps.shell_input_line[gps.shell_input_line_len + 1] = 0
      }
    }

  uc = gps.shell_input_line[gps.shell_input_line_index];

  if uc != 0 {
    gps.shell_input_line_index++;
  }

  /* If UC is NULL, we have reached the end of the current input string.  If
     gps.pushed_string_list is non-empty, it's time to pop to the previous string
     because we have fully consumed the result of the last alias expansion.
     Do it transparently; just return the next character of the string popped
     to. */
pop_alias:
  if (uc == 0 && (gps.pushed_string_list != nil)) {
      gps.pop_string ();
      uc = gps.shell_input_line[gps.shell_input_line_index];
      if uc != 0 {
	  gps.shell_input_line_index++;
      }
  }

  if (uc == '\\' && remove_quoted_newline && gps.shell_input_line[gps.shell_input_line_index] == '\n') {
	gps.line_number++;
	/* XXX - what do we do here if we're expanding an alias whose definition
	   ends with a newline?  Recall that we inhibit the appending of a
	   space in mk_alexpansion() if newline is the last character. */
  
	goto restart_read;
    }

  if (uc == 0 && gps.shell_input_line_terminator == EOF) {
    if gps.shell_input_line_index != 0 {
	return '\n'
    } else {
        return EOF
    }
  }

  return (uc);
}

/* Put C back into the input for the shell. */
func (gps *ParserState) shell_ungetc(c int) {
  if (gps.shell_input_line != nil && gps.shell_input_line_index > 0) {
    gps.shell_input_line_index--
    gps.shell_input_line[gps.shell_input_line_index] = c;
  } else {
    gps.eol_ungetc_lookahead = c;
  }
}

/* Discard input until CHARACTER is seen, then push that character back
   onto the input stream. */
func (gps *ParserState) discard_until(character int) {
  var c int

  for {
    c = gps.shell_getc(false)
    if c == EOF || c == character {
	break
    }
  }

  if (c != EOF) {
    gps.shell_ungetc (c);
  }
}

//void
//execute_variable_command (command, vname)
//     char *command, *vname;
//{
//  char *last_lastarg;
//  sh_parser_state_t ps;
//
//  save_parser_state (&ps);
//  last_lastarg = get_string_value ("_");
//  if (last_lastarg) {
//    last_lastarg = savestring (last_lastarg);
//  }
//
//  parse_and_execute (savestring (command), vname, SEVAL_NONINT|SEVAL_NOHIST);
//
//  restore_parser_state (&ps);
//  bind_variable ("_", last_lastarg, 0);
//
//  if (gps.token_to_read == '\n') {	/* gps.reset_parser was called */
//    gps.token_to_read = 0;
//  }
//}
//
/* Command to read_token () explaining what we want it to do. */
const READ = 0
const RESET = 1

/* Function for yyparse to call.  yylex keeps track of
   the last two tokens read, and calls read_token.  */
func (gps *ParserState) yylex() int {
  gps.two_tokens_ago = gps.token_before_that;
  gps.token_before_that = gps.last_read_token;
  gps.last_read_token = gps.current_token;
  gps.current_token = gps.read_token (READ);

  if ((gps.parser_state & PST_EOFTOKEN != 0) && gps.current_token == gps.shell_eof_token) {
      gps.current_token = yacc_EOF;
// TODO(krasin): decide what to do with bash_input
      //if (bash_input.typ == st_string) {
//	gps.rewind_input_string ();
//	}
  }
  gps.parser_state &= ^PST_EOFTOKEN;

  return (gps.current_token);
}

func (gps *ParserState) gather_here_documents() {
  // TODO(krasin): implement this
  panic("gather_here_documents: not implemented")
//  r := 0;
//  for gps.need_here_doc != 0 {
//      gps.parser_state |= PST_HEREDOC;
//      redir := gps.redir_stack[r]
//      r++
//      gps.make_here_document(redir, gps.line_number)
//      gps.parser_state &= ^PST_HEREDOC;
//      gps.need_here_doc--;
//    }
}


func (gps *ParserState) command_token_position(token int) bool {
  return ((token) == ASSIGNMENT_WORD) || (gps.parser_state&PST_REDIRLIST != 0) ||
   ((token) != SEMI_SEMI && (token) != SEMI_AND && (token) != SEMI_SEMI_AND && reserved_word_acceptable(token))
}

func (gps *ParserState) assignment_acceptable(token int) bool {
  return gps.command_token_position(token) && (gps.parser_state & PST_CASEPAT) == 0
}

/* Check to see if TOKEN is a reserved word and return the token
   value if it is. */
func (wts *wordTokenizerState) CHECK_FOR_RESERVED_WORD(word string) int {
  if wts.dollar_present || wts.quoted || !reserved_word_acceptable(wts.gps.last_read_token) {
    return NO_TOKEN
  }
  tok, ok := word_token_alist[word]
  if !ok {
    return NO_TOKEN
  }
  if (wts.gps.parser_state & PST_CASEPAT != 0) && (tok != ESAC) {
    return NO_TOKEN
  }
  if tok == TIME && !wts.gps.time_command_acceptable () {
    return NO_TOKEN
  }
  switch tok {
  case ESAC:  wts.gps.parser_state &= ^(PST_CASEPAT|PST_CASESTMT)
  case CASE:  wts.gps.parser_state |= PST_CASESTMT
  case COND_END:  wts.gps.parser_state &= ^(PST_CONDCMD|PST_CONDEXPR)
  case COND_START: wts.gps.parser_state |= PST_CONDCMD
  case '{': wts.gps.open_brace_count++;
  case '}': if wts.gps.open_brace_count > 0 { wts.gps.open_brace_count--; }
  }
  return tok
}

//    /* OK, we have a token.  Let's try to alias expand it, if (and only if)
//       it's eligible.
//
//       It is eligible for expansion if EXPAND_ALIASES is set, and
//       the token is unquoted and the last token read was a command
//       separator (or expand_next_token is set), and we are currently
//       processing an alias (gps.pushed_string_list is non-empty) and this
//       token is not the same as the current or any previously
//       processed alias.
//
//       Special cases that disqualify:
//	 In a pattern list in a case statement (gps.parser_state & PST_CASEPAT). */
//
//static char *
//mk_alexpansion (s)
//     char *s;
//{
//  int l;
//  char *r;
//
//  l = strlen (s);
//  r = xmalloc (l + 2);
//  strcpy (r, s);
//  if (r[l -1] != ' ') {
//    r[l++] = ' ';
//  }
//  r[l] = '\0';
//  return r;
//}
//

func alias_expand_token(tokstr string) int {
  // TODO(krasin): implement this
  panic("alias_expand_token: not implemented")
}

//static int
//alias_expand_token (tokstr)
//     char *tokstr;
//{
//  char *expanded;
//  alias_t *ap;
//
//  if (((gps.parser_state & PST_ALEXPNEXT) || gps.command_token_position (gps.last_read_token)) &&
//	(gps.parser_state & PST_CASEPAT) == 0) {
//      ap = find_alias (tokstr);
//
//      /* Currently expanding this token. */
//      if (ap && (ap.flags & AL_BEINGEXPANDED)) {
//	  return (NO_EXPANSION);
//      }
//
//      /* mk_alexpansion puts an extra space on the end of the alias expansion,
//         so the lookahead by the parser works right.  If this gets changed,
//         make sure the code in shell_getc that deals with reaching the end of
//         an expanded alias is changed with it. */
//      expanded = ap ? mk_alexpansion (ap.value) : nil;
//
//      if (expanded) {
//	  gps.push_string (expanded, ap.flags & AL_EXPANDNEXT, ap);
//	  return (RE_READ_TOKEN);
//	} else {
//	/* This is an eligible token that does not have an expansion. */
//	return (NO_EXPANSION);
//      }
//    }
//  return (NO_EXPANSION);
//}

func (gps *ParserState) time_command_acceptable() bool {
  switch gps.last_read_token {
    case 0: fallthrough
    case ';': fallthrough
    case '\n': fallthrough
    case AND_AND: fallthrough
    case OR_OR: fallthrough
    case '&': fallthrough
    case DO: fallthrough
    case THEN: fallthrough
    case ELSE: fallthrough
    case '{': fallthrough /* } */
    case '(':  /* ) */
      return true
  }
  return false
}

/* Handle special cases of token recognition:
	IN is recognized if the last token was WORD and the token
	before that was FOR or CASE or SELECT.

	DO is recognized if the last token was WORD and the token
	before that was FOR or SELECT.

	ESAC is recognized if the last token caused `esacs_needed_count'
	to be set

	`{' is recognized if the last token as WORD and the token
	before that was FUNCTION, or if we just parsed an arithmetic
	`for' command.

	`}' is recognized if there is an unclosed `{' present.

	`-p' is returned as TIMEOPT if the last read token was TIME.

	']]' is returned as COND_END if the parser is currently parsing
	a conditional expression ((gps.parser_state & PST_CONDEXPR) != 0)

	`time' is returned as TIME if and only if it is immediately
	preceded by one of `;', `\n', `||', `&&', or `&'.
*/

func (gps *ParserState) special_case_tokens(tokstr string) int {
  if (gps.last_read_token == WORD) &&
      ((gps.token_before_that == FOR) || (gps.token_before_that == CASE) || (gps.token_before_that == SELECT)) &&
      (tokstr == "in") {
      if (gps.token_before_that == CASE) {
	    gps.parser_state |= PST_CASEPAT;
	    gps.esacs_needed_count++;
	  }
      return (IN);
  }

  if gps.last_read_token == WORD &&
      (gps.token_before_that == FOR || gps.token_before_that == SELECT) &&
      (tokstr == "do") {
    return (DO);
  }

  /* Ditto for ESAC in the CASE case.
     Specifically, this handles "case word in esac", which is a legal
     construct, certainly because someone will pass an empty arg to the
     case construct, and we don't want it to barf.  Of course, we should
     insist that the case construct has at least one pattern in it, but
     the designers disagree. */
  if (gps.esacs_needed_count > 0) {
      gps.esacs_needed_count--;
      if tokstr == "esac" {
	    gps.parser_state &= ^PST_CASEPAT;
	    return (ESAC);
	  }
  }

  /* The start of a shell function definition. */
  if (gps.parser_state & PST_ALLOWOPNBRC) != 0 {
      gps.parser_state &= ^PST_ALLOWOPNBRC;
      if tokstr == "{" {		/* } */
	    gps.open_brace_count++;
	    gps.function_bstart = gps.line_number;
	    return ('{');					/* } */
	  }
  }

  /* We allow a `do' after a for ((...)) without an intervening
     list_terminator */
  if gps.last_read_token == ARITH_FOR_EXPRS && tokstr == "do" {
    return (DO);
  }
  if gps.last_read_token == ARITH_FOR_EXPRS && tokstr == "{" { /* } */
      gps.open_brace_count++;
      return ('{');			/* } */
  }

  if gps.open_brace_count > 0 && reserved_word_acceptable (gps.last_read_token) && tokstr == "}" {
      gps.open_brace_count--;		/* { */
      return ('}');
  }

  /* Handle -p after `time'. */
  if gps.last_read_token == TIME && tokstr == "-p" {
    return (TIMEOPT);
  }

  if (gps.parser_state & PST_CONDEXPR) != 0 && tokstr == "]]" {
    return (COND_END);
  }

  return (-1);
}
//
/* Called from shell.c when Control-C is typed at top level.  Or
   by the error rule at top level. */
func (gps *ParserState) reset_parser() {
  gps.dstack = newIntStack()
  gps.open_brace_count = 0;

  /* Reset to global value of extended glob */
  if (gps.parser_state & PST_EXTPAT != 0) {
    gps.extended_glob = gps.global_extglob;
  }

  gps.parser_state = 0;

  if (gps.pushed_string_list != nil) {
    gps.free_string_list ();
  }

  if (gps.shell_input_line != nil) {
      gps.shell_input_line = nil;
      gps.shell_input_line_size = 0
      gps.shell_input_line_index = 0;
  }

  gps.word_desc_to_read = nil;

  gps.current_token = '\n';		/* XXX */
  gps.last_read_token = '\n';
  gps.token_to_read = '\n';
}

/* Read the next token.  Command can be READ (normal operation) or
   RESET (to normalize state). */
func (gps *ParserState) read_token (command int) (result int) {
  var character int /* Current character. */
  var peek_char int /* Temporary look-ahead character. */

  if (command == RESET) {
      gps.reset_parser ();
      return ('\n');
  }

  if (gps.token_to_read != 0) {
      result = gps.token_to_read;
      if (gps.token_to_read == WORD || gps.token_to_read == ASSIGNMENT_WORD) {
	  gps.yylval.word = gps.word_desc_to_read;
	  gps.word_desc_to_read = nil;
      }
      gps.token_to_read = 0;
      return (result);
  }

  if ((gps.parser_state & (PST_CONDCMD|PST_CONDEXPR)) == PST_CONDCMD) {
      gps.cond_lineno = gps.line_number;
      gps.parser_state |= PST_CONDEXPR;
      gps.yylval.command = gps.parse_cond_command ();
      if (gps.cond_token != COND_END) {
	  gps.cond_error()
	  return (-1);
      }
      gps.token_to_read = COND_END;
      gps.parser_state &= ^(PST_CONDEXPR|PST_CONDCMD);
      return (COND_CMD);
  }

  /* This is a place to jump back to once we have successfully expanded a
     token with an alias and pushed the string with gps.push_string () */
 re_read_token:

  /* Read a single word from input.  Start by skipping blanks. */
  for {
    character = gps.shell_getc (true)
    if character == EOF || !shellblank (character) {
      break
    }
  }

  if (character == EOF) {
      gps.EOF_Reached = true;
      return (yacc_EOF);
  }

  if (character == '#') {
      /* A comment.  Discard until EOL or EOF, and then return a newline. */
      gps.discard_until ('\n');
      gps.shell_getc (false);
      character = '\n';	/* this will take the next if statement and return. */
  }

  if (character == '\n') {
      /* If we're about to return an unquoted newline, we can go and collect
	 the text of any pending here document. */
      if (gps.need_here_doc != 0) {
      	gps.gather_here_documents ();
      }

      gps.parser_state &= ^PST_ALEXPNEXT;

      gps.parser_state &= ^PST_ASSIGNOK;

      return (character);
  }

  if (gps.parser_state & PST_REGEXP != 0) {
    goto tokword;
  }

  /* Shell meta-characters. */
  if (shellmeta (character) && ((gps.parser_state & PST_DBLPAREN) == 0)) {
      /* Turn off alias tokenization iff this character sequence would
	 not leave us ready to read a command. */
      if (character == '<' || character == '>') {
	  gps.parser_state &= ^PST_ALEXPNEXT;
      }

      gps.parser_state &= ^PST_ASSIGNOK;

      peek_char = gps.shell_getc (true);
      switch {
      case character == peek_char:
	  switch (character) {
	    case '<':
	      /* If '<' then we could be at "<<" or at "<<-".  We have to
		 look ahead one more character. */
	      peek_char = gps.shell_getc (true);
              switch {
	      case (peek_char == '-'):
		return (LESS_LESS_MINUS);
	      case (peek_char == '<'):
		return (LESS_LESS_LESS);
	      default:
		  gps.shell_ungetc (peek_char);
		  return (LESS_LESS);
              }

	    case '>':
	      return (GREATER_GREATER);

	    case ';':
	      gps.parser_state |= PST_CASEPAT;
	      gps.parser_state &= ^PST_ALEXPNEXT;

	      peek_char = gps.shell_getc (true);
	      if (peek_char == '&') {
		return (SEMI_SEMI_AND);
	      } else {
		  gps.shell_ungetc (peek_char);
		  return (SEMI_SEMI);
              }

	    case '&':
	      return (AND_AND);

	    case '|':
	      return (OR_OR);

	    case '(':		/* ) */
	      result = gps.parse_dparen (character);
	      if (result == -2) {
	        break;
	      } else {
	        return result;
              }
	    }
      case (character == '<' && peek_char == '&'):
	return (LESS_AND);
      case (character == '>' && peek_char == '&'):
	return (GREATER_AND);
      case (character == '<' && peek_char == '>'):
	return (LESS_GREATER);
      case (character == '>' && peek_char == '|'):
	return (GREATER_BAR);
      case (character == '&' && peek_char == '>'):
	{
	  peek_char = gps.shell_getc (true);
	  if (peek_char == '>') {
	    return (AND_GREATER_GREATER);
	  } else {
	      gps.shell_ungetc (peek_char);
	      return (AND_GREATER);
	  }
	}
      case (character == '|' && peek_char == '&'):
	return (BAR_AND);
      case (character == ';' && peek_char == '&'):
	{
	  gps.parser_state |= PST_CASEPAT;
	  gps.parser_state &= ^PST_ALEXPNEXT;
	  return (SEMI_AND);
	}
      }

      gps.shell_ungetc (peek_char);

      /* If we look like we are reading the start of a function
	 definition, then let the reader know about it so that
	 we will do the right thing with `{'. */
    if (character == ')' && gps.last_read_token == '(' && gps.token_before_that == WORD) {
	  gps.parser_state |= PST_ALLOWOPNBRC;
	  gps.parser_state &= ^PST_ALEXPNEXT;
	  gps.function_dstart = gps.line_number;
	}

      /* case pattern lists may be preceded by an optional left paren.  If
	 we're not trying to parse a case pattern list, the left paren
	 indicates a subshell. */
      switch {
      case (character == '(' && (gps.parser_state & PST_CASEPAT) == 0): /* ) */
	gps.parser_state |= PST_SUBSHELL;
      /*(*/
      case ((gps.parser_state & PST_CASEPAT != 0) && character == ')'):
	gps.parser_state &= ^PST_CASEPAT;
      /*(*/
      case ((gps.parser_state & PST_SUBSHELL != 0) && character == ')'):
	gps.parser_state &= ^PST_SUBSHELL;
      }

      return (character);
    }

  /* Hack <&- (close stdin) case.  Also <&N- (dup and close). */
  if (character == '-' && (gps.last_read_token == LESS_AND || gps.last_read_token == GREATER_AND)) {
    return (character);
  }

tokword:
  /* Okay, if we got this far, we have to read a word.  Read one,
     and then check it against the known ones. */
  result = gps.read_token_word (character);
  if (result == RE_READ_TOKEN) {
    goto re_read_token;
  }
  return result;
}

/*
 * Match a $(...) or other grouping construct.  This has to handle embedded
 * quoted strings ('', ``, "") and nested constructs.  It also must handle
 * reprompting the user, if necessary, after reading a newline, and returning
 * correct error values if it reads EOF.
 */
const P_FIRSTCLOSE = 0x01
const P_ALLOWESC = 0x02
const P_DQUOTE = 0x04
const P_COMMAND = 0x08 /* parsing a command, so look for comments */
const P_BACKQUOTE = 0x10 /* parsing a backquoted command substitution */
const P_ARRAYSUB = 0x20 /* parsing a [...] array subscript for assignment */

/* Lexical state while parsing a grouping construct or $(...). */
const LEX_WASDOL = 0x001
const LEX_CKCOMMENT = 0x002
const LEX_INCOMMENT = 0x004
const LEX_PASSNEXT = 0x008
const LEX_RESWDOK = 0x010
const LEX_CKCASE = 0x020
const LEX_INCASE = 0x040
const LEX_INHEREDOC = 0x080
const LEX_HEREDELIM = 0x100 /* reading here-doc delimiter */
const LEX_STRIPDOC = 0x200 /* <<- strip tabs from here doc delim */
const LEX_INWORD = 0x400

//#define COMSUB_META(ch)		((ch) == ';' || (ch) == '&' || (ch) == '|')
//
//#define CHECK_NESTRET_ERROR() \
//  do { \
//    if (nestret == &matched_pair_error) \
//      { \
//	return &matched_pair_error; \
//      } \
//  } while (0)
//
//#define APPEND_NESTRET() \
//  do { \
//    if (nestlen) \
//      { \
//	RESIZE_MALLOCED_BUFFER (ret, retind, nestlen, retsize, 64); \
//	strcpy (ret + retind, nestret); \
//	retind += nestlen; \
//      } \
//  } while (0)
//
//static char matched_pair_error;
//

func parse_matched_pair(qc int, open int, cloze int, flags int) (ret *StringBuilder, err os.Error) {
	// TODO(krasin): implement this
	panic("parse_matched_pair: not implemented")
}

//static char *
//parse_matched_pair (qc, open, close, lenp, flags)
//     int qc;	/* `"' if this construct is within double quotes */
//     int open, close;
//     int *lenp, flags;
//{
//  int count, ch, tflags;
//  int nestlen, ttranslen, start_lineno;
//  char *ret, *nestret, *ttrans;
//  int retind, retsize, rflags;
//
///*itrace("parse_matched_pair[%d]: open = %c close = %c flags = %d", gps.line_number, open, close, flags);*/
//  count = 1;
//  tflags = 0;
//
//  if ((flags & P_COMMAND) && qc != '`' && qc != '\'' && qc != '"' && (flags & P_DQUOTE) == 0) {
//    tflags |= LEX_CKCOMMENT;
//  }
//
//  /* RFLAGS is the set of flags we want to pass to recursive calls. */
//  rflags = (qc == '"') ? P_DQUOTE : (flags & P_DQUOTE);
//
//  ret = (char *)xmalloc (retsize = 64);
//  retind = 0;
//
//  start_lineno = gps.line_number;
//  while (count)
//    {
//      ch = gps.shell_getc (qc != '\'' && (tflags & LEX_PASSNEXT) == 0);
//
//      if (ch == EOF) {
//	  gps.parser_error (start_lineno, _("unexpected EOF while looking for matching `%c'"), close);
//	  gps.EOF_Reached = true;	/* XXX */
//	  return (&matched_pair_error);
//	}
//
//      /* Don't bother counting parens or doing anything else if in a comment
//	 or part of a case statement */
//      if (tflags & LEX_INCOMMENT) {
//	  /* Add this character. */
//	  RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//	  ret[retind++] = ch;
//
//	  if (ch == '\n') {
//	    tflags &= ^LEX_INCOMMENT;
//        }
//
//	  continue;
//	}
//
//      /* Not exactly right yet, should handle shell metacharacters, too.  If
//	 any changes are made to this test, make analogous changes to subst.c:
//	 extract_delimited_string(). */
//      else if ((tflags & LEX_CKCOMMENT) && (tflags & LEX_INCOMMENT) == 0 && ch == '#' && (retind == 0 || ret[retind-1] == '\n' || shellblank (ret[retind - 1])))
//	tflags |= LEX_INCOMMENT;
//
//      if (tflags & LEX_PASSNEXT)		/* last char was backslash */
//	{
//	  tflags &= ^LEX_PASSNEXT;
//	  if (qc != '\'' && ch == '\n')	/* double-quoted \<newline> disappears. */
//	    {
//	      if (retind > 0)
//		retind--;	/* swallow previously-added backslash */
//	      continue;
//	    }
//
//	  RESIZE_MALLOCED_BUFFER (ret, retind, 2, retsize, 64);
//	  if (ch == CTLESC || ch == CTLNUL)
//	    ret[retind++] = CTLESC;
//	  ret[retind++] = ch;
//	  continue;
//	}
//      /* If we're reparsing the input (e.g., from parse_string_to_word_list),
//	 we've already prepended CTLESC to single-quoted results of $'...'.
//	 We may want to do this for other CTLESC-quoted characters in
//	 reparse, too. */
//      else if ((gps.parser_state & PST_REPARSE) && open == '\'' && (ch == CTLESC || ch == CTLNUL))
//	{
//	  RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//	  ret[retind++] = ch;
//	  continue;
//	}
//      else if (ch == CTLESC || ch == CTLNUL)	/* special shell escapes */
//	{
//	  RESIZE_MALLOCED_BUFFER (ret, retind, 2, retsize, 64);
//	  ret[retind++] = CTLESC;
//	  ret[retind++] = ch;
//	  continue;
//	}
//      else if (ch == close)		/* ending delimiter */
//	count--;
//      /* handle nested ${...} specially. */
//      else if (open != close && (tflags & LEX_WASDOL) && open == '{' && ch == open) /* } */
//	count++;
//      else if (((flags & P_FIRSTCLOSE) == 0) && ch == open)	/* nested begin */
//	count++;
//
//      /* Add this character. */
//      RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//      ret[retind++] = ch;
//
//      /* If we just read the ending character, don't bother continuing. */
//      if (count == 0)
//	break;
//
//      if (open == '\'')			/* '' inside grouping construct */
//	{
//	  if ((flags & P_ALLOWESC) && ch == '\\')
//	    tflags |= LEX_PASSNEXT;
//	  continue;
//	}
//
//      if (ch == '\\')			/* backslashes */
//	tflags |= LEX_PASSNEXT;
//
//#if 0
//      /* The big hammer.  Single quotes aren't special in double quotes.  The
//         problem is that Posix says the single quotes are semi-special:
//         within a double-quoted ${...} construct "an even number of
//         unescaped double-quotes or single-quotes, if any, shall occur." */
//      if (open == '{' && (flags & P_DQUOTE) && ch == '\'')	/* } */
//	continue;
//#endif
//
//      /* Could also check open == '`' if we want to parse grouping constructs
//	 inside old-style command substitution. */
//      if (open != close)		/* a grouping construct */
//	{
//	  if (shellquote (ch))
//	    {
//	      /* '', ``, or "" inside $(...) or other grouping construct. */
//	      push_delimiter (dstack, ch);
//	      if ((tflags & LEX_WASDOL) && ch == '\'')	/* $'...' inside group */
//		nestret = parse_matched_pair (ch, ch, ch, &nestlen, P_ALLOWESC|rflags);
//	      else
//		nestret = parse_matched_pair (ch, ch, ch, &nestlen, rflags);
//	      pop_delimiter (dstack);
//	      CHECK_NESTRET_ERROR ();
//
//	      if ((tflags & LEX_WASDOL) && ch == '\'' && (extended_quote || (rflags & P_DQUOTE) == 0))
//		{
//		  /* Translate $'...' here. */
//		  ttrans = ansiexpand (nestret, 0, nestlen - 1, &ttranslen);
//
//		  if ((rflags & P_DQUOTE) == 0)
//		    {
//		      nestret = sh_single_quote (ttrans);
//		      nestlen = strlen (nestret);
//		    }
//		  else
//		    {
//		      nestret = ttrans;
//		      nestlen = ttranslen;
//		    }
//		  retind -= 2;		/* back up before the $' */
//		}
//	      else if ((tflags & LEX_WASDOL) && ch == '"' && (extended_quote || (rflags & P_DQUOTE) == 0))
//		{
//		  /* Locale expand $"..." here. */
//		  ttrans = localeexpand (nestret, 0, nestlen - 1, start_lineno, &ttranslen);
//
//		  nestret = sh_mkdoublequoted (ttrans, ttranslen, 0);
//		  nestlen = ttranslen + 2;
//		  retind -= 2;		/* back up before the $" */
//		}
//
//	      APPEND_NESTRET ();
//	    }
//	  else if ((flags & P_ARRAYSUB) && (tflags & LEX_WASDOL) && (ch == '(' || ch == '{' || ch == '['))	/* ) } ] */
//	    goto parse_dollar_word;
//	}
//      /* Parse an old-style command substitution within double quotes as a
//	 single word. */
//      /* XXX - sh and ksh93 don't do this - XXX */
//      else if (open == '"' && ch == '`')
//	{
//	  nestret = parse_matched_pair (0, '`', '`', &nestlen, rflags);
//
//	  CHECK_NESTRET_ERROR ();
//	  APPEND_NESTRET ();
//
//	}
//      else if (open != '`' && (tflags & LEX_WASDOL) && (ch == '(' || ch == '{' || ch == '['))	/* ) } ] */
//	/* check for $(), $[], or ${} inside quoted string. */
//	{
//parse_dollar_word:
//	  if (open == ch)	/* undo previous increment */
//	    count--;
//	  if (ch == '(')		/* ) */
//	    nestret = parse_comsub (0, '(', ')', &nestlen, (rflags|P_COMMAND) & ^P_DQUOTE);
//	  else if (ch == '{')		/* } */
//	    nestret = parse_matched_pair (0, '{', '}', &nestlen, P_FIRSTCLOSE|rflags);
//	  else if (ch == '[')		/* ] */
//	    nestret = parse_matched_pair (0, '[', ']', &nestlen, rflags);
//
//	  CHECK_NESTRET_ERROR ();
//	  APPEND_NESTRET ();
//
//	}
//      if (ch == '$')
//	tflags |= LEX_WASDOL;
//      else
//	tflags &= ^LEX_WASDOL;
//    }
//
//  ret[retind] = '\0';
//  if (lenp)
//    *lenp = retind;
///*itrace("parse_matched_pair[%d]: returning %s", gps.line_number, ret);*/
//  return ret;
//}

func parse_comsub(qc int, open int, cloze int, flags int) (*StringBuilder, os.Error) {
  // TODO(krasin): impement this
  panic("parse_comsub: not implemented")
}

///* Parse a $(...) command substitution.  This is messier than I'd like, and
//   reproduces a lot more of the token-reading code than I'd like. */
//static char *
//parse_comsub (qc, open, close, lenp, flags)
//     int qc;	/* `"' if this construct is within double quotes */
//     int open, close;
//     int *lenp, flags;
//{
//  int count, ch, peekc, tflags, lex_rwlen, lex_wlen, lex_firstind;
//  int nestlen, ttranslen, start_lineno;
//  char *ret, *nestret, *ttrans, *heredelim;
//  int retind, retsize, rflags, hdlen;
//
///*itrace("parse_comsub: qc = `%c' open = %c close = %c", qc, open, close);*/
//  count = 1;
//  tflags = LEX_RESWDOK;
//
//  if ((flags & P_COMMAND) && qc != '\'' && qc != '"' && (flags & P_DQUOTE) == 0)
//    tflags |= LEX_CKCASE;
//  if (tflags & LEX_CKCASE)
//    tflags |= LEX_CKCOMMENT;
//
//  /* RFLAGS is the set of flags we want to pass to recursive calls. */
//  rflags = (flags & P_DQUOTE);
//
//  ret = (char *)xmalloc (retsize = 64);
//  retind = 0;
//
//  start_lineno = gps.line_number;
//  lex_rwlen = lex_wlen = 0;
//
//  heredelim = 0;
//  lex_firstind = -1;
//
//  while (count)
//    {
//comsub_readchar:
//      ch = gps.shell_getc (qc != '\'' && (tflags & LEX_PASSNEXT) == 0);
//
//      if (ch == EOF)
//	{
//eof_error:
//	  gps.parser_error (start_lineno, _("unexpected EOF while looking for matching `%c'"), close);
//	  gps.EOF_Reached = true;	/* XXX */
//	  return (&matched_pair_error);
//	}
//
//      /* If we hit the end of a line and are reading the contents of a here
//	 document, and it's not the same line that the document starts on,
//	 check for this line being the here doc delimiter.  Otherwise, if
//	 we're in a here document, mark the next character as the beginning
//	 of a line. */
//      if (ch == '\n')
//	{
//	  if ((tflags & LEX_HEREDELIM) && heredelim)
//	    {
//	      tflags &= ^LEX_HEREDELIM;
//	      tflags |= LEX_INHEREDOC;
//	      lex_firstind = retind + 1;
//	    }
//	  else if (tflags & LEX_INHEREDOC)
//	    {
//	      int tind;
//	      tind = lex_firstind;
//	      while ((tflags & LEX_STRIPDOC) && ret[tind] == '\t')
//		tind++;
//	      if (STREQN (ret + tind, heredelim, hdlen))
//		{
//		  tflags &= ^(LEX_STRIPDOC|LEX_INHEREDOC);
///*itrace("parse_comsub:%d: found here doc end `%s'", gps.line_number, ret + tind);*/
//		  heredelim = 0;
//		  lex_firstind = -1;
//		}
//	      else
//		lex_firstind = retind + 1;
//	    }
//	}
//
//      /* XXX -- possibly allow here doc to be delimited by ending right
//	 paren. */
//      if ((tflags & LEX_INHEREDOC) && ch == close && count == 1)
//	{
//	  int tind;
///*itrace("parse_comsub: in here doc, ch == close, retind - firstind = %d hdlen = %d retind = %d", retind-lex_firstind, hdlen, retind);*/
//	  tind = lex_firstind;
//	  while ((tflags & LEX_STRIPDOC) && ret[tind] == '\t')
//	    tind++;
//	  if (retind-tind == hdlen && STREQN (ret + tind, heredelim, hdlen))
//	    {
//	      tflags &= ^(LEX_STRIPDOC|LEX_INHEREDOC);
///*itrace("parse_comsub:%d: found here doc end `%s'", gps.line_number, ret + tind);*/
//	      heredelim = 0;
//	      lex_firstind = -1;
//	    }
//	}
//
//      /* Don't bother counting parens or doing anything else if in a comment */
//      if (tflags & (LEX_INCOMMENT|LEX_INHEREDOC))
//	{
//	  /* Add this character. */
//	  RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//	  ret[retind++] = ch;
//
//	  if ((tflags & LEX_INCOMMENT) && ch == '\n')
//{
///*itrace("parse_comsub:%d: lex_incomment -> 0 ch = `%c'", gps.line_number, ch);*/
//	    tflags &= ^LEX_INCOMMENT;
//}
//
//	  continue;
//	}
//
//      if (tflags & LEX_PASSNEXT)		/* last char was backslash */
//	{
///*itrace("parse_comsub:%d: lex_passnext -> 0 ch = `%c' (%d)", gps.line_number, ch, __LINE__);*/
//	  tflags &= ^LEX_PASSNEXT;
//	  if (qc != '\'' && ch == '\n')	/* double-quoted \<newline> disappears. */
//	    {
//	      if (retind > 0)
//		retind--;	/* swallow previously-added backslash */
//	      continue;
//	    }
//
//	  RESIZE_MALLOCED_BUFFER (ret, retind, 2, retsize, 64);
//	  if (ch == CTLESC || ch == CTLNUL)
//	    ret[retind++] = CTLESC;
//	  ret[retind++] = ch;
//	  continue;
//	}
//
//      /* If this is a shell break character, we are not in a word.  If not,
//	 we either start or continue a word. */
//      if (shellbreak (ch))
//	{
//	  tflags &= ^LEX_INWORD;
///*itrace("parse_comsub:%d: lex_inword -> 0 ch = `%c' (%d)", gps.line_number, ch, __LINE__);*/
//	}
//      else
//	{
//	  if (tflags & LEX_INWORD)
//	    {
//	      lex_wlen++;
///*itrace("parse_comsub:%d: lex_inword == 1 ch = `%c' lex_wlen = %d (%d)", gps.line_number, ch, lex_wlen, __LINE__);*/
//	    }	      
//	  else
//	    {
///*itrace("parse_comsub:%d: lex_inword -> 1 ch = `%c' (%d)", gps.line_number, ch, __LINE__);*/
//	      tflags |= LEX_INWORD;
//	      lex_wlen = 0;
//	    }
//	}
//
//      /* Skip whitespace */
//      if (shellblank (ch) && lex_rwlen == 0)
//        {
//	  /* Add this character. */
//	  RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//	  ret[retind++] = ch;
//	  continue;
//        }
//
//      /* Either we are looking for the start of the here-doc delimiter
//	 (lex_firstind == -1) or we are reading one (lex_firstind >= 0).
//	 If this character is a shell break character and we are reading
//	 the delimiter, save it and note that we are now reading a here
//	 document.  If we've found the start of the delimiter, note it by
//	 setting lex_firstind.  Backslashes can quote shell metacharacters
//	 in here-doc delimiters. */
//      if (tflags & LEX_HEREDELIM)
//	{
//	  if (lex_firstind == -1 && shellbreak (ch) == 0)
//	    lex_firstind = retind;
//#if 0
//	  else if (heredelim && (tflags & LEX_PASSNEXT) == 0 && ch == '\n')
//	    {
//	      tflags |= LEX_INHEREDOC;
//	      tflags &= ^LEX_HEREDELIM;
//	      lex_firstind = retind + 1;
//	    }
//#endif
//	  else if (lex_firstind >= 0 && (tflags & LEX_PASSNEXT) == 0 && shellbreak (ch))
//	    {
//	      if (heredelim == 0)
//		{
//		  nestret = substring (ret, lex_firstind, retind);
//		  heredelim = string_quote_removal (nestret, 0);
//		  hdlen = STRLEN(heredelim);
///*itrace("parse_comsub:%d: found here doc delimiter `%s' (%d)", gps.line_number, heredelim, hdlen);*/
//		}
//	      if (ch == '\n')
//		{
//		  tflags |= LEX_INHEREDOC;
//		  tflags &= ^LEX_HEREDELIM;
//		  lex_firstind = retind + 1;
//		}
//	      else
//		lex_firstind = -1;
//	    }
//	}
//
//      /* Meta-characters that can introduce a reserved word.  Not perfect yet. */
//      if ((tflags & LEX_RESWDOK) == 0 && (tflags & LEX_CKCASE) && (tflags & LEX_INCOMMENT) == 0 && (shellmeta(ch) || ch == '\n'))
//	{
//	  /* Add this character. */
//	  RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//	  ret[retind++] = ch;
//	  peekc = gps.shell_getc (1);
//	  if (ch == peekc && (ch == '&' || ch == '|' || ch == ';'))	/* two-character tokens */
//	    {
//	      RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//	      ret[retind++] = peekc;
///*itrace("parse_comsub:%d: set lex_reswordok = 1, ch = `%c'", gps.line_number, ch);*/
//	      tflags |= LEX_RESWDOK;
//	      lex_rwlen = 0;
//	      continue;
//	    }
//	  else if (ch == '\n' || COMSUB_META(ch))
//	    {
//	      gps.shell_ungetc (peekc);
///*itrace("parse_comsub:%d: set lex_reswordok = 1, ch = `%c'", gps.line_number, ch);*/
//	      tflags |= LEX_RESWDOK;
//	      lex_rwlen = 0;
//	      continue;
//	    }
//	  else if (ch == EOF)
//	    goto eof_error;
//	  else
//	    {
//	      /* `unget' the character we just added and fall through */
//	      retind--;
//	      gps.shell_ungetc (peekc);
//	    }
//	}
//
//      /* If we can read a reserved word, try to read one. */
//      if (tflags & LEX_RESWDOK)
//	{
//	  if (islower (ch))
//	    {
//	      /* Add this character. */
//	      RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//	      ret[retind++] = ch;
//	      lex_rwlen++;
//	      continue;
//	    }
//	  else if (lex_rwlen == 4 && shellbreak (ch))
//	    {
//	      if (STREQN (ret + retind - 4, "case", 4))
//{
//		tflags |= LEX_INCASE;
///*itrace("parse_comsub:%d: found `case', lex_incase -> 1 lex_reswdok -> 0", gps.line_number);*/
//}
//	      else if (STREQN (ret + retind - 4, "esac", 4))
//{
//		tflags &= ^LEX_INCASE;
///*itrace("parse_comsub:%d: found `esac', lex_incase -> 0 lex_reswdok -> 0", gps.line_number);*/
//}	        
//	      tflags &= ^LEX_RESWDOK;
//	    }
//	  else if ((tflags & LEX_CKCOMMENT) && ch == '#' && (lex_rwlen == 0 || ((tflags & LEX_INWORD) && lex_wlen == 0)))
//	    ;	/* don't modify LEX_RESWDOK if we're starting a comment */
//	  else if ((tflags & LEX_INCASE) && ch != '\n')
//	    /* If we can read a reserved word and we're in case, we're at the
//	       point where we can read a new pattern list or an esac.  We
//	       handle the esac case above.  If we read a newline, we want to
//	       leave LEX_RESWDOK alone.  If we read anything else, we want to
//	       turn off LEX_RESWDOK, since we're going to read a pattern list. */
//{
//	    tflags &= ^LEX_RESWDOK;
///*itrace("parse_comsub:%d: lex_incase == 1 found `%c', lex_reswordok -> 0", gps.line_number, ch);*/
//}
//	  else if (shellbreak (ch) == 0)
//{
//	    tflags &= ^LEX_RESWDOK;
///*itrace("parse_comsub:%d: found `%c', lex_reswordok -> 0", gps.line_number, ch);*/
//}
//	}
//
//      /* Might be the start of a here-doc delimiter */
//      if ((tflags & LEX_INCOMMENT) == 0 && (tflags & LEX_CKCASE) && ch == '<')
//	{
//	  /* Add this character. */
//	  RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//	  ret[retind++] = ch;
//	  peekc = gps.shell_getc (1);
//	  if (peekc == EOF)
//	    goto eof_error;
//	  if (peekc == ch)
//	    {
//	      RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//	      ret[retind++] = peekc;
//	      peekc = gps.shell_getc (1);
//	      if (peekc == EOF)
//		goto eof_error;
//	      if (peekc == '-')
//		{
//		  RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//		  ret[retind++] = peekc;
//		  tflags |= LEX_STRIPDOC;
//		}
//	      else
//		gps.shell_ungetc (peekc);
//	      if (peekc != '<')
//		{
//		  tflags |= LEX_HEREDELIM;
//		  lex_firstind = -1;
//		}
//	      continue;
//	    }
//	  else
//	    ch = peekc;		/* fall through and continue XXX */
//	}
//      else if ((tflags & LEX_CKCOMMENT) && (tflags & LEX_INCOMMENT) == 0 && ch == '#' && (((tflags & LEX_RESWDOK) && lex_rwlen == 0) || ((tflags & LEX_INWORD) && lex_wlen == 0)))
//{
///*itrace("parse_comsub:%d: lex_incomment -> 1 (%d)", gps.line_number, __LINE__);*/
//	tflags |= LEX_INCOMMENT;
//}
//
//      if (ch == CTLESC || ch == CTLNUL)	/* special shell escapes */
//	{
//	  RESIZE_MALLOCED_BUFFER (ret, retind, 2, retsize, 64);
//	  ret[retind++] = CTLESC;
//	  ret[retind++] = ch;
//	  continue;
//	}
//#if 0
//      else if ((tflags & LEX_INCASE) && ch == close && close == ')')
//        tflags &= ^LEX_INCASE;		/* XXX */
//#endif
//      else if (ch == close && (tflags & LEX_INCASE) == 0)		/* ending delimiter */
//{
//	count--;
///*itrace("parse_comsub:%d: found close: count = %d", gps.line_number, count);*/
//}
//      else if (((flags & P_FIRSTCLOSE) == 0) && (tflags & LEX_INCASE) == 0 && ch == open)	/* nested begin */
//{
//	count++;
///*itrace("parse_comsub:%d: found open: count = %d", gps.line_number, count);*/
//}
//
//      /* Add this character. */
//      RESIZE_MALLOCED_BUFFER (ret, retind, 1, retsize, 64);
//      ret[retind++] = ch;
//
//      /* If we just read the ending character, don't bother continuing. */
//      if (count == 0)
//	break;
//
//      if (ch == '\\')			/* backslashes */
//	tflags |= LEX_PASSNEXT;
//
//      if (shellquote (ch))
//        {
//          /* '', ``, or "" inside $(...). */
//          push_delimiter (dstack, ch);
//          if ((tflags & LEX_WASDOL) && ch == '\'')	/* $'...' inside group */
//	    nestret = parse_matched_pair (ch, ch, ch, &nestlen, P_ALLOWESC|rflags);
//	  else
//	    nestret = parse_matched_pair (ch, ch, ch, &nestlen, rflags);
//	  pop_delimiter (dstack);
//	  CHECK_NESTRET_ERROR ();
//
//	  if ((tflags & LEX_WASDOL) && ch == '\'' && (extended_quote || (rflags & P_DQUOTE) == 0))
//	    {
//	      /* Translate $'...' here. */
//	      ttrans = ansiexpand (nestret, 0, nestlen - 1, &ttranslen);
//
//	      if ((rflags & P_DQUOTE) == 0)
//		{
//		  nestret = sh_single_quote (ttrans);
//		  nestlen = strlen (nestret);
//		}
//	      else
//		{
//		  nestret = ttrans;
//		  nestlen = ttranslen;
//		}
//	      retind -= 2;		/* back up before the $' */
//	    }
//	  else if ((tflags & LEX_WASDOL) && ch == '"' && (extended_quote || (rflags & P_DQUOTE) == 0))
//	    {
//	      /* Locale expand $"..." here. */
//	      ttrans = localeexpand (nestret, 0, nestlen - 1, start_lineno, &ttranslen);
//
//	      nestret = sh_mkdoublequoted (ttrans, ttranslen, 0);
//	      nestlen = ttranslen + 2;
//	      retind -= 2;		/* back up before the $" */
//	    }
//
//	  APPEND_NESTRET ();
//	}
//      else if ((tflags & LEX_WASDOL) && (ch == '(' || ch == '{' || ch == '['))	/* ) } ] */
//	/* check for $(), $[], or ${} inside command substitution. */
//	{
//	  if ((tflags & LEX_INCASE) == 0 && open == ch)	/* undo previous increment */
//	    count--;
//	  if (ch == '(')		/* ) */
//	    nestret = parse_comsub (0, '(', ')', &nestlen, (rflags|P_COMMAND) & ^P_DQUOTE);
//	  else if (ch == '{')		/* } */
//	    nestret = parse_matched_pair (0, '{', '}', &nestlen, P_FIRSTCLOSE|rflags);
//	  else if (ch == '[')		/* ] */
//	    nestret = parse_matched_pair (0, '[', ']', &nestlen, rflags);
//
//	  CHECK_NESTRET_ERROR ();
//	  APPEND_NESTRET ();
//
//	}
//      if (ch == '$')
//	tflags |= LEX_WASDOL;
//      else
//	tflags &= ^LEX_WASDOL;
//    }
//
//  ret[retind] = '\0';
//  if (lenp)
//    *lenp = retind;
///*itrace("parse_comsub:%d: returning `%s'", gps.line_number, ret);*/
//  return ret;
//}
//
///* XXX - this needs to handle functionality like subst.c:no_longjmp_on_fatal_error;
//   maybe extract_command_subst should handle it. */
//char *
//xparse_dolparen (base, string, indp, flags)
//     char *base;
//     char *string;
//     int *indp;
//     int flags;
//{
//  sh_parser_state_t ps;
//  int orig_ind, nc, sflags;
//  char *ret, *s, *ep, *ostring;
//
//  /*yydebug = 1;*/
//  orig_ind = *indp;
//  ostring = string;
//
//  sflags = SEVAL_NONINT|SEVAL_NOHIST|SEVAL_NOFREE;
//  if (flags & SX_NOLONGJMP)
//    sflags |= SEVAL_NOLONGJMP;
//  save_parser_state (&ps);
//
//  /*(*/
//  gps.parser_state |= PST_CMDSUBST|PST_EOFTOKEN;	/* allow instant ')' */ /*(*/
//  gps.shell_eof_token = ')';
//  parse_string (string, "command substitution", sflags, &ep);
//
//  restore_parser_state (&ps);
//  gps.reset_parser ();
//
//  /* Need to find how many characters parse_and_execute consumed, update
//     *indp, if flags != 0, copy the portion of the string parsed into RET
//     and return it.  If flags & 1 (EX_NOALLOC) we can return NULL. */
//
//  /*(*/
//  if (ep[-1] != ')')
//    {
//      while (ep > ostring && ep[-1] == '\n') ep--;
//    }
//
//  nc = ep - ostring;
//  *indp = ep - base - 1;
//
//  /*(*/
//
//  if (flags & SX_NOALLOC) 
//    return nil;
//
//  if (nc == 0)
//    {
//      ret = xmalloc (1);
//      ret[0] = '\0';
//    }
//  else
//    ret = substring (ostring, 0, nc - 1);
//
//  return ret;
//}

/* Parse a double-paren construct.  It can be either an arithmetic
   command, an arithmetic `for' command, or a nested subshell.  Returns
   the parsed token, -1 on error, or -2 if we didn't do anything and
   should just go on. */
func (gps *ParserState) parse_dparen (c int) int {
  var cmdtyp int
  var wval string
  var wd *word_desc

  if gps.last_read_token == FOR {
      gps.arith_for_lineno = gps.line_number;
      cmdtyp, wval = gps.parse_arith_cmd (false);
      if cmdtyp == 1 {
	  wd = new(word_desc)
	  wd.word = wval;
	  gps.yylval.word_list = makeWordList (wd, nil);
	  return (ARITH_FOR_EXPRS);
      } else {
	return -1;		/* ERROR */
      }
  }

  if (reserved_word_acceptable (gps.last_read_token)) {
      cmdtyp, wval = gps.parse_arith_cmd (false);
      switch {
      case cmdtyp == 1:	/* arithmetic command */
	  wd = new(word_desc)
	  wd.word = wval;
	  wd.flags = W_QUOTED|W_NOSPLIT|W_NOGLOB|W_DQUOTE;
	  gps.yylval.word_list = makeWordList (wd, nil);
	  return (ARITH_CMD);
      case cmdtyp == 0:	/* nested subshell */
	  gps.push_string (wval, false, nil);
	  if ((gps.parser_state & PST_CASEPAT) == 0) {
	    gps.parser_state |= PST_SUBSHELL;
          }
	  return (c);
      default:			/* ERROR */
	return -1;
      }
  }

  return -2;			/* XXX */
}

/* We've seen a `(('.  Look for the matching `))'.  If we get it, return 1.
   If not, assume it's a nested subshell for backwards compatibility and
   return 0.  In any case, put the characters we've consumed into a locally-
   allocated buffer and make *ep point to that buffer.  Return -1 on an
   error, for example EOF. */
func (gps *ParserState) parse_arith_cmd(adddq bool) (rval int, tokstr string) {
  var c int

  ttok, err := parse_matched_pair (0, '(', ')', 0);
  if err != nil {
    return -1, ""
  }
  rval = 1;
  /* Check that the next character is the closing right paren.  If
     not, this is a syntax error. ( */
  c = gps.shell_getc(false);
  if (c != ')') {
    rval = 0;
  }

  /* if ADDDQ != 0 then (( ... )) -> "..." */
  switch {
  case rval == 1 && adddq:	/* arith cmd, add double quotes */
      tokstr = fmt.Sprintf("\"%s\"", ttok.String())
  case rval == 1:		/* arith cmd, don't add double quotes */
      tokstr = ttok.String()
  default:			/* nested subshell */
      tokstr = fmt.Sprintf("(%s)%s", runesToString([]int { c }))
  }

  return
}

func (gps *ParserState) cond_error() {
  if (gps.EOF_Reached && gps.cond_token != COND_ERROR) {		/* [[ */
    gps.parser_error (gps.cond_lineno, "unexpected EOF while looking for `]]'");
  } else {
    if (gps.cond_token != COND_ERROR) {
      etext := error_token_from_token(gps.cond_token)
      if etext != "" {
        gps.parser_error (gps.cond_lineno, "syntax error in conditional expression: unexpected token `%s'", etext);
      } else {
        gps.parser_error (gps.cond_lineno, "syntax error in conditional expression");
      }
    }
  }
}

func (gps *ParserState) cond_expr () *CondCom {
  return gps.cond_or()
}

func (gps *ParserState) cond_or () *CondCom {
  l := gps.cond_and ();
  if gps.cond_token == OR_OR {
      r := gps.cond_or ();
      l = gps.make_cond_node (COND_OR, nil, l, r);
  }
  return l;
}

func (gps *ParserState) cond_and () *CondCom {
  l := gps.cond_term ();
  if gps.cond_token == AND_AND {
      r := gps.cond_and()
      l = gps.make_cond_node (COND_AND, nil, l, r);
  }
  return l;
}

//static int
//cond_skip_newlines ()
//{
//  while ((gps.cond_token = read_token (READ)) == '\n')
//    {
//    }
//  return (gps.cond_token);
//}
//
//#define COND_RETURN_ERROR() \
//  do { gps.cond_token = COND_ERROR; return (nil); } while (0)
//

// TODO(krasin): implement this
func (gps *ParserState) cond_term() *CondCom {
	panic("cond_term: not implemented")
}

//static CondCom *
//cond_term ()
//{
//  word_desc *op;
//  CondCom *term, *tleft, *tright;
//  int tok, lineno;
//  char *etext;
//
//  /* Read a token.  It can be a left paren, a `!', a unary operator, or a
//     word that should be the first argument of a binary operator.  Start by
//     skipping newlines, since this is a compound command. */
//  tok = cond_skip_newlines ();
//  lineno = gps.line_number;
//  if (tok == COND_END)
//    {
//      COND_RETURN_ERROR ();
//    }
//  else if (tok == '(')
//    {
//      term = cond_expr ();
//      if (gps.cond_token != ')')
//	{
//	  if (term)
//	    dispose_cond_node (term);		/* ( */
//	  if (etext = error_token_from_token (gps.cond_token))
//	    {
//	      gps.parser_error (lineno, _("unexpected token `%s', expected `)'"), etext);
//	    }
//	  else
//	    gps.parser_error (lineno, _("expected `)'"));
//	  COND_RETURN_ERROR ();
//	}
//      term = make_cond_node (COND_EXPR, nil, term, nil);
//      (void)cond_skip_newlines ();
//    }
//  else if (tok == BANG || (tok == WORD && (gps.yylval.word.word[0] == '!' && gps.yylval.word.word[1] == '\0')))
//    {
//      if (tok == WORD)
//	dispose_word (gps.yylval.word);	/* not needed */
//      term = cond_term ();
//      if (term)
//	term.flags |= CMD_INVERT_RETURN;
//    }
//  else if (tok == WORD && gps.yylval.word.word[0] == '-' && gps.yylval.word.word[2] == 0 && test_unop (gps.yylval.word.word))
//    {
//      op = gps.yylval.word;
//      tok = read_token (READ);
//      if (tok == WORD)
//	{
//	  tleft = make_cond_node (COND_TERM, gps.yylval.word, nil, nil);
//	  term = make_cond_node (COND_UNARY, op, tleft, nil);
//	}
//      else
//	{
//	  dispose_word (op);
//	  if (etext = error_token_from_token (tok))
//	    {
//	      gps.parser_error (gps.line_number, _("unexpected argument `%s' to conditional unary operator"), etext);
//	    }
//	  else
//	    gps.parser_error (gps.line_number, _("unexpected argument to conditional unary operator"));
//	  COND_RETURN_ERROR ();
//	}
//
//      (void)cond_skip_newlines ();
//    }
//  else if (tok == WORD)		/* left argument to binary operator */
//    {
//      /* lhs */
//      tleft = make_cond_node (COND_TERM, gps.yylval.word, nil, nil);
//
//      /* binop */
//      tok = read_token (READ);
//      if (tok == WORD && test_binop (gps.yylval.word.word))
//	{
//	  op = gps.yylval.word;
//	  if (op.word[0] == '=' && (op.word[1] == '\0' || (op.word[1] == '=' && op.word[2] == '\0')))
//	    gps.parser_state |= PST_EXTPAT;
//	  else if (op.word[0] == '!' && op.word[1] == '=' && op.word[2] == '\0')
//	    gps.parser_state |= PST_EXTPAT;
//	}
//      else if (tok == WORD && STREQ (gps.yylval.word.word, "=~"))
//	{
//	  op = gps.yylval.word;
//	  gps.parser_state |= PST_REGEXP;
//	}
//      else if (tok == '<' || tok == '>')
//	op = make_word_from_token (tok);  /* ( */
//      /* There should be a check before blindly accepting the `)' that we have
//	 seen the opening `('. */
//      else if (tok == COND_END || tok == AND_AND || tok == OR_OR || tok == ')')
//	{
//	  /* Special case.  [[ x ]] is equivalent to [[ -n x ]], just like
//	     the test command.  Similarly for [[ x && expr ]] or
//	     [[ x || expr ]] or [[ (x) ]]. */
//	  op = make_word ("-n");
//	  term = make_cond_node (COND_UNARY, op, tleft, nil);
//	  gps.cond_token = tok;
//	  return (term);
//	}
//      else
//	{
//	  if (etext = error_token_from_token (tok))
//	    {
//	      gps.parser_error (gps.line_number, _("unexpected token `%s', conditional binary operator expected"), etext);
//	    }
//	  else
//	    gps.parser_error (gps.line_number, _("conditional binary operator expected"));
//	  dispose_cond_node (tleft);
//	  COND_RETURN_ERROR ();
//	}
//
//      /* rhs */
//      if (gps.parser_state & PST_EXTPAT)
//	gps.extended_glob = 1;
//      tok = read_token (READ);
//      if (gps.parser_state & PST_EXTPAT)
//	gps.extended_glob = gps.global_extglob;
//      gps.parser_state &= ^(PST_REGEXP|PST_EXTPAT);
//
//      if (tok == WORD)
//	{
//	  tright = make_cond_node (COND_TERM, gps.yylval.word, nil, nil);
//	  term = make_cond_node (COND_BINARY, op, tleft, tright);
//	}
//      else
//	{
//	  if (etext = error_token_from_token (tok))
//	    {
//	      gps.parser_error (gps.line_number, _("unexpected argument `%s' to conditional binary operator"), etext);
//	    }
//	  else
//	    gps.parser_error (gps.line_number, _("unexpected argument to conditional binary operator"));
//	  dispose_cond_node (tleft);
//	  dispose_word (op);
//	  COND_RETURN_ERROR ();
//	}
//
//      (void)cond_skip_newlines ();
//    }
//  else
//    {
//      if (tok < 256)
//	gps.parser_error (gps.line_number, _("unexpected token `%c' in conditional command"), tok);
//      else if (etext = error_token_from_token (tok))
//	{
//	  gps.parser_error (gps.line_number, _("unexpected token `%s' in conditional command"), etext);
//	}
//      else
//	gps.parser_error (gps.line_number, _("unexpected token %d in conditional command"), tok);
//      COND_RETURN_ERROR ();
//    }
//  return (term);
//}      
//
///* This is kind of bogus -- we slip a mini recursive-descent parser in
//   here to handle the conditional statement syntax. */
func (gps *ParserState) parse_cond_command() *Command {
  gps.global_extglob = gps.extended_glob
  cexp := gps.cond_expr()
  return gps.make_cond_command(cexp)
}

func token_is_assignment(t *StringBuilder) bool {
  // TODO(krasin): implement this
  panic("token_is_assignment: not implemented")
}

func token_is_ident(t *StringBuilder) bool {
  // TODO(krasin): implement this
  panic("token_is_ident: not implemented")
}
///* When this is called, it's guaranteed that we don't care about anything
//   in t beyond i.  We do save and restore the chars, though. */
//static int
//token_is_assignment (t, i)
//     char *t;
//     int i;
//{
//  unsigned char c, c1;
//  int r;
//
//  c = t[i]; c1 = t[i+1];
//  t[i] = '='; t[i+1] = '\0';
//  r = assignment (t, (gps.parser_state & PST_COMPASSIGN) != 0);
//  t[i] = c; t[i+1] = c1;
//  return r;
//}
//
///* XXX - possible changes here for `+=' */
//static int
//token_is_ident (t, i)
//     char *t;
//     int i;
//{
//  unsigned char c;
//  int r;
//
//  c = t[i];
//  t[i] = '\0';
//  r = legal_identifier (t);
//  t[i] = c;
//  return r;
//}
//

type wordTokenizerState struct {
  gps *ParserState

  character int

  /* The value for YYLVAL when a WORD is read. */
  the_word *word_desc

  /* The token that we are building. */
  token *StringBuilder

  /* ALL_DIGITS becomes zero when we see a non-digit. */
  all_digit_token bool

  /* DOLLAR_PRESENT becomes non-zero if we see a `$'. */
  dollar_present bool

  /* COMPOUND_ASSIGNMENT becomes non-zero if we are parsing a compound
     assignment. */
  compound_assignment bool

  /* QUOTED becomes non-zero if we see one of ("), ('), (`), or (\). */
  quoted bool

  /* Non-zero means to ignore the value of the next character, and just
     to add it no matter what. */
  pass_next_character bool

  /* The current delimiting character. */
  cd int
  peek_char int
  ttok *StringBuilder
  ttrans *StringBuilder
  ttoklen int
  translen int
}

type readTokenWordState int
const (
  RTS_BAIL_IMMEDIATELY = readTokenWordState(-1)
  RTS_PASS = readTokenWordState(0)
  RTS_GOT_CHARACTER = readTokenWordState(iota)
  RTS_GOT_ESCAPED_CHARACTER = readTokenWordState(iota)
  RTS_NEXT_CHARACTER = readTokenWordState(iota)
  RTS_GOT_TOKEN = readTokenWordState(iota)
)

func (wts *wordTokenizerState) handleBackslashes() readTokenWordState {
  /* Handle backslashes.  Quote lots of things when not inside of
     double-quotes, quote some things inside of double-quotes. */
  if wts.character == '\\' {
    wts.peek_char = wts.gps.shell_getc(false)

    /* Backslash-newline is ignored in all cases except
       when quoted with single quotes. */
    if (wts.peek_char == '\n') {
      wts.character = '\n';
      return RTS_NEXT_CHARACTER
    } else {
      wts.gps.shell_ungetc (wts.peek_char);

      /* If the next character is to be quoted, note it now. */
      if (wts.cd == 0 || wts.cd == '`' ||
          (wts.cd == '"' && wts.peek_char >= 0 && (sh_syntaxtab[wts.peek_char] & CBSDQUOTE != 0))) {
        wts.pass_next_character = true
      }

      wts.quoted = true
      return RTS_GOT_CHARACTER
    }
  }
  return RTS_PASS
}

func (wts *wordTokenizerState) handleShellQuote() readTokenWordState {
      /* Parse a matched pair of quote characters. */
  if (shellquote (wts.character)) {
	  wts.gps.push_delimiter (wts.character);
          flags := 0
          if wts.character == '`' {
            flags = P_COMMAND
          }
      var err os.Error
	  wts.ttok, err = parse_matched_pair (wts.character, wts.character, wts.character, flags)
	  wts.gps.pop_delimiter()
	  if err != nil {
	    return RTS_BAIL_IMMEDIATELY
      }
      wts.token.Add(wts.character)
      wts.token.Append(wts.ttok)
	  wts.all_digit_token = false
	  wts.quoted = true
	  wts.dollar_present = wts.dollar_present || (wts.character == '"' && wts.ttok.HasRune('$'))
      return RTS_NEXT_CHARACTER
  }
  return RTS_PASS
}

func (wts *wordTokenizerState) handleRegexp() readTokenWordState {
      /* When parsing a regexp as a single word inside a conditional command,
	 we need to special-case characters special to both the shell and
	 regular expressions.  Right now, that is only '(' and '|'. */ /*)*/
      if (wts.gps.parser_state & PST_REGEXP != 0) && (wts.character == '(' || wts.character == '|') { /*)*/
	  if wts.character == '|' {
        return RTS_GOT_CHARACTER
      }

	  wts.gps.push_delimiter (wts.character);
      var err os.Error
	  wts.ttok, err = parse_matched_pair (wts.cd, '(', ')', 0);
	  wts.gps.pop_delimiter()
	  if err != nil {
        return RTS_BAIL_IMMEDIATELY
      }
	  wts.token.Add(wts.character)
	  wts.token.Append(wts.ttok)
      wts.dollar_present = false
	  wts.all_digit_token = false
      return RTS_NEXT_CHARACTER
	}
  return RTS_PASS
}

func (wts *wordTokenizerState) handleExtendedGlob() readTokenWordState {
      /* Parse a ksh-style extended pattern matching specification. */
      if wts.gps.extended_glob && PATTERN_CHAR (wts.character) {
	  wts.peek_char = wts.gps.shell_getc(true)
	  if (wts.peek_char == '(') {		/* ) */
	      wts.gps.push_delimiter (wts.peek_char);
          var err os.Error
	      wts.ttok, err = parse_matched_pair (wts.cd, '(', ')', 0);
	      wts.gps.pop_delimiter()
	      if err != nil {
            return RTS_BAIL_IMMEDIATELY
          }
	      wts.token.Add(wts.character)
	      wts.token.Add(wts.peek_char)
          wts.token.Append(wts.ttok)
	      wts.dollar_present = false
          wts.all_digit_token = false
          return RTS_NEXT_CHARACTER
    } else {
      wts.gps.shell_ungetc (wts.peek_char);
    }
  }
  return RTS_PASS
}

func localeexpand(str *StringBuilder, lineno int) *StringBuilder {
  // TODO(krasin): implement this
  panic("localeexpand: not implemented")
}

func sh_mkdoublequoted(str *StringBuilder, flags int) *StringBuilder {
  // TODO(krasin): implement this
  panic("sh_mkdoublequoted: not implemented")
}

func sh_single_quote(str *StringBuilder) *StringBuilder {
  // TODO(krasin): implement this
  panic("sh_single_quote: not implemented")
}


func (wts *wordTokenizerState) handleShellExp() readTokenWordState {
  /* If the delimiter character is not single quote, parse some of
     the shell expansions that must be read as a single word. */
  var err os.Error
  switch {
  case shellexp(wts.character):
    wts.peek_char = wts.gps.shell_getc (true);
    switch {
    /* $(...), <(...), >(...), $((...)), ${...}, and $[...] constructs */
    case (wts.peek_char == '(' ||
        ((wts.peek_char == '{' || wts.peek_char == '[') && wts.character == '$')):    /* ) ] } */
      switch {
      case wts.peek_char == '{':        /* } */
        wts.ttok, err = parse_matched_pair (wts.cd, '{', '}', P_FIRSTCLOSE);
      case wts.peek_char == '(':        /* ) */
        /* XXX - push and pop the `(' as a delimiter for use by
           the command-oriented-history code.  This way newlines
           appearing in the $(...) string get added to the
           history literally rather than causing a possibly-
           incorrect `;' to be added. ) */
        wts.gps.push_delimiter (wts.peek_char);
        wts.ttok, err = parse_comsub (wts.cd, '(', ')', P_COMMAND);
        wts.gps.pop_delimiter()
      default:
        wts.ttok, err = parse_matched_pair (wts.cd, '[', ']', 0);
      }

      if err != nil {
        return RTS_BAIL_IMMEDIATELY
      }
      wts.token.Add(wts.character)
      wts.token.Add(wts.peek_char)
      wts.token.Append(wts.ttok)
      wts.dollar_present = true
      wts.all_digit_token = false
      return RTS_NEXT_CHARACTER
    /* This handles $'...' and $"..." new-style quoted strings. */
    case (wts.character == '$' && (wts.peek_char == '\'' || wts.peek_char == '"')):
      first_line := wts.gps.line_number;
      wts.gps.push_delimiter (wts.peek_char);
      flags := 0
      if wts.peek_char == '\'' {
        flags = P_ALLOWESC
      }
      wts.ttok, err = parse_matched_pair (wts.peek_char, wts.peek_char, wts.peek_char, flags)
      wts.gps.pop_delimiter()
      if err != nil {
        return RTS_BAIL_IMMEDIATELY
      }
      if (wts.peek_char == '\'') {
        wts.ttrans = ansiexpand(wts.ttok)

        /* Insert the single quotes and correctly quote any
           embedded single quotes (allowed because P_ALLOWESC was
           passed to parse_matched_pair). */
        wts.ttok = sh_single_quote (wts.ttrans);
        wts.ttrans = wts.ttok;
      } else {
        /* Try to locale-expand the converted string. */
        wts.ttrans = localeexpand (wts.ttok, first_line)

        /* Add the double quotes back */
        wts.ttok = sh_mkdoublequoted (wts.ttrans, 0)
        wts.ttrans = wts.ttok;
      }

      wts.token.Append(wts.ttrans)
      wts.quoted = true
      wts.all_digit_token = false
      return RTS_NEXT_CHARACTER
    /* This could eventually be extended to recognize all of the
       shell's single-character parameter expansions, and set flags.*/
    case (wts.character == '$' && wts.peek_char == '$'):
      wts.ttok = StringToBuilder("$$")
      wts.token.Append(wts.ttok)
      wts.dollar_present = true
      wts.all_digit_token = false
      return RTS_NEXT_CHARACTER
    default:
      wts.gps.shell_ungetc (wts.peek_char);
    }

  /* Identify possible array subscript assignment; match [...].  If
     wts.gps.parser_state&PST_COMPASSIGN, we need to parse [sub]=words treating
     `sub' as if it were enclosed in double quotes. */
  case (wts.character == '[' &&        /* ] */
       ((wts.token.Len() > 0 && wts.gps.assignment_acceptable(wts.gps.last_read_token) && token_is_ident (wts.token)) ||
       (wts.token.Len() == 0 && (wts.gps.parser_state&PST_COMPASSIGN != 0)))):
    var err os.Error
    wts.ttok, err = parse_matched_pair (wts.cd, '[', ']', P_ARRAYSUB);
    if err != nil {
      return RTS_BAIL_IMMEDIATELY
    }
    wts.token.Add(wts.character)
    wts.token.Append(wts.ttok)
    wts.all_digit_token = false
    return RTS_NEXT_CHARACTER

      /* Identify possible compound array variable assignment. */
  case (wts.character == '=' && wts.token.Len() > 0 && (wts.gps.assignment_acceptable (wts.gps.last_read_token) || (wts.gps.parser_state & PST_ASSIGNOK != 0)) && token_is_assignment (wts.token)):
    wts.peek_char = wts.gps.shell_getc (true)
    if (wts.peek_char == '(') {        /* ) */
      wts.ttok = parse_compound_assignment()
      wts.token.Add('=')
      wts.token.Add('(')
      if (wts.ttok != nil) {
        wts.token.Append(wts.ttok)
      }
      wts.token.Add(')')
      wts.all_digit_token = false
      wts.compound_assignment = true
      return RTS_NEXT_CHARACTER
    } else {
      wts.gps.shell_ungetc (wts.peek_char);
    }
  }
  return RTS_PASS
}

func (wts *wordTokenizerState) handleChar() (rts_state readTokenWordState) {
  if wts.character == EOF {
    return RTS_GOT_TOKEN
  }

  if wts.pass_next_character {
    wts.pass_next_character = false
    return RTS_GOT_ESCAPED_CHARACTER
  }

  wts.cd = wts.gps.current_delimiter()

  if rts_state = wts.handleBackslashes(); rts_state != RTS_PASS {
    return rts_state
  }
  if rts_state = wts.handleShellQuote(); rts_state != RTS_PASS {
    return rts_state
  }
  if rts_state = wts.handleRegexp(); rts_state != RTS_PASS {
    return rts_state
  }
  if rts_state = wts.handleExtendedGlob(); rts_state != RTS_PASS {
    return rts_state
  }
  if rts_state = wts.handleShellExp(); rts_state != RTS_PASS {
    return rts_state
  }

  /* When not parsing a multi-character word construct, shell meta-
     characters break words. */
  if shellbreak (wts.character) {
    wts.gps.shell_ungetc (wts.character);
    return RTS_GOT_TOKEN;
  }
  return RTS_PASS
}

func (gps *ParserState) read_token_word(ch int) int {
  wts := new(wordTokenizerState)
  wts.gps = gps
  wts.character = ch
  wts.token = NewStringBuilder()

  wts.all_digit_token = unicode.IsDigit(wts.character)

  rts_state := RTS_PASS
  rts_loop: for {
    switch rts_state {
    case RTS_PASS:
      rts_state = wts.handleChar()
      if rts_state != RTS_PASS {
        continue
      }
      fallthrough
    case RTS_GOT_CHARACTER:
      if (wts.character == CTLESC || wts.character == CTLNUL) {
        wts.token.Add(CTLESC)
      }
      fallthrough

    case RTS_GOT_ESCAPED_CHARACTER:
      wts.all_digit_token = wts.all_digit_token && unicode.IsDigit(wts.character);
      wts.dollar_present = wts.dollar_present || (wts.character == '$')

      wts.token.Add(wts.character)
      fallthrough

    case RTS_NEXT_CHARACTER:
      /* We want to remove quoted newlines (that is, a \<newline> pair)
         unless we are within single quotes or wts.pass_next_character is
         set (the shell equivalent of literal-next). */
      wts.cd = wts.gps.current_delimiter()
      wts.character = gps.shell_getc (wts.cd != '\'' && !wts.pass_next_character);
    case RTS_GOT_TOKEN:
      break rts_loop

    case RTS_BAIL_IMMEDIATELY:
      return -1
    }
    rts_state = RTS_PASS
  }  /* end for { */

  /* Check to see what thing we should return.  If the gps.last_read_token
     is a `<', or a `&', or the character which ended this token is
     a '>' or '<', then, and ONLY then, is this input token a NUMBER.
     Otherwise, it is just a word, and should be returned as such. */
  if (wts.all_digit_token && (wts.character == '<' || wts.character == '>' ||
      gps.last_read_token == LESS_AND ||
      gps.last_read_token == GREATER_AND)) {
    if lvalue, err := strconv.Atoi64(wts.token.String()); err == nil && int64(int(lvalue)) == lvalue {
	  gps.yylval.number = int(lvalue)
	} else {
	  gps.yylval.number = -1;
    }
	return (NUMBER);
  }

  token_word := wts.token.String()

  /* Check for special case tokens. */
  result := gps.special_case_tokens (token_word)
  if result >= 0 {
    return result;
  }

  /* Posix.2 does not allow reserved words to be aliased, so check for all
     of them, including special cases, before expanding the current token
     as an alias. */
  if (gps.posixly_correct) {
    if tok := wts.CHECK_FOR_RESERVED_WORD (token_word); tok != NO_TOKEN {
      return tok
    }
  }

  /* Aliases are expanded iff EXPAND_ALIASES is non-zero, and quoting
     inhibits alias expansion. */
  if (gps.expand_aliases && !wts.quoted) {
    result = alias_expand_token (token_word);
    if (result == RE_READ_TOKEN) {
      return (RE_READ_TOKEN);
    } else {
      if (result == NO_EXPANSION) {
        gps.parser_state &= ^PST_ALEXPNEXT;
      }
    }
  }

  /* If not in Posix.2 mode, check for reserved words after alias
     expansion. */
  if !gps.posixly_correct {
    if tok := wts.CHECK_FOR_RESERVED_WORD (token_word); tok != NO_TOKEN {
      return tok
    }
  }

  wts.the_word = new(word_desc)
  wts.the_word.word = wts.token.String()
  if (wts.dollar_present) {
    wts.the_word.flags |= W_HASDOLLAR;
  }
  if (wts.quoted) {
    wts.the_word.flags |= W_QUOTED;		/*(*/
  }
  if (wts.compound_assignment && wts.token.AtLast() == ')') {
    wts.the_word.flags |= W_COMPASSIGN;
  }
  /* A word is an assignment if it appears at the beginning of a
     simple command, or after another assignment word.  This is
     context-dependent, so it cannot be handled in the grammar. */
  if (assignment (wts.token.Runes(), (gps.parser_state & PST_COMPASSIGN) != 0) > 0) {
    wts.the_word.flags |= W_ASSIGNMENT;
      /* Don't perform word splitting on assignment statements. */
    if (gps.assignment_acceptable (gps.last_read_token) || (gps.parser_state & PST_COMPASSIGN) != 0) {
      wts.the_word.flags |= W_NOSPLIT;
    }
  }

  if (gps.command_token_position (gps.last_read_token)) {
    b := gps.builtins.builtin_address_internal (token_word, false);
    if b != nil && (b.flags & ASSIGNMENT_BUILTIN != 0) {
      gps.parser_state |= PST_ASSIGNOK;
    } else {
      if token_word == "eval" || token_word == "let" {
        gps.parser_state |= PST_ASSIGNOK;
      }
    }
  }

  gps.yylval.word = wts.the_word;

  if (wts.token.AtFirst() == '{' && wts.token.AtLast() == '}' &&
      (wts.character == '<' || wts.character == '>')) {
      trimmed := wts.token.RangeString(1, wts.token.Len()-1)
      if legal_identifier (trimmed) {
        wts.the_word.word = trimmed
/*itrace("read_token_word: returning REDIR_WORD for %s", wts.the_word->word);*/
	    return (REDIR_WORD);
      }
  }

  if ((wts.the_word.flags & (W_ASSIGNMENT|W_NOSPLIT)) == (W_ASSIGNMENT|W_NOSPLIT)) {
    result = ASSIGNMENT_WORD
  } else {
    result = WORD
  }


  switch (gps.last_read_token) {
    case FUNCTION:
      gps.parser_state |= PST_ALLOWOPNBRC;
      gps.function_dstart = gps.line_number;
    case CASE: fallthrough
    case SELECT: fallthrough
    case FOR:
      if (gps.word_top < MAX_CASE_NEST) {
        gps.word_top++;
        gps.word_lineno[gps.word_top] = gps.line_number;
      }
  }

  return result
}

/* Return 1 if TOKSYM is a token that after being read would allow
   a reserved word to be seen, else 0. */
func reserved_word_acceptable(toksym int) bool {
  switch (toksym) {
    case '\n': fallthrough
    case ';': fallthrough
    case '(': fallthrough
    case ')': fallthrough
    case '|': fallthrough
    case '&': fallthrough
    case '{': fallthrough
    case '}': fallthrough /* XXX */
    case AND_AND: fallthrough
    case BANG: fallthrough
    case BAR_AND: fallthrough
    case DO: fallthrough
    case DONE: fallthrough
    case ELIF: fallthrough
    case ELSE: fallthrough
    case ESAC: fallthrough
    case FI: fallthrough
    case IF: fallthrough
    case OR_OR: fallthrough
    case SEMI_SEMI: fallthrough
    case SEMI_AND: fallthrough
    case SEMI_SEMI_AND: fallthrough
    case THEN: fallthrough
    case TIME: fallthrough
    case TIMEOPT: fallthrough
    case COPROC: fallthrough
    case UNTIL: fallthrough
    case WHILE: fallthrough
    case 0:
      return true
    }
    return false
}

///* Return the index of TOKEN in the alist of reserved words, or -1 if
//   TOKEN is not a shell reserved word. */
//int
//find_reserved_word (tokstr)
//     char *tokstr;
//{
//  int i;
//  for (i = 0; word_token_alist[i].word; i++)
//    if (STREQ (tokstr, word_token_alist[i].word))
//      return i;
//  return -1;
//}
//
///************************************************
// *						*
// *		ERROR HANDLING			*
// *						*
// ************************************************/
//
///* Report a syntax error, and restart the parser.  Call here for fatal
//   errors. */
//int
//yyerror (msg)
//     const char *msg;
//{
//  report_syntax_error (nil);
//  gps.reset_parser ();
//  return (0);
//}
//


// TODO(krasin): implement this
func error_token_from_token(tok int) string {
	panic("error_token_from_token: not implemented")
}

//static char *
//error_token_from_token (tok)
//     int tok;
//{
//  char *t;
//
//  if (t = find_token_in_alist (tok, word_token_alist, 0))
//    return t;
//
//  if (t = find_token_in_alist (tok, other_token_alist, 0))
//    return t;
//
//  t = nil;
//  /* This stuff is dicy and needs closer inspection */
//  switch (gps.current_token)
//    {
//    case WORD: fallthrough
//    case ASSIGNMENT_WORD:
//      if (gps.yylval.word)
//	t = savestring (gps.yylval.word.word);
//      break;
//    case NUMBER:
//      t = itos (gps.yylval.number);
//      break;
//    case ARITH_CMD:
//      if (gps.yylval.word_list)
//        t = string_list (gps.yylval.word_list);
//      break;
//    case ARITH_FOR_EXPRS:
//      if (gps.yylval.word_list)
//	t = string_list_internal (gps.yylval.word_list, " ; ");
//      break;
//    case COND_CMD:
//      t = nil;		/* punt */
//      break;
//    }
//
//  return t;
//}
//
//static char *
//error_token_from_text ()
//{
//  char *msg, *t;
//  int token_end, i;
//
//  t = gps.shell_input_line;
//  i = gps.shell_input_line_index;
//  token_end = 0;
//  msg = nil;
//
//  if (i && t[i] == '\0')
//    i--;
//
//  while (i && (whitespace (t[i]) || t[i] == '\n'))
//    i--;
//
//  if (i)
//    token_end = i + 1;
//
//  while (i && (member (t[i], " \n\t;|&") == 0))
//    i--;
//
//  while (i != token_end && (whitespace (t[i]) || t[i] == '\n'))
//    i++;
//
//  /* Return our idea of the offending token. */
//  if (token_end || (i == 0 && token_end == 0))
//    {
//      if (token_end)
//	msg = substring (t, i, token_end);
//      else	/* one-character token */
//	{
//	  msg = (char *)xmalloc (2);
//	  msg[0] = t[i];
//	  msg[1] = '\0';
//	}
//    }
//
//  return (msg);
//}
//
//static void
//print_offending_line ()
//{
//  char *msg;
//  int token_end;
//
//  msg = savestring (gps.shell_input_line);
//  token_end = strlen (msg);
//  while (token_end && msg[token_end - 1] == '\n')
//    msg[--token_end] = '\0';
//
//  gps.parser_error (gps.line_number, "`%s'", msg);
//}
//
///* Report a syntax error with line numbers, etc.
//   Call here for recoverable errors.  If you have a message to print,
//   then place it in MESSAGE, otherwise pass NULL and this will figure
//   out an appropriate message for you. */
//static void
//report_syntax_error (message)
//     char *message;
//{
//  char *msg;
//
//  if (message)
//    {
//      gps.parser_error (gps.line_number, "%s", message);
//      last_command_exit_value = parse_and_execute_level ? EX_BADSYNTAX : EX_BADUSAGE;
//      return;
//    }
//
//  /* If the line of input we're reading is not null, try to find the
//     objectionable token.  First, try to figure out what token the
//     parser's complaining about by looking at gps.current_token. */
//  if (gps.current_token != 0 && !gps.EOF_Reached && (msg = error_token_from_token (gps.current_token)))
//    {
//      gps.parser_error (gps.line_number, _("syntax error near unexpected token `%s'"), msg);
//
//      print_offending_line ();
//
//      last_command_exit_value = parse_and_execute_level ? EX_BADSYNTAX : EX_BADUSAGE;
//      return;
//    }
//
//  /* If looking at the current token doesn't prove fruitful, try to find the
//     offending token by analyzing the text of the input line near the current
//     input line index and report what we find. */
//  if (gps.shell_input_line && *gps.shell_input_line)
//    {
//      msg = error_token_from_text ();
//      if (msg)
//	{
//	  gps.parser_error (gps.line_number, _("syntax error near `%s'"), msg);
//	}
//
//      print_offending_line ();
//    }
//  else
//    {
//      msg = gps.EOF_Reached ? _("syntax error: unexpected end of file") : _("syntax error");
//      gps.parser_error (gps.line_number, "%s", msg);
//    }
//
//  last_command_exit_value = parse_and_execute_level ? EX_BADSYNTAX : EX_BADUSAGE;
//}
//
///* ??? Needed function. ??? We have to be able to discard the constructs
//   created during parsing.  In the case of error, we want to return
//   allocated objects to the memory pool.  In the case of no error, we want
//   to throw away the information about where the allocated objects live.
//   (dispose_command () will actually free the command.) */
//static void
//discard_parser_constructs (error_p)
//     int error_p;
//{
//}
//
///************************************************
// *						*
// *		EOF HANDLING			*
// *						*
// ************************************************/
//
///* Do that silly `type "bye" to exit' stuff.  You know, "ignoreeof". */
//
///* A flag denoting whether or not ignoreeof is set. */
//int ignoreeof = 0;

/* If we have EOF as the only input unit, this user wants to leave
   the shell.  If the shell is not interactive, then just leave.
   Otherwise, if ignoreeof is set, and we haven't done this the
   required number of times in a row, print a message. */
func (gps *ParserState) handle_eof_input_unit() {
  /* We don't write history files, etc., for non-interactive shells. */
  gps.EOF_Reached = true;
}

///************************************************
// *						*
// *	STRING PARSING FUNCTIONS		*
// *						*
// ************************************************/
//
///* It's very important that these two functions treat the characters
//   between ( and ) identically. */
//
//static word_list parse_string_error;
//
///* Take a string and run it through the shell parser, returning the
//   resultant word list.  Used by compound array assignment. */
//word_list *
//parse_string_to_word_list (s, flags, whom)
//     char *s;
//     int flags;
//     const char *whom;
//{
//  word_list *wl;
//  int tok, orig_current_token, orig_line_number, orig_input_terminator;
//  int orig_line_count;
//  int old_echo_input, old_expand_aliases;
//
//  orig_line_number = gps.line_number;
//  orig_line_count = gps.current_command_line_count;
//  orig_input_terminator = gps.shell_input_line_terminator;
//  old_echo_input = gps.echo_input_at_read;
//  old_expand_aliases = expand_aliases;
//
//  push_stream (1);
//  gps.last_read_token = WORD;		/* WORD to allow reserved words here */
//  gps.current_command_line_count = 0;
//  gps.echo_input_at_read = expand_aliases = 0;
//
//  with_input_from_string (s, whom);
//  wl = nil;
//
//  if (flags & 1)
//    gps.parser_state |= PST_COMPASSIGN|PST_REPARSE;
//
//  while ((tok = read_token (READ)) != yacc_EOF)
//    {
//      if (tok == '\n' && *bash_input.location.string == '\0')
//	break;
//      if (tok == '\n')		/* Allow newlines in compound assignments */
//	continue;
//      if (tok != WORD && tok != ASSIGNMENT_WORD)
//	{
//	  gps.line_number = orig_line_number + gps.line_number - 1;
//	  orig_current_token = gps.current_token;
//	  gps.current_token = tok;
//	  yyerror (NULL);	/* does the right thing */
//	  gps.current_token = orig_current_token;
//	  if (wl)
//	    dispose_words (wl);
//	  wl = &parse_string_error;
//	  break;
//	}
//      wl = makeWordList (gps.yylval.word, wl);
//    }
//  
//  gps.last_read_token = '\n';
//  pop_stream ();
//
//  gps.echo_input_at_read = old_echo_input;
//  expand_aliases = old_expand_aliases;
//
//  gps.current_command_line_count = orig_line_count;
//  gps.shell_input_line_terminator = orig_input_terminator;
//
//  if (flags & 1)
//    gps.parser_state &= ^(PST_COMPASSIGN|PST_REPARSE);
//
//  if (wl == &parse_string_error)
//    {
//      last_command_exit_value = EXECUTION_FAILURE;
//      if (gps.posixly_correct)
//	jump_to_top_level (FORCE_EOF);
//      else
//	jump_to_top_level (DISCARD);
//    }
//
//  return reverseWordList(wl);
//}

func parse_compound_assignment() *StringBuilder {
  // TODO(krasin): implement this
  panic("parse_compound_assignment: not implemented")
}

//static char *
//parse_compound_assignment (retlenp)
//     int *retlenp;
//{
//  word_list *wl, *rl;
//  int tok, orig_line_number, orig_token_size, orig_last_token, assignok;
//  char *saved_token, *ret;
//
//  saved_token = token;
//  orig_token_size = token_buffer_size;
//  orig_line_number = gps.line_number;
//  orig_last_token = gps.last_read_token;
//
//  gps.last_read_token = WORD;	/* WORD to allow reserved words here */
//
//  token = nil;
//  token_buffer_size = 0;
//
//  assignok = gps.parser_state&PST_ASSIGNOK;		/* XXX */
//
//  wl = nil;	/* ( */
//  gps.parser_state |= PST_COMPASSIGN;
//
//  while ((tok = read_token (READ)) != ')')
//    {
//      if (tok == '\n')			/* Allow newlines in compound assignments */
//	{
//	  continue;
//	}
//      if (tok != WORD && tok != ASSIGNMENT_WORD)
//	{
//	  gps.current_token = tok;	/* for error reporting */
//	  if (tok == yacc_EOF)	/* ( */
//	    gps.parser_error (orig_line_number, _("unexpected EOF while looking for matching `)'"));
//	  else
//	    yyerror(NULL);	/* does the right thing */
//	  if (wl)
//	    dispose_words (wl);
//	  wl = &parse_string_error;
//	  break;
//	}
//      wl = makeWordList (gps.yylval.word, wl);
//    }
//
//  token = saved_token;
//  token_buffer_size = orig_token_size;
//
//  gps.parser_state &= ^PST_COMPASSIGN;
//
//  if (wl == &parse_string_error)
//    {
//      last_command_exit_value = EXECUTION_FAILURE;
//      gps.last_read_token = '\n';	/* XXX */
//      if (gps.posixly_correct)
//	jump_to_top_level (FORCE_EOF);
//      else
//	jump_to_top_level (DISCARD);
//    }
//
//  gps.last_read_token = orig_last_token;		/* XXX - was WORD? */
//
//  if (wl != nil)
//    {
//      rl = reverseWordList(wl)
//      ret = string_list (rl);
//      dispose_words (rl);
//    }
//  else
//    ret = nil;
//
//  if (retlenp)
//    *retlenp = (ret && *ret) ? strlen (ret) : 0;
//
//  if (assignok)
//    gps.parser_state |= PST_ASSIGNOK;
//
//  return ret;
//}
//
///************************************************
// *						*
// *   SAVING AND RESTORING PARTIAL PARSE STATE   *
// *						*
// ************************************************/
//
//sh_parser_state_t *
//save_parser_state (ps)
//     sh_parser_state_t *ps;
//{
//  SHELL_VAR *v;
//
//  if (ps == 0)
//    ps = (sh_parser_state_t *)xmalloc (sizeof (sh_parser_state_t));
//  if (ps == 0)
//    return (nil);
//
//  ps.gps.parser_state = gps.parser_state;
//  ps.token_state = save_token_state ();
//
//  ps.input_line_terminator = gps.shell_input_line_terminator;
//
//  ps.current_command_line_count = gps.current_command_line_count;
//
//  ps.last_command_exit_value = last_command_exit_value;
//  v = find_variable ("PIPESTATUS");
//  if (v && array_p (v) && array_cell (v))
//    ps.pipestatus = array_copy (array_cell (v));
//  else
//    ps.pipestatus = nil;
//
//  ps.last_shell_builtin = last_shell_builtin;
//  ps.this_shell_builtin = this_shell_builtin;
//
//  ps.expand_aliases = expand_aliases;
//  ps.echo_input_at_read = gps.echo_input_at_read;
//
//  return (ps);
//}
//
//void
//restore_parser_state (ps)
//     sh_parser_state_t *ps;
//{
//  SHELL_VAR *v;
//
//  if (ps == 0)
//    return;
//
//  gps.parser_state = ps.gps.parser_state;
//  if (ps.token_state)
//    {
//      restore_token_state (ps.token_state);
//    }
//
//  gps.shell_input_line_terminator = ps.input_line_terminator;
//
//  gps.current_command_line_count = ps.current_command_line_count;
//
//  last_command_exit_value = ps.last_command_exit_value;
//  v = find_variable ("PIPESTATUS");
//  if (v && array_p (v) && array_cell (v))
//    {
//      array_dispose (array_cell (v));
//      var_setarray (v, ps.pipestatus);
//    }
//
//  last_shell_builtin = ps.last_shell_builtin;
//  this_shell_builtin = ps.this_shell_builtin;
//
//  expand_aliases = ps.expand_aliases;
//  gps.echo_input_at_read = ps.echo_input_at_read;
//}


