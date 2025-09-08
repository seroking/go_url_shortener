
# Go URL Shortener

[![Go Version](https://img.shields.io/badge/go-1.25-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green)](LICENSE)

A **fast, secure, and lightweight URL shortener** built with Go and Gin.
This application allows users to shorten URLs, track clicks, and manage their links through a simple REST API.

---

## 🛠 Features

- User authentication: signup and signin with JWT
- Create short links from long URLs
- Redirect to original URLs via shortcodes
- Track clicks for each link
- List all links created by a user
- Delete your own links (or by admin)
- Admin panel for managing users

---

## 🖼 Architecture Overview
User --> Shorten URL --> Store in Database --> Generate Shortcode
|
v
Visit Shortcode --> Lookup Original URL --> Redirect & Increment Clicks

## 📦 Tech Stack

- **Backend:** Go, Gin
- **Database:** GORM (MySQL/PostgreSQL supported)
- **Authentication:** JWT
- **Routing & Middleware:** Gin
- **URL encoding:** Base62 (custom helper)

---

## 🚀 Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/yourusername/go-url-shortener.git
cd go-url-shortener

2. Install dependencies
go mod tidy

3. Setup environment

Create a .env file with:
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=url_shortener
JWT_SECRET=your_secret_key


4. Run migrations
go run main.go
5. Start the server
go run main.go

Server runs on: http://localhost:8080

🔗 API Endpoints
Public

POST /api/v1/signup – Create a new user

POST /api/v1/signin – Authenticate user

GET /api/v1/:shortcode – Redirect to original URL

Protected (JWT required)

GET /api/v1/links – List user’s links

POST /api/v1/links – Create a new short link

DELETE /api/v1/links/:id – Delete your link

GET /api/v1/profile – Get user profile

PUT /api/v1/profile – Update user profile

Admin (JWT + admin role)

GET /api/v1/users – List all users

GET /api/v1/users/:id – Get user by ID

PUT /api/v1/users/:id – Update user

DELETE /api/v1/users/:id – Delete user

⚡ Example cURL Requests

Signup
curl -X POST http://localhost:8080/api/v1/signup \
-H "Content-Type: application/json" \
-d '{"username":"sero","email":"sero@example.com","password":"password"}'

Signin
curl -X POST http://localhost:8080/api/v1/signin \
-H "Content-Type: application/json" \
-d '{"email":"sero@example.com","password":"password"}'
Create a link
curl -X POST http://localhost:8080/api/v1/links \
-H "Authorization: Bearer <TOKEN>" \
-H "Content-Type: application/json" \
-d '{"url":"https://example.com"}'

List user links
curl -X GET http://localhost:8080/api/v1/links \
-H "Authorization: Bearer <TOKEN>"

Delete a link
curl -X DELETE http://localhost:8080/api/v1/links/1 \
-H "Authorization: Bearer <TOKEN>"

Redirect via shortcode
curl -L http://localhost:8080/api/v1/abcd1234
📝 Notes

Clicks are tracked automatically when a short URL is visited.

Shortcodes are generated automatically based on the database ID.

Only the owner or an admin can delete links.

Original URLs are validated to start with http:// or https://.

The app is designed for demonstration/portfolio purposes.

📌 License

This project is licensed under the MIT License. See the LICENSE
