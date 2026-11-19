package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetTestInvocationsPageResponseModel struct {
	Meta    *ListResponseMeta                     `json:"meta,omitempty" xml:"meta,omitempty"`
	Results []*TestInvocationSummaryResponseModel `json:"results" xml:"results" required:"true"`
	// Cursor for the next page of results
	NextCursor *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	// Whether there are more results available
	HasMore         bool `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetTestInvocationsPageResponseModel) GetMeta() (value *ListResponseMeta) {
	if g == nil {
		return
	}
	return g.Meta
}

func (g *GetTestInvocationsPageResponseModel) GetResults() (value []*TestInvocationSummaryResponseModel) {
	if g == nil {
		return
	}
	return g.Results
}

func (g *GetTestInvocationsPageResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetTestInvocationsPageResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetTestInvocationsPageResponseModel) SetMeta(meta *ListResponseMeta) {
	g.Meta = meta
}

func (g *GetTestInvocationsPageResponseModel) SetResults(results []*TestInvocationSummaryResponseModel) {
	g.Results = results
}

func (g *GetTestInvocationsPageResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetTestInvocationsPageResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetTestInvocationsPageResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetTestInvocationsPageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTestInvocationsPageResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetTestInvocationsPageResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetTestInvocationsPageResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetTestInvocationsPageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
