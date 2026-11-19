package client

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	agents "github.com/elevenlabs/elevenlabs-go/agents"
	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
	types "github.com/elevenlabs/elevenlabs-go/types"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsToolsClient(t, 200, `{
  "tools": [
    {
      "id": "value",
      "tool_config": {
        "type": "webhook",
        "name": "value",
        "description": "value",
        "response_timeout_secs": 123,
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "tool_error_handling_mode": "auto",
        "dynamic_variables": {
          "dynamic_variable_placeholders": {
            "key": {}
          }
        },
        "execution_mode": "immediate",
        "api_schema": {
          "request_headers": {
            "key": "value"
          },
          "url": "value",
          "method": "GET",
          "path_params_schema": {
            "key": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            }
          },
          "query_params_schema": {
            "properties": {
              "key": {
                "type": "boolean",
                "description": "value",
                "enum": [
                  "value"
                ],
                "is_system_provided": true,
                "dynamic_variable": "value",
                "allowed_values": {
                  "dynamic_variable": "value"
                },
                "allowed_values_dynamic_variable": "value",
                "constant_value": "value",
                "is_omitted": true
              }
            },
            "required": [
              "value"
            ]
          },
          "request_body_schema": {
            "description": "value",
            "dynamic_variable": "value",
            "constant_value": {
              "key": {}
            },
            "is_omitted": true,
            "type": "object",
            "required": [
              "value"
            ],
            "properties": {
              "key": {
                "type": "boolean",
                "description": "value",
                "enum": [
                  "value"
                ],
                "is_system_provided": true,
                "dynamic_variable": "value",
                "allowed_values": {
                  "dynamic_variable": "value"
                },
                "allowed_values_dynamic_variable": "value",
                "constant_value": "value",
                "is_omitted": true
              }
            },
            "required_constraints": {
              "any_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ],
              "all_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ]
            }
          },
          "response_body_schema": {
            "description": "value",
            "dynamic_variable": "value",
            "constant_value": {
              "key": {}
            },
            "is_omitted": true,
            "type": "object",
            "required": [
              "value"
            ],
            "properties": {
              "key": {
                "type": "boolean",
                "description": "value",
                "enum": [
                  "value"
                ],
                "is_system_provided": true,
                "dynamic_variable": "value",
                "allowed_values": {
                  "dynamic_variable": "value"
                },
                "allowed_values_dynamic_variable": "value",
                "constant_value": "value",
                "is_omitted": true
              }
            },
            "required_constraints": {
              "any_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ],
              "all_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ]
            }
          },
          "response_filter": {
            "mode": "all",
            "filters": [
              "value"
            ],
            "content_type": "application/json"
          },
          "content_type": "application/json",
          "auth_resolved_params": [
            "value"
          ],
          "auth_connection": {
            "auth_connection_id": "value"
          }
        },
        "follow_redirects": true,
        "follow_redirects_allowed_domains": [
          "value"
        ]
      },
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "usage_stats": {
        "total_calls": 123,
        "avg_latency_secs": 9.99
      },
      "response_mocks": [
        {
          "parameter_conditions": [
            {
              "eval": {
                "type": "llm",
                "description": "value"
              },
              "path": "value"
            }
          ],
          "mock_result": "value",
          "is_error": true
        }
      ]
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	searchValue := string("value")
	pageSizeValue := int(100)
	showOnlyOwnedDocumentsValue := bool(true)
	createdByUserIdValue := string("value")
	usedByAgentIdValue := string("value")
	cursorValue := string("value")
	request := &agents.ListToolsRequest{
		Search:                 &searchValue,
		PageSize:               &pageSizeValue,
		ShowOnlyOwnedDocuments: &showOnlyOwnedDocumentsValue,
		CreatedByUserId:        &createdByUserIdValue,
		UsedByAgentId:          &usedByAgentIdValue,
		Cursor:                 &cursorValue,
	}
	result, err := client.List(context.Background(), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsToolsClient(gotResponse, []byte(`{
  "tools": [
    {
      "id": "value",
      "tool_config": {
        "type": "webhook",
        "name": "value",
        "description": "value",
        "response_timeout_secs": 123,
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "tool_error_handling_mode": "auto",
        "dynamic_variables": {
          "dynamic_variable_placeholders": {
            "key": {}
          }
        },
        "execution_mode": "immediate",
        "api_schema": {
          "request_headers": {
            "key": "value"
          },
          "url": "value",
          "method": "GET",
          "path_params_schema": {
            "key": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            }
          },
          "query_params_schema": {
            "properties": {
              "key": {
                "type": "boolean",
                "description": "value",
                "enum": [
                  "value"
                ],
                "is_system_provided": true,
                "dynamic_variable": "value",
                "allowed_values": {
                  "dynamic_variable": "value"
                },
                "allowed_values_dynamic_variable": "value",
                "constant_value": "value",
                "is_omitted": true
              }
            },
            "required": [
              "value"
            ]
          },
          "request_body_schema": {
            "description": "value",
            "dynamic_variable": "value",
            "constant_value": {
              "key": {}
            },
            "is_omitted": true,
            "type": "object",
            "required": [
              "value"
            ],
            "properties": {
              "key": {
                "type": "boolean",
                "description": "value",
                "enum": [
                  "value"
                ],
                "is_system_provided": true,
                "dynamic_variable": "value",
                "allowed_values": {
                  "dynamic_variable": "value"
                },
                "allowed_values_dynamic_variable": "value",
                "constant_value": "value",
                "is_omitted": true
              }
            },
            "required_constraints": {
              "any_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ],
              "all_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ]
            }
          },
          "response_body_schema": {
            "description": "value",
            "dynamic_variable": "value",
            "constant_value": {
              "key": {}
            },
            "is_omitted": true,
            "type": "object",
            "required": [
              "value"
            ],
            "properties": {
              "key": {
                "type": "boolean",
                "description": "value",
                "enum": [
                  "value"
                ],
                "is_system_provided": true,
                "dynamic_variable": "value",
                "allowed_values": {
                  "dynamic_variable": "value"
                },
                "allowed_values_dynamic_variable": "value",
                "constant_value": "value",
                "is_omitted": true
              }
            },
            "required_constraints": {
              "any_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ],
              "all_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ]
            }
          },
          "response_filter": {
            "mode": "all",
            "filters": [
              "value"
            ],
            "content_type": "application/json"
          },
          "content_type": "application/json",
          "auth_resolved_params": [
            "value"
          ],
          "auth_connection": {
            "auth_connection_id": "value"
          }
        },
        "follow_redirects": true,
        "follow_redirects_allowed_domains": [
          "value"
        ]
      },
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "usage_stats": {
        "total_calls": 123,
        "avg_latency_secs": 9.99
      },
      "response_mocks": [
        {
          "parameter_conditions": [
            {
              "eval": {
                "type": "llm",
                "description": "value"
              },
              "path": "value"
            }
          ],
          "mock_result": "value",
          "is_error": true
        }
      ]
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "tools": [
    {
      "id": "value",
      "tool_config": {
        "type": "webhook",
        "name": "value",
        "description": "value",
        "response_timeout_secs": 123,
        "disable_interruptions": true,
        "interruption_mode": "allow",
        "force_pre_tool_speech": true,
        "pre_tool_speech": "auto",
        "assignments": [
          {
            "source": "response",
            "dynamic_variable": "value",
            "value_path": "value",
            "sanitize": true,
            "preserve_native_type": true
          }
        ],
        "tool_call_sound": "typing",
        "tool_call_sound_behavior": "auto",
        "tool_error_handling_mode": "auto",
        "dynamic_variables": {
          "dynamic_variable_placeholders": {
            "key": {}
          }
        },
        "execution_mode": "immediate",
        "api_schema": {
          "request_headers": {
            "key": "value"
          },
          "url": "value",
          "method": "GET",
          "path_params_schema": {
            "key": {
              "type": "boolean",
              "description": "value",
              "enum": [
                "value"
              ],
              "is_system_provided": true,
              "dynamic_variable": "value",
              "allowed_values": {
                "dynamic_variable": "value"
              },
              "allowed_values_dynamic_variable": "value",
              "constant_value": "value",
              "is_omitted": true
            }
          },
          "query_params_schema": {
            "properties": {
              "key": {
                "type": "boolean",
                "description": "value",
                "enum": [
                  "value"
                ],
                "is_system_provided": true,
                "dynamic_variable": "value",
                "allowed_values": {
                  "dynamic_variable": "value"
                },
                "allowed_values_dynamic_variable": "value",
                "constant_value": "value",
                "is_omitted": true
              }
            },
            "required": [
              "value"
            ]
          },
          "request_body_schema": {
            "description": "value",
            "dynamic_variable": "value",
            "constant_value": {
              "key": {}
            },
            "is_omitted": true,
            "type": "object",
            "required": [
              "value"
            ],
            "properties": {
              "key": {
                "type": "boolean",
                "description": "value",
                "enum": [
                  "value"
                ],
                "is_system_provided": true,
                "dynamic_variable": "value",
                "allowed_values": {
                  "dynamic_variable": "value"
                },
                "allowed_values_dynamic_variable": "value",
                "constant_value": "value",
                "is_omitted": true
              }
            },
            "required_constraints": {
              "any_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ],
              "all_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ]
            }
          },
          "response_body_schema": {
            "description": "value",
            "dynamic_variable": "value",
            "constant_value": {
              "key": {}
            },
            "is_omitted": true,
            "type": "object",
            "required": [
              "value"
            ],
            "properties": {
              "key": {
                "type": "boolean",
                "description": "value",
                "enum": [
                  "value"
                ],
                "is_system_provided": true,
                "dynamic_variable": "value",
                "allowed_values": {
                  "dynamic_variable": "value"
                },
                "allowed_values_dynamic_variable": "value",
                "constant_value": "value",
                "is_omitted": true
              }
            },
            "required_constraints": {
              "any_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ],
              "all_of": [
                {
                  "required": [
                    "value"
                  ]
                }
              ]
            }
          },
          "response_filter": {
            "mode": "all",
            "filters": [
              "value"
            ],
            "content_type": "application/json"
          },
          "content_type": "application/json",
          "auth_resolved_params": [
            "value"
          ],
          "auth_connection": {
            "auth_connection_id": "value"
          }
        },
        "follow_redirects": true,
        "follow_redirects_allowed_domains": [
          "value"
        ]
      },
      "access_info": {
        "is_creator": true,
        "creator_name": "value",
        "creator_email": "value",
        "role": "admin",
        "anonymous_access_level_override": "admin",
        "access_source": "creator"
      },
      "usage_stats": {
        "total_calls": 123,
        "avg_latency_secs": 9.99
      },
      "response_mocks": [
        {
          "parameter_conditions": [
            {
              "eval": {
                "type": "llm",
                "description": "value"
              },
              "path": "value"
            }
          ],
          "mock_result": "value",
          "is_error": true
        }
      ]
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsToolsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/tools" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tools")
	}
	if got := req.Query.Get("search"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "search", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
	if got := req.Query.Get("show_only_owned_documents"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "show_only_owned_documents", got, "true")
	}
	if got := req.Query.Get("created_by_user_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "created_by_user_id", got, "value")
	}
	if got := req.Query.Get("used_by_agent_id"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "used_by_agent_id", got, "value")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsToolsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	searchValue := string("value")
	pageSizeValue := int(100)
	showOnlyOwnedDocumentsValue := bool(true)
	createdByUserIdValue := string("value")
	usedByAgentIdValue := string("value")
	cursorValue := string("value")
	request := &agents.ListToolsRequest{
		Search:                 &searchValue,
		PageSize:               &pageSizeValue,
		ShowOnlyOwnedDocuments: &showOnlyOwnedDocumentsValue,
		CreatedByUserId:        &createdByUserIdValue,
		UsedByAgentId:          &usedByAgentIdValue,
		Cursor:                 &cursorValue,
	}
	_, err := client.List(context.Background(), request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestCreateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsToolsClient(t, 200, `{
  "id": "value",
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "usage_stats": {
    "total_calls": 123,
    "avg_latency_secs": 9.99
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.ToolRequestModel
	if err := json.Unmarshal([]byte(`{
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "property_kind": "array",
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "property_kind": "array",
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	result, err := client.Create(context.Background(), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsToolsClient(gotResponse, []byte(`{
  "id": "value",
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "usage_stats": {
    "total_calls": 123,
    "avg_latency_secs": 9.99
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "usage_stats": {
    "total_calls": 123,
    "avg_latency_secs": 9.99
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`)
	}

	req := receiveWireAgentsToolsClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/convai/tools" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tools")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsToolsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsToolsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request types.ToolRequestModel
	if err := json.Unmarshal([]byte(`{
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "property_kind": "array",
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "property_kind": "array",
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	_, err := client.Create(context.Background(), &request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestGetWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsToolsClient(t, 200, `{
  "id": "value",
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "usage_stats": {
    "total_calls": 123,
    "avg_latency_secs": 9.99
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	environmentValue := string("value")
	request := &agents.GetToolsRequest{
		Environment: &environmentValue,
	}
	result, err := client.Get(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsToolsClient(gotResponse, []byte(`{
  "id": "value",
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "usage_stats": {
    "total_calls": 123,
    "avg_latency_secs": 9.99
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "usage_stats": {
    "total_calls": 123,
    "avg_latency_secs": 9.99
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`)
	}

	req := receiveWireAgentsToolsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/tools/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tools/value")
	}
	if got := req.Query.Get("environment"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "environment", got, "value")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsToolsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	environmentValue := string("value")
	request := &agents.GetToolsRequest{
		Environment: &environmentValue,
	}
	_, err := client.Get(context.Background(), string("value"), request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestUpdateWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsToolsClient(t, 200, `{
  "id": "value",
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "usage_stats": {
    "total_calls": 123,
    "avg_latency_secs": 9.99
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request types.ToolRequestModel
	if err := json.Unmarshal([]byte(`{
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "property_kind": "array",
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "property_kind": "array",
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsToolsClient(gotResponse, []byte(`{
  "id": "value",
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "usage_stats": {
    "total_calls": 123,
    "avg_latency_secs": 9.99
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "id": "value",
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "usage_stats": {
    "total_calls": 123,
    "avg_latency_secs": 9.99
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`)
	}

	req := receiveWireAgentsToolsClient(t, recorded)
	if req.Method != "PATCH" {
		t.Fatalf("method: got %q want %q", req.Method, "PATCH")
	}
	if req.Path != "/v1/convai/tools/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tools/value")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsAgentsToolsClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsToolsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	var request types.ToolRequestModel
	if err := json.Unmarshal([]byte(`{
  "tool_config": {
    "type": "webhook",
    "name": "value",
    "description": "value",
    "response_timeout_secs": 123,
    "disable_interruptions": true,
    "interruption_mode": "allow",
    "force_pre_tool_speech": true,
    "pre_tool_speech": "auto",
    "assignments": [
      {
        "source": "response",
        "dynamic_variable": "value",
        "value_path": "value",
        "sanitize": true,
        "preserve_native_type": true
      }
    ],
    "tool_call_sound": "typing",
    "tool_call_sound_behavior": "auto",
    "tool_error_handling_mode": "auto",
    "dynamic_variables": {
      "dynamic_variable_placeholders": {
        "key": {}
      }
    },
    "execution_mode": "immediate",
    "api_schema": {
      "request_headers": {
        "key": "value"
      },
      "url": "value",
      "method": "GET",
      "path_params_schema": {
        "key": {
          "type": "boolean",
          "description": "value",
          "enum": [
            "value"
          ],
          "is_system_provided": true,
          "dynamic_variable": "value",
          "allowed_values": {
            "dynamic_variable": "value"
          },
          "allowed_values_dynamic_variable": "value",
          "constant_value": "value",
          "is_omitted": true
        }
      },
      "query_params_schema": {
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required": [
          "value"
        ]
      },
      "request_body_schema": {
        "property_kind": "array",
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_body_schema": {
        "property_kind": "array",
        "description": "value",
        "dynamic_variable": "value",
        "constant_value": {
          "key": {}
        },
        "is_omitted": true,
        "type": "object",
        "required": [
          "value"
        ],
        "properties": {
          "key": {
            "type": "boolean",
            "description": "value",
            "enum": [
              "value"
            ],
            "is_system_provided": true,
            "dynamic_variable": "value",
            "allowed_values": {
              "dynamic_variable": "value"
            },
            "allowed_values_dynamic_variable": "value",
            "constant_value": "value",
            "is_omitted": true
          }
        },
        "required_constraints": {
          "any_of": [
            {
              "required": [
                "value"
              ]
            }
          ],
          "all_of": [
            {
              "required": [
                "value"
              ]
            }
          ]
        }
      },
      "response_filter": {
        "mode": "all",
        "filters": [
          "value"
        ],
        "content_type": "application/json"
      },
      "content_type": "application/json",
      "auth_resolved_params": [
        "value"
      ],
      "auth_connection": {
        "auth_connection_id": "value"
      }
    },
    "follow_redirects": true,
    "follow_redirects_allowed_domains": [
      "value"
    ]
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`), &request); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	_, err := client.Update(context.Background(), string("value"), &request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestDeleteWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsToolsClient(t, 200, `{}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	forceValue := bool(true)
	request := &agents.DeleteToolsRequest{
		Force: &forceValue,
	}
	result, err := client.Delete(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsToolsClient(gotResponse, []byte(`{}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{}`)
	}

	req := receiveWireAgentsToolsClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/convai/tools/value" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tools/value")
	}
	if got := req.Query.Get("force"); got != "true" {
		t.Fatalf("query %s: got %q want %q", "force", got, "true")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsToolsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	forceValue := bool(true)
	request := &agents.DeleteToolsRequest{
		Force: &forceValue,
	}
	_, err := client.Delete(context.Background(), string("value"), request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

func TestGetDependentAgentsWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireAgentsToolsClient(t, 200, `{
  "agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "branches": [
    {
      "agent_id": "value",
      "agent_name": "value",
      "branch_id": "value",
      "branch_name": "value",
      "is_main": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	cursorValue := string("value")
	pageSizeValue := int(100)
	request := &agents.GetDependentAgentsToolsRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
	}
	result, err := client.GetDependentAgents(context.Background(), string("value"), request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualAgentsToolsClient(gotResponse, []byte(`{
  "agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "branches": [
    {
      "agent_id": "value",
      "agent_name": "value",
      "branch_id": "value",
      "branch_name": "value",
      "is_main": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "agents": [
    {
      "referenced_resource_ids": [
        "value"
      ],
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "branches": [
    {
      "agent_id": "value",
      "agent_name": "value",
      "branch_id": "value",
      "branch_name": "value",
      "is_main": true
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`)
	}

	req := receiveWireAgentsToolsClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/convai/tools/value/dependent-agents" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/convai/tools/value/dependent-agents")
	}
	if got := req.Query.Get("cursor"); got != "value" {
		t.Fatalf("query %s: got %q want %q", "cursor", got, "value")
	}
	if got := req.Query.Get("page_size"); got != "100" {
		t.Fatalf("query %s: got %q want %q", "page_size", got, "100")
	}
}

func TestGetDependentAgentsWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireAgentsToolsClient(t, 422, `{
  "detail": [
    {
      "loc": [
        "value"
      ],
      "msg": "value",
      "type": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
		option.WithMaxAttempts(1),
	))

	cursorValue := string("value")
	pageSizeValue := int(100)
	request := &agents.GetDependentAgentsToolsRequest{
		Cursor:   &cursorValue,
		PageSize: &pageSizeValue,
	}
	_, err := client.GetDependentAgents(context.Background(), string("value"), request)
	if err == nil {
		t.Fatal("expected an error for a 422 response")
	}
	var apiErr *core.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expected a *core.APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != 422 {
		t.Fatalf("status: got %d want %d", apiErr.StatusCode, 422)
	}
}

type wireRequestAgentsToolsClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireAgentsToolsClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireAgentsToolsClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestAgentsToolsClient) {
	t.Helper()
	recorded := make(chan wireRequestAgentsToolsClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestAgentsToolsClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
		default:
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if responseBody != "" {
			_, _ = w.Write([]byte(responseBody))
		}
	}))
	t.Cleanup(server.Close)
	return server, recorded
}

func receiveWireAgentsToolsClient(t *testing.T, recorded <-chan wireRequestAgentsToolsClient) wireRequestAgentsToolsClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestAgentsToolsClient{}
	}
}

// wireJSONContainsAgentsToolsClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsAgentsToolsClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualAgentsToolsClient(got, want)
	}
	for key, wantValue := range wantObject {
		if key == "stream" {
			continue
		}
		if gotValue, ok := gotObject[key]; !ok || !reflect.DeepEqual(gotValue, wantValue) {
			return false
		}
	}
	return true
}

// setWireValueAgentsToolsClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueAgentsToolsClient(t *testing.T, target any, value any) {
	t.Helper()
	field := reflect.ValueOf(target).Elem()
	if field.Kind() == reflect.Ptr {
		field.Set(reflect.New(field.Type().Elem()))
		field = field.Elem()
	}
	switch typed := value.(type) {
	case string:
		field.SetString(typed)
	case int:
		field.SetInt(int64(typed))
	case float64:
		field.SetFloat(typed)
	default:
		t.Fatalf("unsupported value %T", value)
	}
}

func wireJSONBytesEqualAgentsToolsClient(got, want []byte) bool {
	var gotValue any
	var wantValue any
	if err := json.Unmarshal(got, &gotValue); err != nil {
		return false
	}
	if err := json.Unmarshal(want, &wantValue); err != nil {
		return false
	}
	return reflect.DeepEqual(gotValue, wantValue)
}
