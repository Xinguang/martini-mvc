package models

type User struct {
	Name     string       `json:"name"`
	User     string       `json:"user"`
	Password string       `json:"password"`
	Host     dbHostConfig `json:"host"`
}
