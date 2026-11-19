package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProjectVideoThumbnailSheetResponseModel struct {
	StartThumbnailIndex int    `json:"start_thumbnail_index" xml:"start_thumbnail_index" required:"true"`
	ThumbnailCount      int    `json:"thumbnail_count" xml:"thumbnail_count" required:"true"`
	SignedCloudUrl      string `json:"signed_cloud_url" xml:"signed_cloud_url" required:"true"`
	extraProperties     map[string]interface{}
}

func (p *ProjectVideoThumbnailSheetResponseModel) GetStartThumbnailIndex() (value int) {
	if p == nil {
		return
	}
	return p.StartThumbnailIndex
}

func (p *ProjectVideoThumbnailSheetResponseModel) GetThumbnailCount() (value int) {
	if p == nil {
		return
	}
	return p.ThumbnailCount
}

func (p *ProjectVideoThumbnailSheetResponseModel) GetSignedCloudUrl() (value string) {
	if p == nil {
		return
	}
	return p.SignedCloudUrl
}

func (p *ProjectVideoThumbnailSheetResponseModel) SetStartThumbnailIndex(startThumbnailIndex int) {
	p.StartThumbnailIndex = startThumbnailIndex
}

func (p *ProjectVideoThumbnailSheetResponseModel) SetThumbnailCount(thumbnailCount int) {
	p.ThumbnailCount = thumbnailCount
}

func (p *ProjectVideoThumbnailSheetResponseModel) SetSignedCloudUrl(signedCloudUrl string) {
	p.SignedCloudUrl = signedCloudUrl
}

func (p *ProjectVideoThumbnailSheetResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectVideoThumbnailSheetResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectVideoThumbnailSheetResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProjectVideoThumbnailSheetResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias ProjectVideoThumbnailSheetResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = ProjectVideoThumbnailSheetResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProjectVideoThumbnailSheetResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProjectVideoThumbnailSheetResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
