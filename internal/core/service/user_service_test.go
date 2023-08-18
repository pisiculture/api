package service_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pisiculture/infra/mock"
	"github.com/pisiculture/internal/core/domain"
	"github.com/pisiculture/internal/core/service"
	"github.com/pisiculture/internal/core/vo"
	"github.com/stretchr/testify/assert"
)

const name = "Djontan Willenz"
const email = "djonatanhorus@gmail.com"
const password = "djontatas"

func TestNewUserService(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock.NewMockUserRepositoryInterface(ctrl)
	serv := service.NewUserService(r)

	t.Run("New_service_user_success", func(t *testing.T) {
		assert.NotNil(t, serv)
	})

	t.Run("New_user_sucess", func(t *testing.T) {

		usr := generateUserVO()
		r.EXPECT().Save(generateUser()).Return(1, nil)

		id, err := serv.Create(usr)
		assert.Equal(t, 1, id)
		assert.Nil(t, err)
	})

	t.Run("New_user_error_name_invalid", func(t *testing.T) {

		usr := generateUserVO()
		usr.Name = ""

		id, err := serv.Create(usr)

		assert.Equal(t, 0, id)
		assert.NotNil(t, err)
	})

	t.Run("New_user_error_email_invalid", func(t *testing.T) {

		usr := generateUserVO()
		usr.Email = ""

		id, err := serv.Create(usr)

		assert.Equal(t, id, 0)
		assert.NotNil(t, err)
	})

	t.Run("New_user_error_password_invalid", func(t *testing.T) {

		usr := generateUserVO()
		usr.Password = ""

		id, err := serv.Create(usr)

		assert.Equal(t, id, 0)
		assert.NotNil(t, err)
	})

	t.Run("New_user_error_database", func(t *testing.T) {

		usr := generateUserVO()
		r.EXPECT().Save(generateUser()).Return(0, errors.New("erro database"))

		id, err := serv.Create(usr)

		assert.Equal(t, id, 0)
		assert.NotNil(t, err)
	})
}

func generateUserVO() vo.UserVO {
	return vo.UserVO{
		Name:     "Djonatan Willenz",
		Email:    "djonatanhorus@gmail.com",
		Password: "werwnroweirw"}
}

func generateUser() domain.User {
	return domain.User{
		Name:     "Djonatan Willenz",
		Email:    "djonatanhorus@gmail.com",
		Password: "werwnroweirw"}
}
