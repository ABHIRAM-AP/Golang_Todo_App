package services

import (
	"errors"
	task_model "todo-app/models"
	"todo-app/repo"
)

func GetTasks() ([]task_model.Task, error) {
	return repo.GetAllTasks()
}

func AddTask(task *task_model.Task) error {
	if task.Title == "" {
		return errors.New("Title cannot be empty")
	}
	return repo.CreateTask(task)
}

func RemoveTask(id int) error {
	return repo.DeleteTask(id)
}
