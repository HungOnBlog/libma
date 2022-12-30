// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/borrowers": {
            "get": {
                "description": "Return all borrowers (active) in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "borrowers"
                ],
                "summary": "Get all borrowers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/borrowers.BorrowerDto"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create borrower",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "borrowers"
                ],
                "summary": "Create borrower",
                "parameters": [
                    {
                        "description": "Borrower",
                        "name": "borrower",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/borrowers.BorrowerDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/borrowers.BorrowerDto"
                        }
                    }
                }
            }
        },
        "/borrowers/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "borrowers"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Borrower",
                        "name": "borrower",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/borrowers.BorrowerLoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/borrowers.BorrowerDto"
                        }
                    }
                }
            }
        },
        "/borrowers/{id}": {
            "get": {
                "description": "Return borrower by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "borrowers"
                ],
                "summary": "Get borrower by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Borrower id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/borrowers.BorrowerDto"
                        }
                    }
                }
            },
            "put": {
                "description": "Update borrower",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "borrowers"
                ],
                "summary": "Update borrower",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Borrower id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Borrower",
                        "name": "borrower",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/borrowers.BorrowerUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/borrowers.BorrowerDto"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete borrower",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "borrowers"
                ],
                "summary": "Delete borrower",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Borrower id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "borrowers.BorrowerDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "borrowers.BorrowerLoginDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "borrowers.BorrowerUpdateDto": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
