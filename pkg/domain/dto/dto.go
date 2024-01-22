package dto

import "time"

type User struct {
	ID              string
	FirstName       string
	MiddleName      string
	LastName        string
	Email           string
	IsEmailVerified string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Password        string
}
