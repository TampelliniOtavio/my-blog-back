package encrypt

import (
	"testing"

	"github.com/TampelliniOtavio/my-blog-back/internal/infrastructure/util"
	"github.com/stretchr/testify/assert"
)

func Test_hashPassword_should_encrypt(t *testing.T) {
	assert := assert.New(t)

	password := util.RandomString(12)

	hashed, err := HashPassword(password)

	assert.Nil(err)
	assert.True(hashed != password)
}

func Test_hashPassword_should_error(t *testing.T) {
	assert := assert.New(t)

	_, err := HashPassword(util.RandomString(100))

	assert.NotNil(err)
}

func Test_verifyPassword_should_pass(t *testing.T) {
	assert := assert.New(t)

	password := util.RandomString(12)
	hashed, err := HashPassword(password)

	assert.Nil(err)
	assert.True(hashed != password)

	verified := VerifyPassword(password, hashed)

	assert.True(verified)
}

func Test_verifyPassword_should_not_pass(t *testing.T) {
	assert := assert.New(t)

	password := util.RandomString(12)

	hashed, err := HashPassword(password)

	assert.Nil(err)
	assert.True(hashed != password)

	verified := VerifyPassword(util.RandomString(2), hashed)

	assert.False(verified)
}
