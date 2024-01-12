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

	})

	t.Run("should log nothing if the invariant is verified", func(t *testing.T) {

	})

	t.Run("should provide utility to test violation", func (t *testing.T) {
		a := 2
		defer Invariant(&a, isPositive).AssertViolated(t)
		a = -1
	})

	t.Run("should provide utility to test that no violation happen", func (t *testing.T) {
		a := 2
		defer Invariant(&a, isPositive).AssertVerified(t)
	})

	t.Run("should check immutability for", func (t *testing.T) {
		t.Run("types", func(t *testing.T) {
			a := 2
			defer Immutable(&a).AssertVerified(t)
			b := 2
			defer Immutable(&b).AssertViolated(t)
			b = -1
		})
	})
}
