package service_test

import (
	"testing"
	"time"

	mock_ports "github.com/pisiculture/infra/mock"
	"github.com/pisiculture/internal/core/service"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

const name = "Djontan Willenz"
const email = "djonatanhorus@gmail.com"
const password = "djontatas"

func TestNewUserService(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	r := mock_ports.NewMockUserRepositoryInterface(ctrl)

	times := time.Now()

	r.EXPECT().Create(name, email, password, times).Return(1, nil)

	serv := service.NewUserService(r)

	assert.NotNil(t, serv)

	t.Run("new_user_sucess", func(t *testing.T) {
		id, err := serv.Create(name, email, password)

		assert.Equal(t, id, 1, "Test for id")
		assert.Nil(t, err, "Test for error is nil")
	})
}
