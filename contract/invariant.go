package contract

import (
	"errors"
	"testing"
)

var (
	ErrorViolatedInvariant = errors.New("violated invariant")
)

// Invariant is an invariant for a given type.
func Invariant[T any, O comparable](i *T, property func(t T) O) invariant {
	var reference O = property(*i)
	return invariant{
		checkFunc: func() ConditionEvaluation {
			if reference == property(*i) {
				return FulfilledfulConditionEvaluation
			}
			return NewConditionEvaluation(ErrorViolatedInvariant)
		},
	}
}

type invariant struct {
	checkFunc func() ConditionEvaluation
}

// PanicOnViolation will panic if the invariant is violated.
func (i invariant) PanicOnViolation() {
	i.checkFunc().PanicOnFailure()
}

// LogOnViolation will log an error if the invariant is violated.
func (i invariant) LogOnViolation() {
	i.checkFunc().LogOnFailure()
}

// AssertViolated will assert that the invariant is violated.
func (i invariant) AssertViolated(t *testing.T) bool {
	return i.checkFunc().AssertUnfulfilled(t)
}

// AssertVerfified will assert that the invariant is verified.
func (i invariant) AssertVerified(t *testing.T) bool {
	return i.checkFunc().AssertFulfilled(t)
}

// Immutable creates an invariant that check that the value is immutable.
func Immutable[T comparable](value *T) invariant {
	return Invariant(value, func(v T) T {
		return v
	})
}
