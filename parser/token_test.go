package parser

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestToken1(t *testing.T) {
	src := `((abc 123 "hello world")
"hello lf:\nworld cr:\r octet:\073\nhex:\x30\nbackslash:\\\nsingle quote:\'\ndouble quote:\"\nquestion: \?"
(+ 1 2 3) ; comment`
	ss, err := newTokenizer("TestToken1", strings.NewReader(src))
	if err != nil {
		t.Error(err)
	}
	// TODO 真面目にss.scan()を呼び出す毎の結果と比較しないと単体テストにならない。
	r, line, col, err := ss.scan()
	for err == nil {
		switch r {
		case symbol:
			fmt.Printf("[%d,%d]symbol: %v\n", line, col, ss.tokentext())
		case stringLiteral:
			fmt.Printf("[%d,%d]string: %v\n", line, col, ss.tokentext())
		case commentText:
			fmt.Printf("[%d,%d]comment: %v\n", line, col, ss.tokentext())
		default:
			fmt.Printf("[%d,%d]char: '%c'\n", line, col, r)
		}
		r, line, col, err = ss.scan()
	}
	if err != nil && err != io.EOF {
		t.Error(err)
	}
}
