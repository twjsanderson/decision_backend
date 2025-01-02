package models

type User struct {
	id   int    `json: id`
	name string `json: name`
}

func (u *User) GetUserById(id int) (User, error) {
	var newUser User

	// look for user in DB
	// if user in DB return it
	// else return error

	return newUser, nil
}
