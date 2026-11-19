package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TransferToAgentToolResultErrorModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	FromAgent       string  `json:"from_agent" xml:"from_agent" required:"true"`
	Error           string  `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TransferToAgentToolResultErrorModel) GetStatus() (value *string) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TransferToAgentToolResultErrorModel) GetFromAgent() (value string) {
	if t == nil {
		return
	}
	return t.FromAgent
}

func (t *TransferToAgentToolResultErrorModel) GetError() (value string) {
	if t == nil {
		return
	}
	return t.Error
}

func (t *TransferToAgentToolResultErrorModel) SetStatus(status *string) {
	t.Status = status
}

func (t *TransferToAgentToolResultErrorModel) SetFromAgent(fromAgent string) {
	t.FromAgent = fromAgent
}

func (t *TransferToAgentToolResultErrorModel) SetError(error string) {
	t.Error = error
}

func (t *TransferToAgentToolResultErrorModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TransferToAgentToolResultErrorModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TransferToAgentToolResultErrorModel to string"
	}
	return string(jsonData)
}

func (t *TransferToAgentToolResultErrorModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TransferToAgentToolResultErrorModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TransferToAgentToolResultErrorModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "result_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TransferToAgentToolResultErrorModel) MarshalJSON() ([]byte, error) {
	type embed TransferToAgentToolResultErrorModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
