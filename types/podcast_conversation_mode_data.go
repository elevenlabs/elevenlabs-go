package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PodcastConversationModeData struct {
	// The ID of the host voice.
	HostVoiceId string `json:"host_voice_id" xml:"host_voice_id" required:"true"`
	// The ID of the guest voice.
	GuestVoiceId    string `json:"guest_voice_id" xml:"guest_voice_id" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PodcastConversationModeData) GetHostVoiceId() (value string) {
	if p == nil {
		return
	}
	return p.HostVoiceId
}

func (p *PodcastConversationModeData) GetGuestVoiceId() (value string) {
	if p == nil {
		return
	}
	return p.GuestVoiceId
}

func (p *PodcastConversationModeData) SetHostVoiceId(hostVoiceId string) {
	p.HostVoiceId = hostVoiceId
}

func (p *PodcastConversationModeData) SetGuestVoiceId(guestVoiceId string) {
	p.GuestVoiceId = guestVoiceId
}

func (p *PodcastConversationModeData) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PodcastConversationModeData) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PodcastConversationModeData to string"
	}
	return string(jsonData)
}

func (p *PodcastConversationModeData) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PodcastConversationModeData
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PodcastConversationModeData(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PodcastConversationModeData) MarshalJSON() ([]byte, error) {
	type embed PodcastConversationModeData
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
