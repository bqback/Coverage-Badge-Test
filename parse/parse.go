package parse

import (
	"calc/tokens"
	"calc/utils"
	"errors"
	"regexp"
	"strings"
)

var (
	ErrEmptyString           = errors.New("input string is empty")
	ErrNumberParseFail       = errors.New("failed to parse number from buffer")
	ErrUnbalancedParenthesis = errors.New("unbalanced parenthesis in expression")
	ErrBadFormatOperators    = errors.New("bad format: two or more non-parenthesis operators provided without being separated by a number")
)

func Parse(input string) ([]string, error) {
	if input == "" {
		return nil, ErrEmptyString
	}
	inputSplit := strings.Split(input, "")
	tokenList, err := tokenize(inputSplit)
	if err != nil {
		return nil, err
	}
	polishNotation, err := convertToPolish(tokenList)
	return polishNotation, err
}

func tokenize(inputSymbols []string) ([]string, error) {
	var (
		buf                  string
		lParCount, rParCount int
		operatorStreak       int
		output               []string
	)
	numericRe := regexp.MustCompile(`^-?\d+(\.\d*)?$`)
	for _, symbol := range inputSymbols {
		if utils.IsOperator(symbol) {
			if len(buf) > 0 {
				if !numericRe.MatchString(buf) {
					return nil, ErrNumberParseFail
				}
				output = append(output, buf)
				buf = ""
				operatorStreak = 0
			}
			switch symbol {
			case tokens.Add:
				output = append(output, tokens.Add)
				operatorStreak++
			case tokens.Subs:
				if len(output) == 0 || output[len(output)-1] == tokens.Lpar {
					output = append(output, "0")
					operatorStreak = 0
				}
				output = append(output, tokens.Subs)
				operatorStreak++
			case tokens.Mult:
				output = append(output, tokens.Mult)
				operatorStreak++
			case tokens.Div:
				output = append(output, tokens.Div)
				operatorStreak++
			case tokens.Lpar:
				output = append(output, tokens.Lpar)
				lParCount++
			case tokens.Rpar:
				output = append(output, tokens.Rpar)
				rParCount++
			}
			if operatorStreak > 1 {
				return nil, ErrBadFormatOperators
			}
		} else {
			buf += symbol
		}
	}

	if len(buf) > 0 {
		if !numericRe.MatchString(buf) {
			return nil, ErrNumberParseFail
		}
		output = append(output, buf)
	}

	if lParCount != rParCount {
		return nil, ErrUnbalancedParenthesis
	}

	return output, nil
}

func convertToPolish(input []string) ([]string, error) {
	var (
		stack  utils.Stack[string]
		output []string
	)
	for _, token := range input {
		if utils.IsOperator(token) {
			switch token {
			case tokens.Lpar:
				stack.Push(token)
			case tokens.Rpar:
				{
					for top, ok := stack.Top(); ok && (top != tokens.Lpar); top, ok = stack.Top() {
						topToken, _ := stack.Pop()
						output = append(output, topToken)
					}
					stack.Pop()
				}
			default:
				{
					top, ok := stack.Top()
					if (tokens.Order[token] < tokens.Order[top]) || !ok {
						stack.Push(token)
						continue
					}
					for top, ok := stack.Top(); ok && (tokens.Order[token] >= tokens.Order[top]); top, ok = stack.Top() {
						topToken, _ := stack.Pop()
						output = append(output, topToken)
					}
					stack.Push(token)
				}
			}
		} else {
			output = append(output, token)
		}
	}

	for !stack.Empty() {
		topToken, _ := stack.Pop()
		output = append(output, topToken)
	}

	return output, nil
}
