package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// With diarize, a committed transcript's words with their final speakers, sent after
// its committed_transcript (which never waits for them), in commit order.
type DiarizedTranscript struct {
	MessageType *string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	// The segment_index of the committed_transcript the words are of.
	SegmentIndex int `json:"segment_index" xml:"segment_index" required:"true"`
	// The committed transcript text the words are of.
	Text string `json:"text" xml:"text" required:"true"`
	// The words, as in batch diarization, each with its speaker_id, null where the speaker could not be identified.
	Words           []*DiarizedTranscriptWordsItem `json:"words" xml:"words" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DiarizedTranscript) GetMessageType() (value *string) {
	if d == nil {
		return
	}
	return d.MessageType
}

func (d *DiarizedTranscript) GetSegmentIndex() (value int) {
	if d == nil {
		return
	}
	return d.SegmentIndex
}

func (d *DiarizedTranscript) GetText() (value string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DiarizedTranscript) GetWords() (value []*DiarizedTranscriptWordsItem) {
	if d == nil {
		return
	}
	return d.Words
}

func (d *DiarizedTranscript) SetMessageType(messageType *string) {
	d.MessageType = messageType
}

func (d *DiarizedTranscript) SetSegmentIndex(segmentIndex int) {
	d.SegmentIndex = segmentIndex
}

func (d *DiarizedTranscript) SetText(text string) {
	d.Text = text
}

func (d *DiarizedTranscript) SetWords(words []*DiarizedTranscriptWordsItem) {
	d.Words = words
}

func (d *DiarizedTranscript) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DiarizedTranscript) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DiarizedTranscript to string"
	}
	return string(jsonData)
}

func (d *DiarizedTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DiarizedTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DiarizedTranscript(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DiarizedTranscript) MarshalJSON() ([]byte, error) {
	type embed DiarizedTranscript
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
