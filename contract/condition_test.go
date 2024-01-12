package contract

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCondition(t *testing.T) {
	fulfilledCondition := Condition[int](func(i int) error {
		return nil
	})
	unfulfilledCondition := Condition[int](func(i int) error {
		return errors.New("the condition is unfulfilled")
	})

	t.Run("combined with other condition", func(t *testing.T) {
		condition := fulfilledCondition.And(fulfilledCondition)
		assert.NoError(t, condition(1))

		condition = fulfilledCondition.And(unfulfilledCondition)
		assert.Error(t, condition(1))

		condition = unfulfilledCondition.And(fulfilledCondition)
		assert.Error(t, condition(1))

		condition = unfulfilledCondition.And(unfulfilledCondition)
		assert.Error(t, condition(1))
	})

	t.Run("have an alternative condition", func(t *testing.T) {
		condition := fulfilledCondition.Or(fulfilledCondition)
		assert.NoError(t, condition(1))

		condition = fulfilledCondition.Or(unfulfilledCondition)
		assert.NoError(t, condition(1))

		condition = unfulfilledCondition.Or(fulfilledCondition)
		assert.NoError(t, condition(1))

		condition = unfulfilledCondition.Or(unfulfilledCondition)
		assert.Error(t, condition(1))
	})
}
