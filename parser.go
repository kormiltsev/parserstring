package parserstring

import (
	"fmt"
	"strings"
)

type Token struct {
	Name        string
	Startstring string
	Endstring   string
}

// only together
type Req struct {
	Tokens []Token
	Result []map[string]string
}

func NewReq() *Req {
	return &Req{
		Tokens: make([]Token, 0),
		Result: make([]map[string]string, 0),
	}
}

func (r *Req) AddRequestToken(name, start, end string) {
	r.Tokens = append(r.Tokens, Token{Name: name, Startstring: start, Endstring: end})
}

func (r *Req) ParseString(originalstring string) {
	s := originalstring
	founded := true
	aft := ""
	done := false
	i := 0
	for {
		an := make(map[string]string, 0)
		for _, t := range r.Tokens {
			_, aft, founded = strings.Cut(s, t.Startstring)
			if !founded {
				done = true
				continue
			}
			an[t.Name], _, _ = strings.Cut(aft, t.Endstring)
			s = aft
		}
		if done {
			return
		}
		r.Result = append(r.Result, an)
		i++
	}
}

func (r *Req) PrintResult() {
	tknames := make([]string, 0)
	for _, tk := range r.Tokens {
		tknames = append(tknames, tk.Name)
	}
	for _, row := range r.Result {
		for i := 1; i < len(r.Tokens); i++ {
			value, ok := row[r.Tokens[i].Name]
			if !ok {
				continue
			}
			fmt.Println(r.Tokens[i].Name, ":", value)
		}

	}
}
