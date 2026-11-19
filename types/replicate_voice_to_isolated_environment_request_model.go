package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ReplicateVoiceToIsolatedEnvironmentRequestModel struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// ID of the workspace to replicate the voice into. It must belong to the same consolidated billing group as the calling workspace; the target's data residency is derived from that link.
	TargetWorkspaceId string `json:"target_workspace_id" xml:"target_workspace_id" required:"true"`
	// When true (default) the replicated voice keeps the same voice ID in the target residency; set to false to assign a new voice ID.
	PreserveVoiceId *bool `json:"preserve_voice_id,omitempty" xml:"preserve_voice_id,omitempty"`
}

func (r *ReplicateVoiceToIsolatedEnvironmentRequestModel) GetVoiceId() (value string) {
	if r == nil {
		return
	}
	return r.VoiceId
}

func (r *ReplicateVoiceToIsolatedEnvironmentRequestModel) GetTargetWorkspaceId() (value string) {
	if r == nil {
		return
	}
	return r.TargetWorkspaceId
}

func (r *ReplicateVoiceToIsolatedEnvironmentRequestModel) GetPreserveVoiceId() (value *bool) {
	if r == nil {
		return
	}
	return r.PreserveVoiceId
}

func (r *ReplicateVoiceToIsolatedEnvironmentRequestModel) SetVoiceId(voiceId string) {
	r.VoiceId = voiceId
}

func (r *ReplicateVoiceToIsolatedEnvironmentRequestModel) SetTargetWorkspaceId(targetWorkspaceId string) {
	r.TargetWorkspaceId = targetWorkspaceId
}

func (r *ReplicateVoiceToIsolatedEnvironmentRequestModel) SetPreserveVoiceId(preserveVoiceId *bool) {
	r.PreserveVoiceId = preserveVoiceId
}

func (r *ReplicateVoiceToIsolatedEnvironmentRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r ReplicateVoiceToIsolatedEnvironmentRequestModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReplicateVoiceToIsolatedEnvironmentRequestModel to string"
	}
	return string(jsonData)
}

func (r *ReplicateVoiceToIsolatedEnvironmentRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReplicateVoiceToIsolatedEnvironmentRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReplicateVoiceToIsolatedEnvironmentRequestModel(tmp)
	return nil
}

func (r *ReplicateVoiceToIsolatedEnvironmentRequestModel) MarshalJSON() ([]byte, error) {
	type embed ReplicateVoiceToIsolatedEnvironmentRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
