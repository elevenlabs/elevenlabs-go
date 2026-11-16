package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RetryPolicy struct {
	// How long to wait before a short retry.
	RetryDelaySeconds *int `json:"retry_delay_seconds,omitempty" xml:"retry_delay_seconds,omitempty"`
	// Retries a calling window allows after a failed call. Once used up, the next retry waits for the following window, which starts afresh.
	MaxShortRetries *int `json:"max_short_retries,omitempty" xml:"max_short_retries,omitempty" min:"0" max:"30"`
	// The most calls a lead gets, counting the first.
	MaxTotalCalls *int `json:"max_total_calls,omitempty" xml:"max_total_calls,omitempty" min:"1" max:"30"`
	// The most callbacks to honor for one number.
	MaxCallbacks *int `json:"max_callbacks,omitempty" xml:"max_callbacks,omitempty" min:"0" max:"30"`
	// The most times a lead follows a request to call another number.
	MaxRedirects      *int               `json:"max_redirects,omitempty" xml:"max_redirects,omitempty" min:"0" max:"30"`
	IntervalPlacement *IntervalPlacement `json:"interval_placement,omitempty" xml:"interval_placement,omitempty"`
	// How far `rotate` placement moves each call. Only used by `rotate`.
	RotateStepSeconds *int `json:"rotate_step_seconds,omitempty" xml:"rotate_step_seconds,omitempty"`
	// Calls to one number that may fail in a row before the lead moves on to its next listed number, or ends when there is none. An answered call, or a redirect to a new number, restarts the count. Unset leaves only `max_total_calls`.
	MaxConsecutiveFailuresPerNumber *int `json:"max_consecutive_failures_per_number,omitempty" xml:"max_consecutive_failures_per_number,omitempty" min:"1" max:"30"`
	extraProperties                 map[string]interface{}
}

func (r *RetryPolicy) GetRetryDelaySeconds() (value *int) {
	if r == nil {
		return
	}
	return r.RetryDelaySeconds
}

func (r *RetryPolicy) GetMaxShortRetries() (value *int) {
	if r == nil {
		return
	}
	return r.MaxShortRetries
}

func (r *RetryPolicy) GetMaxTotalCalls() (value *int) {
	if r == nil {
		return
	}
	return r.MaxTotalCalls
}

func (r *RetryPolicy) GetMaxCallbacks() (value *int) {
	if r == nil {
		return
	}
	return r.MaxCallbacks
}

func (r *RetryPolicy) GetMaxRedirects() (value *int) {
	if r == nil {
		return
	}
	return r.MaxRedirects
}

func (r *RetryPolicy) GetIntervalPlacement() (value *IntervalPlacement) {
	if r == nil {
		return
	}
	return r.IntervalPlacement
}

func (r *RetryPolicy) GetRotateStepSeconds() (value *int) {
	if r == nil {
		return
	}
	return r.RotateStepSeconds
}

func (r *RetryPolicy) GetMaxConsecutiveFailuresPerNumber() (value *int) {
	if r == nil {
		return
	}
	return r.MaxConsecutiveFailuresPerNumber
}

func (r *RetryPolicy) SetRetryDelaySeconds(retryDelaySeconds *int) {
	r.RetryDelaySeconds = retryDelaySeconds
}

func (r *RetryPolicy) SetMaxShortRetries(maxShortRetries *int) {
	r.MaxShortRetries = maxShortRetries
}

func (r *RetryPolicy) SetMaxTotalCalls(maxTotalCalls *int) {
	r.MaxTotalCalls = maxTotalCalls
}

func (r *RetryPolicy) SetMaxCallbacks(maxCallbacks *int) {
	r.MaxCallbacks = maxCallbacks
}

func (r *RetryPolicy) SetMaxRedirects(maxRedirects *int) {
	r.MaxRedirects = maxRedirects
}

func (r *RetryPolicy) SetIntervalPlacement(intervalPlacement *IntervalPlacement) {
	r.IntervalPlacement = intervalPlacement
}

func (r *RetryPolicy) SetRotateStepSeconds(rotateStepSeconds *int) {
	r.RotateStepSeconds = rotateStepSeconds
}

func (r *RetryPolicy) SetMaxConsecutiveFailuresPerNumber(maxConsecutiveFailuresPerNumber *int) {
	r.MaxConsecutiveFailuresPerNumber = maxConsecutiveFailuresPerNumber
}

func (r *RetryPolicy) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RetryPolicy) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RetryPolicy to string"
	}
	return string(jsonData)
}

func (r *RetryPolicy) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RetryPolicy) MarshalJSON() ([]byte, error) {
	type embed RetryPolicy
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
