package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioRefChunk struct {
	// The ID of the song to source the chunk from. You can find the song ID in the response headers when you generate a song.
	SongId          string     `json:"song_id" xml:"song_id" required:"true" maxLength:"100" minLength:"1"`
	Range           *TimeRange `json:"range" xml:"range" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AudioRefChunk) GetSongId() (value string) {
	if a == nil {
		return
	}
	return a.SongId
}

func (a *AudioRefChunk) GetRange() (value *TimeRange) {
	if a == nil {
		return
	}
	return a.Range
}

func (a *AudioRefChunk) SetSongId(songId string) {
	a.SongId = songId
}

func (a *AudioRefChunk) SetRange(range_ *TimeRange) {
	a.Range = range_
}

func (a *AudioRefChunk) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioRefChunk) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioRefChunk to string"
	}
	return string(jsonData)
}

func (a *AudioRefChunk) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AudioRefChunk
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AudioRefChunk(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioRefChunk) MarshalJSON() ([]byte, error) {
	type embed AudioRefChunk
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
