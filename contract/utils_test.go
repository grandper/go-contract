package contract

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Evaluate(t *testing.T) {
	assert.Equal(t, Evaluate(5 < 0, "we expect %d > 0", 5), errors.New("we expect 5 > 0"))
	assert.NoError(t, Evaluate(5 > 0, "we expect %d > 0", 5))
}
