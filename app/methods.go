package main

import (
	"errors"
)

type Method byte

const (
	MethodNotAllowed Method = 0
	MethodGet        Method = Method('G')
	MethodPut        Method = Method('P')
)

func GetMethod(m byte) (Method, error) {
	switch m {
	case byte(MethodGet):
		return MethodGet, nil
	case byte(MethodPut):
		return MethodPut, nil
	default:
		return MethodNotAllowed, errors.New("invalid method provided")
	}
}

var methodName = map[Method]string{
	MethodGet: "GET",
	MethodPut: "PUT",
}

func (m Method) String() string {
	return methodName[m]
}

func (m Method) GetLen() int {
	return len(m.String())
}
