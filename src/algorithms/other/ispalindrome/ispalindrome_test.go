package ispalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, IsPalindrome("aba"))
	assert.Equal(t, true, IsPalindrome("repaper"))
	assert.Equal(t, true, IsPalindrome("step on no pets"))
	assert.Equal(t, true, IsPalindrome("topspot"))

	assert.Equal(t, false, IsPalindrome("hello"))
	assert.Equal(t, false, IsPalindrome("abab"))
	assert.Equal(t, false, IsPalindrome("12343211"))
	assert.Equal(t, false, IsPalindrome("no no"))
}
