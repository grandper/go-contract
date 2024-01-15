package is

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	assert.NoError(t, Empty(""))
	assert.Error(t, Empty("foo"))
}

func TestNotEmpty(t *testing.T) {
	assert.NoError(t, NotEmpty("foo"))
	assert.Error(t, NotEmpty(""))
}

func TestNil(t *testing.T) {
	assert.NoError(t, Nil(nil))
	assert.Error(t, Nil("foo"))
}

func TestNotNil(t *testing.T) {
	assert.NoError(t, NotNil("foo"))
	assert.Error(t, NotNil(nil))
}

func TestZero(t *testing.T) {
	assert.NoError(t, Zero(0))
	assert.NoError(t, Zero(0.0))
	assert.Error(t, Zero(1))
	assert.Error(t, Zero(1.0))
}

func TestNotZero(t *testing.T) {
	assert.NoError(t, NotZero(1))
	assert.NoError(t, NotZero(1.0))
	assert.Error(t, NotZero(0))
	assert.Error(t, NotZero(0.0))
}

func TestPositive(t *testing.T) {
	assert.NoError(t, Positive(1))
	assert.NoError(t, Positive(1.0))
	assert.Error(t, Positive(0))
	assert.Error(t, Positive(0.0))
	assert.Error(t, Positive(-1))
	assert.Error(t, Positive(-1.0))
}

func TestNonpositive(t *testing.T) {
	assert.NoError(t, Nonpositive(0))
	assert.NoError(t, Nonpositive(0.0))
	assert.NoError(t, Nonpositive(-1))
	assert.NoError(t, Nonpositive(-1.0))
	assert.Error(t, Nonpositive(1))
	assert.Error(t, Nonpositive(1.0))
}

func TestNegative(t *testing.T) {
	assert.NoError(t, Negative(-1))
	assert.NoError(t, Negative(-1.0))
	assert.Error(t, Negative(0))
	assert.Error(t, Negative(0.0))
	assert.Error(t, Negative(1))
	assert.Error(t, Negative(1.0))
}

func TestNonnegative(t *testing.T) {
	assert.NoError(t, Nonnegative(0))
	assert.NoError(t, Nonnegative(0.0))
	assert.NoError(t, Nonnegative(1))
	assert.NoError(t, Nonnegative(1.0))
	assert.Error(t, Nonnegative(-1))
	assert.Error(t, Nonnegative(-1.0))
}

func TestBetween(t *testing.T) {
	condition := Between(2, 4)
	assert.NoError(t, condition(2))
	assert.NoError(t, condition(2.0))
	assert.NoError(t, condition(3))
	assert.NoError(t, condition(3.0))
	assert.NoError(t, condition(4))
	assert.NoError(t, condition(4.0))

	assert.Error(t, condition(1))
	assert.Error(t, condition(1.0))
	assert.Error(t, condition(5))
	assert.Error(t, condition(5.0))
}
