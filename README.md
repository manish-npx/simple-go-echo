# 📝 Simple Go Echo Todo API

A simple and clean RESTful **Todo API** built with **Go**, **Echo framework**, and **PostgreSQL**.

---

## 🗂️ Project Structure

```bash
simple-go-echo/
├── cmd/
│   └── server/
│       └── main.go        # Application entry point
├── configs/
│   └── config.yaml        # Configuration file
├── internal/
│   ├── config/
│   │   └── config.go      # Configuration loader
│   ├── databse/
│   │   └── postgres.go    # Database connection and Configuration
│   ├── handlers/          # HTTP request handlers
│   │   └── todo.go        # Todo-related HTTP handlers
│   ├── models/            # Data models/structures
│   │   └── todo.go        # Todo model definition
│   └── storage/           # Database operations
│       └── todo.go        # Todo storage (database layer)
├── pkg/
│
├── go.mod                 # Go module dependencies
└── README.md              # Project documentation

```

---

## 🚀 Quick Start

### Prerequisites

- Go 1.19+
- PostgreSQL
- Git

### Installation

1. **Clone the repository**

   ```
   git clone <your-repo-url>
   cd simple-go-echo
   ```

2. **Set up PostgreSQL database**

   ```
   CREATE DATABASE todo_db;
   CREATE TABLE todos (
     id SERIAL PRIMARY KEY,
     title VARCHAR(255) NOT NULL,
     done BOOLEAN DEFAULT FALSE,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

3. **Configure the application**

   Create `configs/config.yaml`:

   ```
   server:
     port: 8080
     addr: ":8080"

   database:
     host: "localhost"
     port: 5432
     user: "your_username"
     password: "your_password"
     dbname: "todo_db"
     sslmode: "disable"
   ```

4. **Install dependencies**

   ```
   go mod tidy
   ```

5. **Run the application**
   ```
   go run cmd/server/main.go
   ```
   You should see:
   ```
   🚀 Starting server...
   ✅ Connected to PostgreSQL successfully
   🚀 Server running on :8080
   ```

---

## 📚 API Endpoints

| Method | Endpoint         | Description       | Body                                   |
| ------ | ---------------- | ----------------- | -------------------------------------- |
| GET    | `/api/todos`     | Get all todos     | -                                      |
| POST   | `/api/todos`     | Create a new todo | `{ "title": "Task", "done": false }`   |
| GET    | `/api/todos/:id` | Get todo by ID    | -                                      |
| PUT    | `/api/todos/:id` | Update todo by ID | `{ "title": "Updated", "done": true }` |
| DELETE | `/api/todos/:id` | Delete todo by ID | -                                      |

---

## 💻 Example Usage

**Create a todo:**

bash

    curl -X POST http://localhost:8080/api/todos \
    -H "Content-Type: application/json" \
    -d '{"title": "Learn Go", "done": false}'

- **Get all todos:**
  bash

      curl http://localhost:8080/api/todos

- **Update a todo:**
  bash

      curl -X PUT http://localhost:8080/api/todos/1 \
      -H "Content-Type: application/json" \
      -d '{"title": "Learn Go and Echo", "done": true}'

      Delete a todo:

  bash

      curl -X DELETE http://localhost:8080/api/todos/1

---

## 🏗️ Architecture

This project follows a simple and clean architecture:

    HTTP Request → Handler → Storage → PostgreSQL
    HTTP Response ← Handler ← Storage ← PostgreSQL

### Layers Overview

- **Handlers** (`internal/handlers/`)

  - Handle HTTP requests/responses
  - Input validation & error handling

- **Storage** (`internal/storage/`)

  - Database operations & SQL queries
  - Data persistence

- **Models** (`internal/models/`)

  - Data structures & JSON serialization

- **Config** (`internal/config/`)
  - Configuration management & env setup

---

## 🔧 Configuration

The application uses YAML configuration. Example (`configs/config.yaml`):

- yaml

  server:
  port: 8080
  addr: ":8080"

  database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "password"
  dbname: "todo_db"
  sslmode: "disable"

📦 Dependencies

    Echo - High performance web framework

    Pgx - PostgreSQL driver and toolkit

    YAML - Configuration parsing

🧪 Testing

Run the application and test with curl or Postman:
bash

# Test all endpoints

curl http://localhost:8080/api/todos

---

## 🛠️ Development

### Adding New Features

1. Add new model in `internal/models/`
2. Add storage methods in `internal/storage/`
3. Add HTTP handlers in `internal/handlers/`
4. Update routes in `cmd/server/main.go`

#### Example: Adding a User model

- Create `internal/models/user.go`
- Create `internal/storage/user.go`

---

## 📄 License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.
