package contract

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ConditionEvaluation represents the result when a condition is evaluated.
type ConditionEvaluation struct {
	err error
}

// NewConditionEvaluation creates a condition evaluation.
func NewConditionEvaluation(err error) ConditionEvaluation {
	return ConditionEvaluation{
		err: err,
	}
}

// FulfilledfulConditionEvaluation is the result of a condition that was fullfilled.
var FulfilledfulConditionEvaluation = ConditionEvaluation{
	err: nil,
}

// Fulfilled returns true if the condition was fulfilled.
func (c ConditionEvaluation) Fulfilled() bool {
	return c.err == nil
}

// Unfulfilled returns true if the condition was unfulfilled.
func (c ConditionEvaluation) Unfulfilled() bool {
	return c.err != nil
}

// PanicOnFailure will panic if the condition is unfulfilled.
func (c ConditionEvaluation) PanicOnFailure() {
	if c.Unfulfilled() {
		panic(c.err)
	}
}

// ErrorOnFailure will return an error if the condition is unfulfilled.
func (c ConditionEvaluation) ErrorOnFailure() error {
	if c.Unfulfilled() {
		return c.err
	}
	return nil
}

// LogOnFailure will log the error if the condition is unfulfilled.
func (c ConditionEvaluation) LogOnFailure() {
	if c.Unfulfilled() {
		log.Println(c.err)
	}
}

// AssertFulfilled asserts that the condition is fulfilled.
func (c ConditionEvaluation) AssertFulfilled(t *testing.T) bool {
	return assert.True(t, c.Fulfilled(), c.err)
}

// AssertUnulfilled asserts that the condition is unfulfilled.
func (c ConditionEvaluation) AssertUnfulfilled(t *testing.T) bool {
	return assert.True(t, c.Unfulfilled(), c.err)
}
