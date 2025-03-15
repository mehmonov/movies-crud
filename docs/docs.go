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
            "email": "mehmonov.husniddin1@gmail.com"
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
        "/auth/login": {
            "post": {
                "description": "Login with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "description": "Get new access token using refresh token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Refresh token",
                "parameters": [
                    {
                        "description": "Refresh token",
                        "name": "refresh_token",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TokenPair"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a new user with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "description": "Get a list of all movies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Get all movies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Movie"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new movie to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Create a new movie",
                "parameters": [
                    {
                        "description": "Movie information",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateMovieRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/movies/{id}": {
            "get": {
                "description": "Get details of a specific movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Get a movie by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing movie's details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Update a movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Movie information",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateMovieRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a movie from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Delete a movie",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/movies/{id}/file": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Upload a video file for a movie",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Upload movie file",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Movie file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.MovieFileResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/movies/{id}/files/{fileId}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Download a movie file by ID",
                "produces": [
                    "application/octet-stream"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Download movie file",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "File ID",
                        "name": "fileId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/movies/{id}/media": {
            "get": {
                "description": "Get all media files for a specific movie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "movies"
                ],
                "summary": "Get movie media files",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Media type (poster, backdrop, trailer)",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.MovieMedia"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateMovieRequest": {
            "type": "object",
            "required": [
                "director",
                "title",
                "year"
            ],
            "properties": {
                "director": {
                    "type": "string",
                    "maxLength": 100
                },
                "duration": {
                    "description": "Maximum 1000 minutes (16.6 hours)",
                    "type": "integer",
                    "maximum": 1000,
                    "minimum": 0
                },
                "genre": {
                    "type": "string",
                    "maxLength": 50
                },
                "media_files": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MovieMediaRequest"
                    }
                },
                "metadata": {
                    "$ref": "#/definitions/models.MovieMetadataRequest"
                },
                "plot": {
                    "type": "string"
                },
                "rating": {
                    "type": "number",
                    "maximum": 10,
                    "minimum": 0
                },
                "title": {
                    "type": "string",
                    "maxLength": 100
                },
                "year": {
                    "type": "integer",
                    "maximum": 2100,
                    "minimum": 1800
                }
            }
        },
        "models.CreateUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        },
        "models.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "tokens": {
                    "$ref": "#/definitions/models.TokenPair"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "models.Movie": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "director": {
                    "type": "string"
                },
                "duration": {
                    "description": "Minutes",
                    "type": "integer"
                },
                "genre": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "media_files": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MovieMedia"
                    }
                },
                "metadata": {
                    "$ref": "#/definitions/models.MovieMetadata"
                },
                "plot": {
                    "type": "string"
                },
                "rating": {
                    "type": "number"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.MovieFileResponse": {
            "type": "object",
            "properties": {
                "content_type": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "file_name": {
                    "type": "string"
                },
                "file_size": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "movie_id": {
                    "type": "integer"
                }
            }
        },
        "models.MovieMedia": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_main": {
                    "type": "boolean"
                },
                "movie_id": {
                    "type": "integer"
                },
                "type": {
                    "description": "poster, backdrop, trailer",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.MovieMediaRequest": {
            "type": "object",
            "required": [
                "type",
                "url"
            ],
            "properties": {
                "is_main": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "poster",
                        "backdrop",
                        "trailer"
                    ]
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.MovieMetadata": {
            "type": "object",
            "properties": {
                "awards": {
                    "type": "string"
                },
                "cast": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "language": {
                    "type": "string"
                },
                "movie_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.MovieMetadataRequest": {
            "type": "object",
            "properties": {
                "awards": {
                    "type": "string"
                },
                "cast": {
                    "type": "string"
                },
                "country": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                }
            }
        },
        "models.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.TokenPair": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.UpdateMovieRequest": {
            "type": "object",
            "properties": {
                "director": {
                    "type": "string",
                    "maxLength": 100
                },
                "duration": {
                    "type": "integer",
                    "maximum": 1000,
                    "minimum": 0
                },
                "genre": {
                    "type": "string",
                    "maxLength": 50
                },
                "media_files": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MovieMediaRequest"
                    }
                },
                "metadata": {
                    "$ref": "#/definitions/models.MovieMetadataRequest"
                },
                "plot": {
                    "type": "string"
                },
                "rating": {
                    "type": "number",
                    "maximum": 10,
                    "minimum": 0
                },
                "title": {
                    "type": "string",
                    "maxLength": 100
                },
                "year": {
                    "type": "integer",
                    "maximum": 2100,
                    "minimum": 1800
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Movies CRUD API",
	Description:      "This is a sample movies CRUD API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
