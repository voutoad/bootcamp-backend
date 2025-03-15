// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users": {
            "get": {
                "description": "Get user with optional filters",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Type",
                        "name": "type",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Tag",
                        "name": "tag",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.UserResponseDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Create a new user with given body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "patch": {
                "description": "Updating user tags with given username and tags",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update user tags",
                "parameters": [
                    {
                        "description": "User Update",
                        "name": "userUpdate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserUpdateDTO"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/users/{username}": {
            "get": {
                "description": "Get user with given username or get a 404 error",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by his username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UserResponseDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateUserDTO": {
            "type": "object",
            "required": [
                "age",
                "description",
                "interests",
                "password",
                "rating",
                "type",
                "username"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 0
                },
                "description": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                },
                "image_url": {
                    "type": "string"
                },
                "interests": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                },
                "password": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 8
                },
                "rating": {
                    "type": "integer",
                    "maximum": 100,
                    "minimum": 0
                },
                "type": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                },
                "username": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                }
            }
        },
        "dto.UserResponseDTO": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "image_url": {
                    "type": "string"
                },
                "interests": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                },
                "tags": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.UserUpdateDTO": {
            "type": "object",
            "required": [
                "tags",
                "username"
            ],
            "properties": {
                "tags": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 4
                },
                "username": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 2
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:5000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "PickMe API",
	Description:      "Simple backend for PickMe product",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
