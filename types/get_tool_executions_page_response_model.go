package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetToolExecutionsPageResponseModel struct {
	Executions      []*ToolExecutionResponseModel `json:"executions" xml:"executions" required:"true"`
	NextCursor      *string                       `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         bool                          `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetToolExecutionsPageResponseModel) GetExecutions() (value []*ToolExecutionResponseModel) {
	if g == nil {
		return
	}
	return g.Executions
}

func (g *GetToolExecutionsPageResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetToolExecutionsPageResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetToolExecutionsPageResponseModel) SetExecutions(executions []*ToolExecutionResponseModel) {
	g.Executions = executions
}

func (g *GetToolExecutionsPageResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetToolExecutionsPageResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetToolExecutionsPageResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetToolExecutionsPageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetToolExecutionsPageResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetToolExecutionsPageResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetToolExecutionsPageResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetToolExecutionsPageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
