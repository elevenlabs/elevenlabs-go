package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type FileContentPart struct {
	type_           string
	Source          *FileContentPartSource `json:"source" xml:"source" required:"true"`
	extraProperties map[string]interface{}
}

func (f *FileContentPart) GetSource() (value *FileContentPartSource) {
	if f == nil {
		return
	}
	return f.Source
}

func (f *FileContentPart) SetSource(source *FileContentPartSource) {
	f.Source = source
}

func (f *FileContentPart) Type() string {
	return f.type_
}

func (f *FileContentPart) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FileContentPart) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FileContentPart to string"
	}
	return string(jsonData)
}

func (f *FileContentPart) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias FileContentPart
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = FileContentPart(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "file" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", f, "file", *literals.Type)
	}
	f.type_ = "file"
	extraProperties, err := internal.ExtractExtraProperties(data, *f, "type")
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FileContentPart) MarshalJSON() ([]byte, error) {
	type embed FileContentPart
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*f),
		Type:  "file",
	}
	return json.Marshal(marshaler)
}
