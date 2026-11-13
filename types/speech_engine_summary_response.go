package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeechEngineSummaryResponse struct {
	// The speech engine resource ID
	SpeechEngineId string `json:"speech_engine_id" xml:"speech_engine_id" required:"true"`
	// Human-readable name for the speech engine
	Name string `json:"name" xml:"name" required:"true"`
	// Voice ID assigned to this speech engine
	VoiceId string `json:"voice_id" xml:"voice_id" required:"true"`
	// Creation time in Unix seconds
	CreatedAtUnixSecs int `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	// Last update time in Unix seconds
	UpdatedAtUnixSecs int `json:"updated_at_unix_secs" xml:"updated_at_unix_secs" required:"true"`
	// Arbitrary tags for categorization and filtering
	Tags            []string            `json:"tags" xml:"tags" required:"true"`
	AccessInfo      *ResourceAccessInfo `json:"access_info" xml:"access_info" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SpeechEngineSummaryResponse) GetSpeechEngineId() (value string) {
	if s == nil {
		return
	}
	return s.SpeechEngineId
}

func (s *SpeechEngineSummaryResponse) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

func (s *SpeechEngineSummaryResponse) GetVoiceId() (value string) {
	if s == nil {
		return
	}
	return s.VoiceId
}

func (s *SpeechEngineSummaryResponse) GetCreatedAtUnixSecs() (value int) {
	if s == nil {
		return
	}
	return s.CreatedAtUnixSecs
}

func (s *SpeechEngineSummaryResponse) GetUpdatedAtUnixSecs() (value int) {
	if s == nil {
		return
	}
	return s.UpdatedAtUnixSecs
}

func (s *SpeechEngineSummaryResponse) GetTags() (value []string) {
	if s == nil {
		return
	}
	return s.Tags
}

func (s *SpeechEngineSummaryResponse) GetAccessInfo() (value *ResourceAccessInfo) {
	if s == nil {
		return
	}
	return s.AccessInfo
}

func (s *SpeechEngineSummaryResponse) SetSpeechEngineId(speechEngineId string) {
	s.SpeechEngineId = speechEngineId
}

func (s *SpeechEngineSummaryResponse) SetName(name string) {
	s.Name = name
}

func (s *SpeechEngineSummaryResponse) SetVoiceId(voiceId string) {
	s.VoiceId = voiceId
}

func (s *SpeechEngineSummaryResponse) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	s.CreatedAtUnixSecs = createdAtUnixSecs
}

func (s *SpeechEngineSummaryResponse) SetUpdatedAtUnixSecs(updatedAtUnixSecs int) {
	s.UpdatedAtUnixSecs = updatedAtUnixSecs
}

func (s *SpeechEngineSummaryResponse) SetTags(tags []string) {
	s.Tags = tags
}

func (s *SpeechEngineSummaryResponse) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	s.AccessInfo = accessInfo
}

func (s *SpeechEngineSummaryResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeechEngineSummaryResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeechEngineSummaryResponse to string"
	}
	return string(jsonData)
}

func (s *SpeechEngineSummaryResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SpeechEngineSummaryResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SpeechEngineSummaryResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeechEngineSummaryResponse) MarshalJSON() ([]byte, error) {
	type embed SpeechEngineSummaryResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
