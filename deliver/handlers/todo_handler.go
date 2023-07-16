package handlers

import (
	"example.com/v2/domain"
	"example.com/v2/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoHandler interface {
	GetAllTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoHandlerImpl struct {
	tu usecase.TodoUsecase
}

func NewTodoHandler(tu usecase.TodoUsecase) TodoHandler {
	return &todoHandlerImpl{tu}
}

func (th *todoHandlerImpl) GetAllTodos(c echo.Context) error {
	todos, err := th.tu.GetAllTodos()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

func (th *todoHandlerImpl) GetTodoById(c echo.Context) error {
	id := c.Param("todoId")
	todoId, _ := strconv.Atoi(id)
	todo, err := th.tu.GetTodoById(uint(todoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func (th *todoHandlerImpl) CreateTodo(c echo.Context) error {
	todo := domain.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := th.tu.CreateTodo(todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, todo)
}

func (th *todoHandlerImpl) UpdateTodo(c echo.Context) error {
	id := c.Param("todoId")
	todoId, _ := strconv.Atoi(id)

	todo := domain.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := th.tu.UpdateTodo(uint(todoId), todo); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func (th *todoHandlerImpl) DeleteTodo(c echo.Context) error {
	id := c.Param("todoId")
	todoId, _ := strconv.Atoi(id)

	err := th.tu.DeleteTodo(uint(todoId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
