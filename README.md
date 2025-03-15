# Movies CRUD API

Simple RESTful API for managing movies using Go, Gin, GORM, and PostgreSQL.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/mehmonov/movies-crud.git
cd movies-crud
```

2. Create .env file (optional - default values will be used if not provided):
```bash
cp .env.example .env
```

3. Build and run the application:
```bash
docker-compose up --build
```

The API will be available at: http://localhost:8080
Swagger documentation: http://localhost:8080/swagger/index.html

## API Endpoints

### Public Endpoints
- `GET /api/v1/movies` - Get all movies
- `GET /api/v1/movies/:id` - Get movie by ID
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login user

### Protected Endpoints (Requires JWT Token, basic auth)
- `POST /api/v1/movies` - Create new movie
- `PUT /api/v1/movies/:id` - Update movie
- `DELETE /api/v1/movies/:id` - Delete movie

## Authentication

For protected endpoints, include the JWT token in the Authorization header:
```bash
Authorization: Bearer <your-token>
```

## Development

To stop the containers:
```bash
docker-compose down
```

To view logs:
```bash
docker-compose logs -f
```

## Database

To connect to the database:
```bash
docker exec -it movies_db psql -U postgres -d movies_crud
```

Quick Commands:

```bash
# Build and run
docker-compose up --build

# Run only
docker-compose up

# Run in background
docker-compose up -d

# Stop containers
docker-compose down
```



!! Extended version is here
```

git checkout additional

or 

checkout branch with github

```

