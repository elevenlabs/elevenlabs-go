package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// The new revision after a source edit that returns no segment, such as a delete.
type DubbingTranscriptRevisionResponse struct {
	// The project's source-transcript revision after this edit.
	Revision        int `json:"revision" xml:"revision" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingTranscriptRevisionResponse) GetRevision() (value int) {
	if d == nil {
		return
	}
	return d.Revision
}

func (d *DubbingTranscriptRevisionResponse) SetRevision(revision int) {
	d.Revision = revision
}

func (d *DubbingTranscriptRevisionResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTranscriptRevisionResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTranscriptRevisionResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingTranscriptRevisionResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingTranscriptRevisionResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingTranscriptRevisionResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingTranscriptRevisionResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingTranscriptRevisionResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
