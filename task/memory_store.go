package task

import (
	"errors"
	"time"
)

type TaskMemoryStore struct {
	store map[int]Task
	nextID int
}

func NewTaskMemoryStore() *TaskMemoryStore {
	return &TaskMemoryStore{
		store: make(map[int]Task),
		nextID: 1,
	}
}

func (s * TaskMemoryStore) AddTask(task Task) error {
	// Implementation for adding a task to memory store
	task.ID = s.nextID
	s.store[task.ID] = task;
	s.nextID++
	return nil
}
func (s *TaskMemoryStore) GetTasks() []Task {
	tasks := []Task{};
	for _, val :=range s.store {
		tasks = append(tasks, val)
	}
	return tasks;
}
func (s *TaskMemoryStore) DeleteTask(id int) error {
	task, ok := s.store[id];
	if(!ok) {
		return errors.New("task not found");
	}
	delete(s.store, task.ID);
	return nil
}
func (s *TaskMemoryStore) UpdateTask(id int, task Task)error {
	existTask, ok := s.store[id];
	if(!ok) {
		return errors.New("task not found");
	}
	existTask.Title = task.Title;
	existTask.Description = task.Description
	s.store[id] = existTask;
	return nil
}
func (s *TaskMemoryStore) CompletedTask(id int) error {
	task, ok := s.store[id];
	if !ok {
		return errors.New("task not found");
	}
	task.Completed = true;
	task.DueDate = time.Now().String();
	s.store[id] = task;
	return nil
}