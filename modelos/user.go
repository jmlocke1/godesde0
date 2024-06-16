package modelos

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (usuario *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	usuario.Id = id
	usuario.Name = name
	usuario.CreatedAt = createdAt
	usuario.Status = status
}

func AddUser2(id int, name string, createdAt time.Time, status bool) *User {
	return &User{
		Id:        id,
		Name:      name,
		CreatedAt: createdAt,
		Status:    status,
	}
}
