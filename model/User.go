package model

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Movies   []Movie `json:"movies"`
}
