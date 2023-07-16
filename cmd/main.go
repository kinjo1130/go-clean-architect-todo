package main

import (
	"example.com/v2/deliver/handlers"
	"example.com/v2/deliver/routes"
	"example.com/v2/repository"
	"example.com/v2/repository/database"
	"example.com/v2/usecase"
)

func main() {
	db := database.NewDB()
	todoRepository := repository.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	todoHandler := handlers.NewTodoHandler(todoUsecase)
	e := routes.NewTodoRouter(todoHandler)
	e.Logger.Fatal(e.Start(":8080"))
}