package is

import (
	"github.com/grandper/go-contract/contract"
)

// Ordered is a constraint that permits any ordered type.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// Empty ensures that a string is empty.
func Empty(str string) error {
	return contract.Evaluate(str == "", "string must be empty")
}

// NotEmpty ensures that a string is not empty.
func NotEmpty(str string) error {
	return contract.Evaluate(len(str) > 0, "string cannot be empty")
}

// Nil ensures that the value is nil.
func Nil(value any) error {
	return contract.Evaluate(value == nil, "value must be nil")
}

// NotNil ensures that the value is not nil.
func NotNil(value any) error {
	return contract.Evaluate(value != nil, "value cannot be nil")
}

// Zero ensures that a number is equal to zero.
func Zero[T Ordered](value T) error {
	return contract.Evaluate(value == zero[T](), "number must be zero")
}

// NotZero ensures that a number is not equal to zero.
func NotZero[T Ordered](value T) error {
	return contract.Evaluate(value != zero[T](), "number must be different from zero")
}

// Positive ensures that a number is positive.
func Positive[T Ordered](value T) error {
	return contract.Evaluate(value > zero[T](), "number must be positive")
}

// Nonpositive ensures that a number is nonpositive.
func Nonpositive[T Ordered](value T) error {
	return contract.Evaluate(value <= zero[T](), "number must be nonpositive")
}

// Negative ensures that a number is negative.
func Negative[T Ordered](value T) error {
	return contract.Evaluate(value < zero[T](), "number must be negative")
}

// Nonnegative ensures that a number is nonnegative.
func Nonnegative[T Ordered](value T) error {
	return contract.Evaluate(value >= zero[T](), "number must be nonnegative")
}

// Between ensures that a number is between two bounds. The bounds are included
func Between[T Ordered](from, to T) contract.Condition[T] {
	return func(value T) error {
		return contract.Evaluate(from <= value && value <= to, "number must be between %v and %v", from, to)
	}
}

func zero[T Ordered]() T {
	var z T
	return z
}
