package gobash

import (
	"os"
	"log"
	"fmt"
)

type ShellState struct {
	gps   *ParserState
	input BashInput
}

func newShellState() *ShellState {
	return &ShellState{gps: newParserState()}
}

func ExecuteScript(filename string) int {
	shell := newShellState()
	err := shell.openShellScript(filename)
	if err != nil {
		return shell.fatal(fmt.Sprintf("Can't open a shell script file: %s, err: %v", filename, err))
	}
	shell.shutdown()
	return 0
}

func (sh *ShellState) openShellScript(filename string) (err os.Error) {
	file, err := os.Open(filename, os.O_RDONLY, 0)
	if err != nil {
		return
	}
	sh.input = newBufferedBashInput(file)
	return
}

func (sh *ShellState) shutdown() {
	sh.input.Close()
}

func (sh *ShellState) fatal(msg string) int {
	log.Printf("Fatal: %s\n", msg)
	return 1
}
