package solve

import (
	"calc/parse"
	"calc/tokens"
	"calc/utils"
	"errors"
	"strconv"
	"strings"
)

var ErrBadPolishNotation = errors.New("malformed or empty polish notation stack")

func Solve(input string) (float64, error) {
	input = strings.ReplaceAll(input, " ", "")
	polishNotation, err := parse.Parse(input)
	if err != nil {
		return 0, err
	}
	var polishStack utils.Stack[float64]
	for _, token := range polishNotation {
		if !utils.IsOperator(token) {
			converted, _ := strconv.ParseFloat(token, 64)
			polishStack.Push(converted)
		} else {
			operand2, err := polishStack.Pop()
			if err != nil {
				return 0, err
			}
			operand1, err := polishStack.Pop()
			if err != nil {
				return 0, err
			}
			result, _ := tokens.Operate[token](operand1, operand2)
			polishStack.Push(result)
		}
	}
	if polishStack.Length() != 1 {
		return 0, ErrBadPolishNotation
	}
	result, _ := polishStack.Top()
	return result, nil
}
