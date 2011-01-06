package gobash
/* parser.h -- Everything you wanted to know about the parser, but were
   afraid to ask. */

/* Copyright (C) 1995, 2008,2009 Free Software Foundation, Inc.

   This file is part of GNU Bash, the Bourne Again SHell.

   Bash is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   Bash is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with Bash.  If not, see <http://www.gnu.org/licenses/>.
*/

/* Possible states for the parser that require it to do special things. */
const PST_CASEPAT = 0x000001 /* in a case pattern list */
const PST_ALEXPNEXT = 0x000002 /* expand next word for aliases */
const PST_ALLOWOPNBRC = 0x000004 /* allow open brace for function def */
const PST_NEEDCLOSBRC = 0x000008 /* need close brace */
const PST_DBLPAREN = 0x000010 /* double-paren parsing */
const PST_SUBSHELL = 0x000020 /* ( ... ) subshell */
const PST_CMDSUBST = 0x000040 /* $( ... ) command substitution */
const PST_CASESTMT = 0x000080 /* parsing a case statement */
const PST_CONDCMD = 0x000100 /* parsing a [[...]] command */
const PST_CONDEXPR = 0x000200 /* parsing the guts of [[...]] */
const PST_ARITHFOR = 0x000400 /* parsing an arithmetic for command */
const PST_ALEXPAND = 0x000800 /* OK to expand aliases - unused */
const PST_EXTPAT = 0x001000 /* parsing an extended shell pattern */
const PST_COMPASSIGN = 0x002000 /* parsing x=(...) compound assignment */
const PST_ASSIGNOK = 0x004000 /* assignment statement ok in this context */
const PST_EOFTOKEN = 0x008000 /* yylex checks against shell_eof_token */
const PST_REGEXP = 0x010000 /* parsing an ERE/BRE as a single word */
const PST_HEREDOC = 0x020000 /* reading body of here-document */
const PST_REPARSE = 0x040000 /* re-parsing in parse_string_to_word_list */
const PST_REDIRLIST = 0x080000 /* parsing a list of redirctions preceding a simple command name */


/* Definition of the delimiter stack.  Needed by parse.y and bashhist.c. */
type dstack struct {
/* DELIMITERS is a stack of the nested delimiters that we have
   encountered so far. */
//  char *delimiters; // TODO(krasin): decide what to do with dstack.

/* Offset into the stack of delimiters. */
  delimiter_depth int

/* How many slots are allocated to DELIMITERS. */
  delimiter_space int
}

