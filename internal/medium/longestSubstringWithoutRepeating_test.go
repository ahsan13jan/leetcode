package medium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T)  {

	t.Run("example 1", func(t *testing.T) {
		assert.Equal(t,3, lengthOfLongestSubstring("abcabcbb"))
	})

	t.Run("example 2", func(t *testing.T) {
		assert.Equal(t,1, lengthOfLongestSubstring("bbbbb"))
	})

	t.Run("example 3", func(t *testing.T) {
		assert.Equal(t,3, lengthOfLongestSubstring("pwwkew"))
	})

	t.Run("fail 1", func(t *testing.T) {
		assert.Equal(t,1, lengthOfLongestSubstring(" "))
	})

	t.Run("fail 2", func(t *testing.T) {
		assert.Equal(t,3, lengthOfLongestSubstring("dvdf"))
	})

	t.Run("fail 3", func(t *testing.T) {
		assert.Equal(t,55, lengthOfLongestSubstring("hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789hijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))
	})

	t.Run("test", func(t *testing.T) {
		assert.Equal(t,3, lengthOfLongestSubstring("dvvdf"))
	})
}