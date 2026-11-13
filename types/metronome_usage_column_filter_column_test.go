package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMetronomeUsageColumnFilterColumnVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant MetronomeUsageColumnFilterColumn
		wire    string
	}
	cases := []enumCase{
		{variant: MetronomeUsageColumnFilterColumnModel, wire: "\"model\""},
		{variant: MetronomeUsageColumnFilterColumnRequestSource, wire: "\"request_source\""},
		{variant: MetronomeUsageColumnFilterColumnFiatChargeType, wire: "\"fiat_charge_type\""},
		{variant: MetronomeUsageColumnFilterColumnFiatCurrency, wire: "\"fiat_currency\""},
		{variant: MetronomeUsageColumnFilterColumnUserId, wire: "\"user_id\""},
		{variant: MetronomeUsageColumnFilterColumnHashedXiApiKey, wire: "\"hashed_xi_api_key\""},
		{variant: MetronomeUsageColumnFilterColumnRegion, wire: "\"region\""},
		{variant: MetronomeUsageColumnFilterColumnReportingWorkspaceId, wire: "\"reporting_workspace_id\""},
		{variant: MetronomeUsageColumnFilterColumnBillingGroupId, wire: "\"billing_group_id\""},
		{variant: MetronomeUsageColumnFilterColumnProductType, wire: "\"product_type\""},
		{variant: MetronomeUsageColumnFilterColumnSurface, wire: "\"surface\""},
		{variant: MetronomeUsageColumnFilterColumnActor, wire: "\"actor\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded MetronomeUsageColumnFilterColumn
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
