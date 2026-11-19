package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DetectedEntity struct {
	// The text that was identified as an entity.
	Text string `json:"text" xml:"text" required:"true"`
	// The type of entity detected (e.g., 'credit_card', 'email_address', 'person_name').
	EntityType string `json:"entity_type" xml:"entity_type" required:"true"`
	// Start character position in the transcript text.
	StartChar int `json:"start_char" xml:"start_char" required:"true"`
	// End character position in the transcript text.
	EndChar         int `json:"end_char" xml:"end_char" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DetectedEntity) GetText() (value string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DetectedEntity) GetEntityType() (value string) {
	if d == nil {
		return
	}
	return d.EntityType
}

func (d *DetectedEntity) GetStartChar() (value int) {
	if d == nil {
		return
	}
	return d.StartChar
}

func (d *DetectedEntity) GetEndChar() (value int) {
	if d == nil {
		return
	}
	return d.EndChar
}

func (d *DetectedEntity) SetText(text string) {
	d.Text = text
}

func (d *DetectedEntity) SetEntityType(entityType string) {
	d.EntityType = entityType
}

func (d *DetectedEntity) SetStartChar(startChar int) {
	d.StartChar = startChar
}

func (d *DetectedEntity) SetEndChar(endChar int) {
	d.EndChar = endChar
}

func (d *DetectedEntity) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DetectedEntity) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DetectedEntity to string"
	}
	return string(jsonData)
}

func (d *DetectedEntity) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DetectedEntity
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DetectedEntity(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DetectedEntity) MarshalJSON() ([]byte, error) {
	type embed DetectedEntity
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
