package main

import (
	"fmt"
	"example.com/v2/domain"
	"example.com/v2/repository/database"
)

func main() {
	dbConn := database.NewDB()
	defer fmt.Println("マイグレーションが正常に実行されました")
	defer database.CloseDB(dbConn)
	dbConn.AutoMigrate(&domain.Todo{})
}
