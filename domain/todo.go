package domain

import "time"

type Todo struct {
	ID		uint `json:"id" gorm:"primary_key"`
	Title	string `json:"title" gorm:"not null"`
	Status	string `json:"status" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoRepository interface {
	Create(todo *Todo) error
	FindAll() ([]Todo, error)
	GetByID(id int) (Todo, error)
	update(todo *Todo) error
	Delete(id int) error
}