package main

import (
	"fmt"

	"github.com/todo_list/internal/controller"
)

func main() {
	fmt.Println("Welcome to the To-Do list application")
	controller.TodoListController()
}
