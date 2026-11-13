package agents

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestListConversationsRequestExcludeStatusesItemVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ListConversationsRequestExcludeStatusesItem
		wire    string
	}
	cases := []enumCase{
		{variant: ListConversationsRequestExcludeStatusesItemInitiated, wire: "\"initiated\""},
		{variant: ListConversationsRequestExcludeStatusesItemInProgress, wire: "\"in-progress\""},
		{variant: ListConversationsRequestExcludeStatusesItemProcessing, wire: "\"processing\""},
		{variant: ListConversationsRequestExcludeStatusesItemDone, wire: "\"done\""},
		{variant: ListConversationsRequestExcludeStatusesItemFailed, wire: "\"failed\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ListConversationsRequestExcludeStatusesItem
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

func TestGetConversationsRequestFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant GetConversationsRequestFormat
		wire    string
	}
	cases := []enumCase{
		{variant: GetConversationsRequestFormatJson, wire: "\"json\""},
		{variant: GetConversationsRequestFormatOpentelemetry, wire: "\"opentelemetry\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded GetConversationsRequestFormat
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
