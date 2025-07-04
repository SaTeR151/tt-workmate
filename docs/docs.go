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
        "/api/task/delete": {
            "delete": {
                "description": "Delete finished task by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Delete task",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/dto.ResponceId"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/restutils.HTTPError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/restutils.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/task/info": {
            "get": {
                "description": "Get task status, date of creation and processing time by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Get task info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "task id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/taskManager.TaskInfo"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/restutils.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/task/new": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Create a new task",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.ResponceId"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.ResponceId": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "restutils.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                }
            }
        },
        "taskManager.TaskInfo": {
            "type": "object",
            "properties": {
                "create_date": {
                    "type": "string"
                },
                "process_time": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.6.1",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Test Task I/O bound",
	Description:      "Server for create, read and delete tasks",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
