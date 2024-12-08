package day8

import (
	"bufio"
	"encoding/hex"
	"strconv"
	"strings"
	"unicode"
)

func escape(input string) int {
	n := 0
	for _, s := range strings.Split(input, "\n") {
		u, _ := strconv.Unquote(s)
		n += len(s) - len(u)
	}
	return n
}

func unescape(input string) int {
	n := 0
	for _, s := range strings.Split(input, "\n") {
		n += len(strconv.Quote(s)) - len(s)
	}
	return n
}

// first attempt, close but too aggressively escaping? :shrug:

type parser struct {
	isEscaped bool
	openQuote bool
	isHexChar bool

	buf     *bufio.Scanner
	hexChar []byte

	original string
	escaped  string

	originalLines []string
	escapedLines  []string
}

func new(input string) *parser {
	buf := bufio.NewScanner(strings.NewReader(input))
	buf.Split(bufio.ScanRunes)

	return &parser{
		isEscaped:     false,
		openQuote:     false,
		buf:           buf,
		originalLines: []string{},
		escapedLines:  []string{},
	}
}

func escapeV1(input string) int {
	p := new(input)
	p.run()
	o, e := len(p.original), len(p.escaped)
	return o - e
}

func (p *parser) run() {
	var orig strings.Builder
	var escaped strings.Builder
	for p.buf.Scan() {
		tok := p.buf.Text()
		e := p.parse(tok)
		if e != newLineMarker {
			orig.WriteString(tok)
			escaped.WriteString(e)
		} else {
			o, e := orig.String(), escaped.String()
			orig.Reset()
			escaped.Reset()
			p.original += o
			p.escaped += e
			p.originalLines = append(p.originalLines, o)
			p.escapedLines = append(p.escapedLines, e)
		}
	}
	if orig.Len() > 0 {
		o, e := orig.String(), escaped.String()
		p.original += o
		p.escaped += e
		p.originalLines = append(p.originalLines, o)
		p.escapedLines = append(p.escapedLines, e)
	}
}

const newLineMarker = `\n`

func (p *parser) parse(tok string) string {
	switch tok {
	case `\`:
		p.isEscaped = !p.isEscaped
		if p.isEscaped {
			return ""
		}
	case "\n":
		if !p.openQuote {
			p.isEscaped = false
			return newLineMarker
		}
	case `x`:
		if p.isEscaped {
			p.isEscaped = false
			p.isHexChar = true
			return ""
		}
	case `"`:
		if p.isEscaped {
			p.isEscaped = false
			return `"`
		} else {
			p.openQuote = !p.openQuote
			return ""
		}
	default:
		if p.isHexChar {
			p.hexChar = append(p.hexChar, byte(tok[0]))
			if len(p.hexChar) < 2 {
				return ""
			}
			str := decodeHexChar(p.hexChar)
			p.isHexChar = false
			p.hexChar = []byte{}
			return str
		}
		p.isEscaped = false
	}
	return tok
}

func decodeHexChar(src []byte) string {
	h, err := hex.DecodeString(string(src))
	if err != nil {
		return ""
	}

	dst := []byte{}
	for _, b := range h {
		if !unicode.IsGraphic(rune(b)) {
			continue
		}
		dst = append(dst, b)
	}
	return string(dst)
}
