package gobash
/* builtins.h -- What a builtin looks like, and where to find them. */

/* Copyright (C) 1987-2009 Free Software Foundation, Inc.

   This file is part of GNU Bash, the Bourne Again SHell.

   Bash is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published 
   by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

   Bash is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

   You should have received a copy of the GNU General Public License along with Bash.  If not, see
   <http://www.gnu.org/licenses/>. */

/* Flags describing various things about a builtin. */
const BUILTIN_ENABLED = 0x01 /* This builtin is enabled. */
const BUILTIN_DELETED = 0x02 /* This has been deleted with enable -d. */
const STATIC_BUILTIN = 0x04 /* This builtin is not dynamically loaded. */
const SPECIAL_BUILTIN = 0x08 /* This is a Posix `special' builtin. */
const ASSIGNMENT_BUILTIN = 0x10 /* This builtin takes assignment statements. */
const POSIX_BUILTIN = 0x20 /* This builtins is special in the Posix command search order. */

const BASE_INDENT = 4

type sh_builtin_func_t func(*word_list) int

/* The thing that we build the array of builtins out of. */
type builtin struct {
	name string /* The name that the user types. */
	function *sh_builtin_func_t /* The address of the invoked function. */
	flags int /* One of the consts  above. */
	long_doc []string; /* Array of strings. */
	short_doc string /* Short version of documentaion. */
}

