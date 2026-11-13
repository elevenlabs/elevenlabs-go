package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Media passed inline as base64. Stored as an ephemeral asset with no
// guaranteed retention. To keep an input and reuse it, upload via
// “POST /v1/assets“ and pass an “asset“ reference instead.
type InlineBase64Reference struct {
	type_ string
	// The media file's bytes, base64-encoded (standard alphabet). Up to 25MB decoded.
	ContentBase64 string `json:"content_base64" xml:"content_base64" required:"true" maxLength:"34952536"`
	// The MIME type of the encoded media.
	MimeType        InlineBase64ReferenceMimeType `json:"mime_type" xml:"mime_type" required:"true"`
	extraProperties map[string]interface{}
}

func (i *InlineBase64Reference) GetContentBase64() (value string) {
	if i == nil {
		return
	}
	return i.ContentBase64
}

func (i *InlineBase64Reference) GetMimeType() (value InlineBase64ReferenceMimeType) {
	if i == nil {
		return
	}
	return i.MimeType
}

func (i *InlineBase64Reference) SetContentBase64(contentBase64 string) {
	i.ContentBase64 = contentBase64
}

func (i *InlineBase64Reference) SetMimeType(mimeType InlineBase64ReferenceMimeType) {
	i.MimeType = mimeType
}

func (i *InlineBase64Reference) Type() string {
	return i.type_
}

func (i *InlineBase64Reference) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InlineBase64Reference) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InlineBase64Reference to string"
	}
	return string(jsonData)
}

func (i *InlineBase64Reference) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias InlineBase64Reference
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = InlineBase64Reference(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "inline_base64" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", i, "inline_base64", *literals.Type)
	}
	i.type_ = "inline_base64"
	extraProperties, err := internal.ExtractExtraProperties(data, *i, "type")
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InlineBase64Reference) MarshalJSON() ([]byte, error) {
	type embed InlineBase64Reference
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*i),
		Type:  "inline_base64",
	}
	return json.Marshal(marshaler)
}
