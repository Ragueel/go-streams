package dbmodel

import "github.com/google/uuid"

type Movie struct {
	Id   uuid.UUID
	Name string
}

type MovieResource struct {
	Id      uuid.UUID
	MovieId uuid.UUID
	Path    string
}
