package projects

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type BodyCreateChapterV1StudioProjectsProjectIdChaptersPost struct {
	// The ID of the Studio project.
	ProjectId string `json:"-" url:"-"`
	// The name of the chapter, used for identification only.
	Name string `json:"name" xml:"name" required:"true"`
	// An optional URL from which we will extract content to initialize the Studio project. If this is set, 'from_url' and 'from_content' must be null. If neither 'from_url', 'from_document', 'from_content' are provided we will initialize the Studio project as blank.
	FromUrl *string `json:"from_url,omitempty" xml:"from_url,omitempty"`
}

func (b *BodyCreateChapterV1StudioProjectsProjectIdChaptersPost) GetProjectId() (value string) {
	if b == nil {
		return
	}
	return b.ProjectId
}

func (b *BodyCreateChapterV1StudioProjectsProjectIdChaptersPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateChapterV1StudioProjectsProjectIdChaptersPost) GetFromUrl() (value *string) {
	if b == nil {
		return
	}
	return b.FromUrl
}

func (b *BodyCreateChapterV1StudioProjectsProjectIdChaptersPost) SetProjectId(projectId string) {
	b.ProjectId = projectId
}

func (b *BodyCreateChapterV1StudioProjectsProjectIdChaptersPost) SetName(name string) {
	b.Name = name
}

func (b *BodyCreateChapterV1StudioProjectsProjectIdChaptersPost) SetFromUrl(fromUrl *string) {
	b.FromUrl = fromUrl
}

func (b *BodyCreateChapterV1StudioProjectsProjectIdChaptersPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateChapterV1StudioProjectsProjectIdChaptersPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateChapterV1StudioProjectsProjectIdChaptersPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateChapterV1StudioProjectsProjectIdChaptersPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateChapterV1StudioProjectsProjectIdChaptersPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateChapterV1StudioProjectsProjectIdChaptersPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost struct {
	// The ID of the chapter to be used. You can use the [List project chapters](/docs/api-reference/studio/get-chapters) endpoint to list all the available chapters.
	ChapterId string `json:"-" url:"-"`
	// The ID of the project to be used. You can use the [List projects](/docs/api-reference/studio/get-projects) endpoint to list all the available projects.
	ProjectId string `json:"-" url:"-"`
	// The name of the chapter, used for identification only.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The chapter content to use.
	Content *types.ChapterContentInputModel `json:"content,omitempty" xml:"content,omitempty"`
}

func (b *BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) GetChapterId() (value string) {
	if b == nil {
		return
	}
	return b.ChapterId
}

func (b *BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) GetProjectId() (value string) {
	if b == nil {
		return
	}
	return b.ProjectId
}

func (b *BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) GetContent() (value *types.ChapterContentInputModel) {
	if b == nil {
		return
	}
	return b.Content
}

func (b *BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) SetChapterId(chapterId string) {
	b.ChapterId = chapterId
}

func (b *BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) SetProjectId(projectId string) {
	b.ProjectId = projectId
}

func (b *BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) SetName(name *string) {
	b.Name = name
}

func (b *BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) SetContent(content *types.ChapterContentInputModel) {
	b.Content = content
}

func (b *BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateChapterV1StudioProjectsProjectIdChaptersChapterIdPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
