package tokens

const (
	Add  = "+"
	Subs = "-"
	Mult = "*"
	Div  = "/"
	Lpar = "("
	Rpar = ")"
)

var Order = map[string]int{
	Add:  2,
	Subs: 2,
	Mult: 1,
	Div:  1,
	Lpar: 3,
	Rpar: 3,
	"":   10,
}

var Operate = map[string]func(float64, float64) (float64, error){
	Add:  func(operand1 float64, operand2 float64) (float64, error) { return operand1 + operand2, nil },
	Subs: func(operand1 float64, operand2 float64) (float64, error) { return operand1 - operand2, nil },
	Mult: func(operand1 float64, operand2 float64) (float64, error) { return operand1 * operand2, nil },
	Div:  func(operand1 float64, operand2 float64) (float64, error) { return operand1 / operand2, nil },
}
