package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationSignedUrlResponseModel struct {
	SignedUrl       string `json:"signed_url" xml:"signed_url" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationSignedUrlResponseModel) GetSignedUrl() (value string) {
	if c == nil {
		return
	}
	return c.SignedUrl
}

func (c *ConversationSignedUrlResponseModel) SetSignedUrl(signedUrl string) {
	c.SignedUrl = signedUrl
}

func (c *ConversationSignedUrlResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationSignedUrlResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationSignedUrlResponseModel to string"
	}
	return string(jsonData)
}

func (c *ConversationSignedUrlResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationSignedUrlResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationSignedUrlResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationSignedUrlResponseModel) MarshalJSON() ([]byte, error) {
	type embed ConversationSignedUrlResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
