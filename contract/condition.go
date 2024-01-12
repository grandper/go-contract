package contract

// Condition are used to capture assumptions on objects.
// When a condition is not fulfilled it resturns an error.
type Condition[T any] func(T) error

// And is used to combine the condition with an other.
// The resulting condition is fulfilled if both conditions are fulfilled.
func (c Condition[T]) And(other Condition[T]) Condition[T] {
	return func(value T) error {
		if err := c(value); err != nil {
			return err
		}
		return other(value)
	}
}

// Or is used to combine the condition with an other.
// The resulting condition is fulfilled if one of the two conditions is fulfilled.
func (c Condition[T]) Or(other Condition[T]) Condition[T] {
	return func(value T) error {
		if err := c(value); err == nil {
			return nil
		}
		return other(value)
	}
}
