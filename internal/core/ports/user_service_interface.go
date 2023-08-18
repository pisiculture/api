package ports

import (
	"github.com/pisiculture/internal/core/vo"
)

type UserServiceInterface interface {
	Create(vo vo.UserVO) (int, error)
	Update(id int, vo *vo.UserVO) error
	FindByID(id int) (*vo.UserVO, error)
	DeleteById(id int) error
}
