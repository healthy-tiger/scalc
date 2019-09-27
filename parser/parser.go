package parser

import (
	"bytes"
	"io"
	"strconv"
	"strings"
)

// SymbolTable シンボルIDとシンボル名のマップ
type SymbolTable struct {
	symbolMap map[string]SymbolID
}

// NewSymbolTable 新しいSymbolTableを作る。
func NewSymbolTable() *SymbolTable {
	return &SymbolTable{make(map[string]SymbolID)}
}

// GetSymbolID はシンボルnameに対するIDを返す。
// IDが割り当てられていないシンボルに対しては、新たにIDを割り当てて返す。
func (st *SymbolTable) GetSymbolID(name string) SymbolID {
	n, ok := st.symbolMap[name]
	if !ok {
		n = SymbolID(len(st.symbolMap))
		st.symbolMap[name] = SymbolID(n)
	}
	return n
}

// GetSymbolName はシンボルのIDからシンボル名を取得する。
func (st *SymbolTable) GetSymbolName(id SymbolID) (string, error) {
	for k, v := range st.symbolMap {
		if v == id {
			return k, nil
		}
	}
	return "", ErrorUndefinedSymbol
}

// Position ソースコード上の位置を表す
type Position struct {
	Filename string
	Line     int
	Column   int
}

// Parse srcをスキャンしてSTreeを返す。
func Parse(filename string, st *SymbolTable, src io.Reader) ([]*List, error) {
	lists := make([]*List, 0)
	stack := newStack()
	tokenizer, err := newTokenizer(filename, src)
	if err != nil {
		return nil, err
	}
	tok, line, column, err := tokenizer.scan()
	for err == nil {
		switch tok {
		case symbol:
			toktxt := tokenizer.tokentext()
			lst := stack.peek()
			if lst == nil {
				return nil, newError(filename, line, column, ErrorTopLevelElementMustBeAList)
			}
			// IntかFloatとして処理できるか先に確認し、どちらもダメならシンボルにする。
			vi, err := strconv.ParseInt(toktxt, 0, 64)
			if err == nil {
				lst.elements = append(lst.elements, newLiteral(vi, filename, line, column))
			} else {
				vf, err := strconv.ParseFloat(toktxt, 64)
				if err == nil {
					lst.elements = append(lst.elements, newLiteral(vf, filename, line, column))
				} else {
					lst.elements = append(lst.elements, newLiteral(st.GetSymbolID(toktxt), filename, line, column))
				}
			}

		case stringLiteral:
			lst := stack.peek()
			if lst == nil {
				return nil, newError(filename, line, column, ErrorTopLevelElementMustBeAList)
			}
			lst.elements = append(lst.elements, newLiteral(tokenizer.tokentext(), filename, line, column))

		case commentText:

		default:
			if tok == leftParenthesis || tok == leftSquareBracket || tok == leftCurlyBracket {
				lst := stack.peek()
				lstnew := &List{tok, make([]SyntaxElement, 0), Position{filename, line, column}}
				if lst != nil {
					lst.elements = append(lst.elements, lstnew)
				} else {
					lists = append(lists, lstnew)
				}
				stack.push(lstnew)
			} else if tok == rightParenthesis || tok == rightSquareBracket || tok == rightCurlyBracket {
				lst := stack.peek()
				if lst == nil || !lst.isMatchingParen(tok) {
					return nil, newError(filename, line, column, ErrorUnexpectedInputChar)
				}
				stack.pop()
			} else if tok != tab && tok != space {
				return nil, newError(filename, line, column, ErrorUnexpectedInputChar)
			}
		}
		tok, line, column, err = tokenizer.scan()
	}
	// tokenizerのエラー＝字句解析のエラーの場合はパースを途中で止める。
	if err != io.EOF {
		return nil, err
	}
	if stack.peek() != nil {
		return nil, newError(filename, line, column, ErrorMissingClosingParenthesis)
	}
	return lists, nil
}

// ParseString 文字列をスキャンしてSTreeを返す。
func ParseString(filename string, st *SymbolTable, src string) ([]*List, error) {
	return Parse(filename, st, strings.NewReader(src))
}

func (p Position) String() string {
	var b bytes.Buffer
	b.WriteString(p.Filename)
	b.WriteString("[")
	b.WriteString(strconv.FormatInt(int64(p.Line), 10))
	b.WriteString(":")
	b.WriteString(strconv.FormatInt(int64(p.Column), 10))
	b.WriteString("]")
	return b.String()
}
