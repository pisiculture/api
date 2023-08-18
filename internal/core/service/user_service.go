package service

import (
	"errors"
	"time"

	"github.com/pisiculture/internal/core/domain"
	"github.com/pisiculture/internal/core/ports"
)

type UserService struct {
	r ports.UserRepositoryInterface
}

func NewUserService(r ports.UserRepositoryInterface) *UserService {
	return &UserService{r: r}
}

func (usr *UserService) Create(name, email, password string) (int, error) {
	if name == "" {
		return 0, errors.New("Param name is not valid.")
	}
	if email == "" {
		return 0, errors.New("Param e-mail is not valid.")
	}
	if password == "" || len(password) < 8 {
		return 0, errors.New("Param password is not valid.")
	}
	date := time.Now()
	return usr.r.Create(name, email, password, date)
}
func (usr *UserService) Update(id int, name, email, password string) error {
	return nil
}
func (usr *UserService) FindByID(id int) (*domain.UserDomain, error) {
	return domain.NewUserDomain(), nil
}
func (usr *UserService) DeleteById(id int) error {
	return nil
}
