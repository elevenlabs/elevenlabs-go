package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyCreateForcedAlignmentV1ForcedAlignmentPost struct {
	// The file to align. All major audio formats are supported. The file size must be less than 1GB.
	File io.Reader `json:"file" xml:"file" required:"true"`
	// The text to align with the audio. The input text can be in any format, however diarization is not supported at this time.
	Text string `json:"text" xml:"text" required:"true"`
}

func (b *BodyCreateForcedAlignmentV1ForcedAlignmentPost) GetFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodyCreateForcedAlignmentV1ForcedAlignmentPost) GetText() (value string) {
	if b == nil {
		return
	}
	return b.Text
}

func (b *BodyCreateForcedAlignmentV1ForcedAlignmentPost) SetFile(file io.Reader) {
	b.File = file
}

func (b *BodyCreateForcedAlignmentV1ForcedAlignmentPost) SetText(text string) {
	b.Text = text
}

func (b *BodyCreateForcedAlignmentV1ForcedAlignmentPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateForcedAlignmentV1ForcedAlignmentPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateForcedAlignmentV1ForcedAlignmentPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateForcedAlignmentV1ForcedAlignmentPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyCreateForcedAlignmentV1ForcedAlignmentPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyCreateForcedAlignmentV1ForcedAlignmentPost(tmp)
	return nil
}

func (b *BodyCreateForcedAlignmentV1ForcedAlignmentPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateForcedAlignmentV1ForcedAlignmentPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
