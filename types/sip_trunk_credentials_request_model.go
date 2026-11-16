package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipTrunkCredentialsRequestModel struct {
	// SIP trunk username
	Username string `json:"username" xml:"username" required:"true"`
	// SIP trunk password - if not specified, then remain unchanged
	Password        *string `json:"password,omitempty" xml:"password,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SipTrunkCredentialsRequestModel) GetUsername() (value string) {
	if s == nil {
		return
	}
	return s.Username
}

func (s *SipTrunkCredentialsRequestModel) GetPassword() (value *string) {
	if s == nil {
		return
	}
	return s.Password
}

func (s *SipTrunkCredentialsRequestModel) SetUsername(username string) {
	s.Username = username
}

func (s *SipTrunkCredentialsRequestModel) SetPassword(password *string) {
	s.Password = password
}

func (s *SipTrunkCredentialsRequestModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipTrunkCredentialsRequestModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipTrunkCredentialsRequestModel to string"
	}
	return string(jsonData)
}

func (s *SipTrunkCredentialsRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipTrunkCredentialsRequestModel) MarshalJSON() ([]byte, error) {
	type embed SipTrunkCredentialsRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
