package repository

import "time"

type UserRepository struct {
}

func (usr *UserRepository) Create(name string, email string, password string, createdAt time.Time) (int, error) {

	return 1, nil
}
