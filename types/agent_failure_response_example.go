package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentFailureResponseExample struct {
	Response        string `json:"response" xml:"response" required:"true"`
	type_           string
	extraProperties map[string]interface{}
}

func (a *AgentFailureResponseExample) GetResponse() (value string) {
	if a == nil {
		return
	}
	return a.Response
}

func (a *AgentFailureResponseExample) SetResponse(response string) {
	a.Response = response
}

func (a *AgentFailureResponseExample) Type() string {
	return a.type_
}

func (a *AgentFailureResponseExample) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentFailureResponseExample) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentFailureResponseExample to string"
	}
	return string(jsonData)
}

func (a *AgentFailureResponseExample) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentFailureResponseExample
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentFailureResponseExample(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "failure" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "failure", *literals.Type)
	}
	a.type_ = "failure"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentFailureResponseExample) MarshalJSON() ([]byte, error) {
	type embed AgentFailureResponseExample
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "failure",
	}
	return json.Marshal(marshaler)
}
