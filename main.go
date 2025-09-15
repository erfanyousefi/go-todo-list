package main

import (
	"fmt"

	"github.com/erfanyousefi/go-todo-project/task"
	"github.com/erfanyousefi/go-todo-project/utils"
)

func main() {
	taskService := task.NewTaskMemoryStore()
	newTaskList := []task.Task{
		{
			Title:       "New Task",
			Description: "This is a new task",
			Completed:   false,
			DueDate:     "",
		},
		{
			Title:       "New Task 2",
			Description: "This is a new task 2",
			Completed:   false,
			DueDate:     "",
		},
		{
			Title:       "New Task 3",
			Description: "This is a new task 3",
			Completed:   false,
			DueDate:     "",
		},
	}
	for _, task := range newTaskList {
		taskService.AddTask(task)
	}
	taskService.UpdateTask(2, task.Task{
		Title:       "Updated Task 2",
		Description: "This is an updated task 2",
	})
	taskService.CompletedTask(3)
	taskService.CompletedTask(1)
	taskService.DeleteTask(1)
	tasks := taskService.GetTasks()
	result := utils.PrintTask(tasks)
	fmt.Println(result)
}
