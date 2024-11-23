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
        "/appartements": {
            "get": {
                "description": "Récupère la liste de tous les appartements disponibles",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appartements"
                ],
                "summary": "Récupérer tous les appartements",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.AppartementsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/services.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Crée un appartement avec les informations fournies dans le corps de la requête",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appartements"
                ],
                "summary": "Créer un nouvel appartement",
                "parameters": [
                    {
                        "description": "Appartement à créer",
                        "name": "appartement",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Appartement"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.AppartementResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/services.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/appartements/{id}": {
            "get": {
                "description": "Récupère les détails d'un appartement spécifique en utilisant son ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appartements"
                ],
                "summary": "Récupérer un appartement par ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de l'appartement",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.AppartementResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/services.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/services.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Met à jour les informations d'un appartement existant en utilisant son ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appartements"
                ],
                "summary": "Mettre à jour un appartement existant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de l'appartement",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Données de l'appartement à mettre à jour",
                        "name": "appartement",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Appartement"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.AppartementResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/services.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/services.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/services.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Supprime un appartement en utilisant son ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "appartements"
                ],
                "summary": "Supprimer un appartement",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID de l'appartement",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.SuccessMessage"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/services.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/services.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Appartement": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "localisation": {
                    "$ref": "#/definitions/models.Localisation"
                },
                "nombre_chambre": {
                    "type": "string"
                },
                "prix_loyer": {
                    "type": "string"
                },
                "titre": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Localisation": {
            "type": "object",
            "properties": {
                "adresse": {
                    "type": "string"
                },
                "detail_adresse": {
                    "type": "string"
                },
                "latitude": {
                    "type": "string"
                },
                "localisation_id": {
                    "description": "gorm.Model",
                    "type": "integer"
                },
                "longitude": {
                    "type": "string"
                },
                "ville": {
                    "type": "string"
                }
            }
        },
        "services.AppartementResponse": {
            "type": "object",
            "properties": {
                "appartement": {
                    "$ref": "#/definitions/models.Appartement"
                }
            }
        },
        "services.AppartementsResponse": {
            "type": "object",
            "properties": {
                "appartements": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Appartement"
                    }
                }
            }
        },
        "services.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "services.SuccessMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "GMH API",
	Description:      "This is a sample server for GMH API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}