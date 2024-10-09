package model

type Movie struct {
	ID       string `json:"ID"`
	Title    string `json:"title"`
	Director string `json:"director"`
	UserID   string `json:"user_id"`
}
