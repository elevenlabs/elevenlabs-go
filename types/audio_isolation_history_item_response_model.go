package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioIsolationHistoryItemResponseModel struct {
	Id                    string   `json:"id" xml:"id" required:"true"`
	Title                 *string  `json:"title" xml:"title" required:"true"`
	CreatedAtUnix         int      `json:"created_at_unix" xml:"created_at_unix" required:"true"`
	Format                string   `json:"format" xml:"format" required:"true"`
	DurationSeconds       *float64 `json:"duration_seconds" xml:"duration_seconds" required:"true"`
	DownloadUrl           *string  `json:"download_url" xml:"download_url" required:"true"`
	IconUrl               *string  `json:"icon_url" xml:"icon_url" required:"true"`
	SourceVideoUrl        *string  `json:"source_video_url" xml:"source_video_url" required:"true"`
	SupportsVideo         bool     `json:"supports_video" xml:"supports_video" required:"true"`
	Processing            bool     `json:"processing" xml:"processing" required:"true"`
	VideoProcessingFailed bool     `json:"video_processing_failed" xml:"video_processing_failed" required:"true"`
	PreviewB64            *string  `json:"preview_b64" xml:"preview_b64" required:"true"`
	extraProperties       map[string]interface{}
}

func (a *AudioIsolationHistoryItemResponseModel) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AudioIsolationHistoryItemResponseModel) GetTitle() (value *string) {
	if a == nil {
		return
	}
	return a.Title
}

func (a *AudioIsolationHistoryItemResponseModel) GetCreatedAtUnix() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAtUnix
}

func (a *AudioIsolationHistoryItemResponseModel) GetFormat() (value string) {
	if a == nil {
		return
	}
	return a.Format
}

func (a *AudioIsolationHistoryItemResponseModel) GetDurationSeconds() (value *float64) {
	if a == nil {
		return
	}
	return a.DurationSeconds
}

func (a *AudioIsolationHistoryItemResponseModel) GetDownloadUrl() (value *string) {
	if a == nil {
		return
	}
	return a.DownloadUrl
}

func (a *AudioIsolationHistoryItemResponseModel) GetIconUrl() (value *string) {
	if a == nil {
		return
	}
	return a.IconUrl
}

func (a *AudioIsolationHistoryItemResponseModel) GetSourceVideoUrl() (value *string) {
	if a == nil {
		return
	}
	return a.SourceVideoUrl
}

func (a *AudioIsolationHistoryItemResponseModel) GetSupportsVideo() (value bool) {
	if a == nil {
		return
	}
	return a.SupportsVideo
}

func (a *AudioIsolationHistoryItemResponseModel) GetProcessing() (value bool) {
	if a == nil {
		return
	}
	return a.Processing
}

func (a *AudioIsolationHistoryItemResponseModel) GetVideoProcessingFailed() (value bool) {
	if a == nil {
		return
	}
	return a.VideoProcessingFailed
}

func (a *AudioIsolationHistoryItemResponseModel) GetPreviewB64() (value *string) {
	if a == nil {
		return
	}
	return a.PreviewB64
}

func (a *AudioIsolationHistoryItemResponseModel) SetId(id string) {
	a.Id = id
}

func (a *AudioIsolationHistoryItemResponseModel) SetTitle(title *string) {
	a.Title = title
}

func (a *AudioIsolationHistoryItemResponseModel) SetCreatedAtUnix(createdAtUnix int) {
	a.CreatedAtUnix = createdAtUnix
}

func (a *AudioIsolationHistoryItemResponseModel) SetFormat(format string) {
	a.Format = format
}

func (a *AudioIsolationHistoryItemResponseModel) SetDurationSeconds(durationSeconds *float64) {
	a.DurationSeconds = durationSeconds
}

func (a *AudioIsolationHistoryItemResponseModel) SetDownloadUrl(downloadUrl *string) {
	a.DownloadUrl = downloadUrl
}

func (a *AudioIsolationHistoryItemResponseModel) SetIconUrl(iconUrl *string) {
	a.IconUrl = iconUrl
}

func (a *AudioIsolationHistoryItemResponseModel) SetSourceVideoUrl(sourceVideoUrl *string) {
	a.SourceVideoUrl = sourceVideoUrl
}

func (a *AudioIsolationHistoryItemResponseModel) SetSupportsVideo(supportsVideo bool) {
	a.SupportsVideo = supportsVideo
}

func (a *AudioIsolationHistoryItemResponseModel) SetProcessing(processing bool) {
	a.Processing = processing
}

func (a *AudioIsolationHistoryItemResponseModel) SetVideoProcessingFailed(videoProcessingFailed bool) {
	a.VideoProcessingFailed = videoProcessingFailed
}

func (a *AudioIsolationHistoryItemResponseModel) SetPreviewB64(previewB64 *string) {
	a.PreviewB64 = previewB64
}

func (a *AudioIsolationHistoryItemResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioIsolationHistoryItemResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioIsolationHistoryItemResponseModel to string"
	}
	return string(jsonData)
}

func (a *AudioIsolationHistoryItemResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AudioIsolationHistoryItemResponseModel) MarshalJSON() ([]byte, error) {
	type embed AudioIsolationHistoryItemResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
