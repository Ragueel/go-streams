package dbmodel

import "github.com/google/uuid"

type User struct {
	id       uuid.UUID
	username string
	password string
	isAdmin  bool
}
