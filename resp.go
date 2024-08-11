package main

import "bufio"

const (
	STRING  = "+"
	ERROR   = "-"
	INTEGER = ":"
	BULK    = "$"
	ARRAY   = "*"
)

type Value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []Value
}

type Resp struct {
	reader *bufio.Reader
}

func (r *Resp) NewResp(reader bufio.Reader) *Resp {
	return &Resp{reader: &reader}
}
