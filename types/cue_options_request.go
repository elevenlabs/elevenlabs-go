package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CueOptionsRequest struct {
	// Minimum duration of each cue in milliseconds.
	MinDurationMs *int `json:"min_duration_ms,omitempty" xml:"min_duration_ms,omitempty" max:"2000"`
	// Maximum duration of each cue in milliseconds.
	MaxDurationMs *int `json:"max_duration_ms,omitempty" xml:"max_duration_ms,omitempty" min:"4000"`
	// Maximum number of lines per cue.
	MaxLinesPerCue *int `json:"max_lines_per_cue,omitempty" xml:"max_lines_per_cue,omitempty" min:"1" max:"3"`
	// Maximum number of characters per line in a cue.
	MaxCharsPerLine *int `json:"max_chars_per_line,omitempty" xml:"max_chars_per_line,omitempty" min:"16" max:"50"`
	// Maximum characters per second reading speed. If not set, no reading speed limit is applied.
	MaxCharsPerS *int `json:"max_chars_per_s,omitempty" xml:"max_chars_per_s,omitempty" min:"15" max:"30"`
	// Minimum gap between consecutive cues in frames. If not set, no minimum gap is enforced.
	MinGapBetweenCuesFrames *int `json:"min_gap_between_cues_frames,omitempty" xml:"min_gap_between_cues_frames,omitempty" min:"0" max:"6"`
	extraProperties         map[string]interface{}
}

func (c *CueOptionsRequest) GetMinDurationMs() (value *int) {
	if c == nil {
		return
	}
	return c.MinDurationMs
}

func (c *CueOptionsRequest) GetMaxDurationMs() (value *int) {
	if c == nil {
		return
	}
	return c.MaxDurationMs
}

func (c *CueOptionsRequest) GetMaxLinesPerCue() (value *int) {
	if c == nil {
		return
	}
	return c.MaxLinesPerCue
}

func (c *CueOptionsRequest) GetMaxCharsPerLine() (value *int) {
	if c == nil {
		return
	}
	return c.MaxCharsPerLine
}

func (c *CueOptionsRequest) GetMaxCharsPerS() (value *int) {
	if c == nil {
		return
	}
	return c.MaxCharsPerS
}

func (c *CueOptionsRequest) GetMinGapBetweenCuesFrames() (value *int) {
	if c == nil {
		return
	}
	return c.MinGapBetweenCuesFrames
}

func (c *CueOptionsRequest) SetMinDurationMs(minDurationMs *int) {
	c.MinDurationMs = minDurationMs
}

func (c *CueOptionsRequest) SetMaxDurationMs(maxDurationMs *int) {
	c.MaxDurationMs = maxDurationMs
}

func (c *CueOptionsRequest) SetMaxLinesPerCue(maxLinesPerCue *int) {
	c.MaxLinesPerCue = maxLinesPerCue
}

func (c *CueOptionsRequest) SetMaxCharsPerLine(maxCharsPerLine *int) {
	c.MaxCharsPerLine = maxCharsPerLine
}

func (c *CueOptionsRequest) SetMaxCharsPerS(maxCharsPerS *int) {
	c.MaxCharsPerS = maxCharsPerS
}

func (c *CueOptionsRequest) SetMinGapBetweenCuesFrames(minGapBetweenCuesFrames *int) {
	c.MinGapBetweenCuesFrames = minGapBetweenCuesFrames
}

func (c *CueOptionsRequest) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CueOptionsRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CueOptionsRequest to string"
	}
	return string(jsonData)
}

func (c *CueOptionsRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CueOptionsRequest) MarshalJSON() ([]byte, error) {
	type embed CueOptionsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
