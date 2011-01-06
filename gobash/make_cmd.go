package gobash

/* make_cmd.c -- Functions for making instances of the various
   parser constructs. */

/* Copyright (C) 1989-2009 Free Software Foundation, Inc.

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

import (
	"fmt"
	"os"
)

//extern int line_number, current_command_line_count, parser_state;
//extern int last_command_exit_value;
//
///* Object caching */
//sh_obj_cache_t wdcache = {0, 0, 0};
//sh_obj_cache_t wlcache = {0, 0, 0};
//
//#define WDCACHESIZE	60
//#define WLCACHESIZE	60
//
//static Command *make_for_or_select __P((enum command_type, word_desc *, word_list *, Command *, int));
//#if defined (ArithForCommand)
//static word_list *make_arith_for_expr __P((char *));
//#endif
//static Command *make_until_or_while __P((enum command_type, Command *, Command *));
//
//void
//cmd_init ()
//{
//  ocache_create (wdcache, word_desc, WDCACHESIZE);
//  ocache_create (wlcache, word_list, WLCACHESIZE);
//}
//
//word_desc *
//alloc_word_desc ()
//{
//  word_desc *temp;
//
//  ocache_alloc (wdcache, word_desc, temp);
//  temp.flags = 0;
//  temp.word = 0;
//  return temp;
//}
//
//word_desc *
//make_bare_word (string)
//     const char *string;
//{
//  word_desc *temp;
//
//  temp = alloc_word_desc ();
//
//  if (*string)
//    temp.word = savestring (string);
//  else
//    {
//      temp.word = (char *)xmalloc (1);
//      temp.word[0] = '\0';
//    }
//
//  return (temp);
//}
//
//word_desc *
//make_word_flags (w, string)
//     word_desc *w;
//     const char *string;
//{
//  register int i;
//  size_t slen;
//  DECLARE_MBSTATE;
//
//  i = 0;
//  slen = strlen (string);
//  while (i < slen)
//    {
//      switch (string[i])
//	{
//	case '$':
//	  w.flags |= W_HASDOLLAR;
//	  break;
//	case '\\':
//	  break;	/* continue the loop */
//	case '\'':
//	case '`':
//	case '"':
//	  w.flags |= W_QUOTED;
//	  break;
//	}
//
//      ADVANCE_CHAR (string, slen, i);
//    }
//
//  return (w);
//}
//
//word_desc *
//make_word (string)
//     const char *string;
//{
//  word_desc *temp;
//
//  temp = make_bare_word (string);
//  return (make_word_flags (temp, string));
//}
//
//word_desc *
//make_word_from_token (token)
//     int token;
//{
//  char tokenizer[2];
//
//  tokenizer[0] = token;
//  tokenizer[1] = '\0';
//
//  return (make_word (tokenizer));
//}
//
//word_list *
//make_word_list (word, wlink)
//     word_desc *word;
//     word_list *wlink;
//{
//  word_list *temp;
//
//  ocache_alloc (wlcache, word_list, temp);
//
//  temp.word = word;
//  temp.next = wlink;
//  return (temp);
//}
//
//Command *
//make_command (type, pointer)
//     enum command_type type;
//     SimpleCom *pointer;
//{
//  Command *temp;
//
//  temp = (Command *)xmalloc (sizeof (Command));
//  temp.type = type;
//  temp.value.Simple = pointer;
//  temp.value.Simple.flags = temp.flags = 0;
//  temp.redirects = nil;
//  return (temp);
//}
//
//Command *
//command_connect (com1, com2, connector)
//     Command *com1, *com2;
//     int connector;
//{
//  CONNECTION *temp;
//
//  temp = (CONNECTION *)xmalloc (sizeof (CONNECTION));
//  temp.connector = connector;
//  temp.first = com1;
//  temp.second = com2;
//  return (make_command (cm_connection, (SimpleCom *)temp));
//}
//
//static Command *
//make_for_or_select (type, name, map_list, action, lineno)
//     enum command_type type;
//     word_desc *name;
//     word_list *map_list;
//     Command *action;
//     int lineno;
//{
//  ForCom *temp;
//
//  temp = (ForCom *)xmalloc (sizeof (ForCom));
//  temp.flags = 0;
//  temp.name = name;
//  temp.line = lineno;
//  temp.map_list = map_list;
//  temp.action = action;
//  return (make_command (type, (SimpleCom *)temp));
//}
//
//Command *
//make_for_command (name, map_list, action, lineno)
//     word_desc *name;
//     word_list *map_list;
//     Command *action;
//     int lineno;
//{
//  return (make_for_or_select (cm_for, name, map_list, action, lineno));
//}
//
//Command *
//make_select_command (name, map_list, action, lineno)
//     word_desc *name;
//     word_list *map_list;
//     Command *action;
//     int lineno;
//{
//#if defined (SelectCommand)
//  return (make_for_or_select (cm_select, name, map_list, action, lineno));
//#else
//  last_command_exit_value = 2;
//  return (nil);
//#endif
//}
//
//#if defined (ArithForCommand)
//static word_list *
//make_arith_for_expr (s)
//     char *s;
//{
//  word_list *result;
//  word_desc *wd;
//
//  if (s == 0 || *s == '\0')
//    return ((word_list *)NULL);
//  wd = make_word (s);
//  wd.flags |= W_NOGLOB|W_NOSPLIT|W_QUOTED|W_DQUOTE;	/* no word splitting or globbing */
//  result = make_word_list (wd, (word_list *)NULL);
//  return result;
//}
//#endif
//
///* Note that this function calls dispose_words on EXPRS, since it doesn't
//   use the word list directly.  We free it here rather than at the caller
//   because no other function in this file requires that the caller free
//   any arguments. */
//Command *
//make_arith_for_command (exprs, action, lineno)
//     word_list *exprs;
//     Command *action;
//     int lineno;
//{
//#if defined (ArithForCommand)
//  ArithForCOM *temp;
//  word_list *init, *test, *step;
//  char *s, *t, *start;
//  int nsemi;
//
//  init = test = step = (word_list *)NULL;
//  /* Parse the string into the three component sub-expressions. */
//  start = t = s = exprs.word.word;
//  for (nsemi = 0; ;)
//    {
//      /* skip whitespace at the start of each sub-expression. */
//      while (whitespace (*s))
//	s++;
//      start = s;
//      /* skip to the semicolon or EOS */
//      while (*s && *s != ';')
//	s++;
//
//      t = (s > start) ? substring (start, 0, s - start) : (char *)NULL;
//
//      nsemi++;
//      switch (nsemi)
//	{
//	case 1:
//	  init = make_arith_for_expr (t);
//	  break;
//	case 2:
//	  test = make_arith_for_expr (t);
//	  break;
//	case 3:
//	  step = make_arith_for_expr (t);
//	  break;
//	}
//
//      FREE (t);
//      if (*s == '\0')
//	break;
//      s++;	/* skip over semicolon */
//    }
//
//  if (nsemi != 3)
//    {
//      if (nsemi < 3)
//	parser_error (lineno, _("syntax error: arithmetic expression required"));
//      else
//	parser_error (lineno, _("syntax error: `;' unexpected"));
//      parser_error (lineno, _("syntax error: `((%s))'"), exprs.word.word);
//      last_command_exit_value = 2;
//      return (nil);
//    }
//
//  temp = (ArithForCOM *)xmalloc (sizeof (ARITH_ForCom));
//  temp.flags = 0;
//  temp.line = lineno;
//  temp.init = init ? init : make_arith_for_expr ("1");
//  temp.test = test ? test : make_arith_for_expr ("1");
//  temp.step = step ? step : make_arith_for_expr ("1");
//  temp.action = action;
//
//  dispose_words (exprs);
//  return (make_command (cm_arith_for, (SimpleCom *)temp));
//#else
//  dispose_words (exprs);
//  last_command_exit_value = 2;
//  return (nil);
//#endif /* ArithForCommand */
//}
//
//Command *
//make_group_command (command)
//     Command *command;
//{
//  GroupCom *temp;
//
//  temp = (GroupCom *)xmalloc (sizeof (GroupCom));
//  temp.command = command;
//  return (make_command (cm_group, (SimpleCom *)temp));
//}
//
//Command *
//make_case_command (word, clauses, lineno)
//     word_desc *word;
//     PatternList *clauses;
//     int lineno;
//{
//  CaseCom *temp;
//
//  temp = (CaseCom *)xmalloc (sizeof (CaseCom));
//  temp.flags = 0;
//  temp.line = lineno;
//  temp.word = word;
//  temp.clauses = REVERSE_LIST (clauses, PatternList *);
//  return (make_command (cm_case, (SimpleCom *)temp));
//}
//
//PatternList *
//make_pattern_list (patterns, action)
//     word_list *patterns;
//     Command *action;
//{
//  PatternList *temp;
//
//  temp = (PatternList *)xmalloc (sizeof (PatternList));
//  temp.patterns = REVERSE_LIST (patterns, word_list *);
//  temp.action = action;
//  temp.next = NULL;
//  temp.flags = 0;
//  return (temp);
//}
//
//Command *
//make_if_command (test, true_case, false_case)
//     Command *test, *true_case, *false_case;
//{
//  IF_COM *temp;
//
//  temp = (IF_COM *)xmalloc (sizeof (IF_COM));
//  temp.flags = 0;
//  temp.test = test;
//  temp.true_case = true_case;
//  temp.false_case = false_case;
//  return (make_command (cm_if, (SimpleCom *)temp));
//}
//
//static Command *
//make_until_or_while (which, test, action)
//     enum command_type which;
//     Command *test, *action;
//{
//  WHILE_COM *temp;
//
//  temp = (WHILE_COM *)xmalloc (sizeof (WHILE_COM));
//  temp.flags = 0;
//  temp.test = test;
//  temp.action = action;
//  return (make_command (which, (SimpleCom *)temp));
//}
//
//Command *
//make_while_command (test, action)
//     Command *test, *action;
//{
//  return (make_until_or_while (cm_while, test, action));
//}
//
//Command *
//make_until_command (test, action)
//     Command *test, *action;
//{
//  return (make_until_or_while (cm_until, test, action));
//}
//
//Command *
//make_arith_command (exp)
//     word_list *exp;
//{
//#if defined (DPAREN_ARITHMETIC)
//  Command *command;
//  ARITH_COM *temp;
//
//  command = (Command *)xmalloc (sizeof (Command));
//  command.value.Arith = temp = (ARITH_COM *)xmalloc (sizeof (ARITH_COM));
//
//  temp.flags = 0;
//  temp.line = line_number;
//  temp.exp = exp;
//
//  command.type = cm_arith;
//  command.redirects = nil;
//  command.flags = 0;
//
//  return (command);
//#else
//  last_command_exit_value = 2;
//  return (nil);
//#endif
//}
//
//#if defined (COND_Command)
//struct cond_com *
//make_cond_node (type, op, left, right)
//     int type;
//     word_desc *op;
//     struct cond_com *left, *right;
//{
//  COND_COM *temp;
//
//  temp = (COND_COM *)xmalloc (sizeof (COND_COM));
//  temp.flags = 0;
//  temp.line = line_number;
//  temp.type = type;
//  temp.op = op;
//  temp.left = left;
//  temp.right = right;
//
//  return (temp);
//}
//#endif
//
//Command *
//make_cond_command (cond_node)
//     COND_COM *cond_node;
//{
//#if defined (COND_Command)
//  Command *command;
//
//  command = (Command *)xmalloc (sizeof (Command));
//  command.value.Cond = cond_node;
//
//  command.type = cm_cond;
//  command.redirects = nil;
//  command.flags = 0;
//  command.line = cond_node ? cond_node.line : 0;
//
//  return (command);
//#else
//  last_command_exit_value = 2;
//  return (nil);
//#endif
//}
//
//Command *
//make_bare_simple_command ()
//{
//  Command *command;
//  SimpleCom *temp;
//
//  command = (Command *)xmalloc (sizeof (Command));
//  command.value.Simple = temp = (SimpleCom *)xmalloc (sizeof (SimpleCom));
//
//  temp.flags = 0;
//  temp.line = line_number;
//  temp.words = (word_list *)NULL;
//  temp.redirects = nil;
//
//  command.type = cm_simple;
//  command.redirects = nil;
//  command.flags = 0;
//
//  return (command);
//}
//
///* Return a command which is the connection of the word or redirection
//   in ELEMENT, and the command * or NULL in Command. */
//Command *
//make_simple_command (element, command)
//     ELEMENT element;
//     Command *command;
//{
//  /* If we are starting from scratch, then make the initial command
//     structure.  Also note that we have to fill in all the slots, since
//     malloc doesn't return zeroed space. */
//  if (command == 0)
//    {
//      command = make_bare_simple_command ();
//      parser_state |= PST_REDIRLIST;
//    }
//
//  if (element.word)
//    {
//      command.value.Simple.words = make_word_list (element.word, command.value.Simple.words);
//      parser_state &= ~PST_REDIRLIST;
//    }
//  else if (element.redirect)
//    {
//      Redirect *r = element.redirect;
//      /* Due to the way <> is implemented, there may be more than a single
//	 redirection in element.redirect.  We just follow the chain as far
//	 as it goes, and hook onto the end. */
//      while (r.next)
//	r = r.next;
//      r.next = command.value.Simple.redirects;
//      command.value.Simple.redirects = element.redirect;
//    }
//
//  return (command);
//}
//
///* Because we are Bourne compatible, we read the input for this
//   << or <<- redirection now, from wherever input is coming from.
//   We store the input read into a word_desc.  Replace the text of
//   the redirectee.word with the new input text.  If <<- is on,
//   then remove leading TABS from each line. */
//void
//make_here_document (temp, lineno)
//     Redirect *temp;
//     int lineno;
//{
//  int kill_leading, redir_len;
//  char *redir_word, *document, *full_line;
//  int document_index, document_size, delim_unquoted;
//
//  if (temp.instruction != r_deblank_reading_until &&
//      temp.instruction != r_reading_until)
//    {
//      internal_error (_("make_here_document: bad instruction type %d"), temp.instruction);
//      return;
//    }
//
//  kill_leading = temp.instruction == r_deblank_reading_until;
//
//  document = (char *)NULL;
//  document_index = document_size = 0;
//
//  /* Quote removal is the only expansion performed on the delimiter
//     for here documents, making it an extremely special case. */
//  redir_word = string_quote_removal (temp.redirectee.filename.word, 0);
//
//  /* redirection_expand will return NULL if the expansion results in
//     multiple words or no words.  Check for that here, and just abort
//     this here document if it does. */
//  if (redir_word)
//    redir_len = strlen (redir_word);
//  else
//    {
//      temp.here_doc_eof = (char *)xmalloc (1);
//      temp.here_doc_eof[0] = '\0';
//      goto document_done;
//    }
//
//  free (temp.redirectee.filename.word);
//  temp.here_doc_eof = redir_word;
//
//  /* Read lines from wherever lines are coming from.
//     For each line read, if kill_leading, then kill the
//     leading tab characters.
//     If the line matches redir_word exactly, then we have
//     manufactured the document.  Otherwise, add the line to the
//     list of lines in the document. */
//
//  /* If the here-document delimiter was quoted, the lines should
//     be read verbatim from the input.  If it was not quoted, we
//     need to perform backslash-quoted newline removal. */
//  delim_unquoted = (temp.redirectee.filename.flags & W_QUOTED) == 0;
//  while (full_line = read_secondary_line (delim_unquoted))
//    {
//      register char *line;
//      int len;
//
//      line = full_line;
//      line_number++;
//
//      /* If set -v is in effect, echo the line read.  read_secondary_line/
//	 read_a_line leaves the newline at the end, so don't print another. */
//      if (echo_input_at_read)
//	fprintf (stderr, "%s", line);
//
//      if (kill_leading && *line)
//	{
//	  /* Hack:  To be compatible with some Bourne shells, we
//	     check the word before stripping the whitespace.  This
//	     is a hack, though. */
//	  if (STREQN (line, redir_word, redir_len) && line[redir_len] == '\n')
//	    goto document_done;
//
//	  while (*line == '\t')
//	    line++;
//	}
//
//      if (*line == 0)
//	continue;
//
//      if (STREQN (line, redir_word, redir_len) && line[redir_len] == '\n')
//	goto document_done;
//
//      len = strlen (line);
//      if (len + document_index >= document_size)
//	{
//	  document_size = document_size ? 2 * (document_size + len) : len + 2;
//	  document = (char *)xrealloc (document, document_size);
//	}
//
//      /* len is guaranteed to be > 0 because of the check for line
//	 being an empty string before the call to strlen. */
//      FASTCOPY (line, document + document_index, len);
//      document_index += len;
//    }
//
//  if (full_line == 0)
//    internal_warning (_("here-document at line %d delimited by end-of-file (wanted `%s')"), lineno, redir_word);
//
//document_done:
//  if (document)
//    document[document_index] = '\0';
//  else
//    {
//      document = (char *)xmalloc (1);
//      document[0] = '\0';
//    }
//  temp.redirectee.filename.word = document;
//}
//
/* Generate a Redirect from SOURCE, DEST, and INSTRUCTION.
   INSTRUCTION is the instruction type, SOURCE is a file descriptor,
   and DEST is a file descriptor or a word_desc *. */
func makeRedirection(source Redirectee, instruction r_instruction, dest_and_filename Redirectee, flags int) *Redirect {

	temp := new(Redirect)

  /* First do the common cases. */
  temp.redirector = source;
  temp.redirectee = dest_and_filename;
  temp.instruction = instruction;
  temp.flags = 0;
  temp.rflags = flags;
  temp.next = nil;

  switch (instruction) {

    case r_output_direction:		/* >foo */
    case r_output_force:		/* >| foo */
    case r_err_and_out:			/* &>filename */
      temp.flags = os.O_TRUNC | os.O_WRONLY | os.O_CREAT;
      break;

    case r_appending_to:		/* >>foo */
    case r_append_err_and_out:		/* &>> filename */
      temp.flags = os.O_APPEND | os.O_WRONLY | os.O_CREAT;
      break;

    case r_input_direction:		/* <foo */
    case r_inputa_direction:		/* foo & makes this. */
      temp.flags = os.O_RDONLY;
      break;

    case r_input_output:		/* <>foo */
      temp.flags = os.O_RDWR | os.O_CREAT;
      break;

    case r_deblank_reading_until: 	/* <<-foo */
    case r_reading_until:		/* << foo */
    case r_reading_string:		/* <<< foo */
    case r_close_this:			/* <&- */
    case r_duplicating_input:		/* 1<&2 */
    case r_duplicating_output:		/* 1>&2 */
      break;

    /* the parser doesn't pass these. */
    case r_move_input:			/* 1<&2- */
    case r_move_output:			/* 1>&2- */
    case r_move_input_word:		/* 1<&$foo- */
    case r_move_output_word:		/* 1>&$foo- */
      break;

    /* The way the lexer works we have to do this here. */
    case r_duplicating_input_word:	/* 1<&$foo */
    case r_duplicating_output_word:	/* 1>&$foo */
	panic("r_duplicating_output_word: Not implemented")
//       w = dest_and_filename.filename;
//       wlen = strlen (w.word) - 1;
//       if (w.word[wlen] == '-')		/* Yuck */
//         {
//           w.word[wlen] = '\0';
// 	  if (all_digits (w.word) && legal_number (w.word, &lfd) && lfd == (int)lfd)
// 	    {
// 	      dispose_word (w);
// 	      temp.instruction = (instruction == r_duplicating_input_word) ? r_move_input : r_move_output;
// 	      temp.redirectee.dest = lfd;
// 	    }
// 	  else
// 	    temp.instruction = (instruction == r_duplicating_input_word) ? r_move_input_word : r_move_output_word;
//         }
//           
//       break;
// 
    default:
	panic(fmt.Sprintf("makeRedirection: redirecttion instruction `%s' out of range", instruction))
    }
  return (temp);
}

//Command *
//make_function_def (name, command, lineno, lstart)
//     word_desc *name;
//     Command *command;
//     int lineno, lstart;
//{
//  FunctionDef *temp;
//#if defined (ARRAY_VARS)
//  SHELL_VAR *bash_source_v;
//  ARRAY *bash_source_a;
//#endif
//
//  temp = (FunctionDef *)xmalloc (sizeof (FunctionDef));
//  temp.command = command;
//  temp.name = name;
//  temp.line = lineno;
//  temp.flags = 0;
//  command.line = lstart;
//
//  /* Information used primarily for debugging. */
//  temp.source_file = 0;
//#if defined (ARRAY_VARS)
//  GET_ARRAY_FROM_VAR ("BASH_SOURCE", bash_source_v, bash_source_a);
//  if (bash_source_a && array_num_elements (bash_source_a) > 0)
//    temp.source_file = array_reference (bash_source_a, 0);
//#endif
//#if defined (DEBUGGER)
//  bind_function_def (name.word, temp);
//#endif
//
//  temp.source_file = 0;
//  return (make_command (cm_function_def, (SimpleCom *)temp));
//}
//
//Command *
//make_subshell_command (command)
//     Command *command;
//{
//  SubshellCom *temp;
//
//  temp = (SubshellCom *)xmalloc (sizeof (SubshellCom));
//  temp.command = command;
//  temp.flags = CMD_WANT_SUBSHELL;
//  return (make_command (cm_subshell, (SimpleCom *)temp));
//}
//
//Command *
//make_coproc_command (name, command)
//     char *name;
//     Command *command;
//{
//  CoprocCom *temp;
//
//  temp = (CoprocCom *)xmalloc (sizeof (CoprocCom));
//  temp.name = savestring (name);
//  temp.command = command;
//  temp.flags = CMD_WANT_SUBSHELL|CMD_COPROC_SUBSHELL;
//  return (make_command (cm_coproc, (SimpleCom *)temp));
//}
//
///* Reverse the word list and redirection list in the simple command
//   has just been parsed.  It seems simpler to do this here the one
//   time then by any other method that I can think of. */
//Command *
//clean_simple_command (command)
//     Command *command;
//{
//  if (command.type != cm_simple)
//    command_error ("clean_simple_command", CMDERR_BADTYPE, command.type, 0);
//  else
//    {
//      command.value.Simple.words =
//	REVERSE_LIST (command.value.Simple.words, word_list *);
//      command.value.Simple.redirects =
//	REVERSE_LIST (command.value.Simple.redirects, Redirect *);
//    }
//
//  parser_state &= ~PST_REDIRLIST;
//  return (command);
//}
//
///* The Yacc grammar productions have a problem, in that they take a
//   list followed by an ampersand (`&') and do a simple command connection,
//   making the entire list effectively asynchronous, instead of just
//   the last command.  This means that when the list is executed, all
//   the commands have stdin set to /dev/null when job control is not
//   active, instead of just the last.  This is wrong, and needs fixing
//   up.  This function takes the `&' and applies it to the last command
//   in the list.  This is done only for lists connected by `;'; it makes
//   `;' bind `tighter' than `&'. */
//Command *
//connect_async_list (command, command2, connector)
//     Command *command, *command2;
//     int connector;
//{
//  Command *t, *t1, *t2;
//
//  t1 = command;
//  t = command.value.Connection.second;
//
//  if (!t || (command.flags & CMD_WANT_SUBSHELL) ||
//      command.value.Connection.connector != ';')
//    {
//      t = command_connect (command, command2, connector);
//      return t;
//    }
//
//  /* This is just defensive programming.  The Yacc precedence rules
//     will generally hand this function a command where t points directly
//     to the command we want (e.g. given a ; b ; c ; d &, t1 will point
//     to the `a ; b ; c' list and t will be the `d').  We only want to do
//     this if the list is not being executed as a unit in the background
//     with `( ... )', so we have to check for CMD_WANT_SUBSHELL.  That's
//     the only way to tell. */
//  while (((t.flags & CMD_WANT_SUBSHELL) == 0) && t.type == cm_connection &&
//	 t.value.Connection.connector == ';')
//    {
//      t1 = t;
//      t = t.value.Connection.second;
//    }
//  /* Now we have t pointing to the last command in the list, and
//     t1.value.Connection.second == t. */
//  t2 = command_connect (t, command2, connector);
//  t1.value.Connection.second = t2;
//  return command;
//}
