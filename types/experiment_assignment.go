package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Experiment membership recorded on a conversation.
type ExperimentAssignment struct {
	// Experiment key.
	Key string `json:"key" xml:"key" required:"true"`
	// Variant identifier: the branch id for server_branch, or the client-supplied variant for client_declared.
	Variant string                     `json:"variant" xml:"variant" required:"true"`
	Source  ExperimentAssignmentSource `json:"source" xml:"source" required:"true"`
	// The AgentExperiment id, when the assignment maps to a registered experiment.
	ExperimentId    *string `json:"experiment_id,omitempty" xml:"experiment_id,omitempty"`
	extraProperties map[string]interface{}
}

func (e *ExperimentAssignment) GetKey() (value string) {
	if e == nil {
		return
	}
	return e.Key
}

func (e *ExperimentAssignment) GetVariant() (value string) {
	if e == nil {
		return
	}
	return e.Variant
}

func (e *ExperimentAssignment) GetSource() (value ExperimentAssignmentSource) {
	if e == nil {
		return
	}
	return e.Source
}

func (e *ExperimentAssignment) GetExperimentId() (value *string) {
	if e == nil {
		return
	}
	return e.ExperimentId
}

func (e *ExperimentAssignment) SetKey(key string) {
	e.Key = key
}

func (e *ExperimentAssignment) SetVariant(variant string) {
	e.Variant = variant
}

func (e *ExperimentAssignment) SetSource(source ExperimentAssignmentSource) {
	e.Source = source
}

func (e *ExperimentAssignment) SetExperimentId(experimentId *string) {
	e.ExperimentId = experimentId
}

func (e *ExperimentAssignment) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExperimentAssignment) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExperimentAssignment to string"
	}
	return string(jsonData)
}

func (e *ExperimentAssignment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExperimentAssignment) MarshalJSON() ([]byte, error) {
	type embed ExperimentAssignment
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
