package workspace

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost struct {
	// Start of the time range as a Unix timestamp in milliseconds. Must be at least 2020-01-01.
	StartTime int `json:"start_time" xml:"start_time" required:"true" min:"1577836800000"`
	// End of the time range as a Unix timestamp in milliseconds. Must be at least 2020-01-01.
	EndTime int `json:"end_time" xml:"end_time" required:"true" min:"1577836800000"`
	// Bucket size in seconds. Each row in the response covers this many seconds of the selected time range. For example, pass 3600 for hourly buckets, 86400 for daily buckets, or 604800 for weekly buckets. Whether `time_zone` shifts bucket boundaries depends on this value: week multiples (e.g. 604800) are calendar weeks that start Monday at local midnight; whole-day multiples that are not whole weeks (e.g. 86400) align to local midnight; whole-hour multiples up to 24 hours (e.g. 3600, 14400) align to local hour boundaries from midnight; sub-hour values and other sizes remain UTC-anchored regardless of `time_zone`.
	IntervalSeconds *int                                                                                  `json:"interval_seconds,omitempty" xml:"interval_seconds,omitempty" min:"1"`
	GroupBy         []BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem `json:"group_by,omitempty" xml:"group_by,omitempty"`
	Filters         []*types.ColumnFilter                                                                 `json:"filters,omitempty" xml:"filters,omitempty"`
	// IANA time zone identifier (e.g. 'America/New_York', 'Europe/London', 'UTC') used to align bucket boundaries for eligible `interval_seconds` values. Week multiples (e.g. 604800) are calendar weeks that start Monday at local midnight. Whole-day multiples that are not whole weeks start at local midnight; whole-hour multiples up to 24 hours align to local hour boundaries from midnight. Sub-hour intervals and other bucket sizes remain UTC-anchored regardless of this setting. Defaults to UTC.
	TimeZone *string `json:"time_zone,omitempty" xml:"time_zone,omitempty"`
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) GetStartTime() (value int) {
	if b == nil {
		return
	}
	return b.StartTime
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) GetEndTime() (value int) {
	if b == nil {
		return
	}
	return b.EndTime
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) GetIntervalSeconds() (value *int) {
	if b == nil {
		return
	}
	return b.IntervalSeconds
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) GetGroupBy() (value []BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem) {
	if b == nil {
		return
	}
	return b.GroupBy
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) GetFilters() (value []*types.ColumnFilter) {
	if b == nil {
		return
	}
	return b.Filters
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) GetTimeZone() (value *string) {
	if b == nil {
		return
	}
	return b.TimeZone
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) SetStartTime(startTime int) {
	b.StartTime = startTime
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) SetEndTime(endTime int) {
	b.EndTime = endTime
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) SetIntervalSeconds(intervalSeconds *int) {
	b.IntervalSeconds = intervalSeconds
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) SetGroupBy(groupBy []BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem) {
	b.GroupBy = groupBy
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) SetFilters(filters []*types.ColumnFilter) {
	b.Filters = filters
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) SetTimeZone(timeZone *string) {
	b.TimeZone = timeZone
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost to string"
	}
	return string(jsonData)
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost) MarshalJSON() ([]byte, error) {
	type embed BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem string

const (
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemProductType          BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "product_type"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemModel                BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "model"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemVoiceId              BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "voice_id"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemUserId               BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "user_id"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemFiatCurrency         BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "fiat_currency"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemFiatChargeType       BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "fiat_charge_type"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemRegion               BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "region"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemReportingWorkspaceId BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "reporting_workspace_id"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemRequestSource        BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "request_source"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemResourceId           BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "resource_id"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemSubresourceId        BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "subresource_id"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemRequestQueueType     BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "request_queue_type"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemVoiceMultiplier      BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "voice_multiplier"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemHashedXiApiKey       BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "hashed_xi_api_key"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemBillingGroupId       BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "billing_group_id"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemSurface              BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "surface"
	BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemActor                BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem = "actor"
)

func NewBodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemFromString(s string) (BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem, error) {
	switch s {
	case "product_type":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemProductType, nil
	case "model":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemModel, nil
	case "voice_id":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemVoiceId, nil
	case "user_id":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemUserId, nil
	case "fiat_currency":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemFiatCurrency, nil
	case "fiat_charge_type":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemFiatChargeType, nil
	case "region":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemRegion, nil
	case "reporting_workspace_id":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemReportingWorkspaceId, nil
	case "request_source":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemRequestSource, nil
	case "resource_id":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemResourceId, nil
	case "subresource_id":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemSubresourceId, nil
	case "request_queue_type":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemRequestQueueType, nil
	case "voice_multiplier":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemVoiceMultiplier, nil
	case "hashed_xi_api_key":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemHashedXiApiKey, nil
	case "billing_group_id":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemBillingGroupId, nil
	case "surface":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemSurface, nil
	case "actor":
		return BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItemActor, nil
	}
	var t BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem) Ptr() *BodyGetWorkspaceUsageV1WorkspaceAnalyticsQueryUsageByProductOverTimePostGroupByItem {
	return &e
}

type BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost struct {
	// Start of the time range as a Unix timestamp in milliseconds. Must be at least 2020-01-01.
	StartTime int `json:"start_time" xml:"start_time" required:"true" min:"1577836800000"`
	// End of the time range as a Unix timestamp in milliseconds. Must be at least 2020-01-01.
	EndTime int `json:"end_time" xml:"end_time" required:"true" min:"1577836800000"`
	// Bucket size in seconds, in whole UTC days. Pass 86400 for daily buckets or 604800 for weekly buckets, which start on Monday.
	IntervalSeconds int                                                                 `json:"interval_seconds" xml:"interval_seconds" required:"true" min:"86400" multipleOf:"86400"`
	GroupBy         []BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem `json:"group_by,omitempty" xml:"group_by,omitempty"`
	Filters         []*types.MetronomeUsageColumnFilter                                 `json:"filters,omitempty" xml:"filters,omitempty"`
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) GetStartTime() (value int) {
	if b == nil {
		return
	}
	return b.StartTime
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) GetEndTime() (value int) {
	if b == nil {
		return
	}
	return b.EndTime
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) GetIntervalSeconds() (value int) {
	if b == nil {
		return
	}
	return b.IntervalSeconds
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) GetGroupBy() (value []BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem) {
	if b == nil {
		return
	}
	return b.GroupBy
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) GetFilters() (value []*types.MetronomeUsageColumnFilter) {
	if b == nil {
		return
	}
	return b.Filters
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) SetStartTime(startTime int) {
	b.StartTime = startTime
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) SetEndTime(endTime int) {
	b.EndTime = endTime
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) SetIntervalSeconds(intervalSeconds int) {
	b.IntervalSeconds = intervalSeconds
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) SetGroupBy(groupBy []BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem) {
	b.GroupBy = groupBy
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) SetFilters(filters []*types.MetronomeUsageColumnFilter) {
	b.Filters = filters
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost to string"
	}
	return string(jsonData)
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost) MarshalJSON() ([]byte, error) {
	type embed BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem string

const (
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemModel                BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "model"
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemRequestSource        BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "request_source"
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemFiatChargeType       BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "fiat_charge_type"
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemFiatCurrency         BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "fiat_currency"
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemUserId               BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "user_id"
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemHashedXiApiKey       BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "hashed_xi_api_key"
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemRegion               BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "region"
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemReportingWorkspaceId BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "reporting_workspace_id"
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemBillingGroupId       BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "billing_group_id"
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemProductType          BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "product_type"
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemSurface              BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "surface"
	BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemActor                BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem = "actor"
)

func NewBodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemFromString(s string) (BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem, error) {
	switch s {
	case "model":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemModel, nil
	case "request_source":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemRequestSource, nil
	case "fiat_charge_type":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemFiatChargeType, nil
	case "fiat_currency":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemFiatCurrency, nil
	case "user_id":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemUserId, nil
	case "hashed_xi_api_key":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemHashedXiApiKey, nil
	case "region":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemRegion, nil
	case "reporting_workspace_id":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemReportingWorkspaceId, nil
	case "billing_group_id":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemBillingGroupId, nil
	case "product_type":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemProductType, nil
	case "surface":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemSurface, nil
	case "actor":
		return BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItemActor, nil
	}
	var t BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem) Ptr() *BodyGetWorkspaceCostUsageV1WorkspaceAnalyticsUsagePostGroupByItem {
	return &e
}
