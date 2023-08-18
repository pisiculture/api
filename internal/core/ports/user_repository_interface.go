package ports

import (
	"github.com/pisiculture/internal/core/domain"
)

type UserRepositoryInterface interface {
	Save(usr domain.User) (int, error)
	DeleteById(id int) error
	FindByID(id int) (*domain.User, error)
}
