package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type EditVoiceRequest struct {
	// ID of the voice to be used. You can use the [Get voices](/docs/api-reference/voices/search) endpoint list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// The name that identifies this voice. This will be displayed in the dropdown of the website.
	Name string `json:"name" xml:"name" required:"true"`
	// Audio files to add to the voice
	Files []io.Reader `json:"files,omitempty" xml:"files,omitempty"`
	// If set will remove background noise for voice samples using our audio isolation model. If the samples do not include background noise, it can make the quality worse.
	RemoveBackgroundNoise *bool `json:"remove_background_noise,omitempty" xml:"remove_background_noise,omitempty"`
	// A description of the voice.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Labels for the voice. Keys can be language, accent, gender, or age.
	Labels *EditVoiceRequestLabels `json:"labels,omitempty" xml:"labels,omitempty"`
	// Run synchronous LLM moderation over the voice name and description when they change. Has no effect unless the voice_library_metadata_moderation feature flag is enabled for the user.
	ModerateMetadata *bool `json:"moderate_metadata,omitempty" xml:"moderate_metadata,omitempty"`
}

func (e *EditVoiceRequest) GetVoiceId() (value string) {
	if e == nil {
		return
	}
	return e.VoiceId
}

func (e *EditVoiceRequest) GetName() (value string) {
	if e == nil {
		return
	}
	return e.Name
}

func (e *EditVoiceRequest) GetFiles() (value []io.Reader) {
	if e == nil {
		return
	}
	return e.Files
}

func (e *EditVoiceRequest) GetRemoveBackgroundNoise() (value *bool) {
	if e == nil {
		return
	}
	return e.RemoveBackgroundNoise
}

func (e *EditVoiceRequest) GetDescription() (value *string) {
	if e == nil {
		return
	}
	return e.Description
}

func (e *EditVoiceRequest) GetLabels() (value *EditVoiceRequestLabels) {
	if e == nil {
		return
	}
	return e.Labels
}

func (e *EditVoiceRequest) GetModerateMetadata() (value *bool) {
	if e == nil {
		return
	}
	return e.ModerateMetadata
}

func (e *EditVoiceRequest) SetVoiceId(voiceId string) {
	e.VoiceId = voiceId
}

func (e *EditVoiceRequest) SetName(name string) {
	e.Name = name
}

func (e *EditVoiceRequest) SetFiles(files []io.Reader) {
	e.Files = files
}

func (e *EditVoiceRequest) SetRemoveBackgroundNoise(removeBackgroundNoise *bool) {
	e.RemoveBackgroundNoise = removeBackgroundNoise
}

func (e *EditVoiceRequest) SetDescription(description *string) {
	e.Description = description
}

func (e *EditVoiceRequest) SetLabels(labels *EditVoiceRequestLabels) {
	e.Labels = labels
}

func (e *EditVoiceRequest) SetModerateMetadata(moderateMetadata *bool) {
	e.ModerateMetadata = moderateMetadata
}

func (e *EditVoiceRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (e EditVoiceRequest) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditVoiceRequest to string"
	}
	return string(jsonData)
}

func (e *EditVoiceRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, e)
}

func (e *EditVoiceRequest) MarshalJSON() ([]byte, error) {
	type embed EditVoiceRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
