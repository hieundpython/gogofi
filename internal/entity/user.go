package entity

import "time"

type User struct {
	UserName  string
	UserEmail string
	FirstName string
	LastName  string
	CreateAt  time.Time
	DeleteAt  time.Time
	Active    bool
}
