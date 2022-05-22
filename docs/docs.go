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
        "/admin": {
            "get": {
                "description": "Get all admin with data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin using Token JWT"
                ],
                "summary": "Get All admin",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object",
                                "additionalProperties": true
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "create admin with data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Create admin",
                "parameters": [
                    {
                        "description": "admin",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.AdminSwagger"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/admin.Admin"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/admin/city": {
            "get": {
                "description": "Get all city for admin using jwt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin using Token JWT"
                ],
                "summary": "Get All City",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "object",
                                "additionalProperties": true
                            }
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            },
            "post": {
                "description": "Create data city for admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin using Token JWT"
                ],
                "summary": "Create Data City",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Kota",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/admin/city/{id}": {
            "put": {
                "description": "Update data city for admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin using Token JWT"
                ],
                "summary": "Update Data City",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id city",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Kota",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.Kota"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "Delete data city for admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin using Token JWT"
                ],
                "summary": "Delete Data City",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id city",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/admin/token": {
            "post": {
                "description": "Get token for admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Get Token",
                "parameters": [
                    {
                        "description": "admin",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.InputAdminToken"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/admin/{id}": {
            "get": {
                "description": "Get Admin By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Get Admin By ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "anything id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "put": {
                "description": "update data admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin using Token JWT"
                ],
                "summary": "Update Admin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id admin",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "admin",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.AdminSwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "delete data admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Delete Admin",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id admin",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/cekresi": {
            "post": {
                "description": "Get detail resi from binderbyte",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get ongkir and check resi"
                ],
                "summary": "Get Detail Resi",
                "parameters": [
                    {
                        "description": "Resi",
                        "name": "Resi",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ongkir.Resi"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ongkir.CekResiBinderByte"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/city": {
            "post": {
                "description": "Get data city by name dan tipe for public",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "city"
                ],
                "summary": "Get data city by name dan tipe",
                "parameters": [
                    {
                        "description": "Kota",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.GetCityById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": ""
                    }
                }
            }
        },
        "/cost": {
            "post": {
                "description": "Get Cost Ongkir from raja ongkir",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get ongkir and check resi"
                ],
                "summary": "Get Cost Ongkir",
                "parameters": [
                    {
                        "description": "ongkir",
                        "name": "Ongkir",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ongkir.Ongkir"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ongkir.Results"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "admin.Admin": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.AdminSwagger": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "admin.GetCityById": {
            "type": "object",
            "properties": {
                "nama_kota": {
                    "type": "string"
                },
                "tipe_kota": {
                    "type": "string"
                }
            }
        },
        "admin.InputAdminToken": {
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
        "admin.Kota": {
            "type": "object",
            "required": [
                "kota_nama",
                "postal_code",
                "province_id",
                "rajaongkir_city_id",
                "tipe"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "kota_nama": {
                    "type": "string"
                },
                "postal_code": {
                    "type": "integer"
                },
                "province_id": {
                    "type": "integer"
                },
                "rajaongkir_city_id": {
                    "type": "integer"
                },
                "tipe": {
                    "type": "string"
                }
            }
        },
        "ongkir.CekResiBinderByte": {
            "type": "object",
            "properties": {
                "Message": {
                    "type": "string"
                },
                "data": {
                    "type": "object",
                    "properties": {
                        "history": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ongkir.History"
                            }
                        },
                        "summary": {
                            "type": "object",
                            "properties": {
                                "awb": {
                                    "type": "string"
                                },
                                "courier": {
                                    "type": "string"
                                },
                                "service": {
                                    "type": "string"
                                },
                                "status": {
                                    "type": "string"
                                },
                                "weight": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "ongkir.Cost": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "etd": {
                                "type": "string"
                            },
                            "note": {
                                "type": "string"
                            },
                            "value": {
                                "type": "integer"
                            }
                        }
                    }
                },
                "description": {
                    "type": "string"
                },
                "service": {
                    "type": "string"
                }
            }
        },
        "ongkir.History": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                }
            }
        },
        "ongkir.Ongkir": {
            "type": "object",
            "required": [
                "asal",
                "berat",
                "tipe_asal",
                "tipe_tujuan",
                "tujuan"
            ],
            "properties": {
                "asal": {
                    "type": "string"
                },
                "berat": {
                    "type": "string"
                },
                "tipe_asal": {
                    "type": "string"
                },
                "tipe_tujuan": {
                    "type": "string"
                },
                "tujuan": {
                    "type": "string"
                }
            }
        },
        "ongkir.Resi": {
            "type": "object",
            "required": [
                "kurir",
                "resi"
            ],
            "properties": {
                "kurir": {
                    "type": "string"
                },
                "resi": {
                    "type": "string"
                }
            }
        },
        "ongkir.Results": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "costs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ongkir.Cost"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API Jasa Pengiriman",
	Description:      "Berikut API Jasa Pengiriman",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
