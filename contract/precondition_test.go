package contract

import (
	"errors"
	"testing"
)

func TestPrecondition(t *testing.T) {
	const condition = 3
	isFulfilled := Condition[int](func(i int) error {
		return nil
	})
	isUnfulfilled := Condition[int](func(i int) error {
		return errors.New("the condition is unfulfilled")
	})

	t.Run("should impose a condition on a type", func(t *testing.T) {
		RequiresThat(condition, isFulfilled).AssertFulfilled(t)
		RequiresThat(condition, isUnfulfilled).AssertUnfulfilled(t)
	})

	t.Run("should impose a condition on an expression", func(t *testing.T) {
		Requires(4 > 2).AssertFulfilled(t)
		Requires(2 > 4).AssertUnfulfilled(t)
	})
}
