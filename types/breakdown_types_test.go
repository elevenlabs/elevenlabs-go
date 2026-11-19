package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBreakdownTypesVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BreakdownTypes
		wire    string
	}
	cases := []enumCase{
		{variant: BreakdownTypesNone, wire: "\"none\""},
		{variant: BreakdownTypesVoice, wire: "\"voice\""},
		{variant: BreakdownTypesVoiceMultiplier, wire: "\"voice_multiplier\""},
		{variant: BreakdownTypesUser, wire: "\"user\""},
		{variant: BreakdownTypesGroups, wire: "\"groups\""},
		{variant: BreakdownTypesApiKeys, wire: "\"api_keys\""},
		{variant: BreakdownTypesAllApiKeys, wire: "\"all_api_keys\""},
		{variant: BreakdownTypesProductType, wire: "\"product_type\""},
		{variant: BreakdownTypesModel, wire: "\"model\""},
		{variant: BreakdownTypesResource, wire: "\"resource\""},
		{variant: BreakdownTypesRequestQueue, wire: "\"request_queue\""},
		{variant: BreakdownTypesRegion, wire: "\"region\""},
		{variant: BreakdownTypesSubresourceId, wire: "\"subresource_id\""},
		{variant: BreakdownTypesReportingWorkspaceId, wire: "\"reporting_workspace_id\""},
		{variant: BreakdownTypesHasApiKey, wire: "\"has_api_key\""},
		{variant: BreakdownTypesRequestSource, wire: "\"request_source\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BreakdownTypes
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
