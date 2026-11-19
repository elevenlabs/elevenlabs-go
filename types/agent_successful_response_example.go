package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentSuccessfulResponseExample struct {
	Response        string `json:"response" xml:"response" required:"true"`
	type_           string
	extraProperties map[string]interface{}
}

func (a *AgentSuccessfulResponseExample) GetResponse() (value string) {
	if a == nil {
		return
	}
	return a.Response
}

func (a *AgentSuccessfulResponseExample) SetResponse(response string) {
	a.Response = response
}

func (a *AgentSuccessfulResponseExample) Type() string {
	return a.type_
}

func (a *AgentSuccessfulResponseExample) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentSuccessfulResponseExample) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentSuccessfulResponseExample to string"
	}
	return string(jsonData)
}

func (a *AgentSuccessfulResponseExample) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentSuccessfulResponseExample
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentSuccessfulResponseExample(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "success" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "success", *literals.Type)
	}
	a.type_ = "success"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentSuccessfulResponseExample) MarshalJSON() ([]byte, error) {
	type embed AgentSuccessfulResponseExample
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "success",
	}
	return json.Marshal(marshaler)
}
