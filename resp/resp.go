package resp

import (
	"bufio"
	"strconv"
)

const (
	STRING  = "+"
	ERROR   = "-"
	INTEGER = ":"
	BULK    = "$"
	ARRAY   = "*"
)

type Resp struct {
	reader *bufio.Reader
}

func (r *Resp) NewResp(reader bufio.Reader) *Resp {
	return &Resp{reader: &reader}
}

func (r *Resp) ReadLine() ([]byte, int, error) {
	var line []byte
	var n int
	var err error

	for {
		b, err := r.reader.ReadByte()
		if err != nil {
			return nil, 0, err
		}

		n += 1
		line = append(line, b)
		if len(line) >= 2 && line[len(line)-2] == '\r' {
			break
		}
	}

	return line[:len(line)-2], n, err
}

func (r *Resp) ReadInteger() (int, int, error) {
	var err error

	line, n, err := r.ReadLine()
	if err != nil {
		return 0, 0, err
	}

	i64, err := strconv.ParseInt(string(line), 10, 64)
	if err != nil {
		return 0, 0, err
	}

	return int(i64), n, err
}
