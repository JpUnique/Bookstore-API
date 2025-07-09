# 📚 Bookstore-API

![Go](https://img.shields.io/badge/Go-1.20-blue.svg)
![Build](https://github.com/JpUnique/bookstore-api/actions/workflows/go.yml/badge.svg)
![License](https://img.shields.io/github/license/JpUnique/bookstore-api)
![Issues](https://img.shields.io/github/issues/JpUnique/bookstore-api)


---

```markdown
# 📚 Bookstore-API

A RESTful API for managing books in a bookstore. Built with Go and GORM, the API allows users to **Create**, **Read**, **Update**, and **Delete** (CRUD) books in a MySQL database.

---

## 🚀 Features

- ✅ Create new books
- 🔍 Get all books or a book by ID
- ✏️ Update book information
- ❌ Delete a book
- 🗄️ Connects to a MySQL database using GORM
- 📦 JSON-based request and response format

---

## 📁 Project Structure

```

bookstore-api/
│
├── controllers/         # API logic (handlers)
│   └── book\_controller.go
│
├── models/              # Database models & ORM logic
│   └── book.go
│
├── config/              # Database connection logic
│   └── app.go
│
├── utils/               # Body parsing utilities
│   └── utils.go
│
├── main.go              # Entry point
└── go.mod / go.sum      # Dependencies

````

---

## 🔧 Technologies Used

- **Golang** (Go)
- **GORM** – ORM for Golang
- **MySQL** – Relational database
- **Gorilla Mux** – HTTP routing
- **net/http** – Built-in web server

---

## 🧠 API Endpoints

### 🔹 Get All Books
```http
GET /books
````

### 🔹 Get a Book by ID

```http
GET /books/{bookId}
```

### 🔹 Create a New Book

```http
POST /books
Content-Type: application/json

{
  "name": "The Go Programming Language",
  "author": "Alan A. A. Donovan",
  "publication": "Addison-Wesley"
}
```

### 🔹 Update a Book

```http
PUT /books/{bookId}
Content-Type: application/json

{
  "name": "Updated Title",
  "author": "Updated Author",
  "publication": "Updated Publisher"
}
```

### 🔹 Delete a Book

```http
DELETE /books/{bookId}
```

---

## 🛠️ Setup Instructions

1. **Clone the Repository**

```bash
git clone https://github.com/your-username/bookstore-api.git
cd bookstore-api
```

2. **Install Dependencies**

```bash
go mod tidy
```

3. **Configure Database**

* Set your MySQL connection in `config/app.go`:

```go
gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
```

4. **Run the Server**

```bash
go run main.go
```

---

## 🧪 Sample Response

```json
[
  {
    "ID": 1,
    "Name": "The Go Programming Language",
    "Author": "Alan A. A. Donovan",
    "Publication": "Addison-Wesley",
    "CreatedAt": "2025-07-08T12:00:00Z"
  }
]
```

---

## 📌 Notes

* Error handling can be improved in production.
* Authentication and validation not yet implemented.
* Logging and testing can be added for robustness.

---

## 🧑‍💻 Author

**David Mark**
Backend Developer | Golang Enthusiast
📫 [LinkedIn](https://www.linkedin.com/in/your-link) | 📧 [your.email@example.com](mailto:your.email@example.com)

---

## 📄 License

This project is open-sourced under the MIT License.

```

---

Would you like me to generate a **badge section**, **GitHub Actions CI config**, or a **Dockerfile** to make your repo even more professional?
```
