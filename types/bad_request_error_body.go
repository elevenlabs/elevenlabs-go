package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type BadRequestErrorBody struct {
	Error           *string `json:"error,omitempty" xml:"error,omitempty"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
	extraProperties map[string]interface{}
}

func (b *BadRequestErrorBody) GetError() (value *string) {
	if b == nil {
		return
	}
	return b.Error
}

func (b *BadRequestErrorBody) GetMessage() (value *string) {
	if b == nil {
		return
	}
	return b.Message
}

func (b *BadRequestErrorBody) SetError(error *string) {
	b.Error = error
}

func (b *BadRequestErrorBody) SetMessage(message *string) {
	b.Message = message
}

func (b *BadRequestErrorBody) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BadRequestErrorBody) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BadRequestErrorBody to string"
	}
	return string(jsonData)
}

func (b *BadRequestErrorBody) UnmarshalJSON(data []byte) error {
	type unmarshaler BadRequestErrorBody
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BadRequestErrorBody(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BadRequestErrorBody) MarshalJSON() ([]byte, error) {
	type embed BadRequestErrorBody
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
