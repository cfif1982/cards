// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample User's cards application based on the OpenAPI 2.0 specification.",
    "title": "cards",
    "contact": {},
    "version": "v1"
  },
  "host": "localhost:8080",
  "basePath": "/api",
  "paths": {
    "/bank": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Update an existing bank by uuid",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "bank"
        ],
        "summary": "Update an existing bank",
        "operationId": "updateBank",
        "parameters": [
          {
            "description": "Update an existent bank in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Bank not found"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Add a new bank to the base",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "bank"
        ],
        "summary": "Add a new bank",
        "operationId": "addBank",
        "parameters": [
          {
            "description": "Create a new bank in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewBank"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      }
    },
    "/bank/{bankId}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Returns a single bank",
        "produces": [
          "application/json"
        ],
        "tags": [
          "bank"
        ],
        "summary": "Find bank by uuid",
        "operationId": "getBankByUUID",
        "parameters": [
          {
            "type": "integer",
            "format": "byte16",
            "description": "ID of bank to return",
            "name": "bankId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Bank not found"
          }
        }
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "delete a bank",
        "tags": [
          "bank"
        ],
        "summary": "Deletes a bank",
        "operationId": "deleteBank",
        "parameters": [
          {
            "type": "integer",
            "format": "byte16",
            "description": "Bank uuid to delete",
            "name": "bankId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid bank value"
          }
        }
      }
    },
    "/card": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Update an existing card by uuid",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "card"
        ],
        "summary": "Update an existing card",
        "operationId": "updateCard",
        "parameters": [
          {
            "description": "Update an existent card in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Card"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Card not found"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Add a new card to the base",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "card"
        ],
        "summary": "Add a new card",
        "operationId": "addCard",
        "parameters": [
          {
            "description": "Create a new card in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewCard"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      }
    },
    "/card/{cardNumber}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Returns a single card",
        "produces": [
          "application/json"
        ],
        "tags": [
          "card"
        ],
        "summary": "Find card by number",
        "operationId": "getCardByNumber",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "Number of card to return",
            "name": "cardNumber",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Card not found"
          }
        }
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "delete a card",
        "tags": [
          "card"
        ],
        "summary": "Deletes a card",
        "operationId": "deleteCard",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "Card number to delete",
            "name": "cardNumber",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid card value"
          }
        }
      }
    },
    "/user": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Update an existing user by uuid",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Update an existing user",
        "operationId": "updateUser",
        "parameters": [
          {
            "description": "Update an existent user in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "User not found"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Add a new user to the base",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Add a new user",
        "operationId": "addUser",
        "parameters": [
          {
            "description": "Create a new user in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      }
    },
    "/user/{userId}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Returns a single user",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Find user by uuid",
        "operationId": "getUserByUUID",
        "parameters": [
          {
            "type": "integer",
            "format": "byte16",
            "description": "ID of user to return",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "delete a user",
        "tags": [
          "user"
        ],
        "summary": "Deletes a user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "integer",
            "format": "byte16",
            "description": "User uuid to delete",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid user value"
          }
        }
      }
    }
  },
  "definitions": {
    "Bank": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string",
          "example": "Moscow, Kremlin"
        },
        "bik": {
          "type": "integer",
          "format": "uint32",
          "example": 44525974
        },
        "name": {
          "type": "string",
          "example": "Main Bank"
        },
        "telephone": {
          "type": "string",
          "example": "565-56-56"
        },
        "uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        }
      }
    },
    "Card": {
      "type": "object",
      "properties": {
        "bank_uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        },
        "cvs": {
          "type": "string",
          "example": "109"
        },
        "number": {
          "type": "integer",
          "format": "uint64",
          "example": 4242424242424242
        },
        "password": {
          "type": "string",
          "example": "12345"
        },
        "user_name": {
          "type": "string",
          "example": "JOHN DOE"
        },
        "user_uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        },
        "valid": {
          "type": "string",
          "example": "03/23"
        }
      }
    },
    "NewBank": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string",
          "example": "Moscow, Kremlin"
        },
        "bik": {
          "type": "integer",
          "format": "uint32",
          "example": 44525974
        },
        "name": {
          "type": "string",
          "example": "Main Bank"
        },
        "telephone": {
          "type": "string",
          "example": "565-56-56"
        }
      }
    },
    "NewCard": {
      "type": "object",
      "properties": {
        "bank_uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        },
        "cvs": {
          "type": "string",
          "example": "109"
        },
        "number": {
          "type": "integer",
          "format": "uint64",
          "example": 4242424242424242
        },
        "password": {
          "type": "string",
          "example": "12345"
        },
        "user_name": {
          "type": "string",
          "example": "JOHN DOE"
        },
        "user_uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        },
        "valid": {
          "type": "string",
          "example": "03/23"
        }
      }
    },
    "NewUser": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "example": "john@email.com"
        },
        "first_name": {
          "type": "string",
          "example": "John"
        },
        "last_name": {
          "type": "string",
          "example": "Doe"
        },
        "login": {
          "type": "string",
          "example": "login"
        },
        "password": {
          "type": "string",
          "example": "12345"
        },
        "telephone": {
          "type": "string",
          "example": "12345"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "example": "john@email.com"
        },
        "last_name": {
          "type": "string",
          "example": "Doe"
        },
        "login": {
          "type": "string",
          "example": "login"
        },
        "name": {
          "type": "string",
          "example": "John"
        },
        "password": {
          "type": "string",
          "example": "12345"
        },
        "telephone": {
          "type": "string",
          "example": "12345"
        },
        "uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "users operation",
      "name": "user"
    },
    {
      "description": "banks operation",
      "name": "bank"
    },
    {
      "description": "cards operation",
      "name": "card"
    }
  ],
  "x-components": {}
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a sample User's cards application based on the OpenAPI 2.0 specification.",
    "title": "cards",
    "contact": {},
    "version": "v1"
  },
  "host": "localhost:8080",
  "basePath": "/api",
  "paths": {
    "/bank": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Update an existing bank by uuid",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "bank"
        ],
        "summary": "Update an existing bank",
        "operationId": "updateBank",
        "parameters": [
          {
            "description": "Update an existent bank in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Bank not found"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Add a new bank to the base",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "bank"
        ],
        "summary": "Add a new bank",
        "operationId": "addBank",
        "parameters": [
          {
            "description": "Create a new bank in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewBank"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      }
    },
    "/bank/{bankId}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Returns a single bank",
        "produces": [
          "application/json"
        ],
        "tags": [
          "bank"
        ],
        "summary": "Find bank by uuid",
        "operationId": "getBankByUUID",
        "parameters": [
          {
            "type": "integer",
            "format": "byte16",
            "description": "ID of bank to return",
            "name": "bankId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Bank"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Bank not found"
          }
        }
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "delete a bank",
        "tags": [
          "bank"
        ],
        "summary": "Deletes a bank",
        "operationId": "deleteBank",
        "parameters": [
          {
            "type": "integer",
            "format": "byte16",
            "description": "Bank uuid to delete",
            "name": "bankId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid bank value"
          }
        }
      }
    },
    "/card": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Update an existing card by uuid",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "card"
        ],
        "summary": "Update an existing card",
        "operationId": "updateCard",
        "parameters": [
          {
            "description": "Update an existent card in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Card"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Card not found"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Add a new card to the base",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "card"
        ],
        "summary": "Add a new card",
        "operationId": "addCard",
        "parameters": [
          {
            "description": "Create a new card in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewCard"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      }
    },
    "/card/{cardNumber}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Returns a single card",
        "produces": [
          "application/json"
        ],
        "tags": [
          "card"
        ],
        "summary": "Find card by number",
        "operationId": "getCardByNumber",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "Number of card to return",
            "name": "cardNumber",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Card"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Card not found"
          }
        }
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "delete a card",
        "tags": [
          "card"
        ],
        "summary": "Deletes a card",
        "operationId": "deleteCard",
        "parameters": [
          {
            "type": "integer",
            "format": "uint64",
            "description": "Card number to delete",
            "name": "cardNumber",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid card value"
          }
        }
      }
    },
    "/user": {
      "put": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Update an existing user by uuid",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Update an existing user",
        "operationId": "updateUser",
        "parameters": [
          {
            "description": "Update an existent user in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "User not found"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Add a new user to the base",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Add a new user",
        "operationId": "addUser",
        "parameters": [
          {
            "description": "Create a new user in the base",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewUser"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      }
    },
    "/user/{userId}": {
      "get": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "Returns a single user",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Find user by uuid",
        "operationId": "getUserByUUID",
        "parameters": [
          {
            "type": "integer",
            "format": "byte16",
            "description": "ID of user to return",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      },
      "delete": {
        "security": [
          {
            "api_key": []
          }
        ],
        "description": "delete a user",
        "tags": [
          "user"
        ],
        "summary": "Deletes a user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "integer",
            "format": "byte16",
            "description": "User uuid to delete",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "400": {
            "description": "Invalid user value"
          }
        }
      }
    }
  },
  "definitions": {
    "Bank": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string",
          "example": "Moscow, Kremlin"
        },
        "bik": {
          "type": "integer",
          "format": "uint32",
          "example": 44525974
        },
        "name": {
          "type": "string",
          "example": "Main Bank"
        },
        "telephone": {
          "type": "string",
          "example": "565-56-56"
        },
        "uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        }
      }
    },
    "Card": {
      "type": "object",
      "properties": {
        "bank_uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        },
        "cvs": {
          "type": "string",
          "example": "109"
        },
        "number": {
          "type": "integer",
          "format": "uint64",
          "example": 4242424242424242
        },
        "password": {
          "type": "string",
          "example": "12345"
        },
        "user_name": {
          "type": "string",
          "example": "JOHN DOE"
        },
        "user_uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        },
        "valid": {
          "type": "string",
          "example": "03/23"
        }
      }
    },
    "NewBank": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string",
          "example": "Moscow, Kremlin"
        },
        "bik": {
          "type": "integer",
          "format": "uint32",
          "example": 44525974
        },
        "name": {
          "type": "string",
          "example": "Main Bank"
        },
        "telephone": {
          "type": "string",
          "example": "565-56-56"
        }
      }
    },
    "NewCard": {
      "type": "object",
      "properties": {
        "bank_uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        },
        "cvs": {
          "type": "string",
          "example": "109"
        },
        "number": {
          "type": "integer",
          "format": "uint64",
          "example": 4242424242424242
        },
        "password": {
          "type": "string",
          "example": "12345"
        },
        "user_name": {
          "type": "string",
          "example": "JOHN DOE"
        },
        "user_uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        },
        "valid": {
          "type": "string",
          "example": "03/23"
        }
      }
    },
    "NewUser": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "example": "john@email.com"
        },
        "first_name": {
          "type": "string",
          "example": "John"
        },
        "last_name": {
          "type": "string",
          "example": "Doe"
        },
        "login": {
          "type": "string",
          "example": "login"
        },
        "password": {
          "type": "string",
          "example": "12345"
        },
        "telephone": {
          "type": "string",
          "example": "12345"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string",
          "example": "john@email.com"
        },
        "last_name": {
          "type": "string",
          "example": "Doe"
        },
        "login": {
          "type": "string",
          "example": "login"
        },
        "name": {
          "type": "string",
          "example": "John"
        },
        "password": {
          "type": "string",
          "example": "12345"
        },
        "telephone": {
          "type": "string",
          "example": "12345"
        },
        "uuid": {
          "type": "integer",
          "format": "byte16",
          "example": "53aa35c8-e659-44b2-882f-f6056e443c99"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "users operation",
      "name": "user"
    },
    {
      "description": "banks operation",
      "name": "bank"
    },
    {
      "description": "cards operation",
      "name": "card"
    }
  ],
  "x-components": {}
}`))
}
