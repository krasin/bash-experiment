package gobash
/* eval.c -- reading and evaluating commands. */

/* Copyright (C) 1996-2009 Free Software Foundation, Inc.

   This file is part of GNU Bash, the Bourne Again SHell.

   Bash is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published 
   by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

   Bash is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for more details.

   You should have received a copy of the GNU General Public License along with Bash.  If not, see
   <http://www.gnu.org/licenses/>. */

/* Read and execute commands until EOF is reached.  This assumes that the input source has already been initialized. */
int reader_loop() {
	int our_indirection_level;
	COMMAND *volatile current_command;

	USE_VAR(current_command);

	current_command = (COMMAND *) NULL;

	our_indirection_level = ++indirection_level;

	while (EOF_Reached == 0) {
		int code;

		code = setjmp(top_level);

		if (code != NOT_JUMPED) {
			indirection_level = our_indirection_level;

			switch (code) {
				/* Some kind of throw to top_level has occured. */
			case FORCE_EOF:
			case ERREXIT:
			case EXITPROG:
				current_command = (COMMAND *) NULL;
				if (exit_immediately_on_error)
					variable_context = 0;	/* not in a function */
				EOF_Reached = EOF;
				goto exec_done;

			case DISCARD:
				/* Make sure the exit status is reset to a non-zero value, but leave existing non-zero values (e.g., > 128 on
				   signal) alone. */
				if (last_command_exit_value == 0)
					last_command_exit_value = EXECUTION_FAILURE;
				if (subshell_environment) {
					current_command = (COMMAND *) NULL;
					EOF_Reached = EOF;
					goto exec_done;
				}
				/* Obstack free command elements, etc. */
				if (current_command) {
					dispose_command(current_command);
					current_command = (COMMAND *) NULL;
				}
				break;

			default:
				command_error("reader_loop", CMDERR_BADJUMP, code, 0);
			}
		}

		executing = 0;
		if (temporary_env)
			dispose_used_env_vars();


		if (read_command() == 0) {
			if (interactive_shell == 0 && read_but_dont_execute) {
				last_command_exit_value = EXECUTION_SUCCESS;
				dispose_command(global_command);
				global_command = (COMMAND *) NULL;
			} else if (current_command = global_command) {
				global_command = (COMMAND *) NULL;
				current_command_number++;

				executing = 1;
				stdin_redir = 0;
				execute_command(current_command);

			  exec_done:
				QUIT;

				if (current_command) {
					dispose_command(current_command);
					current_command = (COMMAND *) NULL;
				}
			}
		} else {
			/* Parse error, maybe discard rest of stream if not interactive. */
			if (interactive == 0)
				EOF_Reached = EOF;
		}
		if (just_one_command)
			EOF_Reached = EOF;
	}
	indirection_level--;
	return (last_command_exit_value);
}

/* Call the YACC-generated parser and return the status of the parse. Input is read from the current input stream (bash_input).
   yyparse leaves the parsed command in the global variable GLOBAL_COMMAND. This is where PROMPT_COMMAND is executed. */
int parse_command() {
	int r;
	char *command_to_execute;

	need_here_doc = 0;

	current_command_line_count = 0;
	r = yyparse();

	if (need_here_doc)
		gather_here_documents();

	return (r);
}

/* Read and parse a command, returning the status of the parse.  The command is left in the globval variable GLOBAL_COMMAND for
   use by reader_loop. This is where the shell timeout code is executed. */
int read_command() {
	SHELL_VAR *tmout_var;
	int tmout_len, result;

	set_current_prompt_level(1);
	global_command = (COMMAND *) NULL;

	/* Only do timeouts if interactive. */
	tmout_var = (SHELL_VAR *) NULL;
	tmout_len = 0;
	old_alrm = (SigHandler *) NULL;

	current_command_line_count = 0;
	result = parse_command();

	return (result);
}
