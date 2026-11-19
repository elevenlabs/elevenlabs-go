package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Base64FileSource struct {
	type_ string
	// The name of the file, with its extension.
	Filename  string            `json:"filename" xml:"filename" required:"true" minLength:"1"`
	MediaType FileInputMimeType `json:"media_type" xml:"media_type" required:"true"`
	// The content of the file, base64-encoded.
	Data            string `json:"data" xml:"data" required:"true"`
	extraProperties map[string]interface{}
}

func (b *Base64FileSource) GetFilename() (value string) {
	if b == nil {
		return
	}
	return b.Filename
}

func (b *Base64FileSource) GetMediaType() (value FileInputMimeType) {
	if b == nil {
		return
	}
	return b.MediaType
}

func (b *Base64FileSource) GetData() (value string) {
	if b == nil {
		return
	}
	return b.Data
}

func (b *Base64FileSource) SetFilename(filename string) {
	b.Filename = filename
}

func (b *Base64FileSource) SetMediaType(mediaType FileInputMimeType) {
	b.MediaType = mediaType
}

func (b *Base64FileSource) SetData(data string) {
	b.Data = data
}

func (b *Base64FileSource) Type() string {
	return b.type_
}

func (b *Base64FileSource) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b Base64FileSource) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: Base64FileSource to string"
	}
	return string(jsonData)
}

func (b *Base64FileSource) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias Base64FileSource
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = Base64FileSource(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "base64" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", b, "base64", *literals.Type)
	}
	b.type_ = "base64"
	extraProperties, err := internal.ExtractExtraProperties(data, *b, "type")
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *Base64FileSource) MarshalJSON() ([]byte, error) {
	type embed Base64FileSource
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*b),
		Type:  "base64",
	}
	return json.Marshal(marshaler)
}
