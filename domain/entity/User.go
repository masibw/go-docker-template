package entity

import "errors"

type User struct {
	ID   string
	Name string
}

var (
	ErrUserNotFound = errors.New("user not found")
)
