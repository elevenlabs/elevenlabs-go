package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"io"
)

// Raw multipart/mixed response data that clients must parse. Contains two parts separated by boundary: 1) application/json part with composition_plan and song_metadata 2) audio/* part with binary audio data and Content-Disposition header
type ComposeDetailedMusicResponse struct {
	CompositionPlan map[string]any `json:"composition_plan,omitempty"`
	SongMetadata    map[string]any `json:"song_metadata,omitempty"`
	Audio           io.Reader      `json:"audio,omitempty" xml:"audio,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ComposeDetailedMusicResponse) GetCompositionPlan() (value map[string]any) {
	if c == nil {
		return
	}
	return c.CompositionPlan
}

func (c *ComposeDetailedMusicResponse) GetSongMetadata() (value map[string]any) {
	if c == nil {
		return
	}
	return c.SongMetadata
}

func (c *ComposeDetailedMusicResponse) GetAudio() (value io.Reader) {
	if c == nil {
		return
	}
	return c.Audio
}

func (c *ComposeDetailedMusicResponse) SetCompositionPlan(compositionPlan map[string]any) {
	c.CompositionPlan = compositionPlan
}

func (c *ComposeDetailedMusicResponse) SetSongMetadata(songMetadata map[string]any) {
	c.SongMetadata = songMetadata
}

func (c *ComposeDetailedMusicResponse) SetAudio(audio io.Reader) {
	c.Audio = audio
}

func (c *ComposeDetailedMusicResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ComposeDetailedMusicResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ComposeDetailedMusicResponse to string"
	}
	return string(jsonData)
}

func (c *ComposeDetailedMusicResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ComposeDetailedMusicResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ComposeDetailedMusicResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ComposeDetailedMusicResponse) MarshalJSON() ([]byte, error) {
	type embed ComposeDetailedMusicResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
