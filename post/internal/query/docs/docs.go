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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/admin//post/publish": {
            "post": {
                "description": "PublishPost account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "PublishPost Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "query",
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
            }
        },
        "/v1/admin/category": {
            "get": {
                "description": "CategoriesAll account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CategoriesAll Account",
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
            "put": {
                "description": "UpdateCategories account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "UpdateCategories Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Category ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ParentId in json category",
                        "name": "ParentId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Title in json category",
                        "name": "Title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "MetaTitle in json category",
                        "name": "MetaTitle",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Slug in json category",
                        "name": "Slug",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Content in json category",
                        "name": "Content",
                        "in": "query"
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
                "description": "CreateCategories account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CreateCategories Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ParentId in json category",
                        "name": "ParentId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Title in json category",
                        "name": "Title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "MetaTitle in json category",
                        "name": "MetaTitle",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Slug in json category",
                        "name": "Slug",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Content in json category",
                        "name": "Content",
                        "in": "query"
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
            "delete": {
                "description": "DeleteCategories account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "DeleteCategories Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "category ID",
                        "name": "id",
                        "in": "query",
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
            }
        },
        "/v1/admin/post-metas": {
            "get": {
                "description": "PostMetasAll account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "PostMetasAll Account",
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
            "put": {
                "description": "UpdatePostMetas account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "UpdatePostMetas Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post-metas ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PostId in json post-metas",
                        "name": "PostId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "TextKey in json post-metas",
                        "name": "TextKey",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Content in json post-metas",
                        "name": "Content",
                        "in": "query"
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
                "description": "CreatePostMetas account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CreatePostMetas Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "PostId in json post-metas",
                        "name": "PostId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "TextKey in json post-metas",
                        "name": "TextKey",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Content in json post-metas",
                        "name": "Content",
                        "in": "query"
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
            "delete": {
                "description": "DeletePostMetas account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "DeletePostMetas Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post-metas ID",
                        "name": "id",
                        "in": "query",
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
            }
        },
        "/v1/admin/post/list/not/publish": {
            "get": {
                "description": "ListNotPublishPost account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "ListNotPublishPost Account",
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
            }
        },
        "/v1/admin/post/list/publish": {
            "get": {
                "description": "ListPublishPost account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "ListPublishPost Account",
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
            }
        },
        "/v1/admin/profile": {
            "get": {
                "description": "ProfileAll account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "ProfileAll Account",
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
            }
        },
        "/v1/client/comment": {
            "get": {
                "description": "CommentAllPost account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CommentAllPost Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "query",
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
                "description": "CreateComment account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CreateComment Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "PostId in json Post",
                        "name": "PostId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "ParentId in json Post",
                        "name": "ParentId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Title in json Post",
                        "name": "Title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Content in json Post",
                        "name": "Content",
                        "in": "query"
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
            "delete": {
                "description": "RemoveComment account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "RemoveComment Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "query",
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
            }
        },
        "/v1/client/post": {
            "get": {
                "description": "PostsAll account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "PostsAll Account",
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
            }
        },
        "/v1/client/profile": {
            "get": {
                "description": "ProfileById account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "ProfileById Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "profile ID",
                        "name": "id",
                        "in": "query",
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
            "put": {
                "description": "UpdateProfile account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "UpdateProfile Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Profile ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "username in json profile",
                        "name": "UserId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "FirstName in json profile",
                        "name": "FirstName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "MiddleName in json profile",
                        "name": "MiddleName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "LastName in json profile",
                        "name": "LastName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Mobile in json profile",
                        "name": "Mobile",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Email in json profile",
                        "name": "Email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Profile in json profile",
                        "name": "Profile",
                        "in": "query"
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
                "description": "CreateProfile account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CreateProfile Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username in json profile",
                        "name": "UserId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "FirstName in json profile",
                        "name": "FirstName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "MiddleName in json profile",
                        "name": "MiddleName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "LastName in json profile",
                        "name": "LastName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Mobile in json profile",
                        "name": "Mobile",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Email in json profile",
                        "name": "Email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Profile in json profile",
                        "name": "Profile",
                        "in": "query"
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
            }
        },
        "/v1/write/post": {
            "post": {
                "description": "CreatePosts account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "CreatePosts Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "AuthorId in json post",
                        "name": "AuthorId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Title in json post",
                        "name": "Title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "MetaTitle in json post",
                        "name": "MetaTitle",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Slug in json post",
                        "name": "Slug",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Summary in json post",
                        "name": "Summary",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Content in json post",
                        "name": "Content",
                        "in": "query"
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
            "delete": {
                "description": "DeletePosts account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "DeletePosts Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "query",
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
            }
        },
        "/v1/writer/post": {
            "put": {
                "description": "UpdatePosts account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "UpdatePosts Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "AuthorId in json post",
                        "name": "AuthorId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Title in json post",
                        "name": "Title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "MetaTitle in json post",
                        "name": "MetaTitle",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Slug in json post",
                        "name": "Slug",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Summary in json post",
                        "name": "Summary",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Content in json post",
                        "name": "Content",
                        "in": "query"
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
                "description": "PostsById account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "PostsById Account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "id",
                        "in": "query",
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