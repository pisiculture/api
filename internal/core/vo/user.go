package vo

type UserVO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserVO() *UserVO {
	return &UserVO{}
}
