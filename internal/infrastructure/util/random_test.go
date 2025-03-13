package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandomString_size(t *testing.T) {
	assert := assert.New(t)

	for i := range(20) {
		str := RandomString(i)

		assert.Equal(len(str), i)
	}
}

func Test_RandomEmail_size(t *testing.T) {
	assert := assert.New(t)

	for nameSize := range(20) {
		for domainSize := range(20) {
			str := RandomEmail(nameSize, domainSize)

			assert.True(strings.Contains(str, "@"))
			assert.True(strings.Contains(str, ".com"))

			split := strings.Split(str, "@")

			assert.Equal(len(split), 2)

			assert.Equal(nameSize, len(split[0]))
			assert.Equal(domainSize + 4, len(split[1]))
		}
	}
}
