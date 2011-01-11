package gobash

import (
	"bufio"
	"os"
)

type BashInput interface{
	Close() os.Error
	// Returns UTF-8 Rune or gobash.EOF
	Getc() int
	Ungetc()
}

type bufferedBashInput struct {
	file *os.File
	reader *bufio.Reader
}

func (bi *bufferedBashInput) Close() os.Error {
	return bi.file.Close()
}

func newBufferedBashInput(file *os.File) BashInput {
	return &bufferedBashInput{file: file, reader: bufio.NewReader(file) }
}

func (bi *bufferedBashInput) Getc() int {
	rune, _, err := bi.reader.ReadRune()
	if err != nil {
		return EOF
	}
	return rune
}

func (bi *bufferedBashInput) Ungetc() {
	bi.reader.UnreadRune()
}
