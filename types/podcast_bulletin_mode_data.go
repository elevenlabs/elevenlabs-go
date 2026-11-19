package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PodcastBulletinModeData struct {
	// The ID of the host voice.
	HostVoiceId     string `json:"host_voice_id" xml:"host_voice_id" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PodcastBulletinModeData) GetHostVoiceId() (value string) {
	if p == nil {
		return
	}
	return p.HostVoiceId
}

func (p *PodcastBulletinModeData) SetHostVoiceId(hostVoiceId string) {
	p.HostVoiceId = hostVoiceId
}

func (p *PodcastBulletinModeData) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PodcastBulletinModeData) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PodcastBulletinModeData to string"
	}
	return string(jsonData)
}

func (p *PodcastBulletinModeData) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PodcastBulletinModeData
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PodcastBulletinModeData(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PodcastBulletinModeData) MarshalJSON() ([]byte, error) {
	type embed PodcastBulletinModeData
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
