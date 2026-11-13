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

	core "github.com/elevenlabs/elevenlabs-go/core"
	option "github.com/elevenlabs/elevenlabs-go/option"
	projects "github.com/elevenlabs/elevenlabs-go/studio/projects"
)

func TestListWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireStudioProjectsChaptersClient(t, 200, `{
  "chapters": [
    {
      "chapter_id": "value",
      "name": "value",
      "last_conversion_date_unix": 123,
      "conversion_progress": 9.99,
      "can_be_downloaded": true,
      "state": "default",
      "has_video": true,
      "has_visual_content": true,
      "voice_ids": [
        "value"
      ],
      "statistics": {
        "characters_unconverted": 123,
        "characters_converted": 123,
        "paragraphs_converted": 123,
        "paragraphs_unconverted": 123,
        "credits_needed_to_convert": 123,
        "voice_statistics": [
          {
            "project_voice_ref_id": "value",
            "characters_unconverted": 123,
            "characters_converted": 123,
            "credits_needed_to_convert": 123,
            "voice_id": "value"
          }
        ]
      },
      "last_conversion_error": "value"
    }
  ]
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.List(context.Background(), string("value"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsChaptersClient(gotResponse, []byte(`{
  "chapters": [
    {
      "chapter_id": "value",
      "name": "value",
      "last_conversion_date_unix": 123,
      "conversion_progress": 9.99,
      "can_be_downloaded": true,
      "state": "default",
      "has_video": true,
      "has_visual_content": true,
      "voice_ids": [
        "value"
      ],
      "statistics": {
        "characters_unconverted": 123,
        "characters_converted": 123,
        "paragraphs_converted": 123,
        "paragraphs_unconverted": 123,
        "credits_needed_to_convert": 123,
        "voice_statistics": [
          {
            "project_voice_ref_id": "value",
            "characters_unconverted": 123,
            "characters_converted": 123,
            "credits_needed_to_convert": 123,
            "voice_id": "value"
          }
        ]
      },
      "last_conversion_error": "value"
    }
  ]
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "chapters": [
    {
      "chapter_id": "value",
      "name": "value",
      "last_conversion_date_unix": 123,
      "conversion_progress": 9.99,
      "can_be_downloaded": true,
      "state": "default",
      "has_video": true,
      "has_visual_content": true,
      "voice_ids": [
        "value"
      ],
      "statistics": {
        "characters_unconverted": 123,
        "characters_converted": 123,
        "paragraphs_converted": 123,
        "paragraphs_unconverted": 123,
        "credits_needed_to_convert": 123,
        "voice_statistics": [
          {
            "project_voice_ref_id": "value",
            "characters_unconverted": 123,
            "characters_converted": 123,
            "credits_needed_to_convert": 123,
            "voice_id": "value"
          }
        ]
      },
      "last_conversion_error": "value"
    }
  ]
}`)
	}

	req := receiveWireStudioProjectsChaptersClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/studio/projects/value/chapters" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value/chapters")
	}
}

func TestListWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsChaptersClient(t, 422, `{
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

	_, err := client.List(context.Background(), string("value"))
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

	server, recorded := serveWireStudioProjectsChaptersClient(t, 200, `{
  "chapter": {
    "chapter_id": "value",
    "name": "value",
    "last_conversion_date_unix": 123,
    "conversion_progress": 9.99,
    "can_be_downloaded": true,
    "state": "default",
    "has_video": true,
    "has_visual_content": true,
    "voice_ids": [
      "value"
    ],
    "statistics": {
      "characters_unconverted": 123,
      "characters_converted": 123,
      "paragraphs_converted": 123,
      "paragraphs_unconverted": 123,
      "credits_needed_to_convert": 123,
      "voice_statistics": [
        {
          "project_voice_ref_id": "value",
          "characters_unconverted": 123,
          "characters_converted": 123,
          "credits_needed_to_convert": 123,
          "voice_id": "value"
        }
      ]
    },
    "last_conversion_error": "value",
    "content": {
      "blocks": [
        {
          "block_id": "value",
          "nodes": [
            {
              "type": "tts_node",
              "project_voice_ref_id": "value",
              "text": "value",
              "voice_id": "value"
            }
          ]
        }
      ]
    }
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request projects.BodyCreateChapterV1StudioProjectsProjectIdChaptersPost
	if err := json.Unmarshal([]byte(`{
  "project_id": "value",
  "name": "value",
  "from_url": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Create(context.Background(), string("value"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsChaptersClient(gotResponse, []byte(`{
  "chapter": {
    "chapter_id": "value",
    "name": "value",
    "last_conversion_date_unix": 123,
    "conversion_progress": 9.99,
    "can_be_downloaded": true,
    "state": "default",
    "has_video": true,
    "has_visual_content": true,
    "voice_ids": [
      "value"
    ],
    "statistics": {
      "characters_unconverted": 123,
      "characters_converted": 123,
      "paragraphs_converted": 123,
      "paragraphs_unconverted": 123,
      "credits_needed_to_convert": 123,
      "voice_statistics": [
        {
          "project_voice_ref_id": "value",
          "characters_unconverted": 123,
          "characters_converted": 123,
          "credits_needed_to_convert": 123,
          "voice_id": "value"
        }
      ]
    },
    "last_conversion_error": "value",
    "content": {
      "blocks": [
        {
          "block_id": "value",
          "nodes": [
            {
              "type": "tts_node",
              "project_voice_ref_id": "value",
              "text": "value",
              "voice_id": "value"
            }
          ]
        }
      ]
    }
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "chapter": {
    "chapter_id": "value",
    "name": "value",
    "last_conversion_date_unix": 123,
    "conversion_progress": 9.99,
    "can_be_downloaded": true,
    "state": "default",
    "has_video": true,
    "has_visual_content": true,
    "voice_ids": [
      "value"
    ],
    "statistics": {
      "characters_unconverted": 123,
      "characters_converted": 123,
      "paragraphs_converted": 123,
      "paragraphs_unconverted": 123,
      "credits_needed_to_convert": 123,
      "voice_statistics": [
        {
          "project_voice_ref_id": "value",
          "characters_unconverted": 123,
          "characters_converted": 123,
          "credits_needed_to_convert": 123,
          "voice_id": "value"
        }
      ]
    },
    "last_conversion_error": "value",
    "content": {
      "blocks": [
        {
          "block_id": "value",
          "nodes": [
            {
              "type": "tts_node",
              "project_voice_ref_id": "value",
              "text": "value",
              "voice_id": "value"
            }
          ]
        }
      ]
    }
  }
}`)
	}

	req := receiveWireStudioProjectsChaptersClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/studio/projects/value/chapters" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value/chapters")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsStudioProjectsChaptersClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestCreateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsChaptersClient(t, 422, `{
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

	var request projects.BodyCreateChapterV1StudioProjectsProjectIdChaptersPost
	if err := json.Unmarshal([]byte(`{
  "project_id": "value",
  "name": "value",
  "from_url": "value"
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Create(context.Background(), string("value"), &request)
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

	server, recorded := serveWireStudioProjectsChaptersClient(t, 200, `{
  "chapter_id": "value",
  "name": "value",
  "last_conversion_date_unix": 123,
  "conversion_progress": 9.99,
  "can_be_downloaded": true,
  "state": "default",
  "has_video": true,
  "has_visual_content": true,
  "voice_ids": [
    "value"
  ],
  "statistics": {
    "characters_unconverted": 123,
    "characters_converted": 123,
    "paragraphs_converted": 123,
    "paragraphs_unconverted": 123,
    "credits_needed_to_convert": 123,
    "voice_statistics": [
      {
        "project_voice_ref_id": "value",
        "characters_unconverted": 123,
        "characters_converted": 123,
        "credits_needed_to_convert": 123,
        "voice_id": "value"
      }
    ]
  },
  "last_conversion_error": "value",
  "content": {
    "blocks": [
      {
        "block_id": "value",
        "nodes": [
          {
            "type": "tts_node",
            "project_voice_ref_id": "value",
            "text": "value",
            "voice_id": "value"
          }
        ]
      }
    ]
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Get(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsChaptersClient(gotResponse, []byte(`{
  "chapter_id": "value",
  "name": "value",
  "last_conversion_date_unix": 123,
  "conversion_progress": 9.99,
  "can_be_downloaded": true,
  "state": "default",
  "has_video": true,
  "has_visual_content": true,
  "voice_ids": [
    "value"
  ],
  "statistics": {
    "characters_unconverted": 123,
    "characters_converted": 123,
    "paragraphs_converted": 123,
    "paragraphs_unconverted": 123,
    "credits_needed_to_convert": 123,
    "voice_statistics": [
      {
        "project_voice_ref_id": "value",
        "characters_unconverted": 123,
        "characters_converted": 123,
        "credits_needed_to_convert": 123,
        "voice_id": "value"
      }
    ]
  },
  "last_conversion_error": "value",
  "content": {
    "blocks": [
      {
        "block_id": "value",
        "nodes": [
          {
            "type": "tts_node",
            "project_voice_ref_id": "value",
            "text": "value",
            "voice_id": "value"
          }
        ]
      }
    ]
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "chapter_id": "value",
  "name": "value",
  "last_conversion_date_unix": 123,
  "conversion_progress": 9.99,
  "can_be_downloaded": true,
  "state": "default",
  "has_video": true,
  "has_visual_content": true,
  "voice_ids": [
    "value"
  ],
  "statistics": {
    "characters_unconverted": 123,
    "characters_converted": 123,
    "paragraphs_converted": 123,
    "paragraphs_unconverted": 123,
    "credits_needed_to_convert": 123,
    "voice_statistics": [
      {
        "project_voice_ref_id": "value",
        "characters_unconverted": 123,
        "characters_converted": 123,
        "credits_needed_to_convert": 123,
        "voice_id": "value"
      }
    ]
  },
  "last_conversion_error": "value",
  "content": {
    "blocks": [
      {
        "block_id": "value",
        "nodes": [
          {
            "type": "tts_node",
            "project_voice_ref_id": "value",
            "text": "value",
            "voice_id": "value"
          }
        ]
      }
    ]
  }
}`)
	}

	req := receiveWireStudioProjectsChaptersClient(t, recorded)
	if req.Method != "GET" {
		t.Fatalf("method: got %q want %q", req.Method, "GET")
	}
	if req.Path != "/v1/studio/projects/value/chapters/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value/chapters/value2")
	}
}

func TestGetWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsChaptersClient(t, 422, `{
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

	_, err := client.Get(context.Background(), string("value"), string("value2"))
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

	server, recorded := serveWireStudioProjectsChaptersClient(t, 200, `{
  "chapter": {
    "chapter_id": "value",
    "name": "value",
    "last_conversion_date_unix": 123,
    "conversion_progress": 9.99,
    "can_be_downloaded": true,
    "state": "default",
    "has_video": true,
    "has_visual_content": true,
    "voice_ids": [
      "value"
    ],
    "statistics": {
      "characters_unconverted": 123,
      "characters_converted": 123,
      "paragraphs_converted": 123,
      "paragraphs_unconverted": 123,
      "credits_needed_to_convert": 123,
      "voice_statistics": [
        {
          "project_voice_ref_id": "value",
          "characters_unconverted": 123,
          "characters_converted": 123,
          "credits_needed_to_convert": 123,
          "voice_id": "value"
        }
      ]
    },
    "last_conversion_error": "value",
    "content": {
      "blocks": [
        {
          "block_id": "value",
          "nodes": [
            {
              "type": "tts_node",
              "project_voice_ref_id": "value",
              "text": "value",
              "voice_id": "value"
            }
          ]
        }
      ]
    }
  }
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	var request projects.BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost
	if err := json.Unmarshal([]byte(`{
  "chapter_id": "value",
  "project_id": "value",
  "name": "value",
  "content": {
    "blocks": [
      {
        "sub_type": "p",
        "nodes": [
          {
            "type": "tts_node",
            "text": "value",
            "voice_id": "value"
          }
        ],
        "block_id": "value"
      }
    ]
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	result, err := client.Update(context.Background(), string("value"), string("value2"), &request)
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsChaptersClient(gotResponse, []byte(`{
  "chapter": {
    "chapter_id": "value",
    "name": "value",
    "last_conversion_date_unix": 123,
    "conversion_progress": 9.99,
    "can_be_downloaded": true,
    "state": "default",
    "has_video": true,
    "has_visual_content": true,
    "voice_ids": [
      "value"
    ],
    "statistics": {
      "characters_unconverted": 123,
      "characters_converted": 123,
      "paragraphs_converted": 123,
      "paragraphs_unconverted": 123,
      "credits_needed_to_convert": 123,
      "voice_statistics": [
        {
          "project_voice_ref_id": "value",
          "characters_unconverted": 123,
          "characters_converted": 123,
          "credits_needed_to_convert": 123,
          "voice_id": "value"
        }
      ]
    },
    "last_conversion_error": "value",
    "content": {
      "blocks": [
        {
          "block_id": "value",
          "nodes": [
            {
              "type": "tts_node",
              "project_voice_ref_id": "value",
              "text": "value",
              "voice_id": "value"
            }
          ]
        }
      ]
    }
  }
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "chapter": {
    "chapter_id": "value",
    "name": "value",
    "last_conversion_date_unix": 123,
    "conversion_progress": 9.99,
    "can_be_downloaded": true,
    "state": "default",
    "has_video": true,
    "has_visual_content": true,
    "voice_ids": [
      "value"
    ],
    "statistics": {
      "characters_unconverted": 123,
      "characters_converted": 123,
      "paragraphs_converted": 123,
      "paragraphs_unconverted": 123,
      "credits_needed_to_convert": 123,
      "voice_statistics": [
        {
          "project_voice_ref_id": "value",
          "characters_unconverted": 123,
          "characters_converted": 123,
          "credits_needed_to_convert": 123,
          "voice_id": "value"
        }
      ]
    },
    "last_conversion_error": "value",
    "content": {
      "blocks": [
        {
          "block_id": "value",
          "nodes": [
            {
              "type": "tts_node",
              "project_voice_ref_id": "value",
              "text": "value",
              "voice_id": "value"
            }
          ]
        }
      ]
    }
  }
}`)
	}

	req := receiveWireStudioProjectsChaptersClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/studio/projects/value/chapters/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value/chapters/value2")
	}
	wantBody, err := json.Marshal(&request)
	if err != nil {
		t.Fatalf("marshal request body: %v", err)
	}
	if !wireJSONContainsStudioProjectsChaptersClient(req.Body, wantBody) {
		t.Fatalf("request body mismatch\n got: %s\nwant: %s", string(req.Body), string(wantBody))
	}
}

func TestUpdateWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsChaptersClient(t, 422, `{
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

	var request projects.BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost
	if err := json.Unmarshal([]byte(`{
  "chapter_id": "value",
  "project_id": "value",
  "name": "value",
  "content": {
    "blocks": [
      {
        "sub_type": "p",
        "nodes": [
          {
            "type": "tts_node",
            "text": "value",
            "voice_id": "value"
          }
        ],
        "block_id": "value"
      }
    ]
  }
}`), &request); err != nil {
		t.Fatalf("unmarshal request body: %v", err)
	}
	_, err := client.Update(context.Background(), string("value"), string("value2"), &request)
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

	server, recorded := serveWireStudioProjectsChaptersClient(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Delete(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsChaptersClient(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireStudioProjectsChaptersClient(t, recorded)
	if req.Method != "DELETE" {
		t.Fatalf("method: got %q want %q", req.Method, "DELETE")
	}
	if req.Path != "/v1/studio/projects/value/chapters/value2" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value/chapters/value2")
	}
}

func TestDeleteWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsChaptersClient(t, 422, `{
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

	_, err := client.Delete(context.Background(), string("value"), string("value2"))
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

func TestConvertWire(t *testing.T) {
	t.Parallel()

	server, recorded := serveWireStudioProjectsChaptersClient(t, 200, `{
  "status": "value"
}`)

	client := NewClient(core.NewRequestOptions(
		option.WithBaseURL(server.URL),
	))

	result, err := client.Convert(context.Background(), string("value"), string("value2"))
	if err != nil {
		t.Fatalf("call: %v", err)
	}
	gotResponse, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}
	if !wireJSONBytesEqualStudioProjectsChaptersClient(gotResponse, []byte(`{
  "status": "value"
}`)) {
		t.Fatalf("response mismatch\n got: %s\nwant: %s", string(gotResponse), `{
  "status": "value"
}`)
	}

	req := receiveWireStudioProjectsChaptersClient(t, recorded)
	if req.Method != "POST" {
		t.Fatalf("method: got %q want %q", req.Method, "POST")
	}
	if req.Path != "/v1/studio/projects/value/chapters/value2/convert" {
		t.Fatalf("path: got %q want %q", req.Path, "/v1/studio/projects/value/chapters/value2/convert")
	}
}

func TestConvertWireError422(t *testing.T) {
	t.Parallel()

	server, _ := serveWireStudioProjectsChaptersClient(t, 422, `{
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

	_, err := client.Convert(context.Background(), string("value"), string("value2"))
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

type wireRequestStudioProjectsChaptersClient struct {
	Method string
	Path   string
	Query  url.Values
	Header http.Header
	Body   []byte
}

// serveWireStudioProjectsChaptersClient records the request inside the handler, where the body is still readable,
// and answers with the given status and JSON body.
func serveWireStudioProjectsChaptersClient(t *testing.T, status int, responseBody string) (*httptest.Server, <-chan wireRequestStudioProjectsChaptersClient) {
	t.Helper()
	recorded := make(chan wireRequestStudioProjectsChaptersClient, 4)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		select {
		case recorded <- wireRequestStudioProjectsChaptersClient{Method: r.Method, Path: r.URL.Path, Query: r.URL.Query(), Header: r.Header.Clone(), Body: body}:
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

func receiveWireStudioProjectsChaptersClient(t *testing.T, recorded <-chan wireRequestStudioProjectsChaptersClient) wireRequestStudioProjectsChaptersClient {
	t.Helper()
	select {
	case req := <-recorded:
		return req
	default:
		t.Fatal("expected a request to reach the test server")
		return wireRequestStudioProjectsChaptersClient{}
	}
}

// wireJSONContainsStudioProjectsChaptersClient reports whether got holds every field of want with the same value;
// the client owns the stream flag, so it is not compared.
func wireJSONContainsStudioProjectsChaptersClient(got, want []byte) bool {
	var gotObject, wantObject map[string]any
	if json.Unmarshal(got, &gotObject) != nil || json.Unmarshal(want, &wantObject) != nil {
		return wireJSONBytesEqualStudioProjectsChaptersClient(got, want)
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

// setWireValueStudioProjectsChaptersClient sets a named string, int or float field, or the value a pointer field points to.
func setWireValueStudioProjectsChaptersClient(t *testing.T, target any, value any) {
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

func wireJSONBytesEqualStudioProjectsChaptersClient(got, want []byte) bool {
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
