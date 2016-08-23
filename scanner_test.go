package main

import (
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: EOF},
		{s: `#`, tok: ILLEGAL, lit: `#`},
		{s: ` `, tok: WS, lit: " "},
		{s: "\t", tok: WS, lit: "\t"},
		{s: "\n", tok: WS, lit: "\n"},

		// Misc characters
		{s: `*`, tok: ASTERISK, lit: "*"},

		// Identifiers
		{s: `foo`, tok: IDENT, lit: `foo`},
		//TODO: See why this test exepcts us to remove the ending dash
		{s: `Zx12_3U_-`, tok: IDENT, lit: `Zx12_3U_`},

		// Keywords
		{s: `FROM`, tok: FROM, lit: "FROM"},
		{s: `SELECT`, tok: SELECT, lit: "SELECT"},
	}

	for i, tt := range tests {
		s := NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok.ToS(), tok.ToS(), lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
