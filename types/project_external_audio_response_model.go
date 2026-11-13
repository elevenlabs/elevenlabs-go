package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ProjectExternalAudioResponseModel struct {
	ExternalAudioId               string                                          `json:"external_audio_id" xml:"external_audio_id" required:"true"`
	Filename                      string                                          `json:"filename" xml:"filename" required:"true"`
	SignedUrl                     *string                                         `json:"signed_url" xml:"signed_url" required:"true"`
	OffsetMs                      int                                             `json:"offset_ms" xml:"offset_ms" required:"true"`
	DurationMs                    int                                             `json:"duration_ms" xml:"duration_ms" required:"true"`
	StartTimeMs                   int                                             `json:"start_time_ms" xml:"start_time_ms" required:"true"`
	EndTimeMs                     *int                                            `json:"end_time_ms" xml:"end_time_ms" required:"true"`
	Order                         string                                          `json:"order" xml:"order" required:"true"`
	TrackId                       string                                          `json:"track_id" xml:"track_id" required:"true"`
	CreatedAtMs                   int                                             `json:"created_at_ms" xml:"created_at_ms" required:"true"`
	UpdatedAtMs                   int                                             `json:"updated_at_ms" xml:"updated_at_ms" required:"true"`
	VolumeGainDb                  *float64                                        `json:"volume_gain_db,omitempty" xml:"volume_gain_db,omitempty"`
	Muted                         *bool                                           `json:"muted,omitempty" xml:"muted,omitempty"`
	FadeInMs                      *int                                            `json:"fade_in_ms,omitempty" xml:"fade_in_ms,omitempty"`
	FadeOutMs                     *int                                            `json:"fade_out_ms,omitempty" xml:"fade_out_ms,omitempty"`
	SourceExternalAudioId         *string                                         `json:"source_external_audio_id,omitempty" xml:"source_external_audio_id,omitempty"`
	SourceAssetId                 *string                                         `json:"source_asset_id,omitempty" xml:"source_asset_id,omitempty"`
	SourcePlatformAssetId         *string                                         `json:"source_platform_asset_id,omitempty" xml:"source_platform_asset_id,omitempty"`
	PendingBlocksMetadata         *PendingBlocksMetadataModel                     `json:"pending_blocks_metadata,omitempty" xml:"pending_blocks_metadata,omitempty"`
	PendingExternalAudiosMetadata *PendingExternalAudiosMetadataModel             `json:"pending_external_audios_metadata,omitempty" xml:"pending_external_audios_metadata,omitempty"`
	SpeechImported                *bool                                           `json:"speech_imported,omitempty" xml:"speech_imported,omitempty"`
	PendingTask                   *PendingClipTask                                `json:"pending_task,omitempty" xml:"pending_task,omitempty"`
	Error                         *string                                         `json:"error,omitempty" xml:"error,omitempty"`
	CurrentSnapshotId             *string                                         `json:"current_snapshot_id,omitempty" xml:"current_snapshot_id,omitempty"`
	SourceContext                 *ProjectExternalAudioResponseModelSourceContext `json:"source_context,omitempty" xml:"source_context,omitempty"`
	Analysis                      *AudioAnalysis                                  `json:"analysis,omitempty" xml:"analysis,omitempty"`
	Transcription                 *AssetTranscription                             `json:"transcription,omitempty" xml:"transcription,omitempty"`
	Type                          *string                                         `json:"type,omitempty" xml:"type,omitempty"`
	ImportSpeechProgress          *float64                                        `json:"import_speech_progress,omitempty" xml:"import_speech_progress,omitempty"`
	extraProperties               map[string]interface{}
}

func (p *ProjectExternalAudioResponseModel) GetExternalAudioId() (value string) {
	if p == nil {
		return
	}
	return p.ExternalAudioId
}

func (p *ProjectExternalAudioResponseModel) GetFilename() (value string) {
	if p == nil {
		return
	}
	return p.Filename
}

func (p *ProjectExternalAudioResponseModel) GetSignedUrl() (value *string) {
	if p == nil {
		return
	}
	return p.SignedUrl
}

func (p *ProjectExternalAudioResponseModel) GetOffsetMs() (value int) {
	if p == nil {
		return
	}
	return p.OffsetMs
}

func (p *ProjectExternalAudioResponseModel) GetDurationMs() (value int) {
	if p == nil {
		return
	}
	return p.DurationMs
}

func (p *ProjectExternalAudioResponseModel) GetStartTimeMs() (value int) {
	if p == nil {
		return
	}
	return p.StartTimeMs
}

func (p *ProjectExternalAudioResponseModel) GetEndTimeMs() (value *int) {
	if p == nil {
		return
	}
	return p.EndTimeMs
}

func (p *ProjectExternalAudioResponseModel) GetOrder() (value string) {
	if p == nil {
		return
	}
	return p.Order
}

func (p *ProjectExternalAudioResponseModel) GetTrackId() (value string) {
	if p == nil {
		return
	}
	return p.TrackId
}

func (p *ProjectExternalAudioResponseModel) GetCreatedAtMs() (value int) {
	if p == nil {
		return
	}
	return p.CreatedAtMs
}

func (p *ProjectExternalAudioResponseModel) GetUpdatedAtMs() (value int) {
	if p == nil {
		return
	}
	return p.UpdatedAtMs
}

func (p *ProjectExternalAudioResponseModel) GetVolumeGainDb() (value *float64) {
	if p == nil {
		return
	}
	return p.VolumeGainDb
}

func (p *ProjectExternalAudioResponseModel) GetMuted() (value *bool) {
	if p == nil {
		return
	}
	return p.Muted
}

func (p *ProjectExternalAudioResponseModel) GetFadeInMs() (value *int) {
	if p == nil {
		return
	}
	return p.FadeInMs
}

func (p *ProjectExternalAudioResponseModel) GetFadeOutMs() (value *int) {
	if p == nil {
		return
	}
	return p.FadeOutMs
}

func (p *ProjectExternalAudioResponseModel) GetSourceExternalAudioId() (value *string) {
	if p == nil {
		return
	}
	return p.SourceExternalAudioId
}

func (p *ProjectExternalAudioResponseModel) GetSourceAssetId() (value *string) {
	if p == nil {
		return
	}
	return p.SourceAssetId
}

func (p *ProjectExternalAudioResponseModel) GetSourcePlatformAssetId() (value *string) {
	if p == nil {
		return
	}
	return p.SourcePlatformAssetId
}

func (p *ProjectExternalAudioResponseModel) GetPendingBlocksMetadata() (value *PendingBlocksMetadataModel) {
	if p == nil {
		return
	}
	return p.PendingBlocksMetadata
}

func (p *ProjectExternalAudioResponseModel) GetPendingExternalAudiosMetadata() (value *PendingExternalAudiosMetadataModel) {
	if p == nil {
		return
	}
	return p.PendingExternalAudiosMetadata
}

func (p *ProjectExternalAudioResponseModel) GetSpeechImported() (value *bool) {
	if p == nil {
		return
	}
	return p.SpeechImported
}

func (p *ProjectExternalAudioResponseModel) GetPendingTask() (value *PendingClipTask) {
	if p == nil {
		return
	}
	return p.PendingTask
}

func (p *ProjectExternalAudioResponseModel) GetError() (value *string) {
	if p == nil {
		return
	}
	return p.Error
}

func (p *ProjectExternalAudioResponseModel) GetCurrentSnapshotId() (value *string) {
	if p == nil {
		return
	}
	return p.CurrentSnapshotId
}

func (p *ProjectExternalAudioResponseModel) GetSourceContext() (value *ProjectExternalAudioResponseModelSourceContext) {
	if p == nil {
		return
	}
	return p.SourceContext
}

func (p *ProjectExternalAudioResponseModel) GetAnalysis() (value *AudioAnalysis) {
	if p == nil {
		return
	}
	return p.Analysis
}

func (p *ProjectExternalAudioResponseModel) GetTranscription() (value *AssetTranscription) {
	if p == nil {
		return
	}
	return p.Transcription
}

func (p *ProjectExternalAudioResponseModel) GetType() (value *string) {
	if p == nil {
		return
	}
	return p.Type
}

func (p *ProjectExternalAudioResponseModel) GetImportSpeechProgress() (value *float64) {
	if p == nil {
		return
	}
	return p.ImportSpeechProgress
}

func (p *ProjectExternalAudioResponseModel) SetExternalAudioId(externalAudioId string) {
	p.ExternalAudioId = externalAudioId
}

func (p *ProjectExternalAudioResponseModel) SetFilename(filename string) {
	p.Filename = filename
}

func (p *ProjectExternalAudioResponseModel) SetSignedUrl(signedUrl *string) {
	p.SignedUrl = signedUrl
}

func (p *ProjectExternalAudioResponseModel) SetOffsetMs(offsetMs int) {
	p.OffsetMs = offsetMs
}

func (p *ProjectExternalAudioResponseModel) SetDurationMs(durationMs int) {
	p.DurationMs = durationMs
}

func (p *ProjectExternalAudioResponseModel) SetStartTimeMs(startTimeMs int) {
	p.StartTimeMs = startTimeMs
}

func (p *ProjectExternalAudioResponseModel) SetEndTimeMs(endTimeMs *int) {
	p.EndTimeMs = endTimeMs
}

func (p *ProjectExternalAudioResponseModel) SetOrder(order string) {
	p.Order = order
}

func (p *ProjectExternalAudioResponseModel) SetTrackId(trackId string) {
	p.TrackId = trackId
}

func (p *ProjectExternalAudioResponseModel) SetCreatedAtMs(createdAtMs int) {
	p.CreatedAtMs = createdAtMs
}

func (p *ProjectExternalAudioResponseModel) SetUpdatedAtMs(updatedAtMs int) {
	p.UpdatedAtMs = updatedAtMs
}

func (p *ProjectExternalAudioResponseModel) SetVolumeGainDb(volumeGainDb *float64) {
	p.VolumeGainDb = volumeGainDb
}

func (p *ProjectExternalAudioResponseModel) SetMuted(muted *bool) {
	p.Muted = muted
}

func (p *ProjectExternalAudioResponseModel) SetFadeInMs(fadeInMs *int) {
	p.FadeInMs = fadeInMs
}

func (p *ProjectExternalAudioResponseModel) SetFadeOutMs(fadeOutMs *int) {
	p.FadeOutMs = fadeOutMs
}

func (p *ProjectExternalAudioResponseModel) SetSourceExternalAudioId(sourceExternalAudioId *string) {
	p.SourceExternalAudioId = sourceExternalAudioId
}

func (p *ProjectExternalAudioResponseModel) SetSourceAssetId(sourceAssetId *string) {
	p.SourceAssetId = sourceAssetId
}

func (p *ProjectExternalAudioResponseModel) SetSourcePlatformAssetId(sourcePlatformAssetId *string) {
	p.SourcePlatformAssetId = sourcePlatformAssetId
}

func (p *ProjectExternalAudioResponseModel) SetPendingBlocksMetadata(pendingBlocksMetadata *PendingBlocksMetadataModel) {
	p.PendingBlocksMetadata = pendingBlocksMetadata
}

func (p *ProjectExternalAudioResponseModel) SetPendingExternalAudiosMetadata(pendingExternalAudiosMetadata *PendingExternalAudiosMetadataModel) {
	p.PendingExternalAudiosMetadata = pendingExternalAudiosMetadata
}

func (p *ProjectExternalAudioResponseModel) SetSpeechImported(speechImported *bool) {
	p.SpeechImported = speechImported
}

func (p *ProjectExternalAudioResponseModel) SetPendingTask(pendingTask *PendingClipTask) {
	p.PendingTask = pendingTask
}

func (p *ProjectExternalAudioResponseModel) SetError(error *string) {
	p.Error = error
}

func (p *ProjectExternalAudioResponseModel) SetCurrentSnapshotId(currentSnapshotId *string) {
	p.CurrentSnapshotId = currentSnapshotId
}

func (p *ProjectExternalAudioResponseModel) SetSourceContext(sourceContext *ProjectExternalAudioResponseModelSourceContext) {
	p.SourceContext = sourceContext
}

func (p *ProjectExternalAudioResponseModel) SetAnalysis(analysis *AudioAnalysis) {
	p.Analysis = analysis
}

func (p *ProjectExternalAudioResponseModel) SetTranscription(transcription *AssetTranscription) {
	p.Transcription = transcription
}

func (p *ProjectExternalAudioResponseModel) SetType(type_ *string) {
	p.Type = type_
}

func (p *ProjectExternalAudioResponseModel) SetImportSpeechProgress(importSpeechProgress *float64) {
	p.ImportSpeechProgress = importSpeechProgress
}

func (p *ProjectExternalAudioResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p ProjectExternalAudioResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: ProjectExternalAudioResponseModel to string"
	}
	return string(jsonData)
}

func (p *ProjectExternalAudioResponseModel) UnmarshalJSON(data []byte) error {
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

func (p *ProjectExternalAudioResponseModel) MarshalJSON() ([]byte, error) {
	type embed ProjectExternalAudioResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
