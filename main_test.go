package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSuccess(t *testing.T) {
	result := Sum(1, 2)
	expected := 3
	c := assert.New(t)
	c.Equal(expected, result)
	// if result != expected {
	// 	t.Errorf("Expected %d but got %d", expected, result)
	// }
}
