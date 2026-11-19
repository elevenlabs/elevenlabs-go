package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WordTimestamp struct {
	Word            string `json:"word" xml:"word" required:"true"`
	StartMs         int    `json:"start_ms" xml:"start_ms" required:"true"`
	EndMs           int    `json:"end_ms" xml:"end_ms" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WordTimestamp) GetWord() (value string) {
	if w == nil {
		return
	}
	return w.Word
}

func (w *WordTimestamp) GetStartMs() (value int) {
	if w == nil {
		return
	}
	return w.StartMs
}

func (w *WordTimestamp) GetEndMs() (value int) {
	if w == nil {
		return
	}
	return w.EndMs
}

func (w *WordTimestamp) SetWord(word string) {
	w.Word = word
}

func (w *WordTimestamp) SetStartMs(startMs int) {
	w.StartMs = startMs
}

func (w *WordTimestamp) SetEndMs(endMs int) {
	w.EndMs = endMs
}

func (w *WordTimestamp) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WordTimestamp) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WordTimestamp to string"
	}
	return string(jsonData)
}

func (w *WordTimestamp) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WordTimestamp
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WordTimestamp(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WordTimestamp) MarshalJSON() ([]byte, error) {
	type embed WordTimestamp
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
