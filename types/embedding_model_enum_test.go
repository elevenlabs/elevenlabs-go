package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestEmbeddingModelEnumVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant EmbeddingModelEnum
		wire    string
	}
	cases := []enumCase{
		{variant: EmbeddingModelEnumE5Mistral7BInstruct, wire: "\"e5_mistral_7b_instruct\""},
		{variant: EmbeddingModelEnumMultilingualE5LargeInstruct, wire: "\"multilingual_e5_large_instruct\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded EmbeddingModelEnum
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

func TestEmbeddingModelEnumVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant EmbeddingModelEnum
		wire    string
	}
	cases := []enumCase{
		{variant: EmbeddingModelEnumE5Mistral7BInstruct, wire: "\"e5_mistral_7b_instruct\""},
		{variant: EmbeddingModelEnumMultilingualE5LargeInstruct, wire: "\"multilingual_e5_large_instruct\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded EmbeddingModelEnum
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

func TestEmbeddingModelEnumVariantsRoundTrip3(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant EmbeddingModelEnum
		wire    string
	}
	cases := []enumCase{
		{variant: EmbeddingModelEnumE5Mistral7BInstruct, wire: "\"e5_mistral_7b_instruct\""},
		{variant: EmbeddingModelEnumMultilingualE5LargeInstruct, wire: "\"multilingual_e5_large_instruct\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded EmbeddingModelEnum
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

func TestEmbeddingModelEnumVariantsRoundTrip4(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant EmbeddingModelEnum
		wire    string
	}
	cases := []enumCase{
		{variant: EmbeddingModelEnumE5Mistral7BInstruct, wire: "\"e5_mistral_7b_instruct\""},
		{variant: EmbeddingModelEnumMultilingualE5LargeInstruct, wire: "\"multilingual_e5_large_instruct\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded EmbeddingModelEnum
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

func TestEmbeddingModelEnumVariantsRoundTrip5(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant EmbeddingModelEnum
		wire    string
	}
	cases := []enumCase{
		{variant: EmbeddingModelEnumE5Mistral7BInstruct, wire: "\"e5_mistral_7b_instruct\""},
		{variant: EmbeddingModelEnumMultilingualE5LargeInstruct, wire: "\"multilingual_e5_large_instruct\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded EmbeddingModelEnum
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

func TestEmbeddingModelEnumVariantsRoundTrip6(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant EmbeddingModelEnum
		wire    string
	}
	cases := []enumCase{
		{variant: EmbeddingModelEnumE5Mistral7BInstruct, wire: "\"e5_mistral_7b_instruct\""},
		{variant: EmbeddingModelEnumMultilingualE5LargeInstruct, wire: "\"multilingual_e5_large_instruct\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded EmbeddingModelEnum
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

func TestEmbeddingModelEnumVariantsRoundTrip7(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant EmbeddingModelEnum
		wire    string
	}
	cases := []enumCase{
		{variant: EmbeddingModelEnumE5Mistral7BInstruct, wire: "\"e5_mistral_7b_instruct\""},
		{variant: EmbeddingModelEnumMultilingualE5LargeInstruct, wire: "\"multilingual_e5_large_instruct\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded EmbeddingModelEnum
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
