package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListResponseAgentBranchSummary struct {
	Meta            *ListResponseMeta     `json:"meta,omitempty" xml:"meta,omitempty"`
	Results         []*AgentBranchSummary `json:"results" xml:"results" required:"true"`
	extraProperties map[string]interface{}
}

func (l *ListResponseAgentBranchSummary) GetMeta() (value *ListResponseMeta) {
	if l == nil {
		return
	}
	return l.Meta
}

func (l *ListResponseAgentBranchSummary) GetResults() (value []*AgentBranchSummary) {
	if l == nil {
		return
	}
	return l.Results
}

func (l *ListResponseAgentBranchSummary) SetMeta(meta *ListResponseMeta) {
	l.Meta = meta
}

func (l *ListResponseAgentBranchSummary) SetResults(results []*AgentBranchSummary) {
	l.Results = results
}

func (l *ListResponseAgentBranchSummary) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListResponseAgentBranchSummary) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListResponseAgentBranchSummary to string"
	}
	return string(jsonData)
}

func (l *ListResponseAgentBranchSummary) UnmarshalJSON(data []byte) error {
	type unmarshaler ListResponseAgentBranchSummary
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListResponseAgentBranchSummary(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListResponseAgentBranchSummary) MarshalJSON() ([]byte, error) {
	type embed ListResponseAgentBranchSummary
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
