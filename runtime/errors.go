package runtime

import (
	"fmt"

	"github.com/healthy-tiger/scalc/parser"
)

// 共通のランタイムエラーIDの定義
var (
	ErrorTheNumberOfArgumentsDoesNotMatch                           int
	ErrorUndefinedSymbol                                            int
	ErrorAnEmptyListIsNotAllowed                                    int
	ErrorTheFirstElementOfTheListToBeEvaluatedMustBeACallableObject int
	ErrorFunctionCannotBePassedAsFunctionArgument                   int
	ErrorInsufficientNumberOfArguments                              int
	ErrorTooManyArguments                                           int
)

var errorMessages map[int]string = make(map[int]string)

func init() {
	ErrorTheNumberOfArgumentsDoesNotMatch = RegisterEvalError("The number of arguments does not match")
	ErrorUndefinedSymbol = RegisterEvalError("Undefined symbol")
	ErrorAnEmptyListIsNotAllowed = RegisterEvalError("An empty list is not allowed")
	ErrorTheFirstElementOfTheListToBeEvaluatedMustBeACallableObject = RegisterEvalError("The first element of the list to be evaluated must be a callable object: \"%v\" ")
	ErrorFunctionCannotBePassedAsFunctionArgument = RegisterEvalError("Function cannot be passed as function argument")
	ErrorInsufficientNumberOfArguments = RegisterEvalError("Insufficient number of arguments")
	ErrorTooManyArguments = RegisterEvalError("Too many arguments")
}

// EvalError 実行時エラーの構造体
type EvalError struct {
	ErrorLocation parser.Position
	ID            int
	Arg           interface{}
}

func newEvalError(loc parser.Position, id int, arg interface{}) *EvalError {
	if _, ok := errorMessages[id]; !ok {
		panic("Undefined error id")
	}
	return &EvalError{loc, id, arg}
}

func (err *EvalError) Error() string {
	h := fmt.Sprintf("%s:%d:%d ", err.ErrorLocation.Filename, err.ErrorLocation.Line, err.ErrorLocation.Column)
	m := errorMessages[err.ID]
	if err.Arg != nil {
		m = fmt.Sprintf(m, err.Arg)
	}
	return h + m
}

// RegisterEvalError 実行時エラーのエラーメッセージを登録し、エラーメッセージのIDを返す。
func RegisterEvalError(msg string) int {
	n := len(errorMessages)
	errorMessages[n] = msg
	return n
}
