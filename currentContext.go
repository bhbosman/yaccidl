package yaccidl

import (
	"github.com/bhbosman/gocommon"
)

type CurrentContext struct {
	ReaderCloser         gocommon.ByteReaderCloser
	Next                 *CurrentContext
	FileInformationEntry FileInformationEntry
	Row                  int
	Col                  int
	UnusedByte           byte
	UnusedByteAssigned   bool
	FileName             string
	Path                 string
}

func NewCurrentContext(
	fileName, path string,
	readerCloser gocommon.ByteReaderCloser,
	fileInformationEntry FileInformationEntry,
	next *CurrentContext) *CurrentContext {
	return &CurrentContext{
		ReaderCloser:         readerCloser,
		Next:                 next,
		FileInformationEntry: fileInformationEntry,
		Row:                  1,
		Col:                  1,
		FileName:             fileName,
		Path:                 path,
	}
}

func (c CurrentContext) Close() error {
	return c.ReaderCloser.Close()
}

func (self *CurrentContext) SetUnusedByte(b byte) {
	self.UnusedByte = b
	self.UnusedByteAssigned = true
}

func (self *CurrentContext) GetUnusedByte() byte {
	self.UnusedByteAssigned = false
	return self.UnusedByte
}
