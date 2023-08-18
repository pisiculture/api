package domain_test

import (
	"testing"

	"github.com/pisiculture/internal/core/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := domain.NewUser()
	assert.NotNil(t, user)
}

func TestUserValidate(t *testing.T) {

	name := "Djonatan Willenz"
	email := "djonatanhorus@gmail.com"
	password := "senhaTest"

	t.Run("create_user_domain_success", func(t *testing.T) {

		user := domain.NewUser()

		user.Name = name
		user.Email = email
		user.Password = password

		err := user.Validate()

		assert.Nil(t, err)
	})

	t.Run("create_user_domain_name_invalid", func(t *testing.T) {

		user := domain.NewUser()

		user.Name = ""
		user.Email = email
		user.Password = password

		err := user.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Param name is not valid.")
	})

	t.Run("create_user_domain_email_invalid", func(t *testing.T) {

		user := domain.NewUser()

		user.Name = name
		user.Email = ""
		user.Password = password

		err := user.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Param e-mail is not valid.")
	})

	t.Run("create_user_domain_password_invalid", func(t *testing.T) {

		user := domain.NewUser()

		user.Name = name
		user.Email = email
		user.Password = ""

		err := user.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "Param password is not valid.")
	})
}
