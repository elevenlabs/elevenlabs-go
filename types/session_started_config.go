package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SessionStartedConfig struct {
	SampleRate               int                                        `json:"sample_rate" xml:"sample_rate" required:"true"`
	AudioFormat              SessionStartedConfigAudioFormat            `json:"audio_format" xml:"audio_format" required:"true"`
	LanguageCode             *string                                    `json:"language_code" xml:"language_code" required:"true"`
	SecondaryLanguages       []string                                   `json:"secondary_languages,omitempty" xml:"secondary_languages,omitempty"`
	TimestampsGranularity    *SessionStartedConfigTimestampsGranularity `json:"timestamps_granularity,omitempty" xml:"timestamps_granularity,omitempty"`
	VadCommitStrategy        *bool                                      `json:"vad_commit_strategy,omitempty" xml:"vad_commit_strategy,omitempty"`
	VadSilenceThresholdSecs  *float64                                   `json:"vad_silence_threshold_secs,omitempty" xml:"vad_silence_threshold_secs,omitempty"`
	VadThreshold             *float64                                   `json:"vad_threshold,omitempty" xml:"vad_threshold,omitempty"`
	MinSpeechDurationMs      *int                                       `json:"min_speech_duration_ms,omitempty" xml:"min_speech_duration_ms,omitempty"`
	MinSilenceDurationMs     *int                                       `json:"min_silence_duration_ms,omitempty" xml:"min_silence_duration_ms,omitempty"`
	MaxTokensToRecompute     *int                                       `json:"max_tokens_to_recompute,omitempty" xml:"max_tokens_to_recompute,omitempty"`
	ModelId                  *SessionStartedConfigModelId               `json:"model_id,omitempty" xml:"model_id,omitempty"`
	DisableLogging           *bool                                      `json:"disable_logging,omitempty" xml:"disable_logging,omitempty"`
	IncludeTimestamps        *bool                                      `json:"include_timestamps,omitempty" xml:"include_timestamps,omitempty"`
	IncludeLanguageDetection *bool                                      `json:"include_language_detection,omitempty" xml:"include_language_detection,omitempty"`
	FilterBackgroundAudio    *bool                                      `json:"filter_background_audio,omitempty" xml:"filter_background_audio,omitempty"`
	Keyterms                 []string                                   `json:"keyterms,omitempty" xml:"keyterms,omitempty"`
	NoVerbatim               *bool                                      `json:"no_verbatim,omitempty" xml:"no_verbatim,omitempty"`
	EntityDetection          []string                                   `json:"entity_detection,omitempty" xml:"entity_detection,omitempty"`
	EntityRedaction          []string                                   `json:"entity_redaction,omitempty" xml:"entity_redaction,omitempty"`
	EntityRedactionMode      *string                                    `json:"entity_redaction_mode,omitempty" xml:"entity_redaction_mode,omitempty"`
	KeepaliveIntervalMs      *int                                       `json:"keepalive_interval_ms,omitempty" xml:"keepalive_interval_ms,omitempty"`
	extraProperties          map[string]interface{}
}

func (s *SessionStartedConfig) GetSampleRate() (value int) {
	if s == nil {
		return
	}
	return s.SampleRate
}

func (s *SessionStartedConfig) GetAudioFormat() (value SessionStartedConfigAudioFormat) {
	if s == nil {
		return
	}
	return s.AudioFormat
}

func (s *SessionStartedConfig) GetLanguageCode() (value *string) {
	if s == nil {
		return
	}
	return s.LanguageCode
}

func (s *SessionStartedConfig) GetSecondaryLanguages() (value []string) {
	if s == nil {
		return
	}
	return s.SecondaryLanguages
}

func (s *SessionStartedConfig) GetTimestampsGranularity() (value *SessionStartedConfigTimestampsGranularity) {
	if s == nil {
		return
	}
	return s.TimestampsGranularity
}

func (s *SessionStartedConfig) GetVadCommitStrategy() (value *bool) {
	if s == nil {
		return
	}
	return s.VadCommitStrategy
}

func (s *SessionStartedConfig) GetVadSilenceThresholdSecs() (value *float64) {
	if s == nil {
		return
	}
	return s.VadSilenceThresholdSecs
}

func (s *SessionStartedConfig) GetVadThreshold() (value *float64) {
	if s == nil {
		return
	}
	return s.VadThreshold
}

func (s *SessionStartedConfig) GetMinSpeechDurationMs() (value *int) {
	if s == nil {
		return
	}
	return s.MinSpeechDurationMs
}

func (s *SessionStartedConfig) GetMinSilenceDurationMs() (value *int) {
	if s == nil {
		return
	}
	return s.MinSilenceDurationMs
}

func (s *SessionStartedConfig) GetMaxTokensToRecompute() (value *int) {
	if s == nil {
		return
	}
	return s.MaxTokensToRecompute
}

func (s *SessionStartedConfig) GetModelId() (value *SessionStartedConfigModelId) {
	if s == nil {
		return
	}
	return s.ModelId
}

func (s *SessionStartedConfig) GetDisableLogging() (value *bool) {
	if s == nil {
		return
	}
	return s.DisableLogging
}

func (s *SessionStartedConfig) GetIncludeTimestamps() (value *bool) {
	if s == nil {
		return
	}
	return s.IncludeTimestamps
}

func (s *SessionStartedConfig) GetIncludeLanguageDetection() (value *bool) {
	if s == nil {
		return
	}
	return s.IncludeLanguageDetection
}

func (s *SessionStartedConfig) GetFilterBackgroundAudio() (value *bool) {
	if s == nil {
		return
	}
	return s.FilterBackgroundAudio
}

func (s *SessionStartedConfig) GetKeyterms() (value []string) {
	if s == nil {
		return
	}
	return s.Keyterms
}

func (s *SessionStartedConfig) GetNoVerbatim() (value *bool) {
	if s == nil {
		return
	}
	return s.NoVerbatim
}

func (s *SessionStartedConfig) GetEntityDetection() (value []string) {
	if s == nil {
		return
	}
	return s.EntityDetection
}

func (s *SessionStartedConfig) GetEntityRedaction() (value []string) {
	if s == nil {
		return
	}
	return s.EntityRedaction
}

func (s *SessionStartedConfig) GetEntityRedactionMode() (value *string) {
	if s == nil {
		return
	}
	return s.EntityRedactionMode
}

func (s *SessionStartedConfig) GetKeepaliveIntervalMs() (value *int) {
	if s == nil {
		return
	}
	return s.KeepaliveIntervalMs
}

func (s *SessionStartedConfig) SetSampleRate(sampleRate int) {
	s.SampleRate = sampleRate
}

func (s *SessionStartedConfig) SetAudioFormat(audioFormat SessionStartedConfigAudioFormat) {
	s.AudioFormat = audioFormat
}

func (s *SessionStartedConfig) SetLanguageCode(languageCode *string) {
	s.LanguageCode = languageCode
}

func (s *SessionStartedConfig) SetSecondaryLanguages(secondaryLanguages []string) {
	s.SecondaryLanguages = secondaryLanguages
}

func (s *SessionStartedConfig) SetTimestampsGranularity(timestampsGranularity *SessionStartedConfigTimestampsGranularity) {
	s.TimestampsGranularity = timestampsGranularity
}

func (s *SessionStartedConfig) SetVadCommitStrategy(vadCommitStrategy *bool) {
	s.VadCommitStrategy = vadCommitStrategy
}

func (s *SessionStartedConfig) SetVadSilenceThresholdSecs(vadSilenceThresholdSecs *float64) {
	s.VadSilenceThresholdSecs = vadSilenceThresholdSecs
}

func (s *SessionStartedConfig) SetVadThreshold(vadThreshold *float64) {
	s.VadThreshold = vadThreshold
}

func (s *SessionStartedConfig) SetMinSpeechDurationMs(minSpeechDurationMs *int) {
	s.MinSpeechDurationMs = minSpeechDurationMs
}

func (s *SessionStartedConfig) SetMinSilenceDurationMs(minSilenceDurationMs *int) {
	s.MinSilenceDurationMs = minSilenceDurationMs
}

func (s *SessionStartedConfig) SetMaxTokensToRecompute(maxTokensToRecompute *int) {
	s.MaxTokensToRecompute = maxTokensToRecompute
}

func (s *SessionStartedConfig) SetModelId(modelId *SessionStartedConfigModelId) {
	s.ModelId = modelId
}

func (s *SessionStartedConfig) SetDisableLogging(disableLogging *bool) {
	s.DisableLogging = disableLogging
}

func (s *SessionStartedConfig) SetIncludeTimestamps(includeTimestamps *bool) {
	s.IncludeTimestamps = includeTimestamps
}

func (s *SessionStartedConfig) SetIncludeLanguageDetection(includeLanguageDetection *bool) {
	s.IncludeLanguageDetection = includeLanguageDetection
}

func (s *SessionStartedConfig) SetFilterBackgroundAudio(filterBackgroundAudio *bool) {
	s.FilterBackgroundAudio = filterBackgroundAudio
}

func (s *SessionStartedConfig) SetKeyterms(keyterms []string) {
	s.Keyterms = keyterms
}

func (s *SessionStartedConfig) SetNoVerbatim(noVerbatim *bool) {
	s.NoVerbatim = noVerbatim
}

func (s *SessionStartedConfig) SetEntityDetection(entityDetection []string) {
	s.EntityDetection = entityDetection
}

func (s *SessionStartedConfig) SetEntityRedaction(entityRedaction []string) {
	s.EntityRedaction = entityRedaction
}

func (s *SessionStartedConfig) SetEntityRedactionMode(entityRedactionMode *string) {
	s.EntityRedactionMode = entityRedactionMode
}

func (s *SessionStartedConfig) SetKeepaliveIntervalMs(keepaliveIntervalMs *int) {
	s.KeepaliveIntervalMs = keepaliveIntervalMs
}

func (s *SessionStartedConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SessionStartedConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SessionStartedConfig to string"
	}
	return string(jsonData)
}

func (s *SessionStartedConfig) UnmarshalJSON(data []byte) error {
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

func (s *SessionStartedConfig) MarshalJSON() ([]byte, error) {
	type embed SessionStartedConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
