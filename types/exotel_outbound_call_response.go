package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ExotelOutboundCallResponse struct {
	Success        bool    `json:"success" xml:"success" required:"true"`
	Message        string  `json:"message" xml:"message" required:"true"`
	ConversationId *string `json:"conversation_id" xml:"conversation_id" required:"true"`
	// Exotel CallSid for the originated call
	CallSid         *string `json:"callSid" xml:"callSid" required:"true"`
	extraProperties map[string]interface{}
}

func (e *ExotelOutboundCallResponse) GetSuccess() (value bool) {
	if e == nil {
		return
	}
	return e.Success
}

func (e *ExotelOutboundCallResponse) GetMessage() (value string) {
	if e == nil {
		return
	}
	return e.Message
}

func (e *ExotelOutboundCallResponse) GetConversationId() (value *string) {
	if e == nil {
		return
	}
	return e.ConversationId
}

func (e *ExotelOutboundCallResponse) GetCallSid() (value *string) {
	if e == nil {
		return
	}
	return e.CallSid
}

func (e *ExotelOutboundCallResponse) SetSuccess(success bool) {
	e.Success = success
}

func (e *ExotelOutboundCallResponse) SetMessage(message string) {
	e.Message = message
}

func (e *ExotelOutboundCallResponse) SetConversationId(conversationId *string) {
	e.ConversationId = conversationId
}

func (e *ExotelOutboundCallResponse) SetCallSid(callSid *string) {
	e.CallSid = callSid
}

func (e *ExotelOutboundCallResponse) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ExotelOutboundCallResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExotelOutboundCallResponse to string"
	}
	return string(jsonData)
}

func (e *ExotelOutboundCallResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ExotelOutboundCallResponse) MarshalJSON() ([]byte, error) {
	type embed ExotelOutboundCallResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
