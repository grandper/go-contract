package contract

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCondition(t *testing.T) {
	isFulfilled := Condition[int](func(i int) error {
		return nil
	})
	isUnfulfilled := Condition[int](func(i int) error {
		return errors.New("the condition is unfulfilled")
	})

	t.Run("combined with other condition", func(t *testing.T) {
		condition := isFulfilled.And(isFulfilled)
		assert.NoError(t, condition(1))

		condition = isFulfilled.And(isUnfulfilled)
		assert.Error(t, condition(1))

		condition = isUnfulfilled.And(isFulfilled)
		assert.Error(t, condition(1))

		condition = isUnfulfilled.And(isUnfulfilled)
		assert.Error(t, condition(1))
	})

	t.Run("have an alternative condition", func(t *testing.T) {
		condition := isFulfilled.Or(isFulfilled)
		assert.NoError(t, condition(1))

		condition = isFulfilled.Or(isUnfulfilled)
		assert.NoError(t, condition(1))

		condition = isUnfulfilled.Or(isFulfilled)
		assert.NoError(t, condition(1))

		condition = isUnfulfilled.Or(isUnfulfilled)
		assert.Error(t, condition(1))
	})
}
