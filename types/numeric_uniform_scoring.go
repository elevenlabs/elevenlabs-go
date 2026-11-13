package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type NumericUniformScoring struct {
	MaxScore        int     `json:"max_score" xml:"max_score" required:"true"`
	Instructions    *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	extraProperties map[string]interface{}
}

func (n *NumericUniformScoring) GetMaxScore() (value int) {
	if n == nil {
		return
	}
	return n.MaxScore
}

func (n *NumericUniformScoring) GetInstructions() (value *string) {
	if n == nil {
		return
	}
	return n.Instructions
}

func (n *NumericUniformScoring) SetMaxScore(maxScore int) {
	n.MaxScore = maxScore
}

func (n *NumericUniformScoring) SetInstructions(instructions *string) {
	n.Instructions = instructions
}

func (n *NumericUniformScoring) GetExtraProperties() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.extraProperties
}

func (n NumericUniformScoring) String() string {
	jsonData, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return "error converting struct: NumericUniformScoring to string"
	}
	return string(jsonData)
}

func (n *NumericUniformScoring) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, n); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *n, "mode")
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	return nil
}

func (n *NumericUniformScoring) MarshalJSON() ([]byte, error) {
	type embed NumericUniformScoring
	var marshaler = struct {
		embed
	}{
		embed: embed(*n),
	}
	return json.Marshal(marshaler)
}
