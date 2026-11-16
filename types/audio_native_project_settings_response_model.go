package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioNativeProjectSettingsResponseModel struct {
	// The title of the project.
	Title string `json:"title" xml:"title" required:"true"`
	// The image of the project.
	Image string `json:"image" xml:"image" required:"true"`
	// The author of the project.
	Author string `json:"author" xml:"author" required:"true"`
	// Whether the project is small.
	Small bool `json:"small" xml:"small" required:"true"`
	// The text color of the project.
	TextColor string `json:"text_color" xml:"text_color" required:"true"`
	// The background color of the project.
	BackgroundColor string `json:"background_color" xml:"background_color" required:"true"`
	// The sessionization of the project. Specifies for how many minutes to persist the session across page reloads.
	Sessionization int `json:"sessionization" xml:"sessionization" required:"true"`
	// The path of the audio file.
	AudioPath *string `json:"audio_path,omitempty" xml:"audio_path,omitempty"`
	// The URL of the audio file.
	AudioUrl *string `json:"audio_url,omitempty" xml:"audio_url,omitempty"`
	// Current state of the project
	Status          *AudioNativeProjectSettingsResponseModelStatus `json:"status,omitempty" xml:"status,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AudioNativeProjectSettingsResponseModel) GetTitle() (value string) {
	if a == nil {
		return
	}
	return a.Title
}

func (a *AudioNativeProjectSettingsResponseModel) GetImage() (value string) {
	if a == nil {
		return
	}
	return a.Image
}

func (a *AudioNativeProjectSettingsResponseModel) GetAuthor() (value string) {
	if a == nil {
		return
	}
	return a.Author
}

func (a *AudioNativeProjectSettingsResponseModel) GetSmall() (value bool) {
	if a == nil {
		return
	}
	return a.Small
}

func (a *AudioNativeProjectSettingsResponseModel) GetTextColor() (value string) {
	if a == nil {
		return
	}
	return a.TextColor
}

func (a *AudioNativeProjectSettingsResponseModel) GetBackgroundColor() (value string) {
	if a == nil {
		return
	}
	return a.BackgroundColor
}

func (a *AudioNativeProjectSettingsResponseModel) GetSessionization() (value int) {
	if a == nil {
		return
	}
	return a.Sessionization
}

func (a *AudioNativeProjectSettingsResponseModel) GetAudioPath() (value *string) {
	if a == nil {
		return
	}
	return a.AudioPath
}

func (a *AudioNativeProjectSettingsResponseModel) GetAudioUrl() (value *string) {
	if a == nil {
		return
	}
	return a.AudioUrl
}

func (a *AudioNativeProjectSettingsResponseModel) GetStatus() (value *AudioNativeProjectSettingsResponseModelStatus) {
	if a == nil {
		return
	}
	return a.Status
}

func (a *AudioNativeProjectSettingsResponseModel) SetTitle(title string) {
	a.Title = title
}

func (a *AudioNativeProjectSettingsResponseModel) SetImage(image string) {
	a.Image = image
}

func (a *AudioNativeProjectSettingsResponseModel) SetAuthor(author string) {
	a.Author = author
}

func (a *AudioNativeProjectSettingsResponseModel) SetSmall(small bool) {
	a.Small = small
}

func (a *AudioNativeProjectSettingsResponseModel) SetTextColor(textColor string) {
	a.TextColor = textColor
}

func (a *AudioNativeProjectSettingsResponseModel) SetBackgroundColor(backgroundColor string) {
	a.BackgroundColor = backgroundColor
}

func (a *AudioNativeProjectSettingsResponseModel) SetSessionization(sessionization int) {
	a.Sessionization = sessionization
}

func (a *AudioNativeProjectSettingsResponseModel) SetAudioPath(audioPath *string) {
	a.AudioPath = audioPath
}

func (a *AudioNativeProjectSettingsResponseModel) SetAudioUrl(audioUrl *string) {
	a.AudioUrl = audioUrl
}

func (a *AudioNativeProjectSettingsResponseModel) SetStatus(status *AudioNativeProjectSettingsResponseModelStatus) {
	a.Status = status
}

func (a *AudioNativeProjectSettingsResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioNativeProjectSettingsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioNativeProjectSettingsResponseModel to string"
	}
	return string(jsonData)
}

func (a *AudioNativeProjectSettingsResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioNativeProjectSettingsResponseModel) MarshalJSON() ([]byte, error) {
	type embed AudioNativeProjectSettingsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
