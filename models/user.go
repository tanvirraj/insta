package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

// to create a new entry into database we need db connection.
// one idomatic approcah is make a struct for userService and hold
// *db.DB into struct so that when we create new method to add or read
// data from database we can use that connection.
// how question is how Struct will get the connection

// when we bootup the application we create the connection in main.go file
// pass db connection as dependancy into Controller/ service

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(email, password string) (*User, error) {
	// we need grab the plaintext password
	// need to has the password with golang bcrypt lib
	// create new UserStruct type to hold the data (email, password)
	// Queqy db to create New User
	// scan new user and return

	email = strings.ToLower(email)
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	// hashedBytes is byte slice, we need to convert it into string to save in DB
	passwordHash := string(hashedBytes)

	user := User{
		Email:        email,
		PasswordHash: passwordHash,
	}

	row := us.DB.QueryRow(` INSERT INTO users(email, password_hash)
	VALUES($1, $2) RETURNING ID `, email, passwordHash)

	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create User:  %w", err)
	}

	return &user, nil
}
