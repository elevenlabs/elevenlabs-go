package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubRequest struct {
	// A list of file paths to audio recordings intended for voice cloning
	File *string `json:"file,omitempty" xml:"file,omitempty"`
	// CSV file containing transcription/translation metadata
	CsvFile *string `json:"csv_file,omitempty" xml:"csv_file,omitempty"`
	// For use only with csv input
	ForegroundAudioFile *string `json:"foreground_audio_file,omitempty" xml:"foreground_audio_file,omitempty"`
	// For use only with csv input
	BackgroundAudioFile *string `json:"background_audio_file,omitempty" xml:"background_audio_file,omitempty"`
	// Name of the dubbing project.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// URL of the source video/audio file.
	SourceUrl *string `json:"source_url,omitempty" xml:"source_url,omitempty"`
	// Source language. Expects a valid iso639-1 or iso639-3 language code.
	SourceLang *string `json:"source_lang,omitempty" xml:"source_lang,omitempty"`
	// The Target language to dub the content into. Expects a valid iso639-1 or iso639-3 language code.
	TargetLang *string `json:"target_lang" xml:"target_lang" required:"true"`
	// [Experimental] An accent to apply when selecting voices from the library and to use to inform translation of the dialect to prefer.
	TargetAccent *string `json:"target_accent,omitempty" xml:"target_accent,omitempty"`
	// Number of speakers to use for the dubbing. Set to 0 to automatically detect the number of speakers
	NumSpeakers *int `json:"num_speakers,omitempty" xml:"num_speakers,omitempty"`
	// Whether to apply watermark to the output video.
	Watermark *bool `json:"watermark,omitempty" xml:"watermark,omitempty"`
	// Start time of the source video/audio file.
	StartTime *int `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// End time of the source video/audio file.
	EndTime *int `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// Whether to use the highest resolution available.
	HighestResolution *bool `json:"highest_resolution,omitempty" xml:"highest_resolution,omitempty"`
	// An advanced setting. Whether to drop background audio from the final dub. This can improve dub quality where it's known that audio shouldn't have a background track such as for speeches or monologues.
	DropBackgroundAudio *bool `json:"drop_background_audio,omitempty" xml:"drop_background_audio,omitempty"`
	// [BETA] Whether transcripts should have profanities censored with the words '[censored]'
	UseProfanityFilter *bool `json:"use_profanity_filter,omitempty" xml:"use_profanity_filter,omitempty"`
	// Whether to prepare dub for edits in dubbing studio or edits as a dubbing resource.
	DubbingStudio *bool `json:"dubbing_studio,omitempty" xml:"dubbing_studio,omitempty"`
	// Instead of using a voice clone in dubbing, use a similar voice from the ElevenLabs Voice Library. Voices used from the library will contribute towards a workspace's custom voices limit, and if there aren't enough available slots the dub will fail. Using this feature requires the caller to have the 'add_voice_from_voice_library' permission on their workspace to access new voices.
	DisableVoiceCloning *bool `json:"disable_voice_cloning,omitempty" xml:"disable_voice_cloning,omitempty"`
	// The mode in which to run this Dubbing job. Defaults to automatic, use manual if specifically providing a CSV transcript to use. Note that manual mode is experimental and production use is strongly discouraged.
	Mode *CreateDubbingRequestMode `json:"mode,omitempty" xml:"mode,omitempty"`
	// Frames per second to use when parsing a CSV file for dubbing. If not provided, FPS will be inferred from timecodes.
	CsvFps *float64 `json:"csv_fps,omitempty" xml:"csv_fps,omitempty"`
}

func (d *DubRequest) GetFile() (value *string) {
	if d == nil {
		return
	}
	return d.File
}

func (d *DubRequest) GetCsvFile() (value *string) {
	if d == nil {
		return
	}
	return d.CsvFile
}

func (d *DubRequest) GetForegroundAudioFile() (value *string) {
	if d == nil {
		return
	}
	return d.ForegroundAudioFile
}

func (d *DubRequest) GetBackgroundAudioFile() (value *string) {
	if d == nil {
		return
	}
	return d.BackgroundAudioFile
}

func (d *DubRequest) GetName() (value *string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DubRequest) GetSourceUrl() (value *string) {
	if d == nil {
		return
	}
	return d.SourceUrl
}

func (d *DubRequest) GetSourceLang() (value *string) {
	if d == nil {
		return
	}
	return d.SourceLang
}

func (d *DubRequest) GetTargetLang() (value *string) {
	if d == nil {
		return
	}
	return d.TargetLang
}

func (d *DubRequest) GetTargetAccent() (value *string) {
	if d == nil {
		return
	}
	return d.TargetAccent
}

func (d *DubRequest) GetNumSpeakers() (value *int) {
	if d == nil {
		return
	}
	return d.NumSpeakers
}

func (d *DubRequest) GetWatermark() (value *bool) {
	if d == nil {
		return
	}
	return d.Watermark
}

func (d *DubRequest) GetStartTime() (value *int) {
	if d == nil {
		return
	}
	return d.StartTime
}

func (d *DubRequest) GetEndTime() (value *int) {
	if d == nil {
		return
	}
	return d.EndTime
}

func (d *DubRequest) GetHighestResolution() (value *bool) {
	if d == nil {
		return
	}
	return d.HighestResolution
}

func (d *DubRequest) GetDropBackgroundAudio() (value *bool) {
	if d == nil {
		return
	}
	return d.DropBackgroundAudio
}

func (d *DubRequest) GetUseProfanityFilter() (value *bool) {
	if d == nil {
		return
	}
	return d.UseProfanityFilter
}

func (d *DubRequest) GetDubbingStudio() (value *bool) {
	if d == nil {
		return
	}
	return d.DubbingStudio
}

func (d *DubRequest) GetDisableVoiceCloning() (value *bool) {
	if d == nil {
		return
	}
	return d.DisableVoiceCloning
}

func (d *DubRequest) GetMode() (value *CreateDubbingRequestMode) {
	if d == nil {
		return
	}
	return d.Mode
}

func (d *DubRequest) GetCsvFps() (value *float64) {
	if d == nil {
		return
	}
	return d.CsvFps
}

func (d *DubRequest) SetFile(file *string) {
	d.File = file
}

func (d *DubRequest) SetCsvFile(csvFile *string) {
	d.CsvFile = csvFile
}

func (d *DubRequest) SetForegroundAudioFile(foregroundAudioFile *string) {
	d.ForegroundAudioFile = foregroundAudioFile
}

func (d *DubRequest) SetBackgroundAudioFile(backgroundAudioFile *string) {
	d.BackgroundAudioFile = backgroundAudioFile
}

func (d *DubRequest) SetName(name *string) {
	d.Name = name
}

func (d *DubRequest) SetSourceUrl(sourceUrl *string) {
	d.SourceUrl = sourceUrl
}

func (d *DubRequest) SetSourceLang(sourceLang *string) {
	d.SourceLang = sourceLang
}

func (d *DubRequest) SetTargetLang(targetLang *string) {
	d.TargetLang = targetLang
}

func (d *DubRequest) SetTargetAccent(targetAccent *string) {
	d.TargetAccent = targetAccent
}

func (d *DubRequest) SetNumSpeakers(numSpeakers *int) {
	d.NumSpeakers = numSpeakers
}

func (d *DubRequest) SetWatermark(watermark *bool) {
	d.Watermark = watermark
}

func (d *DubRequest) SetStartTime(startTime *int) {
	d.StartTime = startTime
}

func (d *DubRequest) SetEndTime(endTime *int) {
	d.EndTime = endTime
}

func (d *DubRequest) SetHighestResolution(highestResolution *bool) {
	d.HighestResolution = highestResolution
}

func (d *DubRequest) SetDropBackgroundAudio(dropBackgroundAudio *bool) {
	d.DropBackgroundAudio = dropBackgroundAudio
}

func (d *DubRequest) SetUseProfanityFilter(useProfanityFilter *bool) {
	d.UseProfanityFilter = useProfanityFilter
}

func (d *DubRequest) SetDubbingStudio(dubbingStudio *bool) {
	d.DubbingStudio = dubbingStudio
}

func (d *DubRequest) SetDisableVoiceCloning(disableVoiceCloning *bool) {
	d.DisableVoiceCloning = disableVoiceCloning
}

func (d *DubRequest) SetMode(mode *CreateDubbingRequestMode) {
	d.Mode = mode
}

func (d *DubRequest) SetCsvFps(csvFps *float64) {
	d.CsvFps = csvFps
}

func (d *DubRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (d DubRequest) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubRequest to string"
	}
	return string(jsonData)
}

func (d *DubRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, d)
}

func (d *DubRequest) MarshalJSON() ([]byte, error) {
	type embed DubRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
