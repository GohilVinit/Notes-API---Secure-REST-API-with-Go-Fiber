# 📝 Notes API - Secure REST API with Go & Fiber

A production-ready REST API for user authentication and personal notes management built with Go, Fiber framework, and MySQL. Features JWT authentication, CRUD operations, and comprehensive security measures.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-2.52+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-8.0+-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-20.10+-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=json-web-tokens&logoColor=white)

## ✨ Features

- 🔐 **Secure Authentication**: JWT-based auth with bcrypt password hashing
- 📝 **Notes Management**: Full CRUD operations for personal notes
- 🛡️ **User Isolation**: Users can only access their own notes
- 🔍 **Search & Pagination**: Advanced filtering and pagination support
- 🐳 **Docker Ready**: Complete containerization with Docker Compose
- 🧪 **CLI Seeder**: Database seeding utility for testing
- 📊 **Comprehensive Validation**: Input validation and error handling
- 🚀 **Production Ready**: Clean architecture and best practices

## 🏗️ Architecture

```
notes-api/
├── 🐳 docker-compose.yml     # Container orchestration
├── 📦 Dockerfile            # Application container
├── ⚙️  .env                  # Environment configuration
├── 📋 go.mod & go.sum        # Go dependencies
├── 🚀 main.go               # Application entry point
├── 📁 models/               # Data models
│   ├── user.go
│   └── note.go
├── 🎯 handlers/             # Request handlers
│   ├── auth.go
│   └── notes.go
├── 🛡️  middleware/           # Custom middleware
│   └── jwt.go
├── 🛣️  routes/               # Route definitions
│   └── routes.go
├── 🔧 utils/                # Utility functions
│   ├── database.go
│   ├── hash.go
│   └── jwt.go
└── 🌱 cmd/                  # CLI tools
    └── seed.go
```

## 🚀 Quick Start

### Prerequisites

- Docker & Docker Compose
- Go 1.21+ (for local development)
- MySQL 8.0+ (for local development)

### 1. Clone Repository

```bash
git clone https://github.com/yourusername/notes-api.git
cd notes-api
```

### 2. Environment Setup

Copy and configure environment variables:

```bash
cp .env.example .env
```

Update `.env` with your configurations:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=notes_user
DB_PASSWORD=secure_password_123
DB_NAME=notes_db
DB_ROOT_PASSWORD=root_password_123

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_here_make_it_long_and_random

# Server Configuration
PORT=8080
ENV=development
```

### 3. Run with Docker (Recommended)

```bash
# Start all services
docker-compose up --build

# Run in background
docker-compose up -d --build

# View logs
docker-compose logs -f app
```

### 4. Run Locally (Alternative)

```bash
# Install dependencies
go mod tidy

# Start MySQL server locally
# Update .env with local MySQL credentials

# Run application
go run main.go
```

### 5. Seed Database (Optional)

```bash
# Using Docker
docker-compose exec app go run cmd/seed.go

# Local
go run cmd/seed.go
```

## 📡 API Endpoints

### Base URL
```
http://localhost:8080/api/v1
```

### Authentication Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `POST` | `/auth/register` | Register new user | ❌ |
| `POST` | `/auth/login` | Login user | ❌ |

### Notes Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `POST` | `/notes` | Create new note | ✅ |
| `GET` | `/notes` | Get all user notes | ✅ |
| `GET` | `/notes/:id` | Get specific note | ✅ |
| `PUT` | `/notes/:id` | Update note | ✅ |
| `DELETE` | `/notes/:id` | Delete note | ✅ |

### Additional Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| `GET` | `/health` | Health check | ❌ |

## 📋 API Usage Examples

### Register User

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "created_at": "2025-06-01T10:30:00Z",
    "updated_at": "2025-06-01T10:30:00Z"
  }
}
```

### Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Create Note

```bash
curl -X POST http://localhost:8080/api/v1/notes \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "My First Note",
    "content": "This is the content of my note."
  }'
```

### Get Notes with Pagination & Search

```bash
# Get paginated notes
curl -X GET "http://localhost:8080/api/v1/notes?page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"

# Search notes
curl -X GET "http://localhost:8080/api/v1/notes?search=meeting&page=1&limit=10" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Response:**
```json
{
  "notes": [
    {
      "id": 1,
      "title": "My First Note",
      "content": "This is the content of my note.",
      "user_id": 1,
      "created_at": "2025-06-01T10:35:00Z",
      "updated_at": "2025-06-01T10:35:00Z"
    }
  ],
  "total": 1,
  "page": 1,
  "limit": 10
}
```

## 🔧 Development

### Local Development Setup

```bash
# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build application
go build -o bin/notes-api

# Format code
go fmt ./...

# Run linter (if golangci-lint is installed)
golangci-lint run
```

### Database Migrations

GORM handles auto-migrations automatically. To manually migrate:

```bash
# Using Docker
docker-compose exec app go run -tags migrate main.go

# Local
go run -tags migrate main.go
```

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | `localhost` |
| `DB_PORT` | Database port | `3306` |
| `DB_USER` | Database username | `notes_user` |
| `DB_PASSWORD` | Database password | - |
| `DB_NAME` | Database name | `notes_db` |
| `JWT_SECRET` | JWT signing secret | - |
| `PORT` | Server port | `8080` |
| `ENV` | Environment | `development` |

## 🧪 Testing

### Manual Testing with Postman

1. Import the provided Postman collection
2. Set environment variables:
   - `baseUrl`: `http://localhost:8080/api/v1`
   - `authToken`: (will be set automatically)

3. Run tests in this order:
   - Health Check
   - Register User
   - Login User
   - Create Notes
   - CRUD Operations

### Automated Testing

```bash
# Run unit tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### Seeded Test Data

After running the seeder, you can use these test accounts:

| Email | Password | Notes |
|-------|----------|-------|
| `john@example.com` | `password123` | 3 sample notes |
| `jane@example.com` | `password123` | 3 sample notes |
| `bob@example.com` | `password123` | 3 sample notes |

## 🔒 Security Features

- **Password Hashing**: bcrypt with salt rounds
- **JWT Authentication**: Secure token-based auth
- **Input Validation**: Comprehensive request validation
- **SQL Injection Protection**: GORM ORM with prepared statements
- **CORS Configuration**: Configurable cross-origin requests
- **Rate Limiting**: Can be easily added with Fiber middleware
- **User Isolation**: Users can only access their own resources

## 🚀 Production Deployment

### Docker Production Build

```bash
# Build production image
docker build -t notes-api:latest .

# Run production container
docker run -d \
  --name notes-api \
  -p 8080:8080 \
  --env-file .env.production \
  notes-api:latest
```

### Environment Considerations

1. **Database**: Use managed MySQL service (AWS RDS, Google Cloud SQL)
2. **Secrets**: Use secret management service for JWT_SECRET
3. **Monitoring**: Add health checks and logging
4. **SSL/TLS**: Use reverse proxy (nginx) for HTTPS
5. **Scaling**: Use container orchestration (Kubernetes, Docker Swarm)

## 📊 Performance & Monitoring

### Health Check

```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "ok",
  "message": "Notes API is running"
}
```

### Logging

Application uses structured logging with different levels:
- `INFO`: General application flow
- `ERROR`: Error conditions
- `DEBUG`: Detailed information (development only)

### Metrics

Consider adding:
- Prometheus metrics
- Application performance monitoring (APM)
- Database connection pooling metrics
- Request/response time tracking

## 👨‍💻 Author

**Your Name**
- GitHub: [@GohilVinit](https://github.com/GohilVinit)
- LinkedIn: [Vinit gohil](https://www.linkedin.com/in/vinit-gohil-b46104311/)
- Email: gohilvinit03@gmail.com

---

⭐ **Star this repository if you find it helpful!**
