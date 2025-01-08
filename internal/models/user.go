package models

type ClerkUser struct {
	Id        string  `json:"id"`
	Email     string  `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

type User struct {
	ClerkUser
	IsAdmin bool `json:"is_admin"`
}

type UserHttp struct {
	ClerkUser  *ClerkUser
	DbUser     *User
	HttpStatus int
}
