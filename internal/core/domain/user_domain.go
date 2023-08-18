package domain

type UserDomain struct {
	Id       int
	Name     string
	Email    string
	Password string
}

func NewUserDomain() *UserDomain {
	return &UserDomain{}
}
