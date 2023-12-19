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
            "name": "kavinkishore",
            "url": "iamkavin1309@gmqil.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tracks/artist/{artist}": {
            "get": {
                "description": "Retrieves tracks by the specified artist from the database or Spotify.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetTracks"
                ],
                "summary": "Get tracks by artist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Artist name",
                        "name": "artist",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved tracks",
                        "schema": {
                            "$ref": "#/definitions/model.TrackDeatils"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tracks/create": {
            "post": {
                "description": "Create a new track record in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create or Update Tracks"
                ],
                "summary": "Create a new track",
                "parameters": [
                    {
                        "description": "Track details to create",
                        "name": "track",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Track details",
                        "schema": {
                            "$ref": "#/definitions/model.TrackDeatils"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Track with ISRC code already exists",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tracks/update/{isrc}": {
            "put": {
                "description": "Updates a track with the specified ISRC.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create or Update Tracks"
                ],
                "summary": "Update a track by ISRC",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ISRC of the track",
                        "name": "isrc",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated track details",
                        "name": "existingTrack",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TrackDeatils"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Track updated successfully",
                        "schema": {
                            "$ref": "#/definitions/model.TrackDeatils"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Track not found",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tracks/{isrc}": {
            "get": {
                "description": "Search tracks from the database or Spotify by ISRC code",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetTracks"
                ],
                "summary": "Search tracks by ISRC code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ISRC code of the track",
                        "name": "isrc",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Track details",
                        "schema": {
                            "$ref": "#/definitions/model.TrackDeatils"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Track not found for the given ISRC",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.RequestBody": {
            "type": "object",
            "required": [
                "isrc"
            ],
            "properties": {
                "isrc": {
                    "type": "string"
                }
            }
        },
        "model.TrackDeatils": {
            "type": "object",
            "properties": {
                "artists": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "isrc": {
                    "type": "string"
                },
                "popularity": {
                    "type": "integer"
                },
                "primary_key": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "Documentation for this project can be found on GitHub.",
        "url": "https://github.com/yourusername/yourproject/docs"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Spotify ISRC Finder",
	Description:      "This is a RESTful API for managing music tracks using Gin, Gorm, and PostgreSQL.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
