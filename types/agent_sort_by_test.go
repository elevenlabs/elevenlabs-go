package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAgentSortByVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant AgentSortBy
		wire    string
	}
	cases := []enumCase{
		{variant: AgentSortByName, wire: "\"name\""},
		{variant: AgentSortByCreatedAt, wire: "\"created_at\""},
		{variant: AgentSortByUpdatedAt, wire: "\"updated_at\""},
		{variant: AgentSortByCallCount7D, wire: "\"call_count_7d\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded AgentSortBy
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

func TestAgentSortByVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant AgentSortBy
		wire    string
	}
	cases := []enumCase{
		{variant: AgentSortByName, wire: "\"name\""},
		{variant: AgentSortByCreatedAt, wire: "\"created_at\""},
		{variant: AgentSortByUpdatedAt, wire: "\"updated_at\""},
		{variant: AgentSortByCallCount7D, wire: "\"call_count_7d\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded AgentSortBy
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
