package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost struct {
	// The ID of the project to be used. You can use the [List projects](/docs/api-reference/studio/get-projects) endpoint to list all the available projects.
	ProjectId string `json:"-" url:"-"`
	// Either txt or HTML input file containing the article content. HTML should be formatted as follows '&lt;html&gt;&lt;body&gt;&lt;div&gt;&lt;p&gt;Your content&lt;/p&gt;&lt;h5&gt;More of your content&lt;/h5&gt;&lt;p&gt;Some more of your content&lt;/p&gt;&lt;/div&gt;&lt;/body&gt;&lt;/html&gt;'
	File io.Reader `json:"file,omitempty" xml:"file,omitempty"`
	// Whether to auto convert the project to audio or not.
	AutoConvert *bool `json:"auto_convert,omitempty" xml:"auto_convert,omitempty"`
	// Whether to auto publish the new project snapshot after it's converted.
	AutoPublish *bool `json:"auto_publish,omitempty" xml:"auto_publish,omitempty"`
}

func (b *BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) GetProjectId() (value string) {
	if b == nil {
		return
	}
	return b.ProjectId
}

func (b *BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) GetFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) GetAutoConvert() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoConvert
}

func (b *BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) GetAutoPublish() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoPublish
}

func (b *BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) SetProjectId(projectId string) {
	b.ProjectId = projectId
}

func (b *BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) SetFile(file io.Reader) {
	b.File = file
}

func (b *BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) SetAutoConvert(autoConvert *bool) {
	b.AutoConvert = autoConvert
}

func (b *BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) SetAutoPublish(autoPublish *bool) {
	b.AutoPublish = autoPublish
}

func (b *BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost(tmp)
	return nil
}

func (b *BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
