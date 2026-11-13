package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PostWorkspaceSecretResponseModel struct {
	type_           string
	SecretId        string `json:"secret_id" xml:"secret_id" required:"true"`
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PostWorkspaceSecretResponseModel) GetSecretId() (value string) {
	if p == nil {
		return
	}
	return p.SecretId
}

func (p *PostWorkspaceSecretResponseModel) GetName() (value string) {
	if p == nil {
		return
	}
	return p.Name
}

func (p *PostWorkspaceSecretResponseModel) SetSecretId(secretId string) {
	p.SecretId = secretId
}

func (p *PostWorkspaceSecretResponseModel) SetName(name string) {
	p.Name = name
}

func (p *PostWorkspaceSecretResponseModel) Type() string {
	return p.type_
}

func (p *PostWorkspaceSecretResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PostWorkspaceSecretResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PostWorkspaceSecretResponseModel to string"
	}
	return string(jsonData)
}

func (p *PostWorkspaceSecretResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PostWorkspaceSecretResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PostWorkspaceSecretResponseModel(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "stored" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "stored", *literals.Type)
	}
	p.type_ = "stored"
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PostWorkspaceSecretResponseModel) MarshalJSON() ([]byte, error) {
	type embed PostWorkspaceSecretResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "stored",
	}
	return json.Marshal(marshaler)
}
