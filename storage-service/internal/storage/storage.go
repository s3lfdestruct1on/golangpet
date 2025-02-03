package storage

type User struct{

	ID uint `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
}

func (u *User)Username() string{
	return u.username
}

func (u *User) Password() string{
	return u.password
}