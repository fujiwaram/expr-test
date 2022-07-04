package calc

import (
	"fmt"

	"github.com/antonmedv/expr"
)

type Env map[string]interface{}

func Calc(expression string, ev Env) (interface{}, error) {
	program, err := expr.Compile(expression)
	if err != nil {
		return nil, err
	}
	output, err := expr.Run(program, ev)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return output, nil
}
