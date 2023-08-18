package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pisiculture/infra/mock"
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

	t.Run("New_user_service_success", func(t *testing.T) {
		assert.NotNil(t, serv)
	})

	t.Run("New_user_sucess", func(t *testing.T) {

		usr := generateUserVO()
		r.EXPECT().Create(usr).Return(1, nil)

		id, err := serv.Create(usr)

		assert.Equal(t, 1, id)
		assert.Nil(t, err)
	})

	t.Run("New_user_sucess", func(t *testing.T) {

		usr := generateUserVO()
		usr.Name = ""
		r.EXPECT().Create(usr).Return(1, nil)

		id, err := serv.Create(usr)

		assert.Equal(t, id, 0)
		assert.NotNil(t, err)

	})
}

func generateUserVO() *vo.UserVO {
	usr := vo.NewUserVO()
	usr.Name = "Djonatan Willenz"
	usr.Email = "djonatanhorus@gmail.com"
	usr.Password = "werwnroweirw"
	return usr
}
