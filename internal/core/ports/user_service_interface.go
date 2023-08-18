package ports

import "github.com/pisiculture/internal/core/domain"

type UserServiceInterface interface {
	Create(name, email, password string) (int, error)
	Update(id int, name, email, password string) error
	FindByID(id int) (*domain.UserDomain, error)
	DeleteById(id int) error
}
