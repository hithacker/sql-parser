package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

//Scanner is a cool struct
type Scanner struct {
	r *strings.Reader
}

//NewScanner is a cool func
func NewScanner(reader *strings.Reader) Scanner {
	return Scanner{r: reader}
}

//Scan is a cool func
func (s Scanner) Scan() (Token, string) {
	var buffer bytes.Buffer
	ch, _, err := s.r.ReadRune()
	buffer.WriteRune(ch)
	fmt.Println(buffer.String())
	if err == io.EOF {
		return EOF, ""
	}

	for {
		ch, _, err := s.r.ReadRune()
		if err != io.EOF {
			fmt.Println("Writing")
			buffer.WriteRune(ch)
		} else {
			fmt.Println("Breaking")
			break
		}
	}

	switch bufferValue := buffer.String(); {
	case bufferValue == "SELECT":
		return SELECT, "SELECT"
	case bufferValue == "FROM":
		return FROM, "FROM"
	case bufferValue == "*":
		return ASTERISK, "*"
	case isWhitespace(rune(bufferValue[0])):
		return WS, bufferValue
	default:
		return IDENT, buffer.String()
	}

	return ASTERISK, "*"
}

func isWhitespace(ch rune) bool {
	return ch == '\t' || ch == '\n' || ch == '\r' || ch == ' '
}

func (t Token) ToS() string {
	defaultValue := strconv.Itoa(int(t))
	switch t {
	case SELECT:
		return "SELECT"
	case FROM:
		return "FROM"
	case ASTERISK:
		return "ASTERISK"
	case EOF:
		return "EOF"
	case ILLEGAL:
		return "ILLEGAL"
	case IDENT:
		return "IDENT"
	case WS:
		return "WS"
	default:
		return defaultValue

	}
}
