package models

type Users struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	CreatedDay string `json:"created-day"`
	Nickname   string `json:"nickname"`
	Password   string `json:"password"`
}
