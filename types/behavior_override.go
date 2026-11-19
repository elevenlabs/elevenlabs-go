package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BehaviorOverride struct {
	// Verbosity override. Underlying default applies when unset.
	Verbosity *Verbosity `json:"verbosity,omitempty" xml:"verbosity,omitempty"`
	// Output format override. Underlying default applies when unset.
	OutputFormat *OutputFormat `json:"output_format,omitempty" xml:"output_format,omitempty"`
	// Interaction budget override. Underlying default applies when unset.
	InteractionBudget *InteractionBudget `json:"interaction_budget,omitempty" xml:"interaction_budget,omitempty"`
	extraProperties   map[string]interface{}
}

func (b *BehaviorOverride) GetVerbosity() (value *Verbosity) {
	if b == nil {
		return
	}
	return b.Verbosity
}

func (b *BehaviorOverride) GetOutputFormat() (value *OutputFormat) {
	if b == nil {
		return
	}
	return b.OutputFormat
}

func (b *BehaviorOverride) GetInteractionBudget() (value *InteractionBudget) {
	if b == nil {
		return
	}
	return b.InteractionBudget
}

func (b *BehaviorOverride) SetVerbosity(verbosity *Verbosity) {
	b.Verbosity = verbosity
}

func (b *BehaviorOverride) SetOutputFormat(outputFormat *OutputFormat) {
	b.OutputFormat = outputFormat
}

func (b *BehaviorOverride) SetInteractionBudget(interactionBudget *InteractionBudget) {
	b.InteractionBudget = interactionBudget
}

func (b *BehaviorOverride) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BehaviorOverride) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BehaviorOverride to string"
	}
	return string(jsonData)
}

func (b *BehaviorOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BehaviorOverride) MarshalJSON() ([]byte, error) {
	type embed BehaviorOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
