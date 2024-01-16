package contract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvariant(t *testing.T) {
	isPositive := func(i int) bool {
		return i > 0
	}

	t.Run("should panic if the invariant is violated", func(t *testing.T) {
		a := 2
		assert.Panics(t, func() {
			defer Invariant(&a, isPositive).PanicOnViolation()
			a = -1
		})
	})

	t.Run("should not panic when the invariant is Verified", func(t *testing.T) {
		a := 2
		assert.NotPanics(t, func() {
			defer Invariant(&a, isPositive).PanicOnViolation()
		})
	})

	t.Run("should log an error if the invariant is violated", func(t *testing.T) {
		assertLogContains(t, "violated invariant", func() {
			a := 2
			defer Invariant(&a, isPositive).LogOnViolation()
			a = -1
		})
	})

	t.Run("should log nothing if the invariant is verified", func(t *testing.T) {
		assertNoLogs(t, func() {
			a := 2
			defer Invariant(&a, isPositive).LogOnViolation()
		})
	})

	t.Run("should provide utility to test violation", func(t *testing.T) {
		mockT := new(testing.T)
		a := 2
		invariant1 := Invariant(&a, isPositive)
		defer func() {assert.True(t, invariant1.AssertViolated(mockT))}()
		a = -1
		invariant2 := Invariant(&a, isPositive)
		defer func() {assert.False(t, invariant2.AssertViolated(mockT))}()
	})

	t.Run("should provide utility to test that no violation happen", func(t *testing.T) {
		mockT := new(testing.T)
		a := 2
		invariant1 := Invariant(&a, isPositive)
		defer func() {assert.False(t, invariant1.AssertVerified(mockT))}()
		a = -1
		invariant2 := Invariant(&a, isPositive)
		defer func() {assert.True(t, invariant2.AssertVerified(mockT))}()
	})

	t.Run("should check immutability for", func(t *testing.T) {
		t.Run("types", func(t *testing.T) {
			a := 2
			defer Immutable(&a).AssertVerified(t)
			b := 2
			defer Immutable(&b).AssertViolated(t)
			b = -1
		})
		t.Run("struct with unexported fields", func(t *testing.T) {
			type TestStruct struct{
				B int
				b int
			}
			a1 := TestStruct{}
			defer Immutable(&a1).AssertViolated(t)
			a1.b = -1
			a2 := TestStruct{}
			defer Immutable(&a2).AssertViolated(t)
			a2.B = -1
			a3 := TestStruct{}
			defer Immutable(&a3).AssertVerified(t)
		})
	})
}
