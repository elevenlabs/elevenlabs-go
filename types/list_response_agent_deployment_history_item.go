package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListResponseAgentDeploymentHistoryItem struct {
	Meta            *ListResponseMeta             `json:"meta,omitempty" xml:"meta,omitempty"`
	Results         []*AgentDeploymentHistoryItem `json:"results" xml:"results" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListResponseAgentDeploymentHistoryItem) GetMeta() (value *ListResponseMeta) {
	if l == nil {
		return
	}
	return l.Meta
}

func (l *ListResponseAgentDeploymentHistoryItem) GetResults() (value []*AgentDeploymentHistoryItem) {
	if l == nil {
		return
	}
	return l.Results
}

func (l *ListResponseAgentDeploymentHistoryItem) SetMeta(meta *ListResponseMeta) {
	l.Meta = meta
}

func (l *ListResponseAgentDeploymentHistoryItem) SetResults(results []*AgentDeploymentHistoryItem) {
	l.Results = results
}

func (l *ListResponseAgentDeploymentHistoryItem) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListResponseAgentDeploymentHistoryItem) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListResponseAgentDeploymentHistoryItem to string"
	}
	return string(jsonData)
}

func (l *ListResponseAgentDeploymentHistoryItem) UnmarshalJSON(data []byte) error {
	type unmarshaler ListResponseAgentDeploymentHistoryItem
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListResponseAgentDeploymentHistoryItem(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListResponseAgentDeploymentHistoryItem) MarshalJSON() ([]byte, error) {
	type embed ListResponseAgentDeploymentHistoryItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
