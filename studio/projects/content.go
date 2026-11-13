package projects

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost struct {
	// The ID of the project to be used. You can use the [List projects](/docs/api-reference/studio/get-projects) endpoint to list all the available projects.
	ProjectId string `json:"-" url:"-"`
	// An optional URL from which we will extract content to initialize the Studio project. If this is set, 'from_url' and 'from_content' must be null. If neither 'from_url', 'from_document', 'from_content' are provided we will initialize the Studio project as blank.
	FromUrl *string `json:"from_url,omitempty" xml:"from_url,omitempty"`
	// An optional .epub, .pdf, .txt or similar file can be provided. If provided, we will initialize the Studio project with its content. If this is set, 'from_url' and 'from_content' must be null. If neither 'from_url', 'from_document', 'from_content' are provided we will initialize the Studio project as blank.
	FromDocument *string `json:"from_document,omitempty" xml:"from_document,omitempty"`
	//
	// An optional content to initialize the Studio project with. If this is set, 'from_url' and 'from_document' must be null. If neither 'from_url', 'from_document', 'from_content' are provided we will initialize the Studio project as blank.
	//
	// Example:
	// [{"name": "Chapter A", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "A", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "B", "type": "tts_node"}]}, {"sub_type": "h1", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "C", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "D", "type": "tts_node"}]}]}, {"name": "Chapter B", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "E", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "F", "type": "tts_node"}]}, {"sub_type": "h2", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "G", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "H", "type": "tts_node"}]}]}]
	//
	FromContentJson *string `json:"from_content_json,omitempty" xml:"from_content_json,omitempty"`
	// Whether to auto convert the Studio project to audio or not.
	AutoConvert *bool `json:"auto_convert,omitempty" xml:"auto_convert,omitempty"`
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) GetProjectId() (value string) {
	if b == nil {
		return
	}
	return b.ProjectId
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) GetFromUrl() (value *string) {
	if b == nil {
		return
	}
	return b.FromUrl
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) GetFromDocument() (value *string) {
	if b == nil {
		return
	}
	return b.FromDocument
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) GetFromContentJson() (value *string) {
	if b == nil {
		return
	}
	return b.FromContentJson
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) GetAutoConvert() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoConvert
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) SetProjectId(projectId string) {
	b.ProjectId = projectId
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) SetFromUrl(fromUrl *string) {
	b.FromUrl = fromUrl
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) SetFromDocument(fromDocument *string) {
	b.FromDocument = fromDocument
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) SetFromContentJson(fromContentJson *string) {
	b.FromContentJson = fromContentJson
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) SetAutoConvert(autoConvert *bool) {
	b.AutoConvert = autoConvert
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateStudioProjectContentV1StudioProjectsProjectIdContentPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
