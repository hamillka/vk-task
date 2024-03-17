package models

const (
	ADMIN = 0
	USER  = 1
)

type User struct {
	Id       int
	Role     int // 0 - admin, 1 - user
	Login    string
	Password string
}
