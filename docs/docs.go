// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "DeboZero Tech Team",
            "email": "debidarmawan1998@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login to get access token",
                "parameters": [
                    {
                        "description": "Login data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/global.Response-dto_LoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Remove/Invalidate an access token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/global.Response-dto_Message"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Refresh an access token (get a new one)",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/global.Response-dto_RefreshTokenResponse"
                        }
                    }
                }
            }
        },
        "/auth/verify": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Verify an access token to get the user id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Path",
                        "name": "X-Path",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Method",
                        "name": "X-Method",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/global.Response-dto_VerifyResponse"
                        }
                    }
                }
            }
        },
        "/oauth2/client": {
            "post": {
                "description": "Create Oauth2 Client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Oauth2 Client"
                ],
                "summary": "Create Oauth2 Client",
                "parameters": [
                    {
                        "description": "Oauth2 Client data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.Oauth2Client"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/global.Response-dto_Oauth2ClientResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create an account",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/global.Response-dto_Message"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.Login": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@email.com"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                }
            }
        },
        "dto.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.Message": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.Oauth2Client": {
            "type": "object",
            "required": [
                "client_id",
                "domain",
                "name"
            ],
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "domain": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.Oauth2ClientResponse": {
            "type": "object",
            "required": [
                "client_id",
                "name"
            ],
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "secret_key": {
                    "type": "string"
                }
            }
        },
        "dto.RefreshTokenRequest": {
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
        "dto.RefreshTokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "expires_at": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "dto.UserRegisterRequest": {
            "type": "object",
            "required": [
                "confirm_password",
                "email",
                "name",
                "password",
                "phone",
                "username"
            ],
            "properties": {
                "confirm_password": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 6,
                    "example": "password"
                },
                "email": {
                    "type": "string",
                    "example": "william@debozero.id"
                },
                "name": {
                    "type": "string",
                    "maxLength": 60,
                    "minLength": 3,
                    "example": "William"
                },
                "password": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 6,
                    "example": "password"
                },
                "phone": {
                    "type": "string",
                    "example": "08218833123"
                },
                "username": {
                    "type": "string",
                    "maxLength": 60,
                    "minLength": 3,
                    "example": "william"
                }
            }
        },
        "dto.VerifyResponse": {
            "type": "object",
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "global.Response-dto_LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/dto.LoginResponse"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "global.Response-dto_Message": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/dto.Message"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "global.Response-dto_Oauth2ClientResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/dto.Oauth2ClientResponse"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "global.Response-dto_RefreshTokenResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/dto.RefreshTokenResponse"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "global.Response-dto_VerifyResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/dto.VerifyResponse"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "X-API-KEY",
            "in": "header"
        },
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
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "DeboZero Core Service",
	Description:      "This is an API documentation of DeboZero Core Backend Service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
