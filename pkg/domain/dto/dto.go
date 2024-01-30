package dto

import "time"

type User struct {
	ID              string
	FirstName       string
	MiddleName      string
	LastName        string
	Email           string
	IsEmailVerified string
	PhoneNumber     string
	AccountStatus   string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Password        string
	LastLogin       time.Time
}
