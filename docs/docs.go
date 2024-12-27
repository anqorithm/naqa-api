// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/stocks/year/{year}": {
            "get": {
                "description": "Retrieves all stocks for a specific year from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stocks"
                ],
                "summary": "Get stocks by year",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Year of stocks data",
                        "name": "year",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_handlers.StockResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_handlers.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/stocks/year/{year}/search": {
            "get": {
                "description": "Search stocks with various filters",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stocks"
                ],
                "summary": "Search stocks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Year of stocks",
                        "name": "year",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Stock name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Stock code",
                        "name": "code",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Stock sector",
                        "name": "sector",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sharia compliance status",
                        "name": "sharia_opinion",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_handlers.StockResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/internal_handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/internal_handlers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "internal_handlers.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "Failed to fetch stocks"
                }
            }
        },
        "internal_handlers.Stock": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "2222"
                },
                "name": {
                    "type": "string",
                    "example": "Saudi Aramco"
                },
                "sector": {
                    "type": "string",
                    "example": "Energy"
                },
                "sharia_opinion": {
                    "type": "string",
                    "example": "compliant"
                }
            }
        },
        "internal_handlers.StockResponse": {
            "type": "object",
            "properties": {
                "stocks": {
                    "description": "List of stocks\nswagger:allOf",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/internal_handlers.Stock"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "Naqa API",
	Description:      "Stock Market API Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
