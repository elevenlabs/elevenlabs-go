package voices

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type AddVoiceRequest struct {
	// The name that identifies this voice. This will be displayed in the dropdown of the website.
	Name string `json:"name" xml:"name" required:"true"`
	// A list of file paths to audio recordings intended for voice cloning.
	Files []io.Reader `json:"files" xml:"files" required:"true"`
	// If set will remove background noise for voice samples using our audio isolation model. If the samples do not include background noise, it can make the quality worse.
	RemoveBackgroundNoise *bool `json:"remove_background_noise,omitempty" xml:"remove_background_noise,omitempty"`
	// A description of the voice.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Labels for the voice. Keys can be language, accent, gender, or age.
	Labels *AddVoiceRequestLabels `json:"labels,omitempty" xml:"labels,omitempty"`
}

func (a *AddVoiceRequest) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AddVoiceRequest) GetFiles() (value []io.Reader) {
	if a == nil {
		return
	}
	return a.Files
}

func (a *AddVoiceRequest) GetRemoveBackgroundNoise() (value *bool) {
	if a == nil {
		return
	}
	return a.RemoveBackgroundNoise
}

func (a *AddVoiceRequest) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AddVoiceRequest) GetLabels() (value *AddVoiceRequestLabels) {
	if a == nil {
		return
	}
	return a.Labels
}

func (a *AddVoiceRequest) SetName(name string) {
	a.Name = name
}

func (a *AddVoiceRequest) SetFiles(files []io.Reader) {
	a.Files = files
}

func (a *AddVoiceRequest) SetRemoveBackgroundNoise(removeBackgroundNoise *bool) {
	a.RemoveBackgroundNoise = removeBackgroundNoise
}

func (a *AddVoiceRequest) SetDescription(description *string) {
	a.Description = description
}

func (a *AddVoiceRequest) SetLabels(labels *AddVoiceRequestLabels) {
	a.Labels = labels
}

func (a *AddVoiceRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AddVoiceRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddVoiceRequest to string"
	}
	return string(jsonData)
}

func (a *AddVoiceRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}

func (a *AddVoiceRequest) MarshalJSON() ([]byte, error) {
	type embed AddVoiceRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
