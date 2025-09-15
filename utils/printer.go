package utils

import (
	"fmt"

	"github.com/erfanyousefi/go-todo-project/task"
)

func PrintTask(tasks []task.Task) string {
	result := ""
	for _,t := range tasks {
		status := "❌"
		if(t.Completed) {
			status = "✅"
		}
		result += fmt.Sprintf("%d. %s - %s [%s] \n", t.ID, t.Title, t.Description, status)
	}
	return result
}