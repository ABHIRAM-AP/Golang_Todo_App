package repo

import (
	"todo-app/config"
	task_model "todo-app/models"
)

func GetAllTasks(completedFilter *bool) ([]task_model.Task, error) {
	var tasks []task_model.Task
	var err error

	if completedFilter != nil {
		err = config.DB.Select(&tasks, `
			SELECT * FROM task WHERE completed = $1 ORDER BY created_at DESC
		`, *completedFilter)
	} else {
		err = config.DB.Select(&tasks, "SELECT * FROM task ORDER BY created_at DESC")
	}

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

func UpdateTask(task *task_model.Task) (*task_model.Task, error) {
	query := `
		UPDATE task
		SET title = $1, description = $2, completed = $3, updated_at = CURRENT_TIMESTAMP
		WHERE id = $4
		RETURNING id, title, description, completed, created_at, updated_at
	`

	var updated task_model.Task
	err := config.DB.QueryRowx(query, task.Title, task.Description, task.Completed, task.ID).
		StructScan(&updated)

	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func SetTaskCompletion(id int, status bool) (*task_model.Task, error) {
	query := `
		UPDATE task 
		SET completed = $1, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $2 
		RETURNING id, title, description, completed, created_at, updated_at;
	`
	var updated task_model.Task
	err := config.DB.Get(&updated, query, status, id)
	return &updated, err
}
