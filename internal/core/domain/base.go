package domain

import "time"

type Base struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
}
