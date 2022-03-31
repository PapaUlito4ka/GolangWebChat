package dao

type UserDao struct {
	Id       int64
	Username string
	Password string
	Friends  []UserDao
}
