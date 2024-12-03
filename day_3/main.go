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
	LeftParen
	RightParen
	Comma
	Ignore
	EOF
)

func parseTokens(src string) []Token {
	r := []Token{}
	prev := ""
	for _, ch := range src {
		if ch >= '0' && ch <= '9' {
			r = append(r, Token{t: Digit, ch: ch})
		} else if ch == '(' {
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
		} else {
			r = append(r, Token{t: Ignore})
		}
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

	for !p.isAtEnd() {
		if !p.consumeUntil(Mul) {
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

		res = append(res, []int{left, right})
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
