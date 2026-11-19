package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TeamInviteResult struct {
	Email           string            `json:"email" xml:"email" required:"true"`
	Outcome         TeamInviteOutcome `json:"outcome" xml:"outcome" required:"true"`
	ErrorStatus     *string           `json:"error_status,omitempty" xml:"error_status,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TeamInviteResult) GetEmail() (value string) {
	if t == nil {
		return
	}
	return t.Email
}

func (t *TeamInviteResult) GetOutcome() (value TeamInviteOutcome) {
	if t == nil {
		return
	}
	return t.Outcome
}

func (t *TeamInviteResult) GetErrorStatus() (value *string) {
	if t == nil {
		return
	}
	return t.ErrorStatus
}

func (t *TeamInviteResult) SetEmail(email string) {
	t.Email = email
}

func (t *TeamInviteResult) SetOutcome(outcome TeamInviteOutcome) {
	t.Outcome = outcome
}

func (t *TeamInviteResult) SetErrorStatus(errorStatus *string) {
	t.ErrorStatus = errorStatus
}

func (t *TeamInviteResult) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TeamInviteResult) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TeamInviteResult to string"
	}
	return string(jsonData)
}

func (t *TeamInviteResult) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TeamInviteResult) MarshalJSON() ([]byte, error) {
	type embed TeamInviteResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
