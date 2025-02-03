package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct{

	ID int `json:"id"`
	username string `json:"username"`
	password string `json:"password"`
}

func (u *User)Username() string{
	return u.username
}

func (u *User) Password() string{
	return u.password
}

func CreateUser(username, pass, host, port, dbname string){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, pass, host,port, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO users (id, username, password) VALUES(0 'ostap', '123123');")
	
}