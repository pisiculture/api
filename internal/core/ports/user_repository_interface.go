package ports

import "time"

type UserRepositoryInterface interface {
	Create(name, email, password string, createdAt time.Time) (int, error)
}
