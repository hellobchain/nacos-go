package user

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"userName"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
