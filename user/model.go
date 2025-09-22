package user

type User struct {
	ID       int64
	Username string
	Password string // bcrypt 哈希
}
