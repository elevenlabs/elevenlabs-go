package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestToolCallSoundTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ToolCallSoundType
		wire    string
	}
	cases := []enumCase{
		{variant: ToolCallSoundTypeTyping, wire: "\"typing\""},
		{variant: ToolCallSoundTypeElevator1, wire: "\"elevator1\""},
		{variant: ToolCallSoundTypeElevator2, wire: "\"elevator2\""},
		{variant: ToolCallSoundTypeElevator3, wire: "\"elevator3\""},
		{variant: ToolCallSoundTypeElevator4, wire: "\"elevator4\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ToolCallSoundType
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

func TestToolCallSoundTypeVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ToolCallSoundType
		wire    string
	}
	cases := []enumCase{
		{variant: ToolCallSoundTypeTyping, wire: "\"typing\""},
		{variant: ToolCallSoundTypeElevator1, wire: "\"elevator1\""},
		{variant: ToolCallSoundTypeElevator2, wire: "\"elevator2\""},
		{variant: ToolCallSoundTypeElevator3, wire: "\"elevator3\""},
		{variant: ToolCallSoundTypeElevator4, wire: "\"elevator4\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ToolCallSoundType
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

func TestToolCallSoundTypeVariantsRoundTrip3(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ToolCallSoundType
		wire    string
	}
	cases := []enumCase{
		{variant: ToolCallSoundTypeTyping, wire: "\"typing\""},
		{variant: ToolCallSoundTypeElevator1, wire: "\"elevator1\""},
		{variant: ToolCallSoundTypeElevator2, wire: "\"elevator2\""},
		{variant: ToolCallSoundTypeElevator3, wire: "\"elevator3\""},
		{variant: ToolCallSoundTypeElevator4, wire: "\"elevator4\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ToolCallSoundType
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
