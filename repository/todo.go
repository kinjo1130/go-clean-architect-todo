package repository


import (
	"fmt"
	"example.com/v2/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TodoRepository interface {
	GetAllTodos(todos *[]domain.Todo) error
	GetTodoById(todoId uint, todo *domain.Todo) error
	CreateTodo(todo *domain.Todo) error
	UpdateTodo(todoId uint, todo *domain.Todo) error
	DeleteTodo(todoId uint) error
}

type todoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepositoryImpl{db}
}

func (tr *todoRepositoryImpl) GetAllTodos(todos *[]domain.Todo) error {
	if err := tr.db.Find(&todos).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepositoryImpl) GetTodoById(todoId uint, todo *domain.Todo) error {
	if err := tr.db.First(todo, todoId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepositoryImpl) CreateTodo(todo *domain.Todo) error {
	if err := tr.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepositoryImpl) UpdateTodo(todoId uint, todo *domain.Todo) error {
	result := tr.db.Model(todo).Clauses(clause.Returning{}).Where("id=?", todoId).Updates(map[string]interface{}{"title": todo.Title, "status": todo.Status})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *todoRepositoryImpl) DeleteTodo(todoId uint) error {
	result := tr.db.Where("id=?", todoId).Delete(&domain.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
