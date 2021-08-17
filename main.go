package main

import (
	"fmt"
	"go-todo/app/models"
	"go-todo/app/controllers"
)

func main() {
	fmt.Println(models.Db)
	// init関数を呼び出すために記述

	controllers.StartMainServer()
}
