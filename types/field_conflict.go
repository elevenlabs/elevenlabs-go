package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type FieldConflict struct {
	// Identifier of the conflicting field relative to its section: a dot-path within conversation_config/platform_settings (e.g. `tts.voice_id`), a procedure id, or a workflow path: `nodes.<id>`, `edges.<id>`, or the name of a top-level workflow setting.
	Path    string          `json:"path" xml:"path" required:"true"`
	Section ConflictSection `json:"section" xml:"section" required:"true"`
	// Value at the common ancestor (merge base).
	BaseValue any `json:"base_value,omitempty" xml:"base_value,omitempty"`
	// Value on the source branch tip.
	SourceValue any `json:"source_value,omitempty" xml:"source_value,omitempty"`
	// Value on the target branch tip.
	TargetValue     any `json:"target_value,omitempty" xml:"target_value,omitempty"`
	extraProperties map[string]interface{}
}

func (f *FieldConflict) GetPath() (value string) {
	if f == nil {
		return
	}
	return f.Path
}

func (f *FieldConflict) GetSection() (value ConflictSection) {
	if f == nil {
		return
	}
	return f.Section
}

func (f *FieldConflict) GetBaseValue() (value any) {
	if f == nil {
		return
	}
	return f.BaseValue
}

func (f *FieldConflict) GetSourceValue() (value any) {
	if f == nil {
		return
	}
	return f.SourceValue
}

func (f *FieldConflict) GetTargetValue() (value any) {
	if f == nil {
		return
	}
	return f.TargetValue
}

func (f *FieldConflict) SetPath(path string) {
	f.Path = path
}

func (f *FieldConflict) SetSection(section ConflictSection) {
	f.Section = section
}

func (f *FieldConflict) SetBaseValue(baseValue any) {
	f.BaseValue = baseValue
}

func (f *FieldConflict) SetSourceValue(sourceValue any) {
	f.SourceValue = sourceValue
}

func (f *FieldConflict) SetTargetValue(targetValue any) {
	f.TargetValue = targetValue
}

func (f *FieldConflict) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FieldConflict) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FieldConflict to string"
	}
	return string(jsonData)
}

func (f *FieldConflict) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias FieldConflict
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = FieldConflict(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FieldConflict) MarshalJSON() ([]byte, error) {
	type embed FieldConflict
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
