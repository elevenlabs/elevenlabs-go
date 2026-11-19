package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeechEngineConfig struct {
	// The WebSocket URL for the transcript server
	WsUrl string `json:"ws_url" xml:"ws_url" required:"true"`
	// Headers to include in the WebSocket connection request
	RequestHeaders  map[string]*SpeechEngineConfigRequestHeadersValue `json:"request_headers,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SpeechEngineConfig) GetWsUrl() (value string) {
	if s == nil {
		return
	}
	return s.WsUrl
}

func (s *SpeechEngineConfig) GetRequestHeaders() (value map[string]*SpeechEngineConfigRequestHeadersValue) {
	if s == nil {
		return
	}
	return s.RequestHeaders
}

func (s *SpeechEngineConfig) SetWsUrl(wsUrl string) {
	s.WsUrl = wsUrl
}

func (s *SpeechEngineConfig) SetRequestHeaders(requestHeaders map[string]*SpeechEngineConfigRequestHeadersValue) {
	s.RequestHeaders = requestHeaders
}

func (s *SpeechEngineConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeechEngineConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeechEngineConfig to string"
	}
	return string(jsonData)
}

func (s *SpeechEngineConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SpeechEngineConfig
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SpeechEngineConfig(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeechEngineConfig) MarshalJSON() ([]byte, error) {
	type embed SpeechEngineConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
