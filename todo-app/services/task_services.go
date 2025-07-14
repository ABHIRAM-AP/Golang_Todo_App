package services

import (
	"errors"
	task_model "todo-app/models"
	"todo-app/repo"
)

func GetTasks(completedFilter *bool) ([]task_model.Task, error) {
	return repo.GetAllTasks(completedFilter)
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

func UpdateTask(task *task_model.Task) (*task_model.Task, error) {
	return repo.UpdateTask(task)
}

func SetTaskCompletion(id int, status bool) (*task_model.Task, error) {
	return repo.SetTaskCompletion(id, status)
}
