package main

type Parser interface {
	Parse([]byte, int) (*Request, error)
}

type RequestParser struct {
	whiteSpace byte
	backslash  byte
	frontslash byte
	sep        []byte
}

func NewRequestParser() *RequestParser {
	return &RequestParser{
		whiteSpace: byte(' '),
		backslash:  byte('\\'),
		frontslash: byte('/'),
		sep:        []byte("\r\n"),
	}
}

func (rp *RequestParser) Parse(buf []byte, n int) (*Request, error) {
	return nil, nil
}
