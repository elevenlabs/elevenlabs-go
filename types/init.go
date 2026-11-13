package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Init struct {
	type_           string
	ConversationId  string `json:"conversation_id" xml:"conversation_id" required:"true"`
	extraProperties map[string]interface{}
}

func (i *Init) GetConversationId() (value string) {
	if i == nil {
		return
	}
	return i.ConversationId
}

func (i *Init) SetConversationId(conversationId string) {
	i.ConversationId = conversationId
}

func (i *Init) Type() string {
	return i.type_
}

func (i *Init) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i Init) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: Init to string"
	}
	return string(jsonData)
}

func (i *Init) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias Init
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = Init(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "init" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", i, "init", *literals.Type)
	}
	i.type_ = "init"
	extraProperties, err := internal.ExtractExtraProperties(data, *i, "type")
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *Init) MarshalJSON() ([]byte, error) {
	type embed Init
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*i),
		Type:  "init",
	}
	return json.Marshal(marshaler)
}
