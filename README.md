
A simple, lightweight RESTful API built with Go and the Gin web framework to manage a collection of books stored directly in memory. The API supports full CRUD operations, proper validation, error handling, and custom middleware for logging and request timeouts.

🚀 Features
📖 View all books or a specific book by ID

➕ Add new books

✏️ Update existing books

❌ Delete books

🧠 Input validation with clear error responses

📦 In-memory storage using map[int]Book

📃 Request logging middleware

⏱️ Timeout middleware using context.WithTimeout

🛠️ Technology Stack
Language: Go

Framework: Gin









🧱 How to Run the Project
📦 Prerequisites:
Go installed

🔧 Steps:
Clone the repository:

bash
Copy
Edit
git clone https://github.com/your-username/book-management-api.git
cd book-management-api
Download dependencies:

bash
Copy
Edit
go mod tidy
Run the server:

bash
Copy
Edit
go run ./cmd/server
Visit http://localhost:8080/books to test the API.

