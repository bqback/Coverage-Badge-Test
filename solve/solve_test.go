package solve

import (
	"calc/parse"
	"calc/utils"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSuccess(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		input  string
		output float64
	}{
		{ // Addition
			name:   "Addition",
			input:  "1.25+1.75",
			output: 3,
		},
		{ // Substraction
			name:   "Substraction",
			input:  "2-1.75",
			output: 0.25,
		},
		{ // Multiplication
			name:   "Multiplication",
			input:  "2*2",
			output: 4,
		},
		{ // Division
			name:   "Division",
			input:  "5/2",
			output: 2.5,
		},
		{ // Equal order (addition and substraction)
			name:   "Equal order (addition and substraction)",
			input:  "1+2-3",
			output: 0,
		},
		{ // Equal order (multiplication and division)
			name:   "Equal order (multiplication and division)",
			input:  "1.5*2/3",
			output: 1,
		},
		{ // Parenthesis 1
			name:   "Parenthesis 1",
			input:  "(1+2)-3",
			output: 0,
		},
		{ // Parenthesis 2
			name:   "Parenthesis 2",
			input:  "(1+2)*3",
			output: 9,
		},
		{ // Four to the floor
			name:   "Four to the floor",
			input:  "(9+3)/2*5",
			output: 30,
		},
		{ // Multiple parenthesis groups
			name:   "Multiple parenthesis groups",
			input:  "(13+8)/(5+2)*(20-17)",
			output: 9,
		},
		{ // Negative numbers
			name:   "Negative numbers",
			input:  "1*(-3)",
			output: -3,
		},
		{ // A single number input
			name:   "A single number input",
			input:  "1",
			output: 1,
		},
		{ // Division by zero (+Inf) 1
			name:   "Division by zero",
			input:  "1/0",
			output: math.Inf(1),
		},
		{ // Division by zero (-Inf) 1
			name:   "Division by zero",
			input:  "-1/0",
			output: math.Inf(-1),
		},
		{ // Division by zero (+Inf) 2
			name:   "Division by zero",
			input:  "((12*35)+(15/6)/0)*((167+29.23)/(35-35))",
			output: math.Inf(1),
		},
		{ // Bracket negation
			name:   "Bracket negation",
			input:  "-(1+1)",
			output: -2,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			result, err := Solve(test.input)
			require.NoError(t, err)
			require.Equalf(t, test.output, result,
				fmt.Sprintf("results didn't match: expected %f, got %f", test.output, result),
			)
		})
	}
}

func TestFail(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         string
		output        float64
		expectedError error
	}{
		{ // Number parse fail 1
			name:          "Number parse fail 1",
			input:         "1?1",
			output:        0,
			expectedError: parse.ErrNumberParseFail,
		},
		{ // Number parse fail 2
			name:          "Number parse fail 2",
			input:         "2*3+(1?7.25)/(5#10)",
			output:        0,
			expectedError: parse.ErrNumberParseFail,
		},
		{ // Unbalanced parenthesis (left)
			name:          "Unbalanced parenthesis (left)",
			input:         "((1+1)",
			output:        0,
			expectedError: parse.ErrUnbalancedParenthesis,
		},
		{ // Unbalanced parenthesis (right)
			name:          "Unbalanced parenthesis (right)",
			input:         "(1+1))",
			output:        0,
			expectedError: parse.ErrUnbalancedParenthesis,
		},
		{ // Too many operators
			name:          "Too many operators",
			input:         "1/*+-2",
			output:        0,
			expectedError: parse.ErrBadFormatOperators,
		},
		{ // Empty input
			name:          "Empty input",
			input:         "",
			output:        0,
			expectedError: parse.ErrEmptyString,
		},
		{ // Empty input with brackets
			name:          "Empty input with brackets",
			input:         "((()))",
			output:        0,
			expectedError: ErrBadPolishNotation,
		},
		{ // Operator-only input
			name:          "Operator-only input",
			input:         "+",
			output:        0,
			expectedError: utils.ErrEmptyStack,
		},
		{ // Unfinished expression (left hand)
			name:          "Unfinished expression (left hand)",
			input:         "3*5+(17-7)/2+1+",
			output:        0,
			expectedError: utils.ErrEmptyStack,
		},
		{ // Unfinished expression (right hand)
			name:          "Unfinished expression (right hand)",
			input:         "*2",
			output:        0,
			expectedError: utils.ErrEmptyStack,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := Solve(test.input)
			require.ErrorIsf(t, err, test.expectedError, "Unexpected error")
		})
	}
}
