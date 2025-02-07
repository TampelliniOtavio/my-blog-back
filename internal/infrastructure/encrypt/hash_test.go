package encrypt

import (
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	password = "randomPassword"
)

func Test_hashPassword_should_encrypt(t *testing.T) {
	assert := assert.New(t)

	hashed, err := HashPassword(password)

	assert.Nil(err)
	assert.True(hashed != password)
}

func Test_hashPassword_should_error(t *testing.T) {
	assert := assert.New(t)

	_, err := HashPassword(faker.New().Lorem().Text(100))

	assert.NotNil(err)
}
