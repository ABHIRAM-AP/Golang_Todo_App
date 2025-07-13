package repo

import (
	"todo-app/config"
	task_model "todo-app/models"
)

func GetAllTasks() ([]task_model.Task, error) {

	tasks := []task_model.Task{}
	err := config.DB.Select(&tasks, "SELECT * FROM task ORDER BY created_at DESC")

	return tasks, err
}

func CreateTask(task *task_model.Task) error {
	query := `
		INSERT INTO task (title, description, completed)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at`
	return config.DB.QueryRowx(query, task.Title, task.Description, task.Completed).Scan(
		&task.ID, &task.CreatedAt, &task.UpdatedAt,
	)
}

func DeleteTask(id int) error {
	_, err := config.DB.Exec("DELETE FROM task WHERE id=$1", id)
	return err
}
