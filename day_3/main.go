package main

import (
	"os"
	"strconv"
	"strings"
)

type TokenType int
type Token struct {
	t  TokenType
	ch rune
}

const (
	Digit TokenType = iota
	Mul
	Do
	Dont
	LeftParen
	RightParen
	Comma
	Ignore
	EOF
)

func parseTokens(src string) []Token {
	r := []Token{}
	prev := ""
	i := 0
	for i < len(src)-1 {
		ch := rune(src[i])

		if ch >= '0' && ch <= '9' {
			r = append(r, Token{t: Digit, ch: ch})
		} else if ch == '(' {
			if i+1 < len(src)-1 && src[i+1] == ')' {
				if prev == "do" {
					r = append(r, Token{t: Do})
					i++
					prev = ""
				} else if prev == "don't" {
					r = append(r, Token{t: Dont})
					i++
					prev = ""
				}
			}
			r = append(r, Token{t: LeftParen})
		} else if ch == ')' {
			r = append(r, Token{t: RightParen})
		} else if ch == ',' {
			r = append(r, Token{t: Comma})
		} else if ch == 'm' {
			prev = "m"
		} else if ch == 'u' {
			if prev == "m" {
				prev = "mu"
			} else {
				prev = ""
			}
		} else if ch == 'l' {
			if prev == "mu" {
				r = append(r, Token{t: Mul})
				prev = ""
			}
		} else if ch == 'd' {
			prev = "d"
		} else if ch == 'o' {
			if prev == "d" {
				prev = "do"
			} else {
				prev = ""
			}
		} else if ch == 'n' {
			if prev == "do" {
				prev = "don"
			} else {
				prev = ""
			}
		} else if ch == '\'' {
			if prev == "don" {
				prev = "don'"
			} else {
				prev = ""
			}
		} else if ch == 't' {
			if prev == "don'" {
				prev = "don't"
			} else {
				prev = ""
			}
		} else {
			r = append(r, Token{t: Ignore})
		}
		i++
	}
	r = append(r, Token{t: EOF})
	return r
}

type Parser struct {
	tokens  []Token
	current int
}

func (p *Parser) isAtEnd() bool {
	return p.tokens[p.current].t == EOF
}

func (p *Parser) peek() Token {
	return p.tokens[p.current]
}

func (p *Parser) consume() Token {
	t := p.peek()
	p.current++
	return t
}

func (p *Parser) consumeUntil(t TokenType) bool {
	for !p.isAtEnd() && p.peek().t != t {
		p.consume()
	}

	return p.peek().t == Mul
}

func (p *Parser) parseInteger() int {
	str := []string{}

	for !p.isAtEnd() && p.peek().t == Digit {
		t := p.consume()
		str = append(str, string(t.ch))
	}

	i, _ := strconv.Atoi(strings.Join(str, ""))
	return i
}

func (p *Parser) parse() int {
	res := [][]int{}

	do := true
	for !p.isAtEnd() {
		if p.peek().t == Do {
			do = true
			p.consume()
			continue
		}

		if p.peek().t == Dont {
			do = false
			p.consume()
			continue
		}

		if p.peek().t != Mul {
			p.consume()
			continue
		}

		p.consume()

		if p.peek().t != LeftParen {
			continue
		}

		p.consume()

		if p.peek().t != Digit {
			continue
		}

		left := p.parseInteger()

		if p.peek().t != Comma {
			continue
		}
		p.consume()

		if p.peek().t != Digit {
			continue
		}

		right := p.parseInteger()

		if p.peek().t != RightParen {
			continue
		}

		if do {
			res = append(res, []int{left, right})
		}
	}

	result := 0

	for _, pair := range res {
		result += pair[0] * pair[len(pair)-1]
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		println("Usage: go run main.go <input>")
		os.Exit(1)
	}

	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		println("Can't read file: ", os.Args[1])
		os.Exit(1)
	}

	input := string(b)
	tokens := parseTokens(input)
	parser := Parser{tokens: tokens, current: 0}

	println(parser.parse())
}
