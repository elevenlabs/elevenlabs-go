package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestTtsConversationalModelVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant TtsConversationalModel
		wire    string
	}
	cases := []enumCase{
		{variant: TtsConversationalModelElevenTurboV2, wire: "\"eleven_turbo_v2\""},
		{variant: TtsConversationalModelElevenTurboV25, wire: "\"eleven_turbo_v2_5\""},
		{variant: TtsConversationalModelElevenFlashV2, wire: "\"eleven_flash_v2\""},
		{variant: TtsConversationalModelElevenFlashV25, wire: "\"eleven_flash_v2_5\""},
		{variant: TtsConversationalModelElevenMultilingualV2, wire: "\"eleven_multilingual_v2\""},
		{variant: TtsConversationalModelElevenV3Conversational, wire: "\"eleven_v3_conversational\""},
		{variant: TtsConversationalModelElevenV4, wire: "\"eleven_v4\""},
		{variant: TtsConversationalModelElevenV4Turbo, wire: "\"eleven_v4_turbo\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded TtsConversationalModel
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

func TestTtsConversationalModelVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant TtsConversationalModel
		wire    string
	}
	cases := []enumCase{
		{variant: TtsConversationalModelElevenTurboV2, wire: "\"eleven_turbo_v2\""},
		{variant: TtsConversationalModelElevenTurboV25, wire: "\"eleven_turbo_v2_5\""},
		{variant: TtsConversationalModelElevenFlashV2, wire: "\"eleven_flash_v2\""},
		{variant: TtsConversationalModelElevenFlashV25, wire: "\"eleven_flash_v2_5\""},
		{variant: TtsConversationalModelElevenMultilingualV2, wire: "\"eleven_multilingual_v2\""},
		{variant: TtsConversationalModelElevenV3Conversational, wire: "\"eleven_v3_conversational\""},
		{variant: TtsConversationalModelElevenV4, wire: "\"eleven_v4\""},
		{variant: TtsConversationalModelElevenV4Turbo, wire: "\"eleven_v4_turbo\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded TtsConversationalModel
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

func TestTtsConversationalModelVariantsRoundTrip3(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant TtsConversationalModel
		wire    string
	}
	cases := []enumCase{
		{variant: TtsConversationalModelElevenTurboV2, wire: "\"eleven_turbo_v2\""},
		{variant: TtsConversationalModelElevenTurboV25, wire: "\"eleven_turbo_v2_5\""},
		{variant: TtsConversationalModelElevenFlashV2, wire: "\"eleven_flash_v2\""},
		{variant: TtsConversationalModelElevenFlashV25, wire: "\"eleven_flash_v2_5\""},
		{variant: TtsConversationalModelElevenMultilingualV2, wire: "\"eleven_multilingual_v2\""},
		{variant: TtsConversationalModelElevenV3Conversational, wire: "\"eleven_v3_conversational\""},
		{variant: TtsConversationalModelElevenV4, wire: "\"eleven_v4\""},
		{variant: TtsConversationalModelElevenV4Turbo, wire: "\"eleven_v4_turbo\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded TtsConversationalModel
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
