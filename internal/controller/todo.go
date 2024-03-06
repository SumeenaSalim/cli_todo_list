package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/todo_list/internal/consts"
	"github.com/todo_list/internal/view"
)

func TodoListController() {
	var index int
	var task, newtask string

	fmt.Println("1. Add a task \n2. List tasks \n3. Mark a task a completed \n4. Edit a task \n5. Delete a task \n6. Exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		inp := scanner.Text()

		choice, err := strconv.Atoi(inp)
		if err != nil {
			fmt.Println(consts.InvalidChoice)
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter task: ")
			scanner.Scan()
			task = scanner.Text()

			view.AddTask(task)
		case 2:
			view.ListTask()
		case 3:
			fmt.Print(consts.GetIndex)
			scanner.Scan()

			index, _ = strconv.Atoi(scanner.Text())
			view.MarkCompleted(index)
		case 4:
			fmt.Print(consts.GetIndex)
			scanner.Scan()
			index, _ = strconv.Atoi(scanner.Text())

			fmt.Print("Enter task: ")
			scanner.Scan()
			newtask = scanner.Text()	
			view.EditTask(index, newtask)
		case 5:
			fmt.Print(consts.GetIndex)
			scanner.Scan()
			index, _ = strconv.Atoi(scanner.Text())

			view.DeleteTask(index)
		case 6:
			fmt.Println("Exiting the application...")
			os.Exit(0)
		default:
			fmt.Println(consts.InvalidChoice)
		}
	}
}