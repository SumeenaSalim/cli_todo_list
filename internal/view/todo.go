package view

import (
	"fmt"

	"github.com/todo_list/internal/consts"
	"github.com/todo_list/internal/model"
)

var tasks []model.Todo

func AddTask(task string) {
	newTask := model.Todo{
		TaskName: task,
		Completed: false,
	}
	tasks = append(tasks, newTask)
	fmt.Println(consts.TaskAdded)
}

func ListTask() {
	for i, task := range tasks {
		status := "no"
		if task.Completed {
			status = "yes"
		}

		fmt.Printf("%d \t %s \t [%s]\n", i+1, task.TaskName, status)
	}
}

func MarkCompleted(index int) {
	if index >= 1 && index < len(tasks) {
		tasks[index - 1].Completed = true
		fmt.Println(consts.TaskCompleted)
 	}else {
		fmt.Println(consts.InvalidIndex)
	}
}

func EditTask(index int, newTask string) {
	if index >= 1 && index < len(tasks) {
		tasks[index - 1].TaskName = newTask
		fmt.Println(consts.TaskEdited)
 	}else {
		fmt.Println(consts.InvalidIndex)
	}
}

func DeleteTask(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks = append(tasks[:index-1], tasks[index:]...)
		fmt.Println(consts.TaskDeleted)
 	}else {
		fmt.Println(consts.InvalidIndex)
	}
}