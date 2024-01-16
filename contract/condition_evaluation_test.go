package contract

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"sync"
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
		assertNoLogs(t, func() {
			FulfilledfulConditionEvaluation.LogOnFailure()
		})
		assertLogContains(t, "the condition is not fulfilled", func() {
			unfulfilledConditionEvaluation.LogOnFailure()
		})
	})

	t.Run("should assert if the condition is fulfilled", func(t *testing.T) {
		mockT := new(testing.T)
		assert.True(t, FulfilledfulConditionEvaluation.AssertFulfilled(mockT))
		assert.False(t, unfulfilledConditionEvaluation.AssertFulfilled(mockT))
	})

	t.Run("should assert if the condition is unfulfilled", func(t *testing.T) {
		mockT := new(testing.T)
		assert.True(t, unfulfilledConditionEvaluation.AssertUnfulfilled(mockT))
		assert.False(t, FulfilledfulConditionEvaluation.AssertUnfulfilled(mockT))
	})
}

func assertNoLogs(t *testing.T, f func()) {
	logs := readLogs(t, f)
	assert.Empty(t, logs)
}

func assertLogContains(t *testing.T, expectedLog string, f func()) {
	logs := readLogs(t, f)
	assert.True(t, strings.Contains(logs, expectedLog))
}

func readLogs(t *testing.T, f func()) string {
	reader, writer, err := os.Pipe()
	assert.NoError(t, err)

	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}
