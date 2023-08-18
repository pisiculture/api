package ports

import "github.com/pisiculture/internal/core/domain"

type InstallationRepositoryInterface interface {
	Save(installation domain.User) (int, error)
	DeleteById(id int) error
	FindByID(id int) (*domain.Installation, error)
}
