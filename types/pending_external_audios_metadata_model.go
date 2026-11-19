package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PendingExternalAudiosMetadataModel struct {
	TargetGlobalOffsetMs *int     `json:"target_global_offset_ms" xml:"target_global_offset_ms" required:"true"`
	ExternalAudioIds     []string `json:"external_audio_ids" xml:"external_audio_ids" required:"true"`
	extraProperties      map[string]interface{}
}

func (p *PendingExternalAudiosMetadataModel) GetTargetGlobalOffsetMs() (value *int) {
	if p == nil {
		return
	}
	return p.TargetGlobalOffsetMs
}

func (p *PendingExternalAudiosMetadataModel) GetExternalAudioIds() (value []string) {
	if p == nil {
		return
	}
	return p.ExternalAudioIds
}

func (p *PendingExternalAudiosMetadataModel) SetTargetGlobalOffsetMs(targetGlobalOffsetMs *int) {
	p.TargetGlobalOffsetMs = targetGlobalOffsetMs
}

func (p *PendingExternalAudiosMetadataModel) SetExternalAudioIds(externalAudioIds []string) {
	p.ExternalAudioIds = externalAudioIds
}

func (p *PendingExternalAudiosMetadataModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PendingExternalAudiosMetadataModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PendingExternalAudiosMetadataModel to string"
	}
	return string(jsonData)
}

func (p *PendingExternalAudiosMetadataModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PendingExternalAudiosMetadataModel) MarshalJSON() ([]byte, error) {
	type embed PendingExternalAudiosMetadataModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
