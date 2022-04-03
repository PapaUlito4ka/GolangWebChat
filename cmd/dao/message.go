package dao

import "time"

type MessageDao struct {
	User    UserDao
	Message string
	Time    time.Time
}