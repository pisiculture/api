package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pisiculture/infra/mock"
	"github.com/pisiculture/internal/core/domain"
	"github.com/pisiculture/internal/core/service"
	"github.com/pisiculture/internal/core/vo"
	"github.com/stretchr/testify/assert"
)

func TestNewInstallationService(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	r := mock.NewMockInstallationRepositoryInterface(ctrl)

	ser := service.NewInstallationService(r)

	t.Run("New_installation_success", func(t *testing.T) {
		assert.NotNil(t, ser)
	})

	t.Run("Create_installation_success", func(t *testing.T) {

		r.EXPECT().Save(getDomain()).Return(1, nil)

		id, err := ser.Create(getVO())

		assert.Equal(t, 1, id)
		assert.Nil(t, err)
	})

	t.Run("Create_installation_error_request_name", func(t *testing.T) {

	})

}

func getVO() vo.InstallationVO {
	vo := *vo.NewInstallationVO()
	vo.Name = "Installation for test"
	vo.Description = "Installation for test By DJ"
	return vo
}

func getDomain() domain.Installation {
	domain := domain.NewInstallation()
	domain.Name = "Installation for test"
	domain.Description = "Installation for test"
	return *domain
}
