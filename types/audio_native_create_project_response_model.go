package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioNativeCreateProjectResponseModel struct {
	// The ID of the created Audio Native project.
	ProjectId string `json:"project_id" xml:"project_id" required:"true"`
	// Whether the project is currently being converted.
	Converting bool `json:"converting" xml:"converting" required:"true"`
	// The HTML snippet to embed the Audio Native player.
	HtmlSnippet     string `json:"html_snippet" xml:"html_snippet" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AudioNativeCreateProjectResponseModel) GetProjectId() (value string) {
	if a == nil {
		return
	}
	return a.ProjectId
}

func (a *AudioNativeCreateProjectResponseModel) GetConverting() (value bool) {
	if a == nil {
		return
	}
	return a.Converting
}

func (a *AudioNativeCreateProjectResponseModel) GetHtmlSnippet() (value string) {
	if a == nil {
		return
	}
	return a.HtmlSnippet
}

func (a *AudioNativeCreateProjectResponseModel) SetProjectId(projectId string) {
	a.ProjectId = projectId
}

func (a *AudioNativeCreateProjectResponseModel) SetConverting(converting bool) {
	a.Converting = converting
}

func (a *AudioNativeCreateProjectResponseModel) SetHtmlSnippet(htmlSnippet string) {
	a.HtmlSnippet = htmlSnippet
}

func (a *AudioNativeCreateProjectResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioNativeCreateProjectResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioNativeCreateProjectResponseModel to string"
	}
	return string(jsonData)
}

func (a *AudioNativeCreateProjectResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AudioNativeCreateProjectResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AudioNativeCreateProjectResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioNativeCreateProjectResponseModel) MarshalJSON() ([]byte, error) {
	type embed AudioNativeCreateProjectResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
