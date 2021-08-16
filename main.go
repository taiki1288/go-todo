package main

import (
	"fmt"
	"go-todo/app/models"
	"go-todo/app/controllers"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
