package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mainProgram(t *testing.T) {
	t.Run("Test_Assert", func(t *testing.T) {
		assert.Equal(t, 3, mainProgram([]int{7, 1, 2, 3, 8, 6, 3, 2, 1}))
		assert.Equal(t, 2, mainProgram([]int{7, 1, 2, 9, 7, 2, 1}))
		assert.Equal(t, 2, mainProgram([]int{1, 2, 1, 2, 2, 1}))
		assert.Equal(t, 2, mainProgram([]int{1, 2, 3, 4, 5, 2, 1}))
		assert.Equal(t, 3, mainProgram([]int{1, 2, 3, 4, 5, 3, 2, 1}))
		// Pass Ok
	})
	t.Run("Test_BiasaPertama", func(t *testing.T) {
		expected := 3
		actual := mainProgram([]int{7, 1, 2, 3, 8, 6, 3, 2, 1})
		if expected != actual {
			t.Errorf("the output should be %d not %d", expected, actual)
		}
		// Pass Ok
	})
	t.Run("Test_BiasaKedua", func(t *testing.T) {
		expected := 2
		actual := mainProgram([]int{7, 1, 2, 9, 7, 2, 1})
		if expected != actual {
			t.Errorf("the output should be %d not %d", expected, actual)
		}
		// Pass Ok
	})
	t.Run("Test_BiasaKetiga", func(t *testing.T) {
		expected := 2
		actual := mainProgram([]int{1, 2, 1, 2, 2, 1})
		if expected != actual {
			t.Errorf("the output should be %d not %d", expected, actual)
		}
		// Pass Ok
	})
}

type MockData struct {
	ArrayInput       []int
	SequencialResult []int
	ReverseNumber    []int
	MaxArray         int
}

// Unit Test All workFlow
func Test_FlowProgram(t *testing.T) {
	mock := MockData{
		ArrayInput:       []int{7, 1, 2, 3, 8, 6, 3, 2, 1},
		SequencialResult: []int{1, 2, 3},
		ReverseNumber:    []int{3, 2, 1},
		MaxArray:         3,
	}

	t.Run("Test_FindSequencial", func(t *testing.T) {
		assert.Equal(t, mock.SequencialResult, findSequencial(mock.ArrayInput))
	})
	t.Run("Test_ReverseNumber", func(t *testing.T) {
		assert.Equal(t, mock.ReverseNumber, reverseNum(mock.SequencialResult))
	})
	t.Run("Test_FindReverse", func(t *testing.T) {
		assert.Equal(t, true, findReverse(mock.ArrayInput, mock.ReverseNumber))
	})
	t.Run("Test_FinalGetTheMaxNumber", func(t *testing.T) {
		assert.Equal(t, mock.MaxArray, mainProgram(mock.ArrayInput))
	})

}
