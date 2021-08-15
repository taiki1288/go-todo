package models

import "time"

type User struct {
	ID int
	UUID string
	Name string
	Email string
	Password string
	CreateAt time.Time
}