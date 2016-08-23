package main

import (
	"fmt"
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
	ch, size, err := s.r.ReadRune()
	fmt.Println("Ch", ch, "Size", size, "Error", err)
	fmt.Printf("%q\n", ch)
	if size == 0 {
		return EOF, ""
	}

	return ASTERISK, "*"
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
