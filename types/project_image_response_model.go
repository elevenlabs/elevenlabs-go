package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProjectImageResponseModel struct {
	ImageId            string  `json:"image_id" xml:"image_id" required:"true"`
	Filename           string  `json:"filename" xml:"filename" required:"true"`
	SignedUrl          *string `json:"signed_url,omitempty" xml:"signed_url,omitempty"`
	ThumbnailSignedUrl *string `json:"thumbnail_signed_url,omitempty" xml:"thumbnail_signed_url,omitempty"`
	Type               *string `json:"type,omitempty" xml:"type,omitempty"`
	Source             *string `json:"source,omitempty" xml:"source,omitempty"`
	FileSizeBytes      int     `json:"file_size_bytes" xml:"file_size_bytes" required:"true"`
	Width              int     `json:"width" xml:"width" required:"true"`
	Height             int     `json:"height" xml:"height" required:"true"`
	TrackId            *string `json:"track_id,omitempty" xml:"track_id,omitempty"`
	OffsetMs           int     `json:"offset_ms" xml:"offset_ms" required:"true"`
	DurationMs         int     `json:"duration_ms" xml:"duration_ms" required:"true"`
	Order              string  `json:"order" xml:"order" required:"true"`
	// Defines asset positioning and transformation on canvas.
	CanvasPlacement       *CanvasPlacement `json:"canvas_placement" xml:"canvas_placement" required:"true"`
	Animation             *ClipAnimation   `json:"animation,omitempty" xml:"animation,omitempty"`
	Opacity               *float64         `json:"opacity,omitempty" xml:"opacity,omitempty" min:"0" max:"1"`
	CreatedAtMs           int              `json:"created_at_ms" xml:"created_at_ms" required:"true"`
	UpdatedAtMs           int              `json:"updated_at_ms" xml:"updated_at_ms" required:"true"`
	CurrentSnapshotId     *string          `json:"current_snapshot_id,omitempty" xml:"current_snapshot_id,omitempty"`
	SourceAssetId         *string          `json:"source_asset_id,omitempty" xml:"source_asset_id,omitempty"`
	SourcePlatformAssetId *string          `json:"source_platform_asset_id,omitempty" xml:"source_platform_asset_id,omitempty"`
	Error                 *string          `json:"error,omitempty" xml:"error,omitempty"`
	PendingTask           *PendingClipTask `json:"pending_task,omitempty" xml:"pending_task,omitempty"`
	Analysis              *ImageAnalysis   `json:"analysis,omitempty" xml:"analysis,omitempty"`
	extraProperties       map[string]interface{}
}

func (p *ProjectImageResponseModel) GetImageId() (value string) {
	if p == nil {
		return
	}
	return p.ImageId
}

func (p *ProjectImageResponseModel) GetFilename() (value string) {
	if p == nil {
		return
	}
	return p.Filename
}

func (p *ProjectImageResponseModel) GetSignedUrl() (value *string) {
	if p == nil {
		return
	}
	return p.SignedUrl
}

func (p *ProjectImageResponseModel) GetThumbnailSignedUrl() (value *string) {
	if p == nil {
		return
	}
	return p.ThumbnailSignedUrl
}

func (p *ProjectImageResponseModel) GetType() (value *string) {
	if p == nil {
		return
	}
	return p.Type
}

func (p *ProjectImageResponseModel) GetSource() (value *string) {
	if p == nil {
		return
	}
	return p.Source
}

func (p *ProjectImageResponseModel) GetFileSizeBytes() (value int) {
	if p == nil {
		return
	}
	return p.FileSizeBytes
}

func (p *ProjectImageResponseModel) GetWidth() (value int) {
	if p == nil {
		return
	}
	return p.Width
}

func (p *ProjectImageResponseModel) GetHeight() (value int) {
	if p == nil {
		return
	}
	return p.Height
}

func (p *ProjectImageResponseModel) GetTrackId() (value *string) {
	if p == nil {
		return
	}
	return p.TrackId
}

func (p *ProjectImageResponseModel) GetOffsetMs() (value int) {
	if p == nil {
		return
	}
	return p.OffsetMs
}

func (p *ProjectImageResponseModel) GetDurationMs() (value int) {
	if p == nil {
		return
	}
	return p.DurationMs
}

func (p *ProjectImageResponseModel) GetOrder() (value string) {
	if p == nil {
		return
	}
	return p.Order
}

func (p *ProjectImageResponseModel) GetCanvasPlacement() (value *CanvasPlacement) {
	if p == nil {
		return
	}
	return p.CanvasPlacement
}

func (p *ProjectImageResponseModel) GetAnimation() (value *ClipAnimation) {
	if p == nil {
		return
	}
	return p.Animation
}

func (p *ProjectImageResponseModel) GetOpacity() (value *float64) {
	if p == nil {
		return
	}
	return p.Opacity
}

func (p *ProjectImageResponseModel) GetCreatedAtMs() (value int) {
	if p == nil {
		return
	}
	return p.CreatedAtMs
}

func (p *ProjectImageResponseModel) GetUpdatedAtMs() (value int) {
	if p == nil {
		return
	}
	return p.UpdatedAtMs
}

func (p *ProjectImageResponseModel) GetCurrentSnapshotId() (value *string) {
	if p == nil {
		return
	}
	return p.CurrentSnapshotId
}

func (p *ProjectImageResponseModel) GetSourceAssetId() (value *string) {
	if p == nil {
		return
	}
	return p.SourceAssetId
}

func (p *ProjectImageResponseModel) GetSourcePlatformAssetId() (value *string) {
	if p == nil {
		return
	}
	return p.SourcePlatformAssetId
}

func (p *ProjectImageResponseModel) GetError() (value *string) {
	if p == nil {
		return
	}
	return p.Error
}

func (p *ProjectImageResponseModel) GetPendingTask() (value *PendingClipTask) {
	if p == nil {
		return
	}
	return p.PendingTask
}

func (p *ProjectImageResponseModel) GetAnalysis() (value *ImageAnalysis) {
	if p == nil {
		return
	}
	return p.Analysis
}

func (p *ProjectImageResponseModel) SetImageId(imageId string) {
	p.ImageId = imageId
}

func (p *ProjectImageResponseModel) SetFilename(filename string) {
	p.Filename = filename
}

func (p *ProjectImageResponseModel) SetSignedUrl(signedUrl *string) {
	p.SignedUrl = signedUrl
}

func (p *ProjectImageResponseModel) SetThumbnailSignedUrl(thumbnailSignedUrl *string) {
	p.ThumbnailSignedUrl = thumbnailSignedUrl
}

func (p *ProjectImageResponseModel) SetType(type_ *string) {
	p.Type = type_
}

func (p *ProjectImageResponseModel) SetSource(source *string) {
	p.Source = source
}

func (p *ProjectImageResponseModel) SetFileSizeBytes(fileSizeBytes int) {
	p.FileSizeBytes = fileSizeBytes
}

func (p *ProjectImageResponseModel) SetWidth(width int) {
	p.Width = width
}

func (p *ProjectImageResponseModel) SetHeight(height int) {
	p.Height = height
}

func (p *ProjectImageResponseModel) SetTrackId(trackId *string) {
	p.TrackId = trackId
}

func (p *ProjectImageResponseModel) SetOffsetMs(offsetMs int) {
	p.OffsetMs = offsetMs
}

func (p *ProjectImageResponseModel) SetDurationMs(durationMs int) {
	p.DurationMs = durationMs
}

func (p *ProjectImageResponseModel) SetOrder(order string) {
	p.Order = order
}

func (p *ProjectImageResponseModel) SetCanvasPlacement(canvasPlacement *CanvasPlacement) {
	p.CanvasPlacement = canvasPlacement
}

func (p *ProjectImageResponseModel) SetAnimation(animation *ClipAnimation) {
	p.Animation = animation
}

func (p *ProjectImageResponseModel) SetOpacity(opacity *float64) {
	p.Opacity = opacity
}

func (p *ProjectImageResponseModel) SetCreatedAtMs(createdAtMs int) {
	p.CreatedAtMs = createdAtMs
}

func (p *ProjectImageResponseModel) SetUpdatedAtMs(updatedAtMs int) {
	p.UpdatedAtMs = updatedAtMs
}

func (p *ProjectImageResponseModel) SetCurrentSnapshotId(currentSnapshotId *string) {
	p.CurrentSnapshotId = currentSnapshotId
}

func (p *ProjectImageResponseModel) SetSourceAssetId(sourceAssetId *string) {
	p.SourceAssetId = sourceAssetId
}

func (p *ProjectImageResponseModel) SetSourcePlatformAssetId(sourcePlatformAssetId *string) {
	p.SourcePlatformAssetId = sourcePlatformAssetId
}

func (p *ProjectImageResponseModel) SetError(error *string) {
	p.Error = error
}

func (p *ProjectImageResponseModel) SetPendingTask(pendingTask *PendingClipTask) {
	p.PendingTask = pendingTask
}

func (p *ProjectImageResponseModel) SetAnalysis(analysis *ImageAnalysis) {
	p.Analysis = analysis
}

func (p *ProjectImageResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectImageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectImageResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProjectImageResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *ProjectImageResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProjectImageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
