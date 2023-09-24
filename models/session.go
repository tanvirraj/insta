package models

import "database/sql"

type Session struct {
	ID     int
	UserId int

	Token     string
	TokenHash string
}

type SesstionService struct {
	DB *sql.DB
}

func (ss *SesstionService) Create(userID string) (*Session, error) {
	return nil, nil
}

func (ss *SesstionService) User(token string) (*User, error) {
	return nil, nil
}
