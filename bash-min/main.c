/* shell.c -- GNU's idea of the POSIX shell specification. */

/* Copyright (C) 1987-2009 Free Software Foundation, Inc.

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

/*
  Birthdate:
  Sunday, January 10th, 1988.
  Initial author: Brian Fox
*/
#define INSTALL_DEBUG_MODE

#include "config.h"

#include "bashtypes.h"
#if !defined (_MINIX) && defined (HAVE_SYS_FILE_H)
#  include <sys/file.h>
#endif
#include "posixstat.h"
#include "posixtime.h"
#include "bashansi.h"
#include <stdio.h>
#include <signal.h>
#include <errno.h>
#include "filecntl.h"
#include <pwd.h>

#if defined (HAVE_UNISTD_H)
#  include <unistd.h>
#endif

#include "bashintl.h"

#define NEED_SH_SETLINEBUF_DECL		/* used in externs.h */

#include "shell.h"
#include "flags.h"
#include "trap.h"
#include "mailcheck.h"
#include "builtins.h"
#include "builtins/common.h"

#if defined (JOB_CONTROL)
#include "jobs.h"
#endif /* JOB_CONTROL */

#include "input.h"
#include "execute_cmd.h"
#include "findcmd.h"

#if defined (USING_BASH_MALLOC) && defined (DEBUG) && !defined (DISABLE_MALLOC_WRAPPERS)
#  include <malloc/shmalloc.h>
#endif

#if defined (HISTORY)
#  include "bashhist.h"
#  include <readline/history.h>
#endif

#if defined (READLINE)
#  include "bashline.h"
#endif

#include <tilde/tilde.h>
#include <glob/strmatch.h>

#if defined (__OPENNT)
#  include <opennt/opennt.h>
#endif

#if !defined (HAVE_GETPW_DECLS)
extern struct passwd *getpwuid ();
#endif /* !HAVE_GETPW_DECLS */

#if !defined (errno)
extern int errno;
#endif

#if defined (NO_MAIN_ENV_ARG)
extern char **environ;	/* used if no third argument to main() */
#endif

extern char *dist_version, *release_status;
extern int patch_level, build_version;
extern int shell_level;
extern int subshell_environment;
extern int last_command_exit_value;
extern int line_number;
extern int expand_aliases;
extern int array_needs_making;
extern int gnu_error_format;
extern char *primary_prompt, *secondary_prompt;
extern char *this_command_name;

/* Non-zero means that this shell has already been run; i.e. you should
   call shell_reinitialize () if you need to start afresh. */
int shell_initialized = 0;

COMMAND *global_command = (COMMAND *)NULL;

/* Information about the current user. */
struct user_info current_user =
{
  (uid_t)-1, (uid_t)-1, (gid_t)-1, (gid_t)-1,
  (char *)NULL, (char *)NULL, (char *)NULL
};

/* The current host's name. */
char *current_host_name = (char *)NULL;

/* Non-zero means that this shell is a login shell.
   Specifically:
   0 = not login shell.
   1 = login shell from getty (or equivalent fake out)
  -1 = login shell from "--login" (or -l) flag.
  -2 = both from getty, and from flag.
 */
int login_shell = 0;

/* Non-zero means that at this moment, the shell is interactive.  In
   general, this means that the shell is at this moment reading input
   from the keyboard. */
int interactive = 0;

/* Non-zero means that the shell was started as an interactive shell. */
int interactive_shell = 0;

/* Non-zero means to send a SIGHUP to all jobs when an interactive login
   shell exits. */
int hup_on_exit = 0;

/* Non-zero means to list status of running and stopped jobs at shell exit */
int check_jobs_at_exit = 0;

/* Non-zero means to change to a directory name supplied as a command name */
int autocd = 0;

/* Tells what state the shell was in when it started:
	0 = non-interactive shell script
	1 = interactive
	2 = -c command
	3 = wordexp evaluation
   This is a superset of the information provided by interactive_shell.
*/
int startup_state = 0;

/* Special debugging helper. */
int debugging_login_shell = 0;

/* The environment that the shell passes to other commands. */
char **shell_environment;

/* Non-zero when we are executing a top-level command. */
int executing = 0;

/* The number of commands executed so far. */
int current_command_number = 1;

/* Non-zero is the recursion depth for commands. */
int indirection_level = 0;

/* The name of this shell, as taken from argv[0]. */
char *shell_name = (char *)NULL;

/* time in seconds when the shell was started */
time_t shell_start_time;

/* Are we running in an emacs shell window? */
int running_under_emacs;

/* Do we have /dev/fd? */
#ifdef HAVE_DEV_FD
int have_devfd = HAVE_DEV_FD;
#else
int have_devfd = 0;
#endif

/* The name of the .(shell)rc file. */
static char *bashrc_file = "~/.bashrc";

/* Non-zero means to act more like the Bourne shell on startup. */
static int act_like_sh;

/* Non-zero if this shell is being run by `su'. */
static int su_shell;

/* Non-zero if we have already expanded and sourced $ENV. */
static int sourced_env;

/* Is this shell running setuid? */
static int running_setuid;

/* Values for the long-winded argument names. */
static int debugging;			/* Do debugging things. */
static int no_rc;			/* Don't execute ~/.bashrc */
static int no_profile;			/* Don't execute .profile */
static int do_version;			/* Display interesting version info. */
static int make_login_shell;		/* Make this shell be a `-bash' shell. */
static int want_initial_help;		/* --help option */

int debugging_mode = 0;		/* In debugging mode with --debugger */
int no_line_editing = 0;	/* Don't do fancy line editing. */
int dump_translatable_strings;	/* Dump strings in $"...", don't execute. */
int dump_po_strings;		/* Dump strings in $"..." in po format */
int wordexp_only = 0;		/* Do word expansion only */
int protected_mode = 0;		/* No command substitution with --wordexp */

#if defined (STRICT_POSIX)
int posixly_correct = 1;	/* Non-zero means posix.2 superset. */
#else
int posixly_correct = 0;	/* Non-zero means posix.2 superset. */
#endif


/* These are extern so execute_simple_command can set them, and then
   longjmp back to main to execute a shell script, instead of calling
   main () again and resulting in indefinite, possibly fatal, stack
   growth. */
procenv_t subshell_top_level;
int subshell_argc;
char **subshell_argv;
char **subshell_envp;

#if defined (BUFFERED_INPUT)
/* The file descriptor from which the shell is reading input. */
int default_buffered_input = -1;
#endif

/* The following two variables are not static so they can show up in $-. */
int read_from_stdin;		/* -s flag supplied */
int want_pending_command;	/* -c flag supplied */

/* This variable is not static so it can be bound to $BASH_EXECUTION_STRING */
char *command_execution_string;	/* argument to -c option */

int malloc_trace_at_exit = 0;

static int shell_reinitialized = 0;

static FILE *default_input;

static STRING_INT_ALIST *shopt_alist;
static int shopt_ind = 0, shopt_len = 0;

int main(void) {
  printf("laka\n");
}
