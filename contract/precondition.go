package contract

import "fmt"

// RequiresThat is used to impose a precondition on a type using a condition.
func RequiresThat[T any](value T, specification Condition[T]) ConditionEvaluation {
	if err := specification(value); err != nil {
		return ConditionEvaluation{err: fmt.Errorf("precondition failed: %w", err)}
	}
	return FulfilledfulConditionEvaluation
}

// Requires is used to impose a precondition using an expression.
func Requires(constraint bool) ConditionEvaluation {
	return ConditionEvaluation{
		err: Evaluate(constraint, "precondition failed"),
	}
}
