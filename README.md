# NAQA API | نقاء API

[![Go Version](https://img.shields.io/github/go-mod/go-version/anqorithm/naqa-api)](https://golang.org/)
[![Go Report Card](https://goreportcard.com/badge/github.com/anqorithm/naqa-api)](https://goreportcard.com/report/github.com/anqorithm/naqa-api)
[![License](https://img.shields.io/github/license/anqorithm/naqa-api)](LICENSE)
[![Code Size](https://img.shields.io/github/languages/code-size/anqorithm/naqa-api)](https://github.com/anqorithm/naqa-api)
[![Last Commit](https://img.shields.io/github/last-commit/anqorithm/naqa-api)](https://github.com/anqorithm/naqa-api/commits/main)
[![MongoDB](https://img.shields.io/badge/MongoDB-5.0+-green.svg)](https://www.mongodb.com/)
[![Fiber Framework](https://img.shields.io/badge/Fiber-v2.52.5-blue.svg)](https://gofiber.io/)
![GCP Deployment](https://img.shields.io/badge/Deployed%20on-GCP-4285F4?logo=google-cloud&logoColor=white)

<div dir="rtl">

## نظرة عامة | Overview
نقاء API هي خدمة RESTful مصممة لتوفير بيانات عن عملية تطهير الأسهم السعودية. توفر الخدمة معلومات دقيقة عن نسب التطهير للأسهم في السوق السعودي وتصنيفها الشرعي.

### المميزات الرئيسية | Key Features
- 🧮 حساب مبلغ التطهير للأسهم | Stock Purification Amount Calculator
- 🔍 البحث في الأسهم حسب القطاع والتصنيف الشرعي | Search Stocks by Sector and Shariah Classification
- 📊 معلومات محدثة عن الأسهم النقية والمختلطة | Updated Pure and Mixed Stocks Information
- 🌐 دعم كامل للغة العربية | Full Arabic Language Support
- 🔄 واجهة برمجة تطبيقات RESTful سهلة الاستخدام | Easy-to-use RESTful API

### المتطلبات | Requirements
- Go 1.23.4 أو أحدث | or higher
- Git
- MongoDB

### تثبيت المشروع | Project Setup
```bash
git clone https://github.com/anqorithm/naqa-api
cd naqa-api
go mod download
```

### المتغيرات البيئية | Environment Variables
```bash
MONGO_URI=       
MONGO_DATABASE=
PORT=
```

### التصنيف الشرعي | Shariah Classification
- نقية | Pure: أسهم متوافقة تماماً مع الشريعة | Fully Shariah Compliant Stocks
- مختلطة | Mixed: أسهم تحتاج إلى تطهير | Stocks Requiring Purification
- غير متوافقة | Non-Compliant: أسهم غير متوافقة مع الشريعة | Non-Shariah Compliant Stocks

### السنوات المدعومة | Supported Years
> السنوات المتوفرة | Available Years: 2018, 2019, 2020, 2021, 2022, 2023

</div>

---

## Overview
NAQA API is a RESTful service designed to provide data about the purification process of Saudi stocks. The service offers accurate information about stock purification rates in the Saudi market and their Shariah classification.

### Key Features
- 🧮 Calculate stock purification amounts
- 🔍 Search stocks by sector and Shariah classification
- 📊 Updated information about pure and mixed stocks
- 🌐 Full Arabic language support
- 🔄 Easy-to-use RESTful API

### Deployed on GCP

![GCP Deployment](https://img.shields.io/badge/Deployed%20on-GCP-4285F4?logo=google-cloud&logoColor=white)

[Access the API on GCP](https://naqa-api-36462279645.europe-west1.run.app/)

### Prerequisites
- Go 1.23.4 or higher
- Git
- MongoDB

## Version Information
- Version: 1.9.0
- Environment: Development
- Base API Path: `/api/v1`

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
```http
## Environment Variables
@baseUrl = http://localhost:3000  
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
GET {{baseUrl}}/api/v1/stocks/year/2023  
Content-Type: {{contentType}}

### Search Stocks with Parameters
GET {{baseUrl}}/api/v1/stocks/year/2023/search?name=aramco&sector=energy&sharia_opinion=نقية  
Content-Type: {{contentType}}

### Calculate Purification Amount
POST {{baseUrl}}/api/v1/stocks/calculate-purification  
Content-Type: {{contentType}}

```

### Search Examples

#### Search by Name
```http
GET {{baseUrl}}/api/v1/stocks/year/2023/search?name=%D8%A3%D8%B1%D8%A7%D9%85%D9%83%D9%88  
Content-Type: {{contentType}}
```

#### Search by Code
```http
GET {{baseUrl}}/api/v1/stocks/year/2023/search?code=2222  
Content-Type: {{contentType}}
```

#### Search by Sector
```http
GET {{baseUrl}}/api/v1/stocks/year/2023/search?sector=%D8%A7%D9%84%D8%B7%D8%A7%D9%82%D8%A9  
Content-Type: {{contentType}}
```

#### Search by Sharia Opinion
```http
GET {{baseUrl}}/api/v1/stocks/year/2023/search?sharia_opinion=%D9%86%D9%82%D9%8A%D8%A9  
Content-Type: {{contentType}}
```

#### Combined Search
```http
GET {{baseUrl}}/api/v1/stocks/year/2023/search?sector=%D8%A7%D9%84%D8%B7%D8%A7%D9%82%D8%A9&sharia_opinion=%D9%86%D9%82%D9%8A%D8%A9  
Content-Type: {{contentType}}
```

### مصدر البيانات | Data Source

<div dir="rtl">

هذه الخدمة مبنية على قوائم التطهير المنشورة للعامة من قبل فضيلة الشيخ الدكتور محمد بن سعود العصيمي (المشرف العام على موقع المقاصد). تم جمع وتنظيم البيانات من:

- [حاسبة التطهير - موقع المقاصد](https://almaqased.net/cleansing-calculator/%D9%82%D9%88%D8%A7%D8%A6%D9%85-%D8%A7%D9%84%D8%AA%D8%AD%D9%84%D9%8A%D9%84-%D8%A7%D9%84%D9%85%D8%A7%D9%84%D9%8A-%D9%84%D9%84%D8%B4%D8%B1%D9%83%D8%A7%D8%AA/)
- قوائم التحليل المالي للشركات المنشورة للعامة
- حسابات نسب التطهير المعتمدة من قبل الشيخ د. محمد العصيمي

</div>

> **Important Note**: All purification rates and Shariah classifications are based on the publicly available lists and calculations approved by Sheikh Dr. Mohammed Al-Osaimi, the General Supervisor of Almaqased website.

### الاعتماد والإشراف | Supervision

<div dir="rtl">

تحت إشراف:  
**فضيلة الشيخ الدكتور محمد بن سعود العصيمي**  
المشرف العام على موقع المقاصد للاقتصاد الإسلامي

</div>

### Inspiration and Data Source

This API is inspired by [NaqausStocks.com](https://naquastocks.com/).

Data source: [Almaqased Cleansing Calculator](https://almaqased.net/cleansing-calculator/%D9%82%D9%88%D8%A7%D8%A6%D9%85-%D8%A7%D9%84%D8%AA%D8%AD%D9%84%D9%8A%D9%84-%D8%A7%D9%84%D9%85%D8%A7%D9%84%D9%8A-%D9%84%D9%84%D8%B4%D8%B1%D9%83%D8%A7%D8%AA/)  
المشرف العام: د. محمد بن سعود العصيمي

> **Note:** Supported Years: 2018, 2019, 2020, 2021, 2022, 2023


## Development

### Project Structure
```
.
├── api.http
├── cmd
│   └── api
│       └── main.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   ├── config.go
│   │   └── mongodb.go
│   ├── handlers
│   │   ├── errors.go
│   │   ├── handlers.go
│   │   └── stocks.go
│   ├── middleware
│   │   ├── middleware.go
│   │   └── year_validator.go
│   ├── models
│   │   ├── error.go
│   │   └── stock.go
│   ├── routes
│   │   └── routes.go
│   └── utils
│       └── utils.go
├── LICENSE
├── Makefile
└── README.md
```


## حساب مبلغ تطهير الأسهم | Stock Purification Amount Calculation

<div dir="rtl">

### معادلة حساب مبلغ التطهير
يتم حساب مبلغ تطهير الأسهم باستخدام المعادلة التالية:

</div>

$$
\text{مبلغ التطهير} = \frac{\text{عدد الأسهم} \times \text{نسبة التطهير} \times \text{عدد أيام التملك}}{365.0}
$$

<div dir="rtl">

حيث:
- `عدد الأسهم`: إجمالي عدد الأسهم المملوكة
- `نسبة التطهير`: نسبة التطهير السنوية (النسبة المئوية)
- `عدد أيام التملك`: عدد الأيام التي تم فيها امتلاك الأسهم
- `365.0`: عدد أيام السنة (ثابت)

### مثال
لـ 100 سهم بنسبة تطهير 2.5% تم امتلاكها لمدة 180 يوم:

</div>

$$
\text{مبلغ التطهير} = \frac{100 \times 0.025 \times 180}{365.0} = 1.23
$$


## Stock Purification Amount Calculation

The purification amount for stocks is calculated using the following formula:

$$
\text{purificationAmount} = \frac{\text{numberOfStocks} \times \text{purificationRate} \times \text{daysHeld}}{365.0}
$$

Where:
- `numberOfStocks`: Total number of stocks held
- `purificationRate`: Annual purification rate (percentage)
- `daysHeld`: Number of days the stocks were held
- `365.0`: Days in a year (constant)

### Example
For 100 stocks with 2.5% purification rate held for 180 days:

$$
\text{purificationAmount} = \frac{100 \times 0.025 \times 180}{365.0} = 1.23
$$


## Architecture Design & Diagrams

### Class Diagram

```mermaid
classDiagram
    %% Core Classes based on actual project structure
    class Config {
        +MongoDB mongodb
        +Load() error
        +GetEnvVariable(key string, defaultValue string) string
    }

    class MongoDB {
        +string URI
        +string Database
        +Connect() error
        +GetClient() *mongo.Client
    }

    class Stock {
        +string ID
        +string Code
        +string Name
        +string Sector
        +string ShariaOpinion
        +float64 PurificationRate
        +int Year
    }

    class StockHandler {
        +GetStocks(c *fiber.Ctx) error
        +GetStocksByYear(c *fiber.Ctx) error
        +SearchStocks(c *fiber.Ctx) error
        +CalculatePurification(c *fiber.Ctx) error
    }

    class ErrorHandler {
        +HandleError(c *fiber.Ctx, err error) error
        +NewAPIError(message string, status int) *APIError
    }

    class YearValidator {
        +ValidateYear(year string) bool
        +YearMiddleware(c *fiber.Ctx) error
    }

    class Routes {
        +SetupRoutes(app *fiber.App)
        +setupStockRoutes(app *fiber.App)
    }

    %% Relationships
    Config --> MongoDB : uses
    Routes --> StockHandler : registers
    Routes --> YearValidator : uses
    StockHandler --> ErrorHandler : uses
    StockHandler --> Stock : manages
```

### Docker-Compose Diagram

```mermaid
graph LR
    subgraph Docker-Compose
        subgraph naqa-network
            api[API Service]
            mongo[MongoDB]
            api --> |depends_on| mongo
        end
        
        volume[(mongodb_data)]
        mongo --> |volume| volume
        
        api --> |port| port1[3000:3000]
        mongo --> |port| port2[27017:27017]
        
        env1[Environment:<br/>MONGODB_URI<br/>PORT] --> api
    end
```


### Request Flow Diagram

```mermaid
sequenceDiagram
    participant C as Client
    participant F as Fiber App
    participant M as Year Validator
    participant H as Stock Handler
    participant DB as MongoDB

    C->>F: HTTP Request
    F->>M: Validate Year
    alt Invalid Year
        M-->>C: Error Response
    else Valid Year
        M->>H: Process Request
        H->>DB: Query Data
        DB-->>H: Return Results
        H-->>C: JSON Response
    end
```

### Building for Production
```bash
go build -o naqa-api cmd/api/main.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/NewFeature`)
3. Commit your Changes (`git commit -m 'Add some NewFeature'`)
4. Push to the Branch (`git push origin feature/NewFeature`)
5. Open a Pull Request

## Contributors

- **Abdullah Alqahtani** - *Initial work* - [anqorithm](https://github.com/anqorithm)
  - Email: anqorithm@protonmail.com
- **Yasser** - [zYaseer](https://github.com/zYasser)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Credits

### Data Source
- [Almaqased Cleansing Calculator](https://almaqased.net/)
- المشرف العام: د. محمد بن سعود العصيمي

### Inspiration
This API is inspired by [NaqausStocks.com](https://naquastocks.com/)