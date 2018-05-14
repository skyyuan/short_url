package utils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert.Equal(t, 62, Length, "Token length error")
}

func TestIdToString(t *testing.T) {
	id := 72
	expectValue := "1a"
	assert.Equal(t, expectValue, IdToString(id))
}

func TestStringToId(t *testing.T) {
	str := "1a"
	expectValue := 72
	assert.Equal(t, expectValue, StringToId(str))
}
