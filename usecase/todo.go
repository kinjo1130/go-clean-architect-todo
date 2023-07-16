package usecase

import (
	"example.com/v2/domain"
	"example.com/v2/repository"
)

type TodoUsecase interface {
	GetAllTodos() ([]domain.Todo, error)
	GetTodoById(todoId uint) (domain.Todo, error)
	CreateTodo(todo domain.Todo) error
	UpdateTodo(todoId uint, todo domain.Todo) error
	DeleteTodo(todoId uint) error
}

type todoUsecaseImpl struct {
	tr repository.TodoRepository
}

func NewTodoUsecase(tr repository.TodoRepository) TodoUsecase {
	return &todoUsecaseImpl{tr}
}

func (tu *todoUsecaseImpl) GetAllTodos() ([]domain.Todo, error) {
	todos := []domain.Todo{}
	if err := tu.tr.GetAllTodos(&todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (tu *todoUsecaseImpl) GetTodoById(todoId uint) (domain.Todo, error) {
	todo := domain.Todo{}
	if err := tu.tr.GetTodoById(todoId, &todo); err != nil {
		return domain.Todo{}, err
	}

	return todo, nil
}

func (tu *todoUsecaseImpl) CreateTodo(todo domain.Todo) error {
	if err := tu.tr.CreateTodo(&todo); err != nil {
		return err
	}

	return nil
}

func (tu *todoUsecaseImpl) UpdateTodo(todoId uint, todo domain.Todo) error {
	if err := tu.tr.UpdateTodo(todoId, &todo); err != nil {
		return err
	}

	return nil
}

func (tu *todoUsecaseImpl) DeleteTodo(todoId uint) error {
	if err := tu.tr.DeleteTodo(todoId); err != nil {
		return err
	}
	return nil
}
