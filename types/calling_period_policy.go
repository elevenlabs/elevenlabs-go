package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CallingPeriodPolicy struct {
	Weekdays []int `json:"weekdays" xml:"weekdays" required:"true" minItems:"1"`
	// Windows on each of those days. They may not overlap on any day.
	Intervals       []*CallingIntervalPolicy `json:"intervals" xml:"intervals" required:"true" minItems:"1"`
	extraProperties map[string]interface{}
}

func (c *CallingPeriodPolicy) GetWeekdays() (value []int) {
	if c == nil {
		return
	}
	return c.Weekdays
}

func (c *CallingPeriodPolicy) GetIntervals() (value []*CallingIntervalPolicy) {
	if c == nil {
		return
	}
	return c.Intervals
}

func (c *CallingPeriodPolicy) SetWeekdays(weekdays []int) {
	c.Weekdays = weekdays
}

func (c *CallingPeriodPolicy) SetIntervals(intervals []*CallingIntervalPolicy) {
	c.Intervals = intervals
}

func (c *CallingPeriodPolicy) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CallingPeriodPolicy) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CallingPeriodPolicy to string"
	}
	return string(jsonData)
}

func (c *CallingPeriodPolicy) UnmarshalJSON(data []byte) error {
	type unmarshaler CallingPeriodPolicy
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CallingPeriodPolicy(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CallingPeriodPolicy) MarshalJSON() ([]byte, error) {
	type embed CallingPeriodPolicy
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
