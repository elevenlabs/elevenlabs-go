package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBatchCallRecipientStatusVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BatchCallRecipientStatus
		wire    string
	}
	cases := []enumCase{
		{variant: BatchCallRecipientStatusPending, wire: "\"pending\""},
		{variant: BatchCallRecipientStatusDispatched, wire: "\"dispatched\""},
		{variant: BatchCallRecipientStatusInitiated, wire: "\"initiated\""},
		{variant: BatchCallRecipientStatusInProgress, wire: "\"in_progress\""},
		{variant: BatchCallRecipientStatusCompleted, wire: "\"completed\""},
		{variant: BatchCallRecipientStatusFailed, wire: "\"failed\""},
		{variant: BatchCallRecipientStatusCancelled, wire: "\"cancelled\""},
		{variant: BatchCallRecipientStatusVoicemail, wire: "\"voicemail\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BatchCallRecipientStatus
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

func TestBatchCallRecipientStatusVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BatchCallRecipientStatus
		wire    string
	}
	cases := []enumCase{
		{variant: BatchCallRecipientStatusPending, wire: "\"pending\""},
		{variant: BatchCallRecipientStatusDispatched, wire: "\"dispatched\""},
		{variant: BatchCallRecipientStatusInitiated, wire: "\"initiated\""},
		{variant: BatchCallRecipientStatusInProgress, wire: "\"in_progress\""},
		{variant: BatchCallRecipientStatusCompleted, wire: "\"completed\""},
		{variant: BatchCallRecipientStatusFailed, wire: "\"failed\""},
		{variant: BatchCallRecipientStatusCancelled, wire: "\"cancelled\""},
		{variant: BatchCallRecipientStatusVoicemail, wire: "\"voicemail\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BatchCallRecipientStatus
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

func TestBatchCallRecipientStatusVariantsRoundTrip3(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BatchCallRecipientStatus
		wire    string
	}
	cases := []enumCase{
		{variant: BatchCallRecipientStatusPending, wire: "\"pending\""},
		{variant: BatchCallRecipientStatusDispatched, wire: "\"dispatched\""},
		{variant: BatchCallRecipientStatusInitiated, wire: "\"initiated\""},
		{variant: BatchCallRecipientStatusInProgress, wire: "\"in_progress\""},
		{variant: BatchCallRecipientStatusCompleted, wire: "\"completed\""},
		{variant: BatchCallRecipientStatusFailed, wire: "\"failed\""},
		{variant: BatchCallRecipientStatusCancelled, wire: "\"cancelled\""},
		{variant: BatchCallRecipientStatusVoicemail, wire: "\"voicemail\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BatchCallRecipientStatus
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
