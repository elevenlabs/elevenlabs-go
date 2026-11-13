package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SingleUseTokenResponseModel struct {
	// A time bound single use token that expires after 15 minutes. Will be consumed on use.
	Token           string `json:"token" xml:"token" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SingleUseTokenResponseModel) GetToken() (value string) {
	if s == nil {
		return
	}
	return s.Token
}

func (s *SingleUseTokenResponseModel) SetToken(token string) {
	s.Token = token
}

func (s *SingleUseTokenResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SingleUseTokenResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SingleUseTokenResponseModel to string"
	}
	return string(jsonData)
}

func (s *SingleUseTokenResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SingleUseTokenResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SingleUseTokenResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SingleUseTokenResponseModel) MarshalJSON() ([]byte, error) {
	type embed SingleUseTokenResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
