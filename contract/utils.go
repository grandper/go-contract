package contract

import "fmt"

// Evaluates evaluates an inline condition.
func Evaluate(expression bool, errorFormat string, param ...any) error {
	if expression {
		return nil
	}
	return fmt.Errorf(errorFormat, param...)
}
