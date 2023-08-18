package repository

import (
	"github.com/pisiculture/internal/core/vo"
)

type UserRepository struct {
}

func (usr *UserRepository) Save(vo *vo.UserVO) (int, error) {

	return 1, nil
}
