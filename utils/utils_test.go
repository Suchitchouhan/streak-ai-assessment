package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPair(t *testing.T) {
	response := FindPair([]int{1, 2, 3, 4, 5}, 6)
	expectedResponse := [][]int{{1, 4}, {0, 5}}
	assert.Equal(t, expectedResponse, response)

}
