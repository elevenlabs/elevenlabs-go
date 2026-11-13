package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CallingIntervalPolicy struct {
	// Local start time, as `HH:MM`.
	Start string `json:"start" xml:"start" required:"true" pattern:"^([01]\\d|2[0-3]):[0-5]\\d$"`
	// Local end time, as `HH:MM`; `24:00` ends at midnight.
	End             string `json:"end" xml:"end" required:"true" pattern:"^(([01]\\d|2[0-3]):[0-5]\\d|24:00)$"`
	extraProperties map[string]interface{}
}

func (c *CallingIntervalPolicy) GetStart() (value string) {
	if c == nil {
		return
	}
	return c.Start
}

func (c *CallingIntervalPolicy) GetEnd() (value string) {
	if c == nil {
		return
	}
	return c.End
}

func (c *CallingIntervalPolicy) SetStart(start string) {
	c.Start = start
}

func (c *CallingIntervalPolicy) SetEnd(end string) {
	c.End = end
}

func (c *CallingIntervalPolicy) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CallingIntervalPolicy) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CallingIntervalPolicy to string"
	}
	return string(jsonData)
}

func (c *CallingIntervalPolicy) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CallingIntervalPolicy
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CallingIntervalPolicy(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CallingIntervalPolicy) MarshalJSON() ([]byte, error) {
	type embed CallingIntervalPolicy
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
