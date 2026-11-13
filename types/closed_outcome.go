package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ClosedOutcome struct {
	Reason          MergeProposalCloseReason `json:"reason" xml:"reason" required:"true"`
	At              int                      `json:"at" xml:"at" required:"true"`
	ByUserId        *string                  `json:"by_user_id" xml:"by_user_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ClosedOutcome) GetReason() (value MergeProposalCloseReason) {
	if c == nil {
		return
	}
	return c.Reason
}

func (c *ClosedOutcome) GetAt() (value int) {
	if c == nil {
		return
	}
	return c.At
}

func (c *ClosedOutcome) GetByUserId() (value *string) {
	if c == nil {
		return
	}
	return c.ByUserId
}

func (c *ClosedOutcome) SetReason(reason MergeProposalCloseReason) {
	c.Reason = reason
}

func (c *ClosedOutcome) SetAt(at int) {
	c.At = at
}

func (c *ClosedOutcome) SetByUserId(byUserId *string) {
	c.ByUserId = byUserId
}

func (c *ClosedOutcome) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClosedOutcome) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClosedOutcome to string"
	}
	return string(jsonData)
}

func (c *ClosedOutcome) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "status")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClosedOutcome) MarshalJSON() ([]byte, error) {
	type embed ClosedOutcome
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
