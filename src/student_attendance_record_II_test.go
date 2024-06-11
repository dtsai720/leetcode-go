package src_test

import (
	"testing"

	"github.com/leetcode/src"
	"github.com/stretchr/testify/assert"
)

func TestStudentAttendanceRecordII(t *testing.T) {
	t.Parallel()

	t.Run("Test 1", func(t *testing.T) {
		t.Parallel()
		n := 2
		expected := 8
		got := src.StudentAttendanceRecordII(n)
		assert.Equal(t, expected, got)
	})

	t.Run("Test 2", func(t *testing.T) {
		t.Parallel()
		n := 1
		expected := 3
		got := src.StudentAttendanceRecordII(n)
		assert.Equal(t, expected, got)
	})

	t.Run("Test 3", func(t *testing.T) {
		t.Parallel()
		n := 10101
		expected := 183236316
		got := src.StudentAttendanceRecordII(n)
		assert.Equal(t, expected, got)
	})
}
