package ports

import (
	"github.com/pisiculture/internal/core/domain"
)

type UserRepositoryInterface interface {
	Create(usr *domain.User) (int, error)
	Update(usr *domain.User) error
	DeleteById(id int) error
	FindByID(id int) (*domain.User, error)
}
