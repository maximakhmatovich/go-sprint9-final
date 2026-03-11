package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{42}, 42},
		{"multiple elements", []int{1, 5, 3, 9, 2}, 9},
	}
	for _, v := range tests {
		actual := maximum(v.data)

		// Используется require для проверки, так как если функция не пройдет тест,
		// то тест TestMaxChunks не имеет смысла
		require.Equal(t, v.want, actual, "Test with %s", v.name)
	}
}

func TestMaxChunks(t *testing.T) {
	data := generateRandomElements(123456)
	max := maximum(data)

	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{42}, 42},
		{"less than chunks", []int{3, 1, 9}, 9},
		{"exactly chunks", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"not divisible by chunks", []int{1, 5, 2, 8, 3, 9, 4, 6, 7, 10}, 10},
		{"negative numbers", []int{-10, -3, -50, -1}, -1},
		{"all equal", []int{7, 7, 7, 7, 7}, 7},
		{"random data", data, max},
	}

	for _, test := range tests {
		actual := maxChunks(test.data)
		assert.Equal(t, test.want, actual, "Test with %s", test.name)
	}
}
