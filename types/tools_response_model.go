package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ToolsResponseModel struct {
	Tools           []*ToolResponseModel `json:"tools" xml:"tools" required:"true"`
	NextCursor      *string              `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         bool                 `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (t *ToolsResponseModel) GetTools() (value []*ToolResponseModel) {
	if t == nil {
		return
	}
	return t.Tools
}

func (t *ToolsResponseModel) GetNextCursor() (value *string) {
	if t == nil {
		return
	}
	return t.NextCursor
}

func (t *ToolsResponseModel) GetHasMore() (value bool) {
	if t == nil {
		return
	}
	return t.HasMore
}

func (t *ToolsResponseModel) SetTools(tools []*ToolResponseModel) {
	t.Tools = tools
}

func (t *ToolsResponseModel) SetNextCursor(nextCursor *string) {
	t.NextCursor = nextCursor
}

func (t *ToolsResponseModel) SetHasMore(hasMore bool) {
	t.HasMore = hasMore
}

func (t *ToolsResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t ToolsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ToolsResponseModel to string"
	}
	return string(jsonData)
}

func (t *ToolsResponseModel) UnmarshalJSON(data []byte) error {
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

func (t *ToolsResponseModel) MarshalJSON() ([]byte, error) {
	type embed ToolsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
