package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// One page of the caller's public-API template runs, newest first.
type TemplateRunListResponse struct {
	// The runs on this page, newest first. Each item has the same shape as `GET /v1/flows/templates/{template_id}/runs/{run_id}`.
	Runs []*TemplateRunResponse `json:"runs" xml:"runs" required:"true"`
	// Pass as `cursor` to fetch the next page. `null` when there is no further page.
	NextCursor *string `json:"next_cursor" xml:"next_cursor" required:"true"`
	// Whether more runs exist beyond this page.
	HasMore         *bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TemplateRunListResponse) GetRuns() (value []*TemplateRunResponse) {
	if t == nil {
		return
	}
	return t.Runs
}

func (t *TemplateRunListResponse) GetNextCursor() (value *string) {
	if t == nil {
		return
	}
	return t.NextCursor
}

func (t *TemplateRunListResponse) GetHasMore() (value *bool) {
	if t == nil {
		return
	}
	return t.HasMore
}

func (t *TemplateRunListResponse) SetRuns(runs []*TemplateRunResponse) {
	t.Runs = runs
}

func (t *TemplateRunListResponse) SetNextCursor(nextCursor *string) {
	t.NextCursor = nextCursor
}

func (t *TemplateRunListResponse) SetHasMore(hasMore *bool) {
	t.HasMore = hasMore
}

func (t *TemplateRunListResponse) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateRunListResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateRunListResponse to string"
	}
	return string(jsonData)
}

func (t *TemplateRunListResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateRunListResponse) MarshalJSON() ([]byte, error) {
	type embed TemplateRunListResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
