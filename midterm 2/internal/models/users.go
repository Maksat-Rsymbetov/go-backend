package models

import (
	"database/sql"
	"errors"
)

type User struct {
	username string
	password string
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) AddUser(username, password string) (int, error) {
	cmd := `insert into users(username, password) values(?, ?)`
	result, err := u.DB.Exec(cmd, username, password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (u *UserModel) GetUser(username, password string) (*User, error) {
	const cmd string = `select * from users where username = ? and password = ?`
	found := u.DB.QueryRow(cmd, username, password)

	user := &User{}
	err := found.Scan(&user.username, &user.password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return user, nil
}
