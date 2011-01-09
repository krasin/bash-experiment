package gobash
/* error.c -- Functions for handling errors. */

/* Copyright (C) 1993-2009 Free Software Foundation, Inc.

   This file is part of GNU Bash, the Bourne Again SHell.

   Bash is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version
   3 of the License, or (at your option) any later version.

   Bash is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
   See the GNU General Public License for more details.

   You should have received a copy of the GNU General Public License along with Bash.  If not, see <http://www.gnu.org/licenses/>. */

import (
	"fmt"
	"log"
	"os"
)

//extern int executing_line_number __P((void));
//
//extern int last_command_exit_value;
//extern char *shell_name;
//extern const char *const bash_badsub_errmsg;
//
//static void error_prolog __P((int));
//
///* The current maintainer of the shell.  You change this in the Makefile. */
//#if !defined (MAINTAINER)
//#define MAINTAINER "bash-maintainers@gnu.org"
//#endif
//
//const char *const the_current_maintainer = MAINTAINER;
//
//int gnu_error_format = 0;
//
//static void error_prolog(print_lineno)
//	 int print_lineno;
//{
//	char *ename;
//	int line;
//
//	ename = get_name_for_error();
//	line = (print_lineno && interactive_shell == 0) ? executing_line_number() : -1;
//
//	if (line > 0)
//		fmt.Fprintf(os.Stderr, "%s:%s%d: ", ename, gnu_error_format ? "" : _(" line "), line);
//	else
//		fmt.Fprintf(os.Stderr, "%s: ", ename);
//}
//
///* Return the name of the shell or the shell script for error reporting. */
//char *get_name_for_error() {
//	char *name;
//#if defined (ARRAY_VARS)
//	SHELL_VAR *bash_source_v;
//	ARRAY *bash_source_a;
//#endif
//
//	name = (char *)NULL;
//	if (interactive_shell == 0) {
//#if defined (ARRAY_VARS)
//		bash_source_v = find_variable("BASH_SOURCE");
//		if (bash_source_v && array_p(bash_source_v) && (bash_source_a = array_cell(bash_source_v)))
//			name = array_reference(bash_source_a, 0);
//		if (name == 0 || *name == '\0')	/* XXX - was just name == 0 */
//#endif
//			name = dollar_vars[0];
//	}
//	if (name == 0 && shell_name && *shell_name)
//		name = base_pathname(shell_name);
//	if (name == 0)
//#if defined (PROGRAM)
//		name = PROGRAM;
//#else
//		name = "bash";
//#endif
//
//	return (name);
//}
//
///* Report an error having to do with FILENAME.  This does not use sys_error so the filename is not interpreted as a printf-style format string. */
//void file_error(filename)
//	 const char *filename;
//{
//	report_error("%s: %s", filename, strerror(errno));
//}

func programming_error(fmt string, v ...interface{}) {
	log.Panicf(fmt + "\n", v...)
}

///* Print an error message and, if `set -e' has been executed, exit the shell.  Used in this file by file_error and programming_error.  Used outside this file mostly to report
//   substitution and expansion errors, and for bad invocation options. */
//void
//#if defined (PREFER_STDARG)
//report_error(const char *format, ...)
//#else
//report_error(format, va_alist)
//	 const char *format;
//	 va_dcl
//#endif
//{
//	va_list args;
//
//	error_prolog(1);
//
//	SH_VA_START(args, format);
//
//	fmt.Fprintf(os.Stderr, format, args);
//	fmt.Fprintf(os.Stderr, "\n");
//
//	va_end(args);
//	if (exit_immediately_on_error)
//		exit_shell(1);
//}
//
//void
//#if defined (PREFER_STDARG)
//fatal_error(const char *format, ...)
//#else
//fatal_error(format, va_alist)
//	 const char *format;
//	 va_dcl
//#endif
//{
//	va_list args;
//
//	error_prolog(0);
//
//	SH_VA_START(args, format);
//
//	fmt.Fprintf(os.Stderr, format, args);
//	fmt.Fprintf(os.Stderr, "\n");
//
//	va_end(args);
//	sh_exit(2);
//}
//
//void
//#if defined (PREFER_STDARG)
//internal_error(const char *format, ...)
//#else
//internal_error(format, va_alist)
//	 const char *format;
//	 va_dcl
//#endif
//{
//	va_list args;
//
//	error_prolog(1);
//
//	SH_VA_START(args, format);
//
//	fmt.Fprintf(os.Stderr, format, args);
//	fmt.Fprintf(os.Stderr, "\n");
//
//	va_end(args);
//}
//
//void
//#if defined (PREFER_STDARG)
//internal_warning(const char *format, ...)
//#else
//internal_warning(format, va_alist)
//	 const char *format;
//	 va_dcl
//#endif
//{
//	va_list args;
//
//	error_prolog(1);
//	fmt.Fprintf(os.Stderr, _("warning: "));
//
//	SH_VA_START(args, format);
//
//	fmt.Fprintf(os.Stderr, format, args);
//	fmt.Fprintf(os.Stderr, "\n");
//
//	va_end(args);
//}
//
//void
//#if defined (PREFER_STDARG)
//sys_error(const char *format, ...)
//#else
//sys_error(format, va_alist)
//	 const char *format;
//	 va_dcl
//#endif
//{
//	int e;
//	va_list args;
//
//	e = errno;
//	error_prolog(0);
//
//	SH_VA_START(args, format);
//
//	fmt.Fprintf(os.Stderr, format, args);
//	fmt.Fprintf(os.Stderr, ": %s\n", strerror(e));
//
//	va_end(args);
//}
//
/* An error from the parser takes the general form

   shell_name: input file name: line number: message

   The input file name and line number are omitted if the shell is currently interactive.  If the shell is not currently interactive, the input file name is inserted only if it
   is different from the shell name. */
func (gps *ParserState) parser_error(lineno int, format string, args ...interface{}) {
	// TODO(krasin): use get_name_for_error and yy_input_name
	ename := "(NOT IMPLEMENTED)NAME_FOR_ERROR" // get_name_for_error();
	iname := "(NOT IMPLEMENTED)INPUT_NAME" // yy_input_name();

	if ename == iname {
		fmt.Fprintf(os.Stderr, "%s:%d: ", ename, lineno);
	} else {
		fmt.Fprintf(os.Stderr, "%s: %s:%d: ", ename, iname, lineno);
        }

	fmt.Fprintf(os.Stderr, format, args...);
	fmt.Fprintf(os.Stderr, "\n");

        // TODO: use option exit_immediately_on_error here
	if true { //(exit_immediately_on_error) {
		os.Exit(2)
		//TODO(krasin): use exit_shell
		//exit_shell(last_command_exit_value = 2);
        }
}

//#ifdef DEBUG
//void
//#if defined (PREFER_STDARG)
//itrace(const char *format, ...)
//#else
//itrace(format, va_alist)
//	 const char *format;
//	 va_dcl
//#endif
//{
//	va_list args;
//
//	fmt.Fprintf(os.Stderr, "TRACE: pid %ld: ", (long)getpid());
//
//	SH_VA_START(args, format);
//
//	fmt.Fprintf(os.Stderr, format, args);
//	fmt.Fprintf(os.Stderr, "\n");
//
//	va_end(args);
//
//	fflush(stderr);
//}
//
///* A trace function for silent debugging -- doesn't require a control terminal. */
//void
//#if defined (PREFER_STDARG)
//trace(const char *format, ...)
//#else
//trace(format, va_alist)
//	 const char *format;
//	 va_dcl
//#endif
//{
//	va_list args;
//	static FILE *tracefp = (FILE *) NULL;
//
//	if (tracefp == NULL)
//		tracefp = fopen("/tmp/bash-trace.log", "a+");
//
//	if (tracefp == NULL)
//		tracefp = stderr;
//	else
//		fcntl(fileno(tracefp), F_SETFD, 1);	/* close-on-exec */
//
//	fprintf(tracefp, "TRACE: pid %ld: ", (long)getpid());
//
//	SH_VA_START(args, format);
//
//	vfprintf(tracefp, format, args);
//	fprintf(tracefp, "\n");
//
//	va_end(args);
//
//	fflush(tracefp);
//}
//
//#endif /* DEBUG */

/* **************************************************************** */
/* */
/* Common error reporting */
/* */
/* **************************************************************** */


var cmd_error_table = []string {
	"unknown command error",	/* CMDERR_DEFAULT */
	"bad command type",		/* CMDERR_BADTYPE */
	"bad connector",		/* CMDERR_BADCONN */
	"bad jump",			/* CMDERR_BADJUMP */
	"",
}

func command_error(fun string, code int, e command_type, flags int) {
	if (code > CMDERR_LAST) {
		code = CMDERR_DEFAULT;
	}

	programming_error("%s: %s: %d", fun, cmd_error_table[code], e);
}
//
//char *command_errstr(code)
//	 int code;
//{
//	if (code > CMDERR_LAST)
//		code = CMDERR_DEFAULT;
//
//	return (_(cmd_error_table[code]));
//}
//
//#ifdef ARRAY_VARS
//void err_badarraysub(s)
//	 const char *s;
//{
//	report_error("%s: %s", s, _(bash_badsub_errmsg));
//}
//#endif
//
//void err_unboundvar(s)
//	 const char *s;
//{
//	report_error(_("%s: unbound variable"), s);
//}
//
//void err_readonly(s)
//	 const char *s;
//{
//	report_error(_("%s: readonly variable"), s);
//}
