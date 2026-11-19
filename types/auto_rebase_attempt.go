package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AutoRebaseAttempt struct {
	Outcome AutoRebaseOutcome `json:"outcome" xml:"outcome" required:"true"`
	// Unix seconds when the attempt ran
	AttemptedAt     int `json:"attempted_at" xml:"attempted_at" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AutoRebaseAttempt) GetOutcome() (value AutoRebaseOutcome) {
	if a == nil {
		return
	}
	return a.Outcome
}

func (a *AutoRebaseAttempt) GetAttemptedAt() (value int) {
	if a == nil {
		return
	}
	return a.AttemptedAt
}

func (a *AutoRebaseAttempt) SetOutcome(outcome AutoRebaseOutcome) {
	a.Outcome = outcome
}

func (a *AutoRebaseAttempt) SetAttemptedAt(attemptedAt int) {
	a.AttemptedAt = attemptedAt
}

func (a *AutoRebaseAttempt) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AutoRebaseAttempt) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AutoRebaseAttempt to string"
	}
	return string(jsonData)
}

func (a *AutoRebaseAttempt) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AutoRebaseAttempt
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AutoRebaseAttempt(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AutoRebaseAttempt) MarshalJSON() ([]byte, error) {
	type embed AutoRebaseAttempt
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
