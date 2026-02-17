# 📝 Simple Go Echo Todo API

A simple and clean RESTful **Todo API** built with **Go**, **Echo framework**, and **PostgreSQL**. This is a beginner-friendly project to learn Go web development fundamentals.

---

## 🎓 Learning Path

This project teaches you:
1. **Go Basics** - Packages, imports, functions, and error handling
2. **Web Framework** - Using Echo framework for routing and middleware
3. **Database** - PostgreSQL integration with connection pooling (pgx)
4. **Architecture** - Clean layered architecture (Handlers → Storage → Database)
5. **API Development** - RESTful API best practices
6. **Configuration Management** - YAML-based config loading

**Next Step**: Learn [Go Echo with PostgreSQL Advanced Patterns](#-next-learning-goals) for production-ready applications.

---

## 🗂️ Project Structure & Explanation

```bash
simple-go-echo/
├── cmd/
│   └── server/
│       └── main.go              # 🚀 Application entry point - where everything starts
├── config/
│   └── config.yaml              # ⚙️ Configuration file - server and database settings
├── internal/                     # 📦 Private packages (Go convention for internal code)
│   ├── config/
│   │   └── config.go            # 📋 Reads & parses config.yaml into Go structs
│   ├── database/
│   │   └── postgres.go          # 🔌 Establishes PostgreSQL connection pool
│   ├── http/
│   │   └── handlers/
│   │       └── todo.go          # 🎯 Handles HTTP requests, validates input, returns responses
│   ├── models/
│   │   ├── todo.go              # 📄 Defines Todo data structure
│   │   └── blog.go              # 📄 Defines Blog data structure (for future expansion)
│   ├── server/
│   │   └── server.go            # 🌐 Sets up Echo server, middleware, and routes
│   ├── storage/
│   │   ├── todo.go              # 💾 Database queries for todos (CRUD operations)
│   │   └── blog.go              # 💾 Database queries for blogs
│   └── utils/
│       └── response/
│           └── response.go      # 📤 Helper functions for consistent HTTP responses
├── frontend/                     # ⚛️ React + TypeScript frontend (separate project)
├── go.mod                        # 📚 Module definition and dependency list
├── go.sum                        # 🔐 Checksums for dependencies (ensures reproducibility)
└── README.md                     # 📖 This file

```

---

## 🔍 Deep Dive: Understanding Each Component

### 1. **main.go** - The Entry Point
```go
package main  // Special package name - Go looks for this to start the program

func main() {  // Entry function - execution starts here
    cfg := config.LoadConfig()              // 1. Load config
    db := database.NewPostgres(cfg)         // 2. Connect to database
    defer db.Close()                         // 3. Ensure cleanup when program exits
    srv := server.NewServer(cfg, db)        // 4. Create server with routes
    srv.Start()                              // 5. Start listening for requests
}
```

**Key Concepts:**
- `package main` - Tells Go this is an executable program
- `func main()` - Special function where execution begins
- `defer` - Schedules a function to run when the current function exits (cleanup)

---

### 2. **config/config.go** - Configuration Management
```go
type Config struct {
    Server   Server `yaml:"server"`      // Struct tags tell YAML parser which field to map
    Database Database `yaml:"database"`
}

func LoadConfig() *Config {
    data, err := os.ReadFile("config/config.yaml")  // Read file
    yaml.Unmarshal(data, &cfg)                       // Parse YAML into struct
    return &cfg
}
```

**Key Concepts:**
- **Struct Tags** (`yaml:"server"`) - Tells libraries how to map data
- **Pointers** (`*Config`) - Returns memory reference, not a copy
- **Error Handling** - Go requires explicit error checking (different from exceptions)

---

### 3. **database/postgres.go** - Database Connection
```go
func NewPostgres(cfg *config.Config) *pgxpool.Pool {
    dsn := fmt.Sprintf("postgres://%s:%s@...", cfg.Database.User, cfg.Database.Password)
    pool, err := pgxpool.New(context.Background(), dsn)  // Connection pool
    pool.Ping(context.Background())  // Test connection
    return pool
}
```

**Key Concepts:**
- **Connection Pool** - Reuses database connections for efficiency
- **Context** - Controls timeout and cancellation
- **Error Handling** - Check errors immediately

---

### 4. **models/todo.go** - Data Structure
```go
type Todo struct {
    ID    int64  `json:"id"`          // Will be converted to "id" in JSON
    Title string `json:"title"`       // User can't send todos without title
    Done  bool   `json:"done"`
}
```

**Key Concepts:**
- **Struct** - Go's way to define data structures (like classes in other languages)
- **Struct Tags** - `json:"id"` controls how data is serialized to/from JSON
- **Fields must be capitalized** - Go's way of marking things as "public" (exported)

---

### 5. **server/server.go** - HTTP Server & Routing
```go
func NewServer(cfg *config.Config, db *pgxpool.Pool) *Server {
    e := echo.New()  // Create Echo instance
    
    // Add middleware - functions that process every request
    e.Use(middleware.Logger())      // Log all requests
    e.Use(middleware.Recover())     // Catch panics
    e.Use(middleware.CORS())        // Allow cross-origin requests
    
    // Setup routes
    api := e.Group("/api")          // Group routes with /api prefix
    api.GET("/todos", handler.GetAll)       // Map GET /api/todos to GetAll function
    api.POST("/todos/create", handler.Create)
    api.PUT("/todos/update/:id", handler.Update)
    api.DELETE("/todos/:id", handler.Delete)
}
```

**Key Concepts:**
- **Middleware** - Functions that process every request (logging, error handling, security)
- **Route Groups** - Organize routes by prefix
- **Route Handlers** - Map HTTP methods + paths to Go functions

---

### 6. **handlers/todo.go** - HTTP Request Handling
```go
func (h *TodoHandler) Create(c echo.Context) error {
    var todo models.Todo
    c.Bind(&todo)  // Convert JSON request body to Go struct
    
    // Validate input
    if todo.Title == "" {
        return response.BadRequest(c, "Title is required")
    }
    
    // Call storage layer to save
    id, err := h.storage.Create(c.Request().Context(), &todo)
    
    // Return response
    return response.Created(c, todo)
}
```

**Key Concepts:**
- **Receiver Function** (`(h *TodoHandler)`) - Attach function to a struct (like methods in OOP)
- **Binding** - Parse incoming JSON into Go structs
- **Validation** - Check data before processing
- **Layering** - Handler delegates to storage (separation of concerns)

---

### 7. **storage/todo.go** - Database Operations
```go
func (s *TodoStorage) Create(ctx context.Context, todo *models.Todo) (int64, error) {
    var id int64
    // Execute SQL query with parameters ($1, $2 protect against SQL injection)
    err := s.DB.QueryRow(ctx,
        `INSERT INTO todos (title, done) VALUES ($1, $2) RETURNING id`,
        todo.Title, todo.Done,
    ).Scan(&id)  // Extract returned id
    return id, err
}
```

**Key Concepts:**
- **Parameterized Queries** - `$1, $2` prevent SQL injection
- **QueryRow vs Query** - QueryRow for single result, Query for multiple
- **Scan** - Extract values from database results

---

### 8. **utils/response/response.go** - Response Helpers
```go
func OK(c echo.Context, data any) error {
    return c.JSON(http.StatusOK, data)
}

func Created(c echo.Context, data any) error {
    return c.JSON(http.StatusCreated, data)
}
```

**Key Concepts:**
- **Consistency** - Standardized response format for all endpoints
- **HTTP Status Codes** - 200 OK, 201 Created, 400 Bad Request, etc.
- **DRY Principle** - Don't Repeat Yourself (centralized response logic)

---

## 🏗️ Architecture Explained

```
HTTP Request (from client)
    ↓
Server (port 8080)
    ↓
Router (Echo) - "Which function should handle this?"
    ↓
Handler (validation, business logic) - "Is this request valid?"
    ↓
Storage (database queries) - "Get data from PostgreSQL"
    ↓
Database (PostgreSQL) - "Store/retrieve data"
    ↓
Response (JSON) → Client
```

**Why This Structure?**
- **Separation of Concerns** - Each layer has one job
- **Testability** - Mock storage layer for testing handlers
- **Reusability** - Storage layer can be used by other interfaces (CLI, scheduled jobs)
- **Maintainability** - Easy to find and fix bugs

---

## 🔑 Key Go Concepts Explained

### 1. **Packages** (`package main`, `package config`, etc.)
- Go organizes code into packages
- `package main` is special - it's the entry point for executable programs
- Other packages can be imported with `import "path/to/package"`

### 2. **Functions & Receivers**
```go
// Regular function
func LoadConfig() *Config { }

// Receiver function (method attached to TodoHandler struct)
func (h *TodoHandler) Create(c echo.Context) error { }
//     ↑                                    ↑
//     receiver (like 'this' in other languages)
```

### 3. **Pointers** (`*Config`, `*TodoHandler`)
- `*` = pointer (memory address)
- Go uses pointers for efficiency (pass reference instead of copy)
- When you see `*Type`, it means "pointer to Type"

### 4. **Error Handling**
```go
// Go requires explicit error checking
data, err := os.ReadFile("config.yaml")
if err != nil {
    log.Fatal(err)  // Handle error immediately
}
```
- No try/catch like JavaScript
- Errors are return values that you must check

### 5. **Interfaces** (used by Echo & storage)
- Define what functions a type must implement
- Echo routes expect: `func(c echo.Context) error`
- Allows flexibility and testing

### 6. **Goroutines** (not used here, but important for Go)
- Lightweight threads
- `go functionName()` runs it concurrently
- Used by Echo to handle multiple requests simultaneously

---

## 🚀 Quick Start

### Prerequisites

- Go 1.19+
- PostgreSQL
- Git

### Installation

1. **Clone the repository**

   ```bash
   git clone <your-repo-url>
   cd simple-go-echo
   ```

2. **Set up PostgreSQL database**

   ```sql
   CREATE DATABASE todo_db;
   CREATE TABLE todos (
     id SERIAL PRIMARY KEY,
     title VARCHAR(255) NOT NULL,
     done BOOLEAN DEFAULT FALSE,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   );
   ```

3. **Configure the application**

   Edit `config/config.yaml`:

   ```yaml
   server:
     port: 8080
     addr: ":8080"

   database:
     host: "localhost"
     port: 5432
     user: "postgres"
     password: "your_password"
     dbname: "todo_db"
     sslmode: "disable"
   ```

4. **Install dependencies**

   ```bash
   go mod tidy
   ```

5. **Run the application**
   ```bash
   go run cmd/server/main.go
   ```
   You should see:
   ```
   🚀 Starting application...
   ✅ Connected to PostgreSQL successfully
   🚀 Server running on: localhost:8080
   ```

---

## 📚 API Endpoints

| Method | Endpoint                | Description       | Request Body                              | Response                |
| ------ | ----------------------- | ----------------- | ----------------------------------------- | ----------------------- |
| GET    | `/api/todos`            | Get all todos     | -                                         | `[{...}, {...}]`        |
| POST   | `/api/todos/create`     | Create a new todo | `{"title": "Task", "done": false}`        | `{"id": 1, "title": ...}` |
| GET    | `/api/todos/:id`        | Get todo by ID    | -                                         | `{"id": 1, "title": ...}` |
| PUT    | `/api/todos/update/:id` | Update todo by ID | `{"title": "Updated", "done": true}`      | `{"id": 1, "title": ...}` |
| DELETE | `/api/todos/:id`        | Delete todo by ID | -                                         | -                       |

---

## 💻 Example Usage

**Create a todo:**

```bash
curl -X POST http://localhost:8080/api/todos/create \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go", "done": false}'

# Response:
# {"id":1,"title":"Learn Go","done":false}
```

**Get all todos:**

```bash
curl http://localhost:8080/api/todos

# Response:
# [{"id":1,"title":"Learn Go","done":false}]
```

**Update a todo:**

```bash
curl -X PUT http://localhost:8080/api/todos/update/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go and Echo", "done": true}'

# Response:
# {"id":1,"title":"Learn Go and Echo","done":true}
```

**Delete a todo:**

```bash
curl -X DELETE http://localhost:8080/api/todos/1

# Response: (No content, just success status)
```

---

## 📦 Dependencies Explained

### **Echo** (`github.com/labstack/echo/v4`)
- High-performance web framework for Go
- Provides routing, middleware, and request handling
- Similar to Express.js in Node.js or Flask in Python

### **pgx** (`github.com/jackc/pgx/v5`)
- PostgreSQL driver for Go
- `pgxpool` - Connection pooling (reuses database connections)
- Efficiently execute SQL queries and handle results

### **YAML** (`gopkg.in/yaml.v3`)
- Library for parsing YAML files (used for configuration)
- Converts `config.yaml` into Go structs

---

## 🏗️ How Data Flows Through the System

### **Creating a Todo Example:**

```
1. Client sends: POST /api/todos/create
   └─ Body: {"title": "Learn Go", "done": false}

2. Echo Router receives request
   └─ Matches pattern "/api/todos/create"
   └─ Calls: TodoHandler.Create(c)

3. Handler layer (Create function)
   └─ c.Bind(&todo)         ← Parse JSON to Go struct
   └─ Validate: title != ""  ← Check requirements
   └─ Call: storage.Create()

4. Storage layer (database query)
   └─ Execute SQL: INSERT INTO todos...
   └─ Return: generated ID (from database)

5. Handler returns response
   └─ response.Created(c, todo)  ← Return 201 + JSON

6. Echo sends back to client:
   └─ Status: 201 Created
   └─ Body: {"id":1,"title":"Learn Go","done":false}
```

### **Getting All Todos:**

```
GET /api/todos
   ↓
TodoHandler.GetAll()
   ↓
storage.GetAll()  (SELECT * FROM todos)
   ↓
Loop through rows, Scan each todo
   ↓
Return []Todo
   ↓
response.OK(c, todos)  ← Return 200 + JSON array
```

---

## 🧠 Understanding Context

The `context.Context` parameter appears in many functions. Here's why:

```go
// context.Context allows:
// 1. Timeout - "Stop if request takes > 5 seconds"
// 2. Cancellation - "Stop if client closes connection"
// 3. Values - "Pass data between functions"

func Create(ctx context.Context, todo *Todo) (int64, error) {
    // If ctx is cancelled or times out, database query stops immediately
    err := s.DB.QueryRow(ctx, "INSERT ...").Scan(&id)
    if err != nil {
        return 0, err  // Return immediately if cancelled
    }
}
```

---

## 🧪 Testing Endpoints

You can use any of these tools:

### **curl** (command line)
```bash
curl -X GET http://localhost:8080/api/todos
```

### **Postman** (GUI)
- Create new request
- Select GET/POST/PUT/DELETE
- Enter URL: `http://localhost:8080/api/todos`
- Add JSON in Body tab

### **Thunder Client** (VS Code extension)
- Install extension
- Create requests visually

---

## 🛠️ Adding New Features

### Example: Adding a "User" Model

**Step 1: Create Model** (`internal/models/user.go`)
```go
package models

type User struct {
    ID    int64  `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

**Step 2: Create Storage** (`internal/storage/user.go`)
```go
package storage

func (s *UserStorage) Create(ctx context.Context, user *User) (int64, error) {
    var id int64
    err := s.DB.QueryRow(ctx,
        `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`,
        user.Name, user.Email,
    ).Scan(&id)
    return id, err
}
```

**Step 3: Create Handler** (`internal/http/handlers/user.go`)
```go
package handlers

func (h *UserHandler) Create(c echo.Context) error {
    var user models.User
    c.Bind(&user)
    id, err := h.storage.Create(c.Request().Context(), &user)
    user.ID = id
    return response.Created(c, user)
}
```

**Step 4: Add Routes** (in `internal/server/server.go`)
```go
userStorage := storage.NewUserStorage(db)
userHandler := handlers.NewUserHandler(userStorage)

api.POST("/users/create", userHandler.Create)
api.GET("/users", userHandler.GetAll)
```

---

## 🌟 Next Learning Goals

After mastering this project, learn:

### **1. Go Echo + PostgreSQL Advanced Patterns**
- [ ] Transaction handling (rollback on error)
- [ ] Prepared statements for security
- [ ] Database migrations (schema version control)
- [ ] Connection pooling optimization
- [ ] Query logging and monitoring

### **2. Authentication & Authorization**
- [ ] JWT tokens for user authentication
- [ ] Role-based access control (RBAC)
- [ ] Password hashing (bcrypt)
- [ ] Protected endpoints

### **3. Input Validation**
- [ ] Use `go-playground/validator` library
- [ ] Custom validation rules
- [ ] Error messages for invalid fields

### **4. Testing**
- [ ] Unit tests for handlers
- [ ] Integration tests with test database
- [ ] Table-driven tests

### **5. Error Handling**
- [ ] Custom error types
- [ ] Error wrapping for debugging
- [ ] HTTP error responses

### **6. Advanced Echo Features**
- [ ] Custom middleware
- [ ] WebSockets
- [ ] File uploads
- [ ] Rate limiting

### **7. Deployment**
- [ ] Docker containerization
- [ ] Environment variables (.env files)
- [ ] CI/CD pipelines (GitHub Actions)
- [ ] Production server setup

---

## 📄 License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.

