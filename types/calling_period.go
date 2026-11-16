package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CallingPeriod struct {
	// Days the intervals apply on, from `0` (Monday) to `6` (Sunday).
	Weekdays []int `json:"weekdays" xml:"weekdays" required:"true" minItems:"1"`
	// Windows on each of those days. They may not overlap on any day.
	Intervals       []*CallingIntervalPolicy `json:"intervals" xml:"intervals" required:"true" minItems:"1"`
	extraProperties map[string]interface{}
}

func (c *CallingPeriod) GetWeekdays() (value []int) {
	if c == nil {
		return
	}
	return c.Weekdays
}

func (c *CallingPeriod) GetIntervals() (value []*CallingIntervalPolicy) {
	if c == nil {
		return
	}
	return c.Intervals
}

func (c *CallingPeriod) SetWeekdays(weekdays []int) {
	c.Weekdays = weekdays
}

func (c *CallingPeriod) SetIntervals(intervals []*CallingIntervalPolicy) {
	c.Intervals = intervals
}

func (c *CallingPeriod) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CallingPeriod) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CallingPeriod to string"
	}
	return string(jsonData)
}

func (c *CallingPeriod) UnmarshalJSON(data []byte) error {
	type unmarshaler CallingPeriod
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CallingPeriod(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CallingPeriod) MarshalJSON() ([]byte, error) {
	type embed CallingPeriod
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
