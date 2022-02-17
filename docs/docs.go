// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate_swagger = `{
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
        "/company/post": {
            "get": {
                "description": "All Post from the app that has been posted",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "All Post from the app that has been posted",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/company/post/:id": {
            "get": {
                "description": "All Post from the app that has been posted",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "All Post from the app that has been posted",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/company/user/:id": {
            "get": {
                "description": "Retrieves detail one user",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Retrieves detail one user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Retrieves update point on one post",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Retrieves update point on one post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/company/user/login": {
            "get": {
                "description": "Retrieves the list of users who has to login in the app",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Retrieves the list of users who has to login in the app",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/company/user/point": {
            "get": {
                "description": "Retrieves all users point",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Retrieves all users point",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/company/user/point/:id": {
            "get": {
                "description": "Retrieves detail point from one user",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Retrieves detail point from one user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/company/user/register": {
            "get": {
                "description": "Retrieves the list of users who has to register in the app",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Retrieves the list of users who has to register in the app",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/form/login": {
            "post": {
                "description": "Login for user and company company in the app",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Form"
                ],
                "summary": "Login for user and company in the app",
                "parameters": [
                    {
                        "description": "Company",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bearer {token}\"}",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/form/register/company": {
            "post": {
                "description": "Register company in the app",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Form"
                ],
                "summary": "Register company in the app",
                "parameters": [
                    {
                        "description": "Company",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Company"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Company"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/form/register/user": {
            "post": {
                "description": "Register user in the app",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Form"
                ],
                "summary": "Register user in the app",
                "parameters": [
                    {
                        "description": "User",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/post": {
            "post": {
                "description": "Retrieves the list of users who has to register in the app",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Retrieves the list of users who has to register in the app",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/post/:id": {
            "get": {
                "description": "All Post from the app that has been posted",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "All Post from the app that has been posted",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Post"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "All Post from the app that has been posted",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "All Post from the app that has been posted",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Post"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/point": {
            "get": {
                "description": "Retrieves all users point",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Retrieves all users point",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.LoginDto": {
            "type": "object",
            "required": [
                "email",
                "password",
                "role"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "entity.Company": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "Company Email\n@Column(type:varchar(400), nullable:false)",
                    "type": "string"
                },
                "id": {
                    "description": "Company ID\n@PrimaryKey\n@Column(type:uuid, unique: true)",
                    "type": "string"
                },
                "is_active": {
                    "description": "Company IsActive\n@Column(type:boolean, default:true)",
                    "type": "boolean"
                },
                "name": {
                    "description": "Company Name\n@Column(type:varchar(400), nullable:false)",
                    "type": "string"
                },
                "password": {
                    "description": "Company Phone\n@Column(type:varchar(400), nullable:false)",
                    "type": "string"
                }
            }
        },
        "entity.Post": {
            "type": "object",
            "properties": {
                "author_id": {
                    "description": "Post Author ID\n@Column(type:uuid, nullable:false)",
                    "type": "string"
                },
                "content": {
                    "description": "Post Content\n@Column(type:varchar(400), nullable:false)",
                    "type": "string"
                },
                "hastag": {
                    "description": "Post Hastag\n@Column(type:varchar(400), nullable:true)",
                    "type": "string"
                },
                "id": {
                    "description": "Post ID\n@PrimaryKey\n@Column(type:int, unique: true, autoincrement: true)",
                    "type": "integer"
                },
                "photo": {
                    "description": "Post Photo\n@Column(type:text, nullable:false)",
                    "type": "string"
                },
                "point": {
                    "description": "Post Point\n@Column(type:int, nullable:false, default:0)",
                    "type": "integer"
                },
                "source": {
                    "description": "Post Source\n@Column(type:varchar(400), nullable:true)",
                    "type": "string"
                },
                "thumbnail": {
                    "description": "Post Thumbnaiil\n@Column(type:text, nullable:false)",
                    "type": "string"
                },
                "title": {
                    "description": "Post Title\n@Column(type:varchar(400), nullable:false)",
                    "type": "string"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "description": "Company Email\n@Column(type:varchar(400), nullable:false)",
                    "type": "string"
                },
                "id": {
                    "description": "Company ID\n@PrimaryKey\n@Column(type:uuid, unique: true)",
                    "type": "string"
                },
                "is_active": {
                    "description": "Company IsActive\n@Column(type:boolean, default:true)",
                    "type": "boolean"
                },
                "name": {
                    "description": "Company Name\n@Column(type:varchar(400), nullable:false)",
                    "type": "string"
                },
                "password": {
                    "description": "Company Phone\n@Column(type:varchar(400), nullable:false)",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo_swagger holds exported Swagger Info so clients can modify it
var SwaggerInfo_swagger = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Gin Swagger Example API",
	Description:      "This is a sample server server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate_swagger,
}

func init() {
	swag.Register(SwaggerInfo_swagger.InstanceName(), SwaggerInfo_swagger)
}