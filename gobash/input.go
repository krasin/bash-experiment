package gobash

import (
	"os"
)

type BashInput interface{
	Close() os.Error
}

type bufferedBashInput struct {
	file *os.File
}

func (bi *bufferedBashInput) Close() os.Error {
	return bi.file.Close()
}

func newBufferedBashInput(file *os.File) BashInput {
	return &bufferedBashInput{file: file}
}
