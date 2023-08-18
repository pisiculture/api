package service

import (
	"github.com/pisiculture/internal/core/ports"
	"github.com/pisiculture/internal/core/vo"
)

type InstallationService struct {
	r ports.InstallationRepositoryInterface
}

func NewInstallationService(r ports.InstallationRepositoryInterface) *InstallationService {
	return &InstallationService{r: r}
}

func (inst *InstallationService) Create(vo vo.InstallationVO) (int, error) {
	return 1, nil
}
