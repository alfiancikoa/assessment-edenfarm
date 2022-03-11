package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArray(t *testing.T) {
	assert.Equal(t, 3, maxArray([]int{7, 1, 2, 3, 8, 6, 3, 2, 1}))
	assert.Equal(t, 2, maxArray([]int{7, 1, 2, 9, 7, 2, 1}))
	assert.Equal(t, 2, maxArray([]int{1, 2, 1, 2, 2, 1}))
}
