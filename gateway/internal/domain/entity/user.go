package entity

type User interface {
	GetId() string
	GetUserName() string
	GetHashedPassword() string
	GetRole() string
}
