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
	"strings"
	"unicode"
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

func make_bare_word(str string) *word_desc {
  temp := new(word_desc)
  temp.word = str
  return temp
}

func make_word_flags(w *word_desc, str string) *word_desc {
  for _, ch := range str {
      switch ch {
	case '$':
	  w.flags |= W_HASDOLLAR;
	case '\\':
	  break;	/* continue the loop */
	case '\'': fallthrough
	case '`': fallthrough
	case '"':
	  w.flags |= W_QUOTED;
	}
  }

  return w
}

func make_word_from_token(token int) *word_desc {
  return make_word(runesToString([]int { token }))
}

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

func (gps *ParserState) make_command(typ command_type, value interface{}) *Command {
  temp := new(Command)
  temp.typ = typ;
  switch typ {
  case cm_for: temp.value.For = value.(*ForCom)
	case cm_case: temp.value.Case = value.(*CaseCom)
	case cm_while: temp.value.While = value.(*WhileCom)
	case cm_if: temp.value.If = value.(*IfCom)
	case cm_simple: temp.value.Simple = value.(*SimpleCom)
	case cm_select: temp.value.Select = value.(*SelectCom)
	case cm_connection: temp.value.Connection = value.(*Connection)
	case cm_function_def: temp.value.Function_def = value.(*FunctionDef)
	case cm_until: temp.value.While = value.(*WhileCom)
	case cm_group: temp.value.Group = value.(*GroupCom)
	case cm_arith: temp.value.Arith = value.(*ArithCom)
	case cm_cond: temp.value.Cond = value.(*CondCom)
	case cm_arith_for: temp.value.ArithFor = value.(*ArithForCom)
	case cm_subshell: temp.value.Subshell = value.(*SubshellCom)
	case cm_coproc: temp.value.Coproc = value.(*CoprocCom)

  }
  temp.flags = 0
  temp.redirects = nil;
  return temp
}

func (gps *ParserState) command_connect(com1, com2 *Command, connector int) *Command {
  temp := new(Connection)
  temp.connector = connector;
  temp.first = com1;
  temp.second = com2;
  return gps.make_command (cm_connection, temp)
}

func (gps *ParserState) make_for_or_select(typ command_type, name *word_desc, map_list *word_list, action *Command, lineno int) *Command {
  temp := new(ForCom)
  temp.name = name;
  temp.line = lineno;
  temp.map_list = map_list;
  temp.action = action;
  return gps.make_command(typ, temp)
}

func (gps *ParserState) make_for_command(name *word_desc, map_list *word_list, action *Command, lineno int) *Command {
  return gps.make_for_or_select(cm_for, name, map_list, action, lineno)
}

func (gps *ParserState) make_select_command(name *word_desc, map_list *word_list, action *Command, lineno int) *Command {
	return gps.make_for_or_select(cm_select, name, map_list, action, lineno)
}

func make_arith_for_expr(s string) *word_list {
  if s == "" {
    return nil
  }
  wd := make_word(s)
  wd.flags |= W_NOGLOB|W_NOSPLIT|W_QUOTED|W_DQUOTE;	/* no word splitting or globbing */
  result := make_word_list (wd, nil);
  return result;
}

func (gps *ParserState) make_arith_for_command(exprs *word_list, action *Command, lineno int) *Command {
  chunks := strings.Split(exprs.word.word, ";", -1)
  if len(chunks) != 3 {
    if len(chunks) < 3 {
      gps.parser_error(lineno, "syntax error: arithmetic expression required")
    } else {
      gps.parser_error(lineno, "syntax error: `;' unexpected")
    }
    gps.parser_error (lineno, "syntax error: `((%s))'", exprs.word.word)
    gps.last_command_exit_value = 2;
    return nil
  }

  make_expr := func(x string) *word_list {
    expr := make_arith_for_expr(strings.TrimLeftFunc(x, unicode.IsSpace))
    if expr == nil {
      expr = make_arith_for_expr("1")
    }
    return expr
  }

  temp := new(ArithForCom)
  temp.line = lineno;
  temp.init = make_expr(chunks[0])
  temp.test = make_expr(chunks[1])
  temp.step = make_expr(chunks[2])
  temp.action = action;

  return gps.make_command(cm_arith_for, temp)
}

func (gps *ParserState) make_group_command(command *Command) *Command {
  temp := new(GroupCom)
  temp.command = command;
  return gps.make_command(cm_group, temp)
}

func (gps *ParserState) make_case_command(word *word_desc, clauses *PatternList, lineno int) *Command {
  temp := new(CaseCom)
  temp.line = lineno
  temp.word = word
  temp.clauses = reversePatternListList(clauses)
  return gps.make_command(cm_case, temp)
}

func (gps *ParserState) make_pattern_list(patterns *word_list, action *Command) *PatternList {
  temp := new(PatternList)
  temp.patterns = reverseWordList(patterns)
  temp.action = action;
  return temp
}

func (gps *ParserState) make_if_command(test *Command, true_case *Command, false_case *Command) *Command {
  temp := new(IfCom)
  temp.test = test;
  temp.true_case = true_case;
  temp.false_case = false_case;
  return gps.make_command(cm_if, temp)
}

func (gps *ParserState) make_until_or_while(which command_type, test *Command, action *Command) *Command {
  temp := new(WhileCom)
  temp.flags = 0;
  temp.test = test;
  temp.action = action;
  return gps.make_command(which, temp)
}

func (gps *ParserState) make_while_command (test *Command, action *Command) *Command {
  return gps.make_until_or_while (cm_while, test, action)
}

func (gps *ParserState) make_until_command (test *Command, action *Command) *Command {
  return gps.make_until_or_while (cm_until, test, action)
}

func (gps *ParserState) make_arith_command(exp *word_list) *Command {
  temp := new(ArithCom)
  temp.line = gps.line_number;
  temp.exp = exp;

  command := new(Command)
  command.typ = cm_arith;
  command.value.Arith = temp

  return command
}

func (gps *ParserState) make_cond_node (typ int, op *word_desc, left *CondCom, right *CondCom) *CondCom {
  temp := new(CondCom)
  temp.line = gps.line_number;
  temp.typ = typ;
  temp.op = op;
  temp.left = left;
  temp.right = right;

  return temp
}

func (gps *ParserState) make_cond_command (cond_node *CondCom) (command *Command) {
  command = new(Command)
  command.value.Cond = cond_node;

  command.typ = cm_cond;
  if cond_node != nil {
    command.line = cond_node.line
  }

  return
}

func (gps *ParserState) make_bare_simple_command () *Command {
  command := new(Command)
  temp := new(SimpleCom)
  command.value.Simple = temp

  temp.flags = 0;
  temp.line = gps.line_number;
  temp.words = nil;
  temp.redirects = nil;

  command.typ = cm_simple;
  command.redirects = nil;
  command.flags = 0;

  return (command);
}

/* Return a command which is the connection of the word or redirection
   in ELEMENT, and the command * or NULL in Command. */
func (gps *ParserState) make_simple_command (element ELEMENT, command *Command) *Command {
  /* If we are starting from scratch, then make the initial command
     structure.  Also note that we have to fill in all the slots, since
     malloc doesn't return zeroed space. */
  if command == nil {
      command = gps.make_bare_simple_command ();
      gps.parser_state |= PST_REDIRLIST;
    }

  if element.word != nil {
      command.value.Simple.words = make_word_list (element.word, command.value.Simple.words);
      gps.parser_state &= ^PST_REDIRLIST;
    }  else if element.redirect != nil {
      r:= element.redirect;
      /* Due to the way <> is implemented, there may be more than a single
	 redirection in element.redirect.  We just follow the chain as far
	 as it goes, and hook onto the end. */
      for r.next != nil {
	r = r.next;
      }
      r.next = command.value.Simple.redirects;
      command.value.Simple.redirects = element.redirect;
    }

  return (command);
}

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
//  document = nil;
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
//  for (full_line = read_secondary_line (delim_unquoted))
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
//	  for (*line == '\t')
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

    case r_output_direction: fallthrough /* >foo */
    case r_output_force: fallthrough	/* >| foo */
    case r_err_and_out:			/* &>filename */
      temp.flags = os.O_TRUNC | os.O_WRONLY | os.O_CREAT;
      break;

    case r_appending_to: fallthrough	/* >>foo */
    case r_append_err_and_out:		/* &>> filename */
      temp.flags = os.O_APPEND | os.O_WRONLY | os.O_CREAT;
      break;

    case r_input_direction: fallthrough	/* <foo */
    case r_inputa_direction: /* foo & makes this. */
      temp.flags = os.O_RDONLY;
      break;

    case r_input_output:		/* <>foo */
      temp.flags = os.O_RDWR | os.O_CREAT;
      break;

    case r_deblank_reading_until: fallthrough /* <<-foo */
    case r_reading_until: fallthrough /* << foo */
    case r_reading_string: fallthrough	/* <<< foo */
    case r_close_this: fallthrough	/* <&- */
    case r_duplicating_input: fallthrough /* 1<&2 */
    case r_duplicating_output:		/* 1>&2 */
      break;

    /* the parser doesn't pass these. */
    case r_move_input: fallthrough	/* 1<&2- */
    case r_move_output:	fallthrough	/* 1>&2- */
    case r_move_input_word: fallthrough	/* 1<&$foo- */
    case r_move_output_word:		/* 1>&$foo- */
      break;

    /* The way the lexer works we have to do this here. */
    case r_duplicating_input_word: fallthrough	/* 1<&$foo */
    case r_duplicating_output_word:	/* 1>&$foo */
	panic("r_duplicating_output_word: Not implemented")
//       w = dest_and_filename.filename;
//       wlen = strlen (w.word) - 1;
//       if (w.word[wlen] == '-')		/* Yuck */
//         {
//           w.word[wlen] = '\0';
// 	  if (all_digits (w.word) && legal_number (w.word, &lfd) && lfd == (int)lfd)
// 	    {
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

func (gps *ParserState) make_function_def(name *word_desc, command *Command, lineno int, lstart int) *Command {
	// TODO(krasin): implement this.
	panic("make_function_def is not implemented. Sorry")
}

func (gps *ParserState) make_subshell_command(command *Command) *Command {
  temp := new(SubshellCom)
  temp.command = command;
  return gps.make_command(cm_subshell, temp)
}

func (gps *ParserState) make_coproc_command(name string, command *Command) *Command {
  temp := new(CoprocCom)
  temp.name = name
  temp.command = command;
  return gps.make_command(cm_coproc, temp)
}

/* Reverse the word list and redirection list in the simple command
   has just been parsed.  It seems simpler to do this here the one
   time then by any other method that I can think of. */
func (gps *ParserState) clean_simple_command (command *Command) *Command {
  if command.typ != cm_simple {
    command_error ("clean_simple_command", CMDERR_BADTYPE, command.typ, 0);
  } else {
    command.value.Simple.words = reverseWordList(command.value.Simple.words)
    command.value.Simple.redirects = reverseRedirectList(command.value.Simple.redirects)
 }

  gps.parser_state &= ^PST_REDIRLIST;
  return (command);
}

/* The Yacc grammar productions have a problem, in that they take a
   list followed by an ampersand (`&') and do a simple command connection,
   making the entire list effectively asynchronous, instead of just
   the last command.  This means that when the list is executed, all
   the commands have stdin set to /dev/null when job control is not
   active, instead of just the last.  This is wrong, and needs fixing
   up.  This function takes the `&' and applies it to the last command
   in the list.  This is done only for lists connected by `;'; it makes
   `;' bind `tighter' than `&'. */
func (gps *ParserState) connect_async_list(command *Command, command2 *Command, connector int) *Command {
  t1 := command;
  t := command.value.Connection.second;

  if t == nil || (command.flags & CMD_WANT_SUBSHELL != 0) ||
      command.value.Connection.connector != ';'  {
      t = gps.command_connect (command, command2, connector);
      return t;
    }

  /* This is just defensive programming.  The Yacc precedence rules
     will generally hand this function a command where t points directly
     to the command we want (e.g. given a ; b ; c ; d &, t1 will point
     to the `a ; b ; c' list and t will be the `d').  We only want to do
     this if the list is not being executed as a unit in the background
     with `( ... )', so we have to check for CMD_WANT_SUBSHELL.  That's
     the only way to tell. */
  for ((t.flags & CMD_WANT_SUBSHELL) == 0) && t.typ == cm_connection &&
	 t.value.Connection.connector == ';'    {
      t1 = t;
      t = t.value.Connection.second;
    }
  /* Now we have t pointing to the last command in the list, and
     t1.value.Connection.second == t. */
  t2 := gps.command_connect (t, command2, connector);
  t1.value.Connection.second = t2;
  return command;
}
