package models

type User struct {
	UUID      string `json:"uuid"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Birthdate string `json:"birthdate"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}
