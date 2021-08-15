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

func (u *User) CreateUser() (err error) {
	cmd := `insert into users(
        uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`
	
	_, err := Db.Exec(cmd)
}