package controller

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}
