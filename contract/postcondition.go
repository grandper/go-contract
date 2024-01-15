package contract

import "fmt"

// EnsuresThat is used to impose a postcondition on a type using a condition.
func EnsuresThat[T any](value T, specification Condition[T]) ConditionEvaluation {
	if err := specification(value); err != nil {
		return ConditionEvaluation{err: fmt.Errorf("postcondition failed: %w", err)}
	}
	return FulfilledfulConditionEvaluation
}

// Ensures is used to impose a postcondition using an expression.
func Ensures(constraint bool) ConditionEvaluation {
	return ConditionEvaluation{
		err: Evaluate(constraint, "postcondition failed"),
	}
}
