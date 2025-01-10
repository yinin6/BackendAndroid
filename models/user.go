package models

import (
	"database/sql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func CreateUser(db *sql.DB, user *User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := db.Exec(query, user.Username, user.Password)
	return err
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	row := db.QueryRow(query, username)

	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
