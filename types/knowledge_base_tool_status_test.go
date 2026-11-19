package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestKnowledgeBaseToolStatusVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant KnowledgeBaseToolStatus
		wire    string
	}
	cases := []enumCase{
		{variant: KnowledgeBaseToolStatusSuccess, wire: "\"success\""},
		{variant: KnowledgeBaseToolStatusNoMatchingDocuments, wire: "\"no_matching_documents\""},
		{variant: KnowledgeBaseToolStatusNoResults, wire: "\"no_results\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded KnowledgeBaseToolStatus
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
