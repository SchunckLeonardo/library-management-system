package value_objects_test

import (
	value_objects "github.com/SchunckLeonardo/library-management-system/pkg/value-objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHashPassword(t *testing.T) {
	hashPassword, err := value_objects.NewHashPassword("password")
	assert.Nil(t, err)

	assert.NotEmpty(t, hashPassword.GetValue())
	assert.NotEqual(t, hashPassword.GetValue(), "password")
}

func TestHashPassword_Compare(t *testing.T) {
	hashPassword, err := value_objects.NewHashPassword("password")
	assert.Nil(t, err)

	err = hashPassword.Compare("password")
	assert.Nil(t, err)

	err = hashPassword.Compare("wrong_password")
	assert.NotNil(t, err)
}

func TestHashPassword_GetValue(t *testing.T) {
	hashPassword, err := value_objects.NewHashPassword("password")
	assert.Nil(t, err)

	assert.NotEmpty(t, hashPassword.GetValue())
	assert.NotEqual(t, hashPassword.GetValue(), "password")
}
