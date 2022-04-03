package dao

type ChatDao struct {
	Id     int64
	User   UserDao
	Friend UserDao
}
