package contract

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConditionEvaluation(t *testing.T) {
	err := errors.New("the condition is not fulfilled")
	unfulfilledConditionEvaluation := NewConditionEvaluation(err)

	t.Run("should tell if the condition is fulfilled", func(t *testing.T) {
		assert.True(t, FulfilledfulConditionEvaluation.Fulfilled())
		assert.False(t, unfulfilledConditionEvaluation.Fulfilled())
	})

	t.Run("should tell if the condition is unfulfilled", func(t *testing.T) {
		assert.False(t, FulfilledfulConditionEvaluation.Unfulfilled())
		assert.True(t, unfulfilledConditionEvaluation.Unfulfilled())
	})

	t.Run("should panic if the condition is unfulfilled", func(t *testing.T) {
		assert.NotPanics(t, func() {
			FulfilledfulConditionEvaluation.PanicOnFailure()
		})
		assert.Panics(t, func() {
			unfulfilledConditionEvaluation.PanicOnFailure()
		})
	})

	t.Run("should return an error if the condition is unfulfilled", func(t *testing.T) {
		assert.NoError(t, FulfilledfulConditionEvaluation.ErrorOnFailure())
		assert.Error(t, unfulfilledConditionEvaluation.ErrorOnFailure())
	})

	t.Run("should log an error if the condition is unfulfilled", func(t *testing.T) {
		
	})

	t.Run("should assert if the condition is fulfilled", func(t *testing.T) {
		
	})

	t.Run("should assert if the condition is unfulfilled", func(t *testing.T) {
		
	})
}
