package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryErrorCommonModel struct {
	Code            int     `json:"code" xml:"code" required:"true"`
	Reason          *string `json:"reason,omitempty" xml:"reason,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationHistoryErrorCommonModel) GetCode() (value int) {
	if c == nil {
		return
	}
	return c.Code
}

func (c *ConversationHistoryErrorCommonModel) GetReason() (value *string) {
	if c == nil {
		return
	}
	return c.Reason
}

func (c *ConversationHistoryErrorCommonModel) SetCode(code int) {
	c.Code = code
}

func (c *ConversationHistoryErrorCommonModel) SetReason(reason *string) {
	c.Reason = reason
}

func (c *ConversationHistoryErrorCommonModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryErrorCommonModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryErrorCommonModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryErrorCommonModel) UnmarshalJSON(data []byte) error {
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

func (c *ConversationHistoryErrorCommonModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryErrorCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
