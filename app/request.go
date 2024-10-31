package main

import "fmt"

type Request struct {
	Method
	Content []byte
}

func NewRequest(m Method, c []byte) *Request {
	return &Request{
		Method:  m,
		Content: c,
	}
}

func (m *Request) String() string {
	return fmt.Sprintf("Method = %s\nContent = %s", m.Method, string(m.Content))
}
