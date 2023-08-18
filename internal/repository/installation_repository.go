package repository

import "github.com/pisiculture/internal/core/domain"

type InstallationRepository struct {
}

func (ins *InstallationRepository) Save(installation domain.User) (int, error) {
	return 0, nil
}

func (ins *InstallationRepository) DeleteById(id int) error {
	return nil
}

func (ins *InstallationRepository) FindByID(id int) (*domain.Installation, error) {
	return domain.NewInstallation(), nil
}
