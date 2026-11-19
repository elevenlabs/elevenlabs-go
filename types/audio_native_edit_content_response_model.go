package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioNativeEditContentResponseModel struct {
	// The ID of the project.
	ProjectId string `json:"project_id" xml:"project_id" required:"true"`
	// Whether the project is currently being converted.
	Converting bool `json:"converting" xml:"converting" required:"true"`
	// Whether the project is currently being published.
	Publishing bool `json:"publishing" xml:"publishing" required:"true"`
	// The HTML snippet to embed the Audio Native player.
	HtmlSnippet     string `json:"html_snippet" xml:"html_snippet" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AudioNativeEditContentResponseModel) GetProjectId() (value string) {
	if a == nil {
		return
	}
	return a.ProjectId
}

func (a *AudioNativeEditContentResponseModel) GetConverting() (value bool) {
	if a == nil {
		return
	}
	return a.Converting
}

func (a *AudioNativeEditContentResponseModel) GetPublishing() (value bool) {
	if a == nil {
		return
	}
	return a.Publishing
}

func (a *AudioNativeEditContentResponseModel) GetHtmlSnippet() (value string) {
	if a == nil {
		return
	}
	return a.HtmlSnippet
}

func (a *AudioNativeEditContentResponseModel) SetProjectId(projectId string) {
	a.ProjectId = projectId
}

func (a *AudioNativeEditContentResponseModel) SetConverting(converting bool) {
	a.Converting = converting
}

func (a *AudioNativeEditContentResponseModel) SetPublishing(publishing bool) {
	a.Publishing = publishing
}

func (a *AudioNativeEditContentResponseModel) SetHtmlSnippet(htmlSnippet string) {
	a.HtmlSnippet = htmlSnippet
}

func (a *AudioNativeEditContentResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioNativeEditContentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioNativeEditContentResponseModel to string"
	}
	return string(jsonData)
}

func (a *AudioNativeEditContentResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AudioNativeEditContentResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AudioNativeEditContentResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioNativeEditContentResponseModel) MarshalJSON() ([]byte, error) {
	type embed AudioNativeEditContentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
