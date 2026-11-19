package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyCheckWatermarkV1ContentProvenanceWatermarkPost struct {
	// Audio file to check for an embedded watermark.
	File io.Reader `json:"file" xml:"file" required:"true"`
	// Watermark decoder version. Defaults to latest.
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (b *BodyCheckWatermarkV1ContentProvenanceWatermarkPost) GetFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodyCheckWatermarkV1ContentProvenanceWatermarkPost) GetVersion() (value *string) {
	if b == nil {
		return
	}
	return b.Version
}

func (b *BodyCheckWatermarkV1ContentProvenanceWatermarkPost) SetFile(file io.Reader) {
	b.File = file
}

func (b *BodyCheckWatermarkV1ContentProvenanceWatermarkPost) SetVersion(version *string) {
	b.Version = version
}

func (b *BodyCheckWatermarkV1ContentProvenanceWatermarkPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCheckWatermarkV1ContentProvenanceWatermarkPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCheckWatermarkV1ContentProvenanceWatermarkPost to string"
	}
	return string(jsonData)
}

func (b *BodyCheckWatermarkV1ContentProvenanceWatermarkPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCheckWatermarkV1ContentProvenanceWatermarkPost) MarshalJSON() ([]byte, error) {
	type embed BodyCheckWatermarkV1ContentProvenanceWatermarkPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
