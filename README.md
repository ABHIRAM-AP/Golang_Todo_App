# 📝 To-Do App - REST API with Go, PostgreSQL, and Docker

This is a simple, containerized **To-Do REST API** built using **Golang** for the backend and **PostgreSQL** as the database, following clean architecture principles.

##  Tech Stack

- **Go (Golang)** - Backend API
- **PostgreSQL** - Relational Database
- **Docker + Docker Compose** - Containerization
- **pgAdmin** - PostgreSQL GUI

---

## 🚀 Features

-  Create a new task
-  Get all tasks (with filter options)
-  Update an existing task
-  Delete a task
-  Mark tasks as complete/incomplete
-  Input validation & error handling

---

## 🧱 Architecture Overview (Clean Architecture)

```
/app
│
├── controllers    # Handles HTTP request/response
├── services       # Business logic
├── repositories   # Database interactions (CRUD)
├── models         # Task structure and types
├── routes         # API route handlers
├── config         # Database config & environment setup
└── main.go        # Entry point
```

- **Controllers**: Receive HTTP requests, call services, return responses.
- **Services**: Contain business rules like validation or logic.
- **Repositories**: Deal with raw SQL or ORM to perform database operations.
- **Models**: Define the Task struct and other DTOs.

---

##  API Endpoints

###  Create a Task
```http
POST /tasks
Body (JSON):
{
  "title": "Learn Go",
  "description": "Finish the Golang tutorial"
}
```

###  Get Tasks (with filter)
```http
GET /tasks
Query Params:
- completed=true/false
```

###  Update a Task
```http
PUT /tasks/{id}
Body (JSON):
{
  "title": "Learn GoLang",
  "description": "Update the tutorial",
  "completed": true
}
```

###  Mark Task as Complete/Incomplete
```http
PATCH /tasks/{id}/complete
Body (JSON):
{
  "completed": true
}
```

###  Delete a Task
```http
DELETE /tasks/{id}
```

---

## 🐳 Docker Setup

All services are containerized using Docker Compose.

### 🔧 docker-compose.yml
- `go-todo-app`: Go application
- `postgres`: PostgreSQL DB
- `pgadmin`: Admin dashboard


---

## ▶️ Running the App

### Step 1: Clone the repository
```bash
git clone https://github.com/your-username/go-todo-app.git
cd go-todo-app
```

### Step 2: Create `.env` file
```bash
cp .env.example .env
```

### Step 3: Start the app
```bash
docker-compose up --build
```

- API: `http://localhost:8080`
- pgAdmin: `http://localhost:8081`  
  - Login using email and password from `.env`
  - Add server using:
    - Host: `postgres`
    - Port: `5432`
    - Username: `postgres`
    - Password: from `.env`

---

##  Testing

Test all endpoints using **Postman** or **cURL**. Examples are provided above for each endpoint.

---

##  Task JSON Structure

```json
{
  "id": 1,
  "title": "Learn Go",
  "description": "Finish the Golang tutorial",
  "completed": false,
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

---

## 📂 Folder Structure

```
go-todo-app/
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
```

