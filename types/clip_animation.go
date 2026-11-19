package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ClipAnimation struct {
	EnterEffect     *ClipAnimationEnterEffect `json:"enter_effect,omitempty" xml:"enter_effect,omitempty"`
	EnterDurationMs *int                      `json:"enter_duration_ms,omitempty" xml:"enter_duration_ms,omitempty"`
	ExitEffect      *ClipAnimationExitEffect  `json:"exit_effect,omitempty" xml:"exit_effect,omitempty"`
	ExitDurationMs  *int                      `json:"exit_duration_ms,omitempty" xml:"exit_duration_ms,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ClipAnimation) GetEnterEffect() (value *ClipAnimationEnterEffect) {
	if c == nil {
		return
	}
	return c.EnterEffect
}

func (c *ClipAnimation) GetEnterDurationMs() (value *int) {
	if c == nil {
		return
	}
	return c.EnterDurationMs
}

func (c *ClipAnimation) GetExitEffect() (value *ClipAnimationExitEffect) {
	if c == nil {
		return
	}
	return c.ExitEffect
}

func (c *ClipAnimation) GetExitDurationMs() (value *int) {
	if c == nil {
		return
	}
	return c.ExitDurationMs
}

func (c *ClipAnimation) SetEnterEffect(enterEffect *ClipAnimationEnterEffect) {
	c.EnterEffect = enterEffect
}

func (c *ClipAnimation) SetEnterDurationMs(enterDurationMs *int) {
	c.EnterDurationMs = enterDurationMs
}

func (c *ClipAnimation) SetExitEffect(exitEffect *ClipAnimationExitEffect) {
	c.ExitEffect = exitEffect
}

func (c *ClipAnimation) SetExitDurationMs(exitDurationMs *int) {
	c.ExitDurationMs = exitDurationMs
}

func (c *ClipAnimation) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClipAnimation) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClipAnimation to string"
	}
	return string(jsonData)
}

func (c *ClipAnimation) UnmarshalJSON(data []byte) error {
	type unmarshaler ClipAnimation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ClipAnimation(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClipAnimation) MarshalJSON() ([]byte, error) {
	type embed ClipAnimation
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
