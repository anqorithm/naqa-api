# Naqa API

Naqa API is a RESTful service designed to provide data on the purification process of Saudi stocks.

## Version Information
- Version: 1.0.0
- Environment: Development
- Base API Path: `/api/v1`

## Prerequisites

- Go 1.23.4 or higher
- Git
- MongoDB

## Quick Start

1. Clone the repository:
```bash
git clone <repository-url>
cd naqa-api
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configurations
```

## Running the Application

### Local Development
```bash
go run cmd/api/main.go
```
Server will start at `http://localhost:3000`

### Docker Setup

#### Using Docker Compose (Recommended)
```bash
# Start services
docker-compose up -d

# Stop services
docker-compose down
```

#### Manual Docker Build
```bash
# Build image
docker build -t naqa-api .

# Run container
docker run -d -p 3000:3000 naqa-api
```

## API Endpoints

### Base URL
`http://localhost:3000/api/v1`

### Available Endpoints
- `GET /` - API information
- `GET /health` - Health check
- `GET /users` - Users endpoints
- `GET /tasks` - Tasks endpoints

## Development

### Project Structure
```
.
├── cmd/
│   └── api/          # Application entrypoint
├── internal/
│   ├── config/       # Configuration
│   ├── handlers/     # Request handlers
│   ├── middleware/   # Custom middleware
│   └── routes/       # Route definitions
└── docker-compose.yml
```

### Building for Production
```bash
go build -o naqa-api cmd/api/main.go
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
