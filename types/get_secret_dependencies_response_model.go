package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetSecretDependenciesResponseModel struct {
	Dependencies *GetSecretDependenciesResponseModelDependencies `json:"dependencies" xml:"dependencies" required:"true"`
	// Cursor for fetching the next page of dependencies
	NextCursor      *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetSecretDependenciesResponseModel) GetDependencies() (value *GetSecretDependenciesResponseModelDependencies) {
	if g == nil {
		return
	}
	return g.Dependencies
}

func (g *GetSecretDependenciesResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetSecretDependenciesResponseModel) SetDependencies(dependencies *GetSecretDependenciesResponseModelDependencies) {
	g.Dependencies = dependencies
}

func (g *GetSecretDependenciesResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetSecretDependenciesResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetSecretDependenciesResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetSecretDependenciesResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetSecretDependenciesResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetSecretDependenciesResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetSecretDependenciesResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
