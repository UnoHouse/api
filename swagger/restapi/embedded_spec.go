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
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Service to write/read Web panel data",
    "title": "UnoHouse API",
    "version": "0.0.1"
  },
  "basePath": "/",
  "paths": {
    "/healthcheck": {
      "get": {
        "summary": "Get Health status",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/badRequest"
          },
          "500": {
            "$ref": "#/responses/internalServerError"
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "location",
        "locationType",
        "reason"
      ],
      "properties": {
        "location": {
          "description": "The name of the field/property where the issue occurs",
          "type": "string"
        },
        "locationType": {
          "description": "Place in request where error origin can be found - possible values are header, body, query",
          "type": "string"
        },
        "message": {
          "type": "string"
        },
        "reason": {
          "type": "string"
        }
      }
    },
    "ErrorResponse": {
      "type": "object",
      "required": [
        "code",
        "reason"
      ],
      "properties": {
        "code": {
          "description": "http status code",
          "type": "integer"
        },
        "errors": {
          "description": "array of error objects containing more detail information",
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/Error"
          }
        },
        "message": {
          "description": "short error description",
          "type": "string"
        },
        "reason": {
          "description": "http status reason",
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "companyId": {
      "type": "string",
      "description": "company id filter - Values are separated by | (e.g: ?companyId=1|2|3)",
      "name": "companyId",
      "in": "query"
    },
    "id": {
      "type": "string",
      "description": "resource unique identifier",
      "name": "id",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "badRequest": {
      "description": "Bad request",
      "schema": {
        "$ref": "#/definitions/ErrorResponse"
      }
    },
    "internalServerError": {
      "description": "Internal server error",
      "schema": {
        "$ref": "#/definitions/ErrorResponse"
      }
    },
    "notFound": {
      "description": "Not found",
      "schema": {
        "$ref": "#/definitions/ErrorResponse"
      }
    },
    "serviceUnavailable": {
      "description": "Service unavailable",
      "schema": {
        "$ref": "#/definitions/ErrorResponse"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Service to write/read Web panel data",
    "title": "UnoHouse API",
    "version": "0.0.1"
  },
  "basePath": "/",
  "paths": {
    "/healthcheck": {
      "get": {
        "summary": "Get Health status",
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "location",
        "locationType",
        "reason"
      ],
      "properties": {
        "location": {
          "description": "The name of the field/property where the issue occurs",
          "type": "string"
        },
        "locationType": {
          "description": "Place in request where error origin can be found - possible values are header, body, query",
          "type": "string"
        },
        "message": {
          "type": "string"
        },
        "reason": {
          "type": "string"
        }
      }
    },
    "ErrorResponse": {
      "type": "object",
      "required": [
        "code",
        "reason"
      ],
      "properties": {
        "code": {
          "description": "http status code",
          "type": "integer"
        },
        "errors": {
          "description": "array of error objects containing more detail information",
          "type": "array",
          "minItems": 1,
          "items": {
            "$ref": "#/definitions/Error"
          }
        },
        "message": {
          "description": "short error description",
          "type": "string"
        },
        "reason": {
          "description": "http status reason",
          "type": "string"
        }
      }
    }
  },
  "parameters": {
    "companyId": {
      "type": "string",
      "description": "company id filter - Values are separated by | (e.g: ?companyId=1|2|3)",
      "name": "companyId",
      "in": "query"
    },
    "id": {
      "type": "string",
      "description": "resource unique identifier",
      "name": "id",
      "in": "path",
      "required": true
    }
  },
  "responses": {
    "badRequest": {
      "description": "Bad request",
      "schema": {
        "$ref": "#/definitions/ErrorResponse"
      }
    },
    "internalServerError": {
      "description": "Internal server error",
      "schema": {
        "$ref": "#/definitions/ErrorResponse"
      }
    },
    "notFound": {
      "description": "Not found",
      "schema": {
        "$ref": "#/definitions/ErrorResponse"
      }
    },
    "serviceUnavailable": {
      "description": "Service unavailable",
      "schema": {
        "$ref": "#/definitions/ErrorResponse"
      }
    }
  }
}`))
}
