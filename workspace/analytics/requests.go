package analytics

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost struct {
	// Start of the time range as a Unix timestamp in milliseconds.
	StartTime *int `json:"start_time,omitempty" xml:"start_time,omitempty" min:"1577836800000"`
	// End of the time range as a Unix timestamp in milliseconds.
	EndTime *int `json:"end_time,omitempty" xml:"end_time,omitempty" min:"1577836800000"`
	Limit   *int `json:"limit,omitempty" xml:"limit,omitempty" min:"1" max:"1000"`
	// Optional timestamp sort direction. If omitted, defaults to desc when end_time is provided, otherwise asc.
	Sort    *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort `json:"sort,omitempty" xml:"sort,omitempty"`
	Filters []*types.ColumnFilter                                    `json:"filters,omitempty" xml:"filters,omitempty"`
	Search  *string                                                  `json:"search,omitempty" xml:"search,omitempty"`
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) GetStartTime() (value *int) {
	if b == nil {
		return
	}
	return b.StartTime
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) GetEndTime() (value *int) {
	if b == nil {
		return
	}
	return b.EndTime
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) GetLimit() (value *int) {
	if b == nil {
		return
	}
	return b.Limit
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) GetSort() (value *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort) {
	if b == nil {
		return
	}
	return b.Sort
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) GetFilters() (value []*types.ColumnFilter) {
	if b == nil {
		return
	}
	return b.Filters
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) GetSearch() (value *string) {
	if b == nil {
		return
	}
	return b.Search
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) SetStartTime(startTime *int) {
	b.StartTime = startTime
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) SetEndTime(endTime *int) {
	b.EndTime = endTime
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) SetLimit(limit *int) {
	b.Limit = limit
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) SetSort(sort *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort) {
	b.Sort = sort
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) SetFilters(filters []*types.ColumnFilter) {
	b.Filters = filters
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) SetSearch(search *string) {
	b.Search = search
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost to string"
	}
	return string(jsonData)
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost) MarshalJSON() ([]byte, error) {
	type embed BodyListApiRequestsV1WorkspaceAnalyticsRequestsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort string

const (
	BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSortAsc  BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort = "asc"
	BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSortDesc BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort = "desc"
)

func NewBodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSortFromString(s string) (BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort, error) {
	switch s {
	case "asc":
		return BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSortAsc, nil
	case "desc":
		return BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSortDesc, nil
	}
	var t BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort) Ptr() *BodyListApiRequestsV1WorkspaceAnalyticsRequestsPostSort {
	return &e
}
