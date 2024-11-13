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
        "/krathong": {
            "get": {
                "description": "Retrieve the latest 50 Krathong entries",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Krathong"
                ],
                "summary": "Get Krathongs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetKrathongListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.KrathongDataResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Krathong entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Krathong"
                ],
                "summary": "Create Krathong",
                "parameters": [
                    {
                        "description": "Krathong Data",
                        "name": "krathong",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Krathong"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handlers.KrathongDataResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.KrathongDataResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.KrathongDataResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.GetKrathongListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Krathong"
                    }
                },
                "response_code": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                }
            }
        },
        "handlers.KrathongDataResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.Krathong"
                },
                "response_code": {
                    "type": "string"
                },
                "response_message": {
                    "type": "string"
                }
            }
        },
        "models.Krathong": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "emp_department": {
                    "type": "string"
                },
                "emp_name": {
                    "type": "string"
                },
                "emp_wish": {
                    "type": "string"
                },
                "krathong_id": {
                    "type": "integer"
                },
                "krathong_type": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "10.144.130.86:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Loy Krathong API",
	Description:      "This is a API for the Loy Krathong festival.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
