# Movies CRUD API (Extended Version)

Enhanced version of the Movies CRUD API with additional features for media file management

## Additional Features

### Media Management
- Upload movie files (video)
- Download movie files
- Manage movie media files (posters, backdrops, trailers)
- Movie metadata support

### File Upload Features
- Secure file storage with hash-based filenames
- File type validation (MP4, MKV, MOV)
- File size limits (max 2GB)
- Separate storage for each movie

### Security Features
- Protected file upload/download endpoints
- Content-Type validation
- File hash verification
- Secure file paths

## API Endpoints

### Public Endpoints
- `GET /api/v1/movies` - Get all movies
- `GET /api/v1/movies/:id` - Get movie by ID
- `GET /api/v1/movies/:id/media` - Get movie media files
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login user
- `POST /api/v1/auth/refresh` - Refresh access token

### Protected Endpoints (Requires JWT Token)
- `POST /api/v1/movies` - Create new movie
- `PUT /api/v1/movies/:id` - Update movie
- `DELETE /api/v1/movies/:id` - Delete movie
- `POST /api/v1/movies/:id/file` - Upload movie file
- `GET /api/v1/movies/:id/files/:fileId` - Download movie file

## File Upload Guide

### Supported File Types
- MP4 (video/mp4)
- MKV (video/x-matroska)
- MOV (video/quicktime)

### Upload Example (using Postman)
1. Create a POST request to `/api/v1/movies/{id}/file`
2. Set Authorization header with JWT token
3. Use form-data body:
   - Key: `file`
   - Type: `File`
   - Value: Select video file

### Download Example
```bash
GET /api/v1/movies/{id}/files/{fileId}
Authorization: Bearer <your-jwt-token>
```

## Project Structure
```
.
├── config/             # Configuration
├── internal/
│   ├── api/
│   │   ├── controllers/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   └── routes/
│   ├── models/        # Data models
│   └── services/      # Business logic
├── pkg/
│   ├── auth/          # JWT authentication
│   └── filestore/     # File handling
└── uploads/           # File storage
```

## Environment Variables
```env
# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=1234
DB_NAME=movies-crud

# Server
SERVER_PORT=8080

# JWT
JWT_SECRET=your-secret-key
ACCESS_TOKEN_SECRET=your-access-token-secret
REFRESH_TOKEN_SECRET=your-refresh-token-secret

# File Upload
UPLOAD_PATH=./uploads
```

## Getting Started

1. Clone and switch to extended branch:
```bash
git clone https://github.com/mehmonov/movies-crud.git
cd movies-crud
git checkout extended
```

2. Set up environment:
```bash
cp .env.example .env
# Edit .env file with your settings
```

3. Create uploads directory:
```bash
mkdir uploads
chmod 755 uploads
```

4. Run with Docker:
```bash
docker-compose up --build
```

## References

- [OWASP File Upload Security](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html)
- [Content-Type Headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type)
- [SQL Injection Prevention](https://go.dev/doc/database/sql-injection)



