package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetTestsPageResponseModel struct {
	Tests           []*UnitTestSummaryResponseModel `json:"tests" xml:"tests" required:"true"`
	NextCursor      *string                         `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	HasMore         bool                            `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetTestsPageResponseModel) GetTests() (value []*UnitTestSummaryResponseModel) {
	if g == nil {
		return
	}
	return g.Tests
}

func (g *GetTestsPageResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetTestsPageResponseModel) GetHasMore() (value bool) {
	if g == nil {
		return
	}
	return g.HasMore
}

func (g *GetTestsPageResponseModel) SetTests(tests []*UnitTestSummaryResponseModel) {
	g.Tests = tests
}

func (g *GetTestsPageResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetTestsPageResponseModel) SetHasMore(hasMore bool) {
	g.HasMore = hasMore
}

func (g *GetTestsPageResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetTestsPageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTestsPageResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetTestsPageResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetTestsPageResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetTestsPageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
