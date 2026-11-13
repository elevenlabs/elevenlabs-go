package workspace

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_time": 1577836800000,
  "end_time": 1577836800000,
  "interval_seconds": 123,
  "group_by": [
    "product_type"
  ],
  "filters": [
    {
      "column": "value",
      "operation": "in",
      "values": [
        "value"
      ]
    }
  ],
  "time_zone": "value"
}`
	assertRoundTrip(t, wire, new(BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost))
}

func TestBodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_time": 1577836800000,
  "end_time": 1577836800000,
  "interval_seconds": 86400,
  "group_by": [
    "model"
  ],
  "filters": [
    {
      "column": "model",
      "operation": "eq",
      "values": [
        "value"
      ]
    }
  ]
}`
	assertRoundTrip(t, wire, new(BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost))
}

func TestBodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem
		wire    string
	}
	cases := []enumCase{
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemProductType, wire: "\"product_type\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemModel, wire: "\"model\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemVoiceId, wire: "\"voice_id\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemUserId, wire: "\"user_id\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemFiatCurrency, wire: "\"fiat_currency\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemFiatChargeType, wire: "\"fiat_charge_type\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemRegion, wire: "\"region\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemReportingWorkspaceId, wire: "\"reporting_workspace_id\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemRequestSource, wire: "\"request_source\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemResourceId, wire: "\"resource_id\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemSubresourceId, wire: "\"subresource_id\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemRequestQueueType, wire: "\"request_queue_type\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemVoiceMultiplier, wire: "\"voice_multiplier\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemHashedXiApiKey, wire: "\"hashed_xi_api_key\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemBillingGroupId, wire: "\"billing_group_id\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemSurface, wire: "\"surface\""},
		{variant: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemActor, wire: "\"actor\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem
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

func TestBodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem
		wire    string
	}
	cases := []enumCase{
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemModel, wire: "\"model\""},
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemRequestSource, wire: "\"request_source\""},
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemFiatChargeType, wire: "\"fiat_charge_type\""},
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemFiatCurrency, wire: "\"fiat_currency\""},
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemUserId, wire: "\"user_id\""},
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemHashedXiApiKey, wire: "\"hashed_xi_api_key\""},
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemRegion, wire: "\"region\""},
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemReportingWorkspaceId, wire: "\"reporting_workspace_id\""},
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemBillingGroupId, wire: "\"billing_group_id\""},
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemProductType, wire: "\"product_type\""},
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemSurface, wire: "\"surface\""},
		{variant: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemActor, wire: "\"actor\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem
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
