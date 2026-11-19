package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MatchAnythingParameterEvaluationStrategy struct {
	type_           string
	extraProperties map[string]interface{}
}

func (m *MatchAnythingParameterEvaluationStrategy) Type() string {
	return m.type_
}

func (m *MatchAnythingParameterEvaluationStrategy) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MatchAnythingParameterEvaluationStrategy) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MatchAnythingParameterEvaluationStrategy to string"
	}
	return string(jsonData)
}

func (m *MatchAnythingParameterEvaluationStrategy) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MatchAnythingParameterEvaluationStrategy
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MatchAnythingParameterEvaluationStrategy(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "anything" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", m, "anything", *literals.Type)
	}
	m.type_ = "anything"
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "type")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MatchAnythingParameterEvaluationStrategy) MarshalJSON() ([]byte, error) {
	type embed MatchAnythingParameterEvaluationStrategy
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*m),
		Type:  "anything",
	}
	return json.Marshal(marshaler)
}
