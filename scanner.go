package main

import "strings"

//Scanner is a cool struct
type Scanner struct {
}

//NewScanner is a cool func
func NewScanner(reader *strings.Reader) Scanner {
	return Scanner{}
}

//Scan is a cool func
func (s Scanner) Scan() (Token, string) {
	return ASTERISK, ""
}
