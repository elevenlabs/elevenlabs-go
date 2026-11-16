package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentOutputError struct {
	Type            string `json:"type" xml:"type" required:"true"`
	Code            string `json:"code" xml:"code" required:"true"`
	Message         string `json:"message" xml:"message" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentOutputError) GetType() (value string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentOutputError) GetCode() (value string) {
	if a == nil {
		return
	}
	return a.Code
}

func (a *AgentOutputError) GetMessage() (value string) {
	if a == nil {
		return
	}
	return a.Message
}

func (a *AgentOutputError) SetType(type_ string) {
	a.Type = type_
}

func (a *AgentOutputError) SetCode(code string) {
	a.Code = code
}

func (a *AgentOutputError) SetMessage(message string) {
	a.Message = message
}

func (a *AgentOutputError) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentOutputError) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentOutputError to string"
	}
	return string(jsonData)
}

func (a *AgentOutputError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentOutputError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentOutputError(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentOutputError) MarshalJSON() ([]byte, error) {
	type embed AgentOutputError
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
