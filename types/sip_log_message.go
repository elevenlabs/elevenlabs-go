package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SipLogMessage struct {
	CallId             string                 `json:"call_id" xml:"call_id" required:"true"`
	PhoneNumbers       []string               `json:"phone_numbers" xml:"phone_numbers" required:"true"`
	LocalAddress       string                 `json:"local_address" xml:"local_address" required:"true"`
	RemoteAddress      string                 `json:"remote_address" xml:"remote_address" required:"true"`
	Transport          string                 `json:"transport" xml:"transport" required:"true"`
	RawMessage         string                 `json:"raw_message" xml:"raw_message" required:"true"`
	ErrorMessage       string                 `json:"error_message" xml:"error_message" required:"true"`
	Direction          SipLogMessageDirection `json:"direction" xml:"direction" required:"true"`
	CreatedAtUnixMicro int                    `json:"created_at_unix_micro" xml:"created_at_unix_micro" required:"true"`
	extraProperties    map[string]interface{}
}

func (s *SipLogMessage) GetCallId() (value string) {
	if s == nil {
		return
	}
	return s.CallId
}

func (s *SipLogMessage) GetPhoneNumbers() (value []string) {
	if s == nil {
		return
	}
	return s.PhoneNumbers
}

func (s *SipLogMessage) GetLocalAddress() (value string) {
	if s == nil {
		return
	}
	return s.LocalAddress
}

func (s *SipLogMessage) GetRemoteAddress() (value string) {
	if s == nil {
		return
	}
	return s.RemoteAddress
}

func (s *SipLogMessage) GetTransport() (value string) {
	if s == nil {
		return
	}
	return s.Transport
}

func (s *SipLogMessage) GetRawMessage() (value string) {
	if s == nil {
		return
	}
	return s.RawMessage
}

func (s *SipLogMessage) GetErrorMessage() (value string) {
	if s == nil {
		return
	}
	return s.ErrorMessage
}

func (s *SipLogMessage) GetDirection() (value SipLogMessageDirection) {
	if s == nil {
		return
	}
	return s.Direction
}

func (s *SipLogMessage) GetCreatedAtUnixMicro() (value int) {
	if s == nil {
		return
	}
	return s.CreatedAtUnixMicro
}

func (s *SipLogMessage) SetCallId(callId string) {
	s.CallId = callId
}

func (s *SipLogMessage) SetPhoneNumbers(phoneNumbers []string) {
	s.PhoneNumbers = phoneNumbers
}

func (s *SipLogMessage) SetLocalAddress(localAddress string) {
	s.LocalAddress = localAddress
}

func (s *SipLogMessage) SetRemoteAddress(remoteAddress string) {
	s.RemoteAddress = remoteAddress
}

func (s *SipLogMessage) SetTransport(transport string) {
	s.Transport = transport
}

func (s *SipLogMessage) SetRawMessage(rawMessage string) {
	s.RawMessage = rawMessage
}

func (s *SipLogMessage) SetErrorMessage(errorMessage string) {
	s.ErrorMessage = errorMessage
}

func (s *SipLogMessage) SetDirection(direction SipLogMessageDirection) {
	s.Direction = direction
}

func (s *SipLogMessage) SetCreatedAtUnixMicro(createdAtUnixMicro int) {
	s.CreatedAtUnixMicro = createdAtUnixMicro
}

func (s *SipLogMessage) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SipLogMessage) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SipLogMessage to string"
	}
	return string(jsonData)
}

func (s *SipLogMessage) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SipLogMessage
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SipLogMessage(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SipLogMessage) MarshalJSON() ([]byte, error) {
	type embed SipLogMessage
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
