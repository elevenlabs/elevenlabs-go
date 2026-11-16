package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReplicateVoiceToIsolatedEnvironmentResponseModel struct {
	// Voice ID of the replicated voice in the target residency.
	VoiceId         string `json:"voice_id" xml:"voice_id" required:"true"`
	extraProperties map[string]interface{}
}

func (r *ReplicateVoiceToIsolatedEnvironmentResponseModel) GetVoiceId() (value string) {
	if r == nil {
		return
	}
	return r.VoiceId
}

func (r *ReplicateVoiceToIsolatedEnvironmentResponseModel) SetVoiceId(voiceId string) {
	r.VoiceId = voiceId
}

func (r *ReplicateVoiceToIsolatedEnvironmentResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReplicateVoiceToIsolatedEnvironmentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReplicateVoiceToIsolatedEnvironmentResponseModel to string"
	}
	return string(jsonData)
}

func (r *ReplicateVoiceToIsolatedEnvironmentResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReplicateVoiceToIsolatedEnvironmentResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReplicateVoiceToIsolatedEnvironmentResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReplicateVoiceToIsolatedEnvironmentResponseModel) MarshalJSON() ([]byte, error) {
	type embed ReplicateVoiceToIsolatedEnvironmentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
