package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestReviewResponseModelRejectReasonsItemVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ReviewResponseModelRejectReasonsItem
		wire    string
	}
	cases := []enumCase{
		{variant: ReviewResponseModelRejectReasonsItemLacksStructure, wire: "\"lacks_structure\""},
		{variant: ReviewResponseModelRejectReasonsItemDoesntOpen, wire: "\"doesnt_open\""},
		{variant: ReviewResponseModelRejectReasonsItemNotLiteraryWork, wire: "\"not_literary_work\""},
		{variant: ReviewResponseModelRejectReasonsItemLanguageNotSupported, wire: "\"language_not_supported\""},
		{variant: ReviewResponseModelRejectReasonsItemTooShort, wire: "\"too_short\""},
		{variant: ReviewResponseModelRejectReasonsItemDuplicate, wire: "\"duplicate\""},
		{variant: ReviewResponseModelRejectReasonsItemPromotional, wire: "\"promotional\""},
		{variant: ReviewResponseModelRejectReasonsItemFormattingIssues, wire: "\"formatting_issues\""},
		{variant: ReviewResponseModelRejectReasonsItemLowQuality, wire: "\"low_quality\""},
		{variant: ReviewResponseModelRejectReasonsItemMetadataIncomplete, wire: "\"metadata_incomplete\""},
		{variant: ReviewResponseModelRejectReasonsItemMetadataInaccurate, wire: "\"metadata_inaccurate\""},
		{variant: ReviewResponseModelRejectReasonsItemTypos, wire: "\"typos\""},
		{variant: ReviewResponseModelRejectReasonsItemReviewError, wire: "\"review_error\""},
		{variant: ReviewResponseModelRejectReasonsItemSpam, wire: "\"spam\""},
		{variant: ReviewResponseModelRejectReasonsItemLegalViolation, wire: "\"legal_violation\""},
		{variant: ReviewResponseModelRejectReasonsItemContentPolicy, wire: "\"content_policy\""},
		{variant: ReviewResponseModelRejectReasonsItemPublicDomain, wire: "\"public_domain\""},
		{variant: ReviewResponseModelRejectReasonsItemOther, wire: "\"other\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ReviewResponseModelRejectReasonsItem
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
