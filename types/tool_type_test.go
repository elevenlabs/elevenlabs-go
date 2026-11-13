package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestToolTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ToolType
		wire    string
	}
	cases := []enumCase{
		{variant: ToolTypeSystem, wire: "\"system\""},
		{variant: ToolTypeWebhook, wire: "\"webhook\""},
		{variant: ToolTypeClient, wire: "\"client\""},
		{variant: ToolTypeMcp, wire: "\"mcp\""},
		{variant: ToolTypeWorkflow, wire: "\"workflow\""},
		{variant: ToolTypeApiIntegrationWebhook, wire: "\"api_integration_webhook\""},
		{variant: ToolTypeSmb, wire: "\"smb\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ToolType
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}

func TestToolTypeVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ToolType
		wire    string
	}
	cases := []enumCase{
		{variant: ToolTypeSystem, wire: "\"system\""},
		{variant: ToolTypeWebhook, wire: "\"webhook\""},
		{variant: ToolTypeClient, wire: "\"client\""},
		{variant: ToolTypeMcp, wire: "\"mcp\""},
		{variant: ToolTypeWorkflow, wire: "\"workflow\""},
		{variant: ToolTypeApiIntegrationWebhook, wire: "\"api_integration_webhook\""},
		{variant: ToolTypeSmb, wire: "\"smb\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ToolType
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
