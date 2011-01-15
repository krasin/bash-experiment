package gobash

import (
	"os"
	"log"
	"fmt"
)

const NO_PIPE = -1
const REDIRECT_BOTH = -2

const NO_VARIABLE = -1

/* Values that can be returned by execute_command (). */
const EXECUTION_FAILURE = 1
const EXECUTION_SUCCESS = 0

/* Usage messages by builtins result in a return status of 2. */
const EX_BADUSAGE = 2

/* Special exit statuses used by the shell, internally and externally. */
const EX_RETRYFAIL = 124
const EX_WEXPCOMSUB = 125
const EX_BINARY_FILE = 126
const EX_NOEXEC = 126
const EX_NOINPUT = 126
const EX_NOTFOUND = 127

const EX_SHERRBASE = 256 /* all special error values are > this. */

const EX_BADSYNTAX = 257 /* shell syntax error */
const EX_USAGE = 258 /* syntax error in usage */
const EX_REDIRFAIL = 259 /* redirection failed */
const EX_BADASSIGN = 260 /* variable assignment error */
const EX_EXPFAIL = 261 /* word expansion failed */

/* Flag values that control parameter pattern substitution. */
const MATCH_ANY = 0x000
const MATCH_BEG = 0x001
const MATCH_END = 0x002

const MATCH_TYPEMASK = 0x003

const MATCH_GLOBREP = 0x010
const MATCH_QUOTED = 0x020
const MATCH_STARSUB = 0x040


type ShellState struct {
	gps   *ParserState
	input BashInput
}

func newShellState() *ShellState {
	return new(ShellState)
}

func ExecuteScript(filename string) int {
	shell := newShellState()
	defer shell.shutdown()
	err := shell.openShellScript(filename)
	if err != nil {
		return shell.fatal(fmt.Sprintf("Can't open a shell script file: %s, err: %v", filename, err))
	}
	return shell.readerLoop()
}

func (sh *ShellState) openShellScript(filename string) (err os.Error) {
	file, err := os.Open(filename, os.O_RDONLY, 0)
	if err != nil {
		return
	}
	sh.input = newBufferedBashInput(file)
	sh.gps = newParserState(sh.input)
	return
}

func (sh *ShellState) readerLoop() int {
	for !sh.gps.EOF_Reached {
		sh.gps.Yyparse()
		fmt.Fprintf(os.Stderr, "global_command: %v\n", sh.gps.global_command)
	}
	return 0
}

func (sh *ShellState) shutdown() {
	if sh.input != nil {
		sh.input.Close()
        }
}

func (sh *ShellState) fatal(msg string) int {
	log.Printf("Fatal: %s\n", msg)
	return 1
}
