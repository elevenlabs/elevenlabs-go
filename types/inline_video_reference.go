package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A video passed inline as base64. The video is stored as an ephemeral
// asset with no guaranteed retention: it may be deleted at any time after
// the generation completes. To keep an input and reuse it across
// generations, upload it via the assets API (`POST /v1/assets`) and pass
// an `asset` reference instead.
type InlineVideoReference struct {
	type_ string
	// The media file's bytes, base64-encoded (standard alphabet). Up to 25MB decoded.
	ContentBase64 string `json:"content_base64" xml:"content_base64" required:"true" maxLength:"34952536"`
	// The MIME type of the encoded video.
	MimeType        InlineVideoReferenceMimeType `json:"mime_type" xml:"mime_type" required:"true"`
	extraProperties map[string]interface{}
}

func (i *InlineVideoReference) GetContentBase64() (value string) {
	if i == nil {
		return
	}
	return i.ContentBase64
}

func (i *InlineVideoReference) GetMimeType() (value InlineVideoReferenceMimeType) {
	if i == nil {
		return
	}
	return i.MimeType
}

func (i *InlineVideoReference) SetContentBase64(contentBase64 string) {
	i.ContentBase64 = contentBase64
}

func (i *InlineVideoReference) SetMimeType(mimeType InlineVideoReferenceMimeType) {
	i.MimeType = mimeType
}

func (i *InlineVideoReference) Type() string {
	return i.type_
}

func (i *InlineVideoReference) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InlineVideoReference) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InlineVideoReference to string"
	}
	return string(jsonData)
}

func (i *InlineVideoReference) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias InlineVideoReference
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = InlineVideoReference(tmp)
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

func (i *InlineVideoReference) MarshalJSON() ([]byte, error) {
	type embed InlineVideoReference
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*i),
		Type:  "inline_base64",
	}
	return json.Marshal(marshaler)
}
