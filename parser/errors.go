package parser

import (
	"bytes"
	"errors"
	"strconv"
)

// 内部エラーの定義
var (
	ErrorArgumentIsNil            = errors.New("Argument is nil")
	ErrorValueTypeIsNotAsExpected = errors.New("Value type is not as expected")
	ErrorUndefinedSymbol          = errors.New("Undefined symbol")
)

// 構文解析、字句解析のエラーメッセージの定義
const (
	ErrorUnmatchedParenthesis           = iota
	ErrorUnexpectedToken                = iota
	ErrorUnexpectedInputChar            = iota
	ErrorInsufficientInput              = iota
	ErrorFirstElementTypeMustBeASymbol  = iota
	ErrorStringLiteralMustBeASingleLine = iota
	ErrorIllegalEscapeSequence          = iota
	ErrorNotStringLiteral               = iota
	ErrorTopLevelElementMustBeAList     = iota
	ErrorMissingClosingParenthesis      = iota
)

var errorMessages map[int]string

func init() {
	errorMessages = map[int]string{
		ErrorUnmatchedParenthesis:           "Unmatched parenthesis",
		ErrorUnexpectedToken:                "Unexpected token",
		ErrorUnexpectedInputChar:            "Unexpected input char",
		ErrorInsufficientInput:              "Insufficient input",
		ErrorFirstElementTypeMustBeASymbol:  "First element type must be a symbol",
		ErrorStringLiteralMustBeASingleLine: "String literal must be a single line",
		ErrorIllegalEscapeSequence:          "Illegal escape sequence",
		ErrorNotStringLiteral:               "Not string literal",
		ErrorTopLevelElementMustBeAList:     "Top-level element must be a list",
		ErrorMissingClosingParenthesis:      "Missing closing parenthesis",
	}
}

// ParseError パース時のエラーメッセージを格納する
type ParseError struct {
	ErrorLocation Position
	ID            int
}

func (err *ParseError) Error() string {
	var b bytes.Buffer
	b.WriteString(err.ErrorLocation.Filename)
	b.WriteString("[")
	b.WriteString(strconv.FormatInt(int64(err.ErrorLocation.Line), 10))
	b.WriteString(",")
	b.WriteString(strconv.FormatInt(int64(err.ErrorLocation.Column), 10))
	b.WriteString("] ")
	b.WriteString(errorMessages[err.ID])
	return b.String()
}

func newError(filename string, line int, column int, messageid int) *ParseError {
	return &ParseError{Position{filename, line, column}, messageid}
}
