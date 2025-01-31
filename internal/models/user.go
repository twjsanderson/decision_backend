package models

type ClerkUser struct {
	Id        string  `json:"id"`
	Email     string  `json:"email"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}

type User struct {
	ClerkUser
	IsAdmin bool `json:"isAdmin"`
}

type UserHttp struct {
	ClerkUser  *ClerkUser
	DbUser     *User
	HttpStatus int
}

type UserPermissions struct {
	Id     int    `json:"id"`
	UserId string `json:"userId"`
	Max    int    `json:"max"`
}
