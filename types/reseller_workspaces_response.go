package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ResellerWorkspacesResponse struct {
	Workspaces      []*ResellerWorkspaceSummary `json:"workspaces" xml:"workspaces" required:"true"`
	NextCursor      *string                     `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         *bool                       `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ResellerWorkspacesResponse) GetWorkspaces() (value []*ResellerWorkspaceSummary) {
	if r == nil {
		return
	}
	return r.Workspaces
}

func (r *ResellerWorkspacesResponse) GetNextCursor() (value *string) {
	if r == nil {
		return
	}
	return r.NextCursor
}

func (r *ResellerWorkspacesResponse) GetHasMore() (value *bool) {
	if r == nil {
		return
	}
	return r.HasMore
}

func (r *ResellerWorkspacesResponse) SetWorkspaces(workspaces []*ResellerWorkspaceSummary) {
	r.Workspaces = workspaces
}

func (r *ResellerWorkspacesResponse) SetNextCursor(nextCursor *string) {
	r.NextCursor = nextCursor
}

func (r *ResellerWorkspacesResponse) SetHasMore(hasMore *bool) {
	r.HasMore = hasMore
}

func (r *ResellerWorkspacesResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerWorkspacesResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerWorkspacesResponse to string"
	}
	return string(jsonData)
}

func (r *ResellerWorkspacesResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ResellerWorkspacesResponse) MarshalJSON() ([]byte, error) {
	type embed ResellerWorkspacesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
