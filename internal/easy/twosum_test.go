package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {

	t.Run("example 1", func(t *testing.T) {
		arr := []int{2, 7, 11, 15}
		expected := []int{0, 1}
		target := 9
		assert.Equal(t, expected, twoSum(arr, target))
	})

	t.Run("example 2", func(t *testing.T) {
		arr := []int{3, 2, 4}
		expected := []int{1, 2}
		target := 6
		assert.Equal(t, expected, twoSum(arr, target))
	})

	t.Run("example 3", func(t *testing.T) {
		arr := []int{3, 3}
		expected := []int{0, 1}
		target := 6
		assert.Equal(t, expected, twoSum(arr, target))
	})
}
