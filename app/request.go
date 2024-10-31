package main

import "fmt"

type Request struct {
	Method
	Route   string
	Content []byte
}

func NewRequest(m Method) *Request {
	return &Request{
		Method: m,
	}
}

func (req *Request) SetRoute(route string) {
	req.Route = route
}

func (req *Request) CreateContentBuf(n int) {
	req.Content = make([]byte, n)
}

func (req *Request) String() string {
	return fmt.Sprintf("Method = %s\nRoute = %s\nContent = %s\n", req.Method, req.Route, string(req.Content))
}

func (req *Request) IsValidRoute(r Router) bool {
	return r.Has(req.Route)
}
