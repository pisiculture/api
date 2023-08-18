package service

import (
	"errors"

	"github.com/pisiculture/internal/core/domain"
	"github.com/pisiculture/internal/core/ports"
	"github.com/pisiculture/internal/core/util"
	"github.com/pisiculture/internal/core/vo"
)

type UserService struct {
	r ports.UserRepositoryInterface
}

func NewUserService(r ports.UserRepositoryInterface) *UserService {
	return &UserService{r: r}
}

func (usr *UserService) Create(vo vo.UserVO) (int, error) {

	user := domain.NewUser()
	user.Name = vo.Name
	user.Email = vo.Email
	user.Password = vo.Password

	err := user.Validate()
	if err != nil {
		return 0, err
	}

	err = util.ValidateEmail(user.Email)
	if err != nil {
		return 0, err
	}

	return usr.r.Save(*user)
}
func (usr *UserService) Update(id int, vo vo.UserVO) error {

	if id == 0 {
		return errors.New("Param id not is valid")
	}

	register, err := usr.r.FindByID(id)

	if err != nil {
		return errors.New("Register not found")
	}

	pass, err := util.Crypto(vo.Password)

	if err != nil {
		return err
	}

	register.Name = vo.Name
	register.Password = pass

	err = register.Validate()

	if err != nil {
		return err
	}
	_, err = usr.r.Save(*register)

	return err
}
func (usr *UserService) FindByID(id int) (*domain.User, error) {
	return domain.NewUser(), nil
}
func (usr *UserService) DeleteById(id int) error {
	return nil
}
