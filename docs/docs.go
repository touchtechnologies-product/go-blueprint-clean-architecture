// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/staff": {
            "post": {
                "description": "CreateStaff",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create Statff",
                "parameters": [
                    {
                        "description": "Request Ex.",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/staff.CreateStaffInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/staff.CreateStaffOutput"
                        }
                    }
                }
            }
        },
        "/staffsByCompany": {
            "get": {
                "description": "Get all the existing Staff",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "List Staff By Company",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/staff.GetStaffsByCompanyOutput"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "staff.CreateStaffInput": {
            "type": "object",
            "required": [
                "companyId",
                "name"
            ],
            "properties": {
                "companyId": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tel": {
                    "type": "string"
                }
            }
        },
        "staff.CreateStaffOutput": {
            "type": "object",
            "properties": {
                "staff": {
                    "type": "object",
                    "$ref": "#/definitions/staff.Staff"
                }
            }
        },
        "staff.GetStaffsByCompanyOutput": {
            "type": "object",
            "properties": {
                "staffs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/staff.Staff"
                    }
                }
            }
        },
        "staff.Staff": {
            "type": "object",
            "properties": {
                "companyId": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "tel": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
