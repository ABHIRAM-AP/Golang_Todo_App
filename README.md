# 📝 To-Do REST API (Go + PostgreSQL + Docker)

This is a simple **To-Do List REST API** built using **Go**, **PostgreSQL**, and **Docker**, following a clean architecture approach. It demonstrates full CRUD operations, task completion toggling, and containerized deployment using Docker Compose.

---

## 🚀 Features

- Add a new task
- Get all tasks (with optional filter by completion)
- Update a task
- Delete a task
- Mark a task as complete/incomplete
- Clean architecture: Controllers, Services, Repositories
- PostgreSQL integration
- Dockerized setup (Go app + PostgreSQL + pgAdmin)
- Environment variable support with `.env` file

---

## 📁 Project Structure

- go-todo-app/
├── app/
│   ├── controllers/
│   ├── services/
│   ├── repositories/
│   ├── models/
│   ├── routes/
│   └── config/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go


- **Controllers**: Receive HTTP requests, call services, return responses.
- **Services**: Contain business rules like validation or logic.
- **Repositories**: Deal with raw SQL or ORM to perform database operations.
- **Models**: Define the Task struct and other DTOs.

---

## 🧪 API Endpoints

### ➕ Create a Task
```http
POST /tasks
Body (JSON):
{
  "title": "Learn Go",
  "description": "Finish the Golang tutorial"
}
GET /tasks
Query Params:
- completed=true/false
PUT /tasks/{id}
Body (JSON):
{
  "title": "Learn GoLang",
  "description": "Update the tutorial",
  "completed": true
}
PATCH /tasks/{id}/complete
Body (JSON):
{
  "completed": true
}
DELETE /tasks/{id}
