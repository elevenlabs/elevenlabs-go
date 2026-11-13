package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TokenResponseModel struct {
	Token           string `json:"token" xml:"token" required:"true"`
	ConversationId  string `json:"conversation_id" xml:"conversation_id" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TokenResponseModel) GetToken() (value string) {
	if t == nil {
		return
	}
	return t.Token
}

func (t *TokenResponseModel) GetConversationId() (value string) {
	if t == nil {
		return
	}
	return t.ConversationId
}

func (t *TokenResponseModel) SetToken(token string) {
	t.Token = token
}

func (t *TokenResponseModel) SetConversationId(conversationId string) {
	t.ConversationId = conversationId
}

func (t *TokenResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TokenResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TokenResponseModel to string"
	}
	return string(jsonData)
}

func (t *TokenResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TokenResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TokenResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TokenResponseModel) MarshalJSON() ([]byte, error) {
	type embed TokenResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
