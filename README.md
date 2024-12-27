# NAQA API

[![Go Version](https://img.shields.io/github/go-mod/go-version/anqorithm/naqa-api)](https://golang.org/)
[![Build Status](https://img.shields.io/github/workflow/status/anqorithm/naqa-api/main)](https://github.com/anqorithm/naqa-api/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/anqorithm/naqa-api)](https://goreportcard.com/report/github.com/anqorithm/naqa-api)
[![License](https://img.shields.io/github/license/anqorithm/naqa-api)](LICENSE)
[![Code Size](https://img.shields.io/github/languages/code-size/anqorithm/naqa-api)](https://github.com/anqorithm/naqa-api)
[![Last Commit](https://img.shields.io/github/last-commit/anqorithm/naqa-api)](https://github.com/anqorithm/naqa-api/commits/main)
[![Docker Image Size](https://img.shields.io/docker/image-size/anqorithm/naqa-api)](https://hub.docker.com/r/anqorithm/naqa-api)
[![Coverage](https://img.shields.io/codecov/c/github/anqorithm/naqa-api)](https://codecov.io/gh/anqorithm/naqa-api)
[![API Documentation](https://img.shields.io/badge/api-documentation-blue)](https://naqa-api.docs.stoplight.io)
[![MongoDB](https://img.shields.io/badge/MongoDB-4.4+-green.svg)](https://www.mongodb.com/)
[![Fiber Framework](https://img.shields.io/badge/Fiber-2.x-blue.svg)](https://gofiber.io/)

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
git clone https://github.com/anqorithm/naqa-api
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

## Installation
1. Clone the project.
2. Install Go dependencies:
   ```
   go mod download
   ```

## Environment
Set the following environment variables or edit the .env file:
- MONGO_URI
- MONGO_DATABASE
- PORT

## Usage
1. Run the server:
   ```
   go run cmd/api/main.go
   ```
2. Access the API at:
   ```
   http://localhost:3000/api/v1
   ```

## Environment Variables

Before running the application, make sure to set up your environment variables:

```bash
# Copy the example env file
cp .env.example .env

# Fill in your environment variables in .env file:
API_VERSION      # API version
ENVIRONMENT      # development, production, etc.
PORT            # Server port
APP_NAME        # Application name
APP_DESCRIPTION # Application description
MONGO_URI      # MongoDB connection URI
MONGO_DATABASE # MongoDB database name
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

### Environment Variables
@baseUrl = http://localhost:3030  
@contentType = application/json

### Health Check
GET {{baseUrl}}/api/health  
Content-Type: {{contentType}}

### Metrics Dashboard
GET {{baseUrl}}/metrics  
Content-Type: {{contentType}}

### Root Endpoint
GET {{baseUrl}}/  
Content-Type: {{contentType}}

### Get Stocks by Year
GET {{baseUrl}}/api/v1/stocks/year/2025  
Content-Type: {{contentType}}

### Search Stocks with Parameters
GET {{baseUrl}}/api/v1/stocks/year/2023/search?name=aramco&sector=energy&sharia_opinion=نقية  
Content-Type: {{contentType}}

### Calculate Purification Amount
POST {{baseUrl}}/api/v1/stocks/calculate-purification  
Content-Type: {{contentType}}

```json
{
    "start_date": "2023-01-01",
    "end_date": "2023-12-31",
    "number_of_stocks": 100,
    "stock_code": "1111"
}
```

### Search Examples

#### Search by Name
GET {{baseUrl}}/api/v1/stocks/year/2023/search?name=%D8%A3%D8%B1%D8%A7%D9%85%D9%83%D9%88  
Content-Type: {{contentType}}

#### Search by Code
GET {{baseUrl}}/api/v1/stocks/year/2023/search?code=2222  
Content-Type: {{contentType}}

#### Search by Sector
GET {{baseUrl}}/api/v1/stocks/year/2023/search?sector=%D8%A7%D9%84%D8%B7%D8%A7%D9%82%D8%A9  
Content-Type: {{contentType}}

#### Search by Sharia Opinion
GET {{baseUrl}}/api/v1/stocks/year/2023/search?sharia_opinion=%D9%86%D9%82%D9%8A%D8%A9  
Content-Type: {{contentType}}

#### Combined Search
GET {{baseUrl}}/api/v1/stocks/year/2023/search?sector=%D8%A7%D9%84%D8%B7%D8%A7%D9%82%D8%A9&sharia_opinion=%D9%86%D9%82%D9%8A%D8%A9  
Content-Type: {{contentType}}

### Inspiration and Data Source

This API is inspired by [NaqausStocks.com](https://naquastocks.com/).

Data source: [Almaqased Cleansing Calculator](https://almaqased.net/cleansing-calculator/%D9%82%D9%88%D8%A7%D8%A6%D9%85-%D8%A7%D9%84%D8%AA%D8%AD%D9%84%D9%8A%D9%84-%D8%A7%D9%84%D9%85%D8%A7%D9%84%D9%8A-%D9%84%D9%84%D8%B4%D8%B1%D9%83%D8%A7%D8%AA/)  
المشرف العام: د. محمد بن سعود العصيمي

## Development

### Project Structure
```
.
├── cmd
│   └── api
│       └── main.go
├── docker-compose.yml
├── Dockerfile
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   ├── config.go
│   │   └── mongodb.go
│   ├── handlers
│   │   ├── handlers.go
│   │   └── stocks.go
│   ├── middleware
│   │   ├── middleware.go
│   │   └── year_validator.go
│   └── routes
│       └── routes.go
├── LICENSE
├── Makefile
└── README.md
```

## Architecture Design & Diagrams

### Class Diagram

```mermaid
classDiagram
    class Config {
        +MongoDB mongodb
        +Load()
    }

    class MongoDB {
        +string URI
        +string Database
        +string Collection
        +Connect()
    }

    class Handlers {
        +HandleError(w ResponseWriter, err error, status int)
        +GetStocks(w ResponseWriter, r *Request)
        +GetStocksByYear(w ResponseWriter, r *Request)
    }

    class Middleware {
        +YearValidator(next http.Handler) http.Handler
        +ValidateYear(year string) bool
    }

    class Routes {
        +SetupRoutes(r *mux.Router)
    }

    class main {
        +main()
    }

    Config --> MongoDB : contains
    Routes --> Handlers : uses
    Routes --> Middleware : uses
    main --> Routes : initializes
    main --> Config : loads
    Handlers --> MongoDB : uses
```

### Building for Production
```bash
go build -o naqa-api cmd/api/main.go
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
