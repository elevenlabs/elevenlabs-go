package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetTestsSummariesByIdsResponseModel struct {
	// Dictionary mapping test IDs to their summary information
	Tests           map[string]*UnitTestSummaryResponseModel `json:"tests" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetTestsSummariesByIdsResponseModel) GetTests() (value map[string]*UnitTestSummaryResponseModel) {
	if g == nil {
		return
	}
	return g.Tests
}

func (g *GetTestsSummariesByIdsResponseModel) SetTests(tests map[string]*UnitTestSummaryResponseModel) {
	g.Tests = tests
}

func (g *GetTestsSummariesByIdsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetTestsSummariesByIdsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetTestsSummariesByIdsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetTestsSummariesByIdsResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler GetTestsSummariesByIdsResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetTestsSummariesByIdsResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetTestsSummariesByIdsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetTestsSummariesByIdsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
