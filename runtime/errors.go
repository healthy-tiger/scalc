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
	ErrorInvalidOperation                                           int
	ErrorValueOutOfRange                                            int
)

var errorMessages map[int]string = make(map[int]string)

func init() {
	ErrorTheNumberOfArgumentsDoesNotMatch = RegisterEvalError("The number of arguments does not match(%v given, %v need)")
	ErrorUndefinedSymbol = RegisterEvalError("Undefined symbol %v")
	ErrorAnEmptyListIsNotAllowed = RegisterEvalError("An empty list is not allowed")
	ErrorTheFirstElementOfTheListToBeEvaluatedMustBeACallableObject = RegisterEvalError("The first element of the list to be evaluated must be a callable object: %v ")
	ErrorFunctionCannotBePassedAsFunctionArgument = RegisterEvalError("Function cannot be passed as function argument")
	ErrorInsufficientNumberOfArguments = RegisterEvalError("Insufficient number of arguments(%v given, %v need)")
	ErrorInvalidOperation = RegisterEvalError("Invalid Operation")
	ErrorValueOutOfRange = RegisterEvalError("Value out of range %v(%v to %v)")
}

// EvalError 実行時エラーの構造体
type EvalError struct {
	ErrorLocation parser.Position
	Message       string
}

// NewEvalError 式の評価の際に発生したエラーを表すオブジェクトを生成する。
func NewEvalError(loc parser.Position, id int, args ...interface{}) *EvalError {
	msg, ok := errorMessages[id]
	if !ok {
		panic("Undefined error id")
	}
	e := new(EvalError)
	e.ErrorLocation = loc
	e.Message = fmt.Sprintf(msg, args...)
	return e
}

func (err *EvalError) Error() string {
	h := fmt.Sprintf("%s:%d:%d ", err.ErrorLocation.Filename, err.ErrorLocation.Line, err.ErrorLocation.Column)
	return h + err.Message
}

// RegisterEvalError 実行時エラーのエラーメッセージを登録し、エラーメッセージのIDを返す。
func RegisterEvalError(msg string) int {
	n := len(errorMessages)
	errorMessages[n] = msg
	return n
}
