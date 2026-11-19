package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TwilioOutboundCallResponse struct {
	Success         bool    `json:"success" xml:"success" required:"true"`
	Message         string  `json:"message" xml:"message" required:"true"`
	ConversationId  *string `json:"conversation_id" xml:"conversation_id" required:"true"`
	CallSid         *string `json:"callSid" xml:"callSid" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TwilioOutboundCallResponse) GetSuccess() (value bool) {
	if t == nil {
		return
	}
	return t.Success
}

func (t *TwilioOutboundCallResponse) GetMessage() (value string) {
	if t == nil {
		return
	}
	return t.Message
}

func (t *TwilioOutboundCallResponse) GetConversationId() (value *string) {
	if t == nil {
		return
	}
	return t.ConversationId
}

func (t *TwilioOutboundCallResponse) GetCallSid() (value *string) {
	if t == nil {
		return
	}
	return t.CallSid
}

func (t *TwilioOutboundCallResponse) SetSuccess(success bool) {
	t.Success = success
}

func (t *TwilioOutboundCallResponse) SetMessage(message string) {
	t.Message = message
}

func (t *TwilioOutboundCallResponse) SetConversationId(conversationId *string) {
	t.ConversationId = conversationId
}

func (t *TwilioOutboundCallResponse) SetCallSid(callSid *string) {
	t.CallSid = callSid
}

func (t *TwilioOutboundCallResponse) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TwilioOutboundCallResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioOutboundCallResponse to string"
	}
	return string(jsonData)
}

func (t *TwilioOutboundCallResponse) UnmarshalJSON(data []byte) error {
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

func (t *TwilioOutboundCallResponse) MarshalJSON() ([]byte, error) {
	type embed TwilioOutboundCallResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
