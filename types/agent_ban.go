package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentBan struct {
	AtUnix          int           `json:"at_unix" xml:"at_unix" required:"true"`
	Reason          *string       `json:"reason,omitempty" xml:"reason,omitempty"`
	ReasonType      BanReasonType `json:"reason_type" xml:"reason_type" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentBan) GetAtUnix() (value int) {
	if a == nil {
		return
	}
	return a.AtUnix
}

func (a *AgentBan) GetReason() (value *string) {
	if a == nil {
		return
	}
	return a.Reason
}

func (a *AgentBan) GetReasonType() (value BanReasonType) {
	if a == nil {
		return
	}
	return a.ReasonType
}

func (a *AgentBan) SetAtUnix(atUnix int) {
	a.AtUnix = atUnix
}

func (a *AgentBan) SetReason(reason *string) {
	a.Reason = reason
}

func (a *AgentBan) SetReasonType(reasonType BanReasonType) {
	a.ReasonType = reasonType
}

func (a *AgentBan) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentBan) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentBan to string"
	}
	return string(jsonData)
}

func (a *AgentBan) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentBan) MarshalJSON() ([]byte, error) {
	type embed AgentBan
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
