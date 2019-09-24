package runtime

import (
	"bytes"
	"strconv"

	"github.com/healthy-tiger/scalc/parser"
)

// 共通のランタイムエラーIDの定義
const (
	ErrorTheNumberOfArgumentsDoesNotMatch                           = iota
	ErrorUndefinedSymbol                                            = iota
	ErrorAnEmptyListIsNotAllowed                                    = iota
	ErrorTheFirstElementOfTheListToBeEvaluatedMustBeACallableObject = iota
	ErrorFunctionCannotBePassedAsFunctionArgument                   = iota
	ErrorInsufficientNumberOfArguments                              = iota
	ErrorTooManyArguments                                           = iota
)

// CommonRuntimeName ランタイムの共通部分を示す文字列
const CommonRuntimeName = "Common Runtime"

var commonRuntimeErrorMessages map[int]string

func init() {
	commonRuntimeErrorMessages = map[int]string{
		ErrorTheNumberOfArgumentsDoesNotMatch:                           "The number of arguments does not match",
		ErrorUndefinedSymbol:                                            "Undefined symbol",
		ErrorAnEmptyListIsNotAllowed:                                    "An empty list is not allowed",
		ErrorTheFirstElementOfTheListToBeEvaluatedMustBeACallableObject: "The first element of the list to be evaluated must be a callable object",
		ErrorFunctionCannotBePassedAsFunctionArgument:                   "Function cannot be passed as function argument",
		ErrorInsufficientNumberOfArguments:                              "Insufficient number of arguments",
		ErrorTooManyArguments:                                           "Too many arguments",
	}
}

// EvalError 実行時エラーの構造体
type EvalError struct {
	ErrorLocation parser.Position
	ExtName       string
	ID            int
	messages      map[int]string
}

func newEvalError(loc parser.Position, id int) *EvalError {
	if _, ok := commonRuntimeErrorMessages[id]; !ok {
		panic("Undefined error id")
	}
	return &EvalError{loc, CommonRuntimeName, id, commonRuntimeErrorMessages}
}

func (err *EvalError) Error() string {
	var b bytes.Buffer
	b.WriteString("[Runtime Error ")
	b.WriteString(err.ExtName)
	b.WriteString(" ")
	b.WriteString(err.ErrorLocation.Filename)
	b.WriteString(" ")
	b.WriteString(strconv.FormatInt(int64(err.ErrorLocation.Line), 10))
	b.WriteString(":")
	b.WriteString(strconv.FormatInt(int64(err.ErrorLocation.Column), 10))
	b.WriteString("] ")
	b.WriteString(err.messages[err.ID])
	return b.String()
}
