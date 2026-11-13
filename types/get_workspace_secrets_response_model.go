package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetWorkspaceSecretsResponseModel struct {
	Secrets []*ConvAiWorkspaceStoredSecretConfig `json:"secrets" xml:"secrets" required:"true"`
	// Cursor for fetching the next page of secrets
	NextCursor      *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetWorkspaceSecretsResponseModel) GetSecrets() (value []*ConvAiWorkspaceStoredSecretConfig) {
	if g == nil {
		return
	}
	return g.Secrets
}

func (g *GetWorkspaceSecretsResponseModel) GetNextCursor() (value *string) {
	if g == nil {
		return
	}
	return g.NextCursor
}

func (g *GetWorkspaceSecretsResponseModel) SetSecrets(secrets []*ConvAiWorkspaceStoredSecretConfig) {
	g.Secrets = secrets
}

func (g *GetWorkspaceSecretsResponseModel) SetNextCursor(nextCursor *string) {
	g.NextCursor = nextCursor
}

func (g *GetWorkspaceSecretsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetWorkspaceSecretsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetWorkspaceSecretsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetWorkspaceSecretsResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetWorkspaceSecretsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetWorkspaceSecretsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
