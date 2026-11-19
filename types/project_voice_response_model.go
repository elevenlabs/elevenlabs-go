package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProjectVoiceResponseModel struct {
	ProjectVoiceRefId string  `json:"project_voice_ref_id" xml:"project_voice_ref_id" required:"true"`
	VoiceId           string  `json:"voice_id" xml:"voice_id" required:"true"`
	Alias             string  `json:"alias" xml:"alias" required:"true"`
	Stability         float64 `json:"stability" xml:"stability" required:"true"`
	SimilarityBoost   float64 `json:"similarity_boost" xml:"similarity_boost" required:"true"`
	Style             float64 `json:"style" xml:"style" required:"true"`
	IsPinned          bool    `json:"is_pinned" xml:"is_pinned" required:"true"`
	UseSpeakerBoost   bool    `json:"use_speaker_boost" xml:"use_speaker_boost" required:"true"`
	VolumeGain        float64 `json:"volume_gain" xml:"volume_gain" required:"true"`
	Speed             float64 `json:"speed" xml:"speed" required:"true"`
	extraProperties   map[string]interface{}
}

func (p *ProjectVoiceResponseModel) GetProjectVoiceRefId() (value string) {
	if p == nil {
		return
	}
	return p.ProjectVoiceRefId
}

func (p *ProjectVoiceResponseModel) GetVoiceId() (value string) {
	if p == nil {
		return
	}
	return p.VoiceId
}

func (p *ProjectVoiceResponseModel) GetAlias() (value string) {
	if p == nil {
		return
	}
	return p.Alias
}

func (p *ProjectVoiceResponseModel) GetStability() (value float64) {
	if p == nil {
		return
	}
	return p.Stability
}

func (p *ProjectVoiceResponseModel) GetSimilarityBoost() (value float64) {
	if p == nil {
		return
	}
	return p.SimilarityBoost
}

func (p *ProjectVoiceResponseModel) GetStyle() (value float64) {
	if p == nil {
		return
	}
	return p.Style
}

func (p *ProjectVoiceResponseModel) GetIsPinned() (value bool) {
	if p == nil {
		return
	}
	return p.IsPinned
}

func (p *ProjectVoiceResponseModel) GetUseSpeakerBoost() (value bool) {
	if p == nil {
		return
	}
	return p.UseSpeakerBoost
}

func (p *ProjectVoiceResponseModel) GetVolumeGain() (value float64) {
	if p == nil {
		return
	}
	return p.VolumeGain
}

func (p *ProjectVoiceResponseModel) GetSpeed() (value float64) {
	if p == nil {
		return
	}
	return p.Speed
}

func (p *ProjectVoiceResponseModel) SetProjectVoiceRefId(projectVoiceRefId string) {
	p.ProjectVoiceRefId = projectVoiceRefId
}

func (p *ProjectVoiceResponseModel) SetVoiceId(voiceId string) {
	p.VoiceId = voiceId
}

func (p *ProjectVoiceResponseModel) SetAlias(alias string) {
	p.Alias = alias
}

func (p *ProjectVoiceResponseModel) SetStability(stability float64) {
	p.Stability = stability
}

func (p *ProjectVoiceResponseModel) SetSimilarityBoost(similarityBoost float64) {
	p.SimilarityBoost = similarityBoost
}

func (p *ProjectVoiceResponseModel) SetStyle(style float64) {
	p.Style = style
}

func (p *ProjectVoiceResponseModel) SetIsPinned(isPinned bool) {
	p.IsPinned = isPinned
}

func (p *ProjectVoiceResponseModel) SetUseSpeakerBoost(useSpeakerBoost bool) {
	p.UseSpeakerBoost = useSpeakerBoost
}

func (p *ProjectVoiceResponseModel) SetVolumeGain(volumeGain float64) {
	p.VolumeGain = volumeGain
}

func (p *ProjectVoiceResponseModel) SetSpeed(speed float64) {
	p.Speed = speed
}

func (p *ProjectVoiceResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectVoiceResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectVoiceResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProjectVoiceResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias ProjectVoiceResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = ProjectVoiceResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProjectVoiceResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProjectVoiceResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
