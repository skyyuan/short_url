package conf

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConfig_Read(t *testing.T) {
	read, e := Read("mongo_url")
	if e != nil {
		panic(e)
	}
	assert.Equal(t, "127.0.0.1:3000", read)
}
