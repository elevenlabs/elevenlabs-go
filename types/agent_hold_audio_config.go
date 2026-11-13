package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Custom hold audio played on loop to callers waiting in the agent's queue.
//
// Set by uploading a file through the agent hold-audio endpoint. Values sent in
// agent create or update requests are ignored.
type AgentHoldAudioConfig struct {
	// Storage path of the uploaded clip
	AudioPath string `json:"audio_path" xml:"audio_path" required:"true"`
	// Public CDN URL of the uploaded clip
	AudioUrl string `json:"audio_url" xml:"audio_url" required:"true"`
	// Filename of the uploaded clip as provided by the user
	OriginalFilename string `json:"original_filename" xml:"original_filename" required:"true"`
	// Duration of the uploaded clip in seconds
	DurationSecs float64 `json:"duration_secs" xml:"duration_secs" required:"true"`
	// Size of the uploaded clip in bytes
	SizeBytes       int `json:"size_bytes" xml:"size_bytes" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentHoldAudioConfig) GetAudioPath() (value string) {
	if a == nil {
		return
	}
	return a.AudioPath
}

func (a *AgentHoldAudioConfig) GetAudioUrl() (value string) {
	if a == nil {
		return
	}
	return a.AudioUrl
}

func (a *AgentHoldAudioConfig) GetOriginalFilename() (value string) {
	if a == nil {
		return
	}
	return a.OriginalFilename
}

func (a *AgentHoldAudioConfig) GetDurationSecs() (value float64) {
	if a == nil {
		return
	}
	return a.DurationSecs
}

func (a *AgentHoldAudioConfig) GetSizeBytes() (value int) {
	if a == nil {
		return
	}
	return a.SizeBytes
}

func (a *AgentHoldAudioConfig) SetAudioPath(audioPath string) {
	a.AudioPath = audioPath
}

func (a *AgentHoldAudioConfig) SetAudioUrl(audioUrl string) {
	a.AudioUrl = audioUrl
}

func (a *AgentHoldAudioConfig) SetOriginalFilename(originalFilename string) {
	a.OriginalFilename = originalFilename
}

func (a *AgentHoldAudioConfig) SetDurationSecs(durationSecs float64) {
	a.DurationSecs = durationSecs
}

func (a *AgentHoldAudioConfig) SetSizeBytes(sizeBytes int) {
	a.SizeBytes = sizeBytes
}

func (a *AgentHoldAudioConfig) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentHoldAudioConfig) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentHoldAudioConfig to string"
	}
	return string(jsonData)
}

func (a *AgentHoldAudioConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentHoldAudioConfig
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentHoldAudioConfig(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentHoldAudioConfig) MarshalJSON() ([]byte, error) {
	type embed AgentHoldAudioConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
