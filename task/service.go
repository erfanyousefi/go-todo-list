package task;

type TaskService interface {
	AddTask(task Task) error;
	GetTasks() []Task;
	CompletedTask(id int) error;
	DeleteTask(id int) error;
	UpdateTask(id int, updatedTask Task) error;
}