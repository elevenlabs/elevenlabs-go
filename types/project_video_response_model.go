package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProjectVideoResponseModel struct {
	VideoId                       string                                     `json:"video_id" xml:"video_id" required:"true"`
	Filename                      string                                     `json:"filename" xml:"filename" required:"true"`
	SignedUrl                     *string                                    `json:"signed_url" xml:"signed_url" required:"true"`
	SignedPreviewUrl              *string                                    `json:"signed_preview_url" xml:"signed_preview_url" required:"true"`
	OffsetMs                      int                                        `json:"offset_ms" xml:"offset_ms" required:"true"`
	DurationMs                    int                                        `json:"duration_ms" xml:"duration_ms" required:"true"`
	VolumeGainDb                  float64                                    `json:"volume_gain_db" xml:"volume_gain_db" required:"true"`
	Muted                         bool                                       `json:"muted" xml:"muted" required:"true"`
	FadeInMs                      *int                                       `json:"fade_in_ms,omitempty" xml:"fade_in_ms,omitempty"`
	FadeOutMs                     *int                                       `json:"fade_out_ms,omitempty" xml:"fade_out_ms,omitempty"`
	Width                         int                                        `json:"width" xml:"width" required:"true"`
	Height                        int                                        `json:"height" xml:"height" required:"true"`
	Codec                         string                                     `json:"codec" xml:"codec" required:"true"`
	Order                         string                                     `json:"order" xml:"order" required:"true"`
	CreatedAtMs                   int                                        `json:"created_at_ms" xml:"created_at_ms" required:"true"`
	UpdatedAtMs                   int                                        `json:"updated_at_ms" xml:"updated_at_ms" required:"true"`
	Error                         *string                                    `json:"error,omitempty" xml:"error,omitempty"`
	ThumbnailIntervalSeconds      float64                                    `json:"thumbnail_interval_seconds" xml:"thumbnail_interval_seconds" required:"true"`
	ThumbnailSize                 []int                                      `json:"thumbnail_size" xml:"thumbnail_size" required:"true" minItems:"2" maxItems:"2"`
	ThumbnailSheets               []*ProjectVideoThumbnailSheetResponseModel `json:"thumbnail_sheets" xml:"thumbnail_sheets" required:"true"`
	StartTimeMs                   int                                        `json:"start_time_ms" xml:"start_time_ms" required:"true"`
	EndTimeMs                     *int                                       `json:"end_time_ms" xml:"end_time_ms" required:"true"`
	AssetPreviewSignedUrl         *string                                    `json:"asset_preview_signed_url,omitempty" xml:"asset_preview_signed_url,omitempty"`
	SourceVideoId                 *string                                    `json:"source_video_id,omitempty" xml:"source_video_id,omitempty"`
	SourceAssetId                 *string                                    `json:"source_asset_id,omitempty" xml:"source_asset_id,omitempty"`
	SourcePlatformAssetId         *string                                    `json:"source_platform_asset_id,omitempty" xml:"source_platform_asset_id,omitempty"`
	PendingBlocksMetadata         *PendingBlocksMetadataModel                `json:"pending_blocks_metadata,omitempty" xml:"pending_blocks_metadata,omitempty"`
	PendingExternalAudiosMetadata *PendingExternalAudiosMetadataModel        `json:"pending_external_audios_metadata,omitempty" xml:"pending_external_audios_metadata,omitempty"`
	SpeechImported                *bool                                      `json:"speech_imported,omitempty" xml:"speech_imported,omitempty"`
	PendingTask                   *PendingClipTask                           `json:"pending_task,omitempty" xml:"pending_task,omitempty"`
	AudioTrackReady               *bool                                      `json:"audio_track_ready,omitempty" xml:"audio_track_ready,omitempty"`
	ExportFormatReady             *bool                                      `json:"export_format_ready,omitempty" xml:"export_format_ready,omitempty"`
	CurrentSnapshotId             *string                                    `json:"current_snapshot_id,omitempty" xml:"current_snapshot_id,omitempty"`
	SourceContext                 *GenerationSourceContext                   `json:"source_context,omitempty" xml:"source_context,omitempty"`
	Analysis                      *VideoAnalysis                             `json:"analysis,omitempty" xml:"analysis,omitempty"`
	Transcription                 *AssetTranscription                        `json:"transcription,omitempty" xml:"transcription,omitempty"`
	Type                          *string                                    `json:"type,omitempty" xml:"type,omitempty"`
	// Defines asset positioning and transformation on canvas.
	CanvasPlacement      *CanvasPlacement `json:"canvas_placement,omitempty" xml:"canvas_placement,omitempty"`
	Animation            *ClipAnimation   `json:"animation,omitempty" xml:"animation,omitempty"`
	PlaybackSpeed        *float64         `json:"playback_speed,omitempty" xml:"playback_speed,omitempty"`
	Opacity              *float64         `json:"opacity,omitempty" xml:"opacity,omitempty" min:"0" max:"1"`
	TrackId              *string          `json:"track_id,omitempty" xml:"track_id,omitempty"`
	PreviewJobProgress   *float64         `json:"preview_job_progress,omitempty" xml:"preview_job_progress,omitempty"`
	ImportSpeechProgress *float64         `json:"import_speech_progress,omitempty" xml:"import_speech_progress,omitempty"`
	extraProperties      map[string]interface{}
}

func (p *ProjectVideoResponseModel) GetVideoId() (value string) {
	if p == nil {
		return
	}
	return p.VideoId
}

func (p *ProjectVideoResponseModel) GetFilename() (value string) {
	if p == nil {
		return
	}
	return p.Filename
}

func (p *ProjectVideoResponseModel) GetSignedUrl() (value *string) {
	if p == nil {
		return
	}
	return p.SignedUrl
}

func (p *ProjectVideoResponseModel) GetSignedPreviewUrl() (value *string) {
	if p == nil {
		return
	}
	return p.SignedPreviewUrl
}

func (p *ProjectVideoResponseModel) GetOffsetMs() (value int) {
	if p == nil {
		return
	}
	return p.OffsetMs
}

func (p *ProjectVideoResponseModel) GetDurationMs() (value int) {
	if p == nil {
		return
	}
	return p.DurationMs
}

func (p *ProjectVideoResponseModel) GetVolumeGainDb() (value float64) {
	if p == nil {
		return
	}
	return p.VolumeGainDb
}

func (p *ProjectVideoResponseModel) GetMuted() (value bool) {
	if p == nil {
		return
	}
	return p.Muted
}

func (p *ProjectVideoResponseModel) GetFadeInMs() (value *int) {
	if p == nil {
		return
	}
	return p.FadeInMs
}

func (p *ProjectVideoResponseModel) GetFadeOutMs() (value *int) {
	if p == nil {
		return
	}
	return p.FadeOutMs
}

func (p *ProjectVideoResponseModel) GetWidth() (value int) {
	if p == nil {
		return
	}
	return p.Width
}

func (p *ProjectVideoResponseModel) GetHeight() (value int) {
	if p == nil {
		return
	}
	return p.Height
}

func (p *ProjectVideoResponseModel) GetCodec() (value string) {
	if p == nil {
		return
	}
	return p.Codec
}

func (p *ProjectVideoResponseModel) GetOrder() (value string) {
	if p == nil {
		return
	}
	return p.Order
}

func (p *ProjectVideoResponseModel) GetCreatedAtMs() (value int) {
	if p == nil {
		return
	}
	return p.CreatedAtMs
}

func (p *ProjectVideoResponseModel) GetUpdatedAtMs() (value int) {
	if p == nil {
		return
	}
	return p.UpdatedAtMs
}

func (p *ProjectVideoResponseModel) GetError() (value *string) {
	if p == nil {
		return
	}
	return p.Error
}

func (p *ProjectVideoResponseModel) GetThumbnailIntervalSeconds() (value float64) {
	if p == nil {
		return
	}
	return p.ThumbnailIntervalSeconds
}

func (p *ProjectVideoResponseModel) GetThumbnailSize() (value []int) {
	if p == nil {
		return
	}
	return p.ThumbnailSize
}

func (p *ProjectVideoResponseModel) GetThumbnailSheets() (value []*ProjectVideoThumbnailSheetResponseModel) {
	if p == nil {
		return
	}
	return p.ThumbnailSheets
}

func (p *ProjectVideoResponseModel) GetStartTimeMs() (value int) {
	if p == nil {
		return
	}
	return p.StartTimeMs
}

func (p *ProjectVideoResponseModel) GetEndTimeMs() (value *int) {
	if p == nil {
		return
	}
	return p.EndTimeMs
}

func (p *ProjectVideoResponseModel) GetAssetPreviewSignedUrl() (value *string) {
	if p == nil {
		return
	}
	return p.AssetPreviewSignedUrl
}

func (p *ProjectVideoResponseModel) GetSourceVideoId() (value *string) {
	if p == nil {
		return
	}
	return p.SourceVideoId
}

func (p *ProjectVideoResponseModel) GetSourceAssetId() (value *string) {
	if p == nil {
		return
	}
	return p.SourceAssetId
}

func (p *ProjectVideoResponseModel) GetSourcePlatformAssetId() (value *string) {
	if p == nil {
		return
	}
	return p.SourcePlatformAssetId
}

func (p *ProjectVideoResponseModel) GetPendingBlocksMetadata() (value *PendingBlocksMetadataModel) {
	if p == nil {
		return
	}
	return p.PendingBlocksMetadata
}

func (p *ProjectVideoResponseModel) GetPendingExternalAudiosMetadata() (value *PendingExternalAudiosMetadataModel) {
	if p == nil {
		return
	}
	return p.PendingExternalAudiosMetadata
}

func (p *ProjectVideoResponseModel) GetSpeechImported() (value *bool) {
	if p == nil {
		return
	}
	return p.SpeechImported
}

func (p *ProjectVideoResponseModel) GetPendingTask() (value *PendingClipTask) {
	if p == nil {
		return
	}
	return p.PendingTask
}

func (p *ProjectVideoResponseModel) GetAudioTrackReady() (value *bool) {
	if p == nil {
		return
	}
	return p.AudioTrackReady
}

func (p *ProjectVideoResponseModel) GetExportFormatReady() (value *bool) {
	if p == nil {
		return
	}
	return p.ExportFormatReady
}

func (p *ProjectVideoResponseModel) GetCurrentSnapshotId() (value *string) {
	if p == nil {
		return
	}
	return p.CurrentSnapshotId
}

func (p *ProjectVideoResponseModel) GetSourceContext() (value *GenerationSourceContext) {
	if p == nil {
		return
	}
	return p.SourceContext
}

func (p *ProjectVideoResponseModel) GetAnalysis() (value *VideoAnalysis) {
	if p == nil {
		return
	}
	return p.Analysis
}

func (p *ProjectVideoResponseModel) GetTranscription() (value *AssetTranscription) {
	if p == nil {
		return
	}
	return p.Transcription
}

func (p *ProjectVideoResponseModel) GetType() (value *string) {
	if p == nil {
		return
	}
	return p.Type
}

func (p *ProjectVideoResponseModel) GetCanvasPlacement() (value *CanvasPlacement) {
	if p == nil {
		return
	}
	return p.CanvasPlacement
}

func (p *ProjectVideoResponseModel) GetAnimation() (value *ClipAnimation) {
	if p == nil {
		return
	}
	return p.Animation
}

func (p *ProjectVideoResponseModel) GetPlaybackSpeed() (value *float64) {
	if p == nil {
		return
	}
	return p.PlaybackSpeed
}

func (p *ProjectVideoResponseModel) GetOpacity() (value *float64) {
	if p == nil {
		return
	}
	return p.Opacity
}

func (p *ProjectVideoResponseModel) GetTrackId() (value *string) {
	if p == nil {
		return
	}
	return p.TrackId
}

func (p *ProjectVideoResponseModel) GetPreviewJobProgress() (value *float64) {
	if p == nil {
		return
	}
	return p.PreviewJobProgress
}

func (p *ProjectVideoResponseModel) GetImportSpeechProgress() (value *float64) {
	if p == nil {
		return
	}
	return p.ImportSpeechProgress
}

func (p *ProjectVideoResponseModel) SetVideoId(videoId string) {
	p.VideoId = videoId
}

func (p *ProjectVideoResponseModel) SetFilename(filename string) {
	p.Filename = filename
}

func (p *ProjectVideoResponseModel) SetSignedUrl(signedUrl *string) {
	p.SignedUrl = signedUrl
}

func (p *ProjectVideoResponseModel) SetSignedPreviewUrl(signedPreviewUrl *string) {
	p.SignedPreviewUrl = signedPreviewUrl
}

func (p *ProjectVideoResponseModel) SetOffsetMs(offsetMs int) {
	p.OffsetMs = offsetMs
}

func (p *ProjectVideoResponseModel) SetDurationMs(durationMs int) {
	p.DurationMs = durationMs
}

func (p *ProjectVideoResponseModel) SetVolumeGainDb(volumeGainDb float64) {
	p.VolumeGainDb = volumeGainDb
}

func (p *ProjectVideoResponseModel) SetMuted(muted bool) {
	p.Muted = muted
}

func (p *ProjectVideoResponseModel) SetFadeInMs(fadeInMs *int) {
	p.FadeInMs = fadeInMs
}

func (p *ProjectVideoResponseModel) SetFadeOutMs(fadeOutMs *int) {
	p.FadeOutMs = fadeOutMs
}

func (p *ProjectVideoResponseModel) SetWidth(width int) {
	p.Width = width
}

func (p *ProjectVideoResponseModel) SetHeight(height int) {
	p.Height = height
}

func (p *ProjectVideoResponseModel) SetCodec(codec string) {
	p.Codec = codec
}

func (p *ProjectVideoResponseModel) SetOrder(order string) {
	p.Order = order
}

func (p *ProjectVideoResponseModel) SetCreatedAtMs(createdAtMs int) {
	p.CreatedAtMs = createdAtMs
}

func (p *ProjectVideoResponseModel) SetUpdatedAtMs(updatedAtMs int) {
	p.UpdatedAtMs = updatedAtMs
}

func (p *ProjectVideoResponseModel) SetError(error *string) {
	p.Error = error
}

func (p *ProjectVideoResponseModel) SetThumbnailIntervalSeconds(thumbnailIntervalSeconds float64) {
	p.ThumbnailIntervalSeconds = thumbnailIntervalSeconds
}

func (p *ProjectVideoResponseModel) SetThumbnailSize(thumbnailSize []int) {
	p.ThumbnailSize = thumbnailSize
}

func (p *ProjectVideoResponseModel) SetThumbnailSheets(thumbnailSheets []*ProjectVideoThumbnailSheetResponseModel) {
	p.ThumbnailSheets = thumbnailSheets
}

func (p *ProjectVideoResponseModel) SetStartTimeMs(startTimeMs int) {
	p.StartTimeMs = startTimeMs
}

func (p *ProjectVideoResponseModel) SetEndTimeMs(endTimeMs *int) {
	p.EndTimeMs = endTimeMs
}

func (p *ProjectVideoResponseModel) SetAssetPreviewSignedUrl(assetPreviewSignedUrl *string) {
	p.AssetPreviewSignedUrl = assetPreviewSignedUrl
}

func (p *ProjectVideoResponseModel) SetSourceVideoId(sourceVideoId *string) {
	p.SourceVideoId = sourceVideoId
}

func (p *ProjectVideoResponseModel) SetSourceAssetId(sourceAssetId *string) {
	p.SourceAssetId = sourceAssetId
}

func (p *ProjectVideoResponseModel) SetSourcePlatformAssetId(sourcePlatformAssetId *string) {
	p.SourcePlatformAssetId = sourcePlatformAssetId
}

func (p *ProjectVideoResponseModel) SetPendingBlocksMetadata(pendingBlocksMetadata *PendingBlocksMetadataModel) {
	p.PendingBlocksMetadata = pendingBlocksMetadata
}

func (p *ProjectVideoResponseModel) SetPendingExternalAudiosMetadata(pendingExternalAudiosMetadata *PendingExternalAudiosMetadataModel) {
	p.PendingExternalAudiosMetadata = pendingExternalAudiosMetadata
}

func (p *ProjectVideoResponseModel) SetSpeechImported(speechImported *bool) {
	p.SpeechImported = speechImported
}

func (p *ProjectVideoResponseModel) SetPendingTask(pendingTask *PendingClipTask) {
	p.PendingTask = pendingTask
}

func (p *ProjectVideoResponseModel) SetAudioTrackReady(audioTrackReady *bool) {
	p.AudioTrackReady = audioTrackReady
}

func (p *ProjectVideoResponseModel) SetExportFormatReady(exportFormatReady *bool) {
	p.ExportFormatReady = exportFormatReady
}

func (p *ProjectVideoResponseModel) SetCurrentSnapshotId(currentSnapshotId *string) {
	p.CurrentSnapshotId = currentSnapshotId
}

func (p *ProjectVideoResponseModel) SetSourceContext(sourceContext *GenerationSourceContext) {
	p.SourceContext = sourceContext
}

func (p *ProjectVideoResponseModel) SetAnalysis(analysis *VideoAnalysis) {
	p.Analysis = analysis
}

func (p *ProjectVideoResponseModel) SetTranscription(transcription *AssetTranscription) {
	p.Transcription = transcription
}

func (p *ProjectVideoResponseModel) SetType(type_ *string) {
	p.Type = type_
}

func (p *ProjectVideoResponseModel) SetCanvasPlacement(canvasPlacement *CanvasPlacement) {
	p.CanvasPlacement = canvasPlacement
}

func (p *ProjectVideoResponseModel) SetAnimation(animation *ClipAnimation) {
	p.Animation = animation
}

func (p *ProjectVideoResponseModel) SetPlaybackSpeed(playbackSpeed *float64) {
	p.PlaybackSpeed = playbackSpeed
}

func (p *ProjectVideoResponseModel) SetOpacity(opacity *float64) {
	p.Opacity = opacity
}

func (p *ProjectVideoResponseModel) SetTrackId(trackId *string) {
	p.TrackId = trackId
}

func (p *ProjectVideoResponseModel) SetPreviewJobProgress(previewJobProgress *float64) {
	p.PreviewJobProgress = previewJobProgress
}

func (p *ProjectVideoResponseModel) SetImportSpeechProgress(importSpeechProgress *float64) {
	p.ImportSpeechProgress = importSpeechProgress
}

func (p *ProjectVideoResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectVideoResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectVideoResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProjectVideoResponseModel) UnmarshalJSON(data []byte) error {
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

func (p *ProjectVideoResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProjectVideoResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
