package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ChapterResponse struct {
	// The ID of the chapter.
	ChapterId string `json:"chapter_id" xml:"chapter_id" required:"true"`
	// The name of the chapter.
	Name string `json:"name" xml:"name" required:"true"`
	// The last conversion date of the chapter.
	LastConversionDateUnix *int `json:"last_conversion_date_unix,omitempty" xml:"last_conversion_date_unix,omitempty"`
	// The conversion progress of the chapter.
	ConversionProgress *float64 `json:"conversion_progress,omitempty" xml:"conversion_progress,omitempty"`
	// Whether the chapter can be downloaded.
	CanBeDownloaded bool `json:"can_be_downloaded" xml:"can_be_downloaded" required:"true"`
	// The state of the chapter.
	State ChapterState `json:"state" xml:"state" required:"true"`
	// Whether the chapter has a video.
	HasVideo *bool `json:"has_video,omitempty" xml:"has_video,omitempty"`
	// Whether the chapter has any visual content (video, image, or text clips).
	HasVisualContent *bool `json:"has_visual_content,omitempty" xml:"has_visual_content,omitempty"`
	// List of voice ids used by the chapter
	VoiceIds []string `json:"voice_ids,omitempty" xml:"voice_ids,omitempty"`
	// The statistics of the chapter.
	Statistics *ChapterStatisticsResponse `json:"statistics,omitempty" xml:"statistics,omitempty"`
	// The last conversion error of the chapter.
	LastConversionError *string `json:"last_conversion_error,omitempty" xml:"last_conversion_error,omitempty"`
	extraProperties     map[string]interface{}
}

func (c *ChapterResponse) GetChapterId() (value string) {
	if c == nil {
		return
	}
	return c.ChapterId
}

func (c *ChapterResponse) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *ChapterResponse) GetLastConversionDateUnix() (value *int) {
	if c == nil {
		return
	}
	return c.LastConversionDateUnix
}

func (c *ChapterResponse) GetConversionProgress() (value *float64) {
	if c == nil {
		return
	}
	return c.ConversionProgress
}

func (c *ChapterResponse) GetCanBeDownloaded() (value bool) {
	if c == nil {
		return
	}
	return c.CanBeDownloaded
}

func (c *ChapterResponse) GetState() (value ChapterState) {
	if c == nil {
		return
	}
	return c.State
}

func (c *ChapterResponse) GetHasVideo() (value *bool) {
	if c == nil {
		return
	}
	return c.HasVideo
}

func (c *ChapterResponse) GetHasVisualContent() (value *bool) {
	if c == nil {
		return
	}
	return c.HasVisualContent
}

func (c *ChapterResponse) GetVoiceIds() (value []string) {
	if c == nil {
		return
	}
	return c.VoiceIds
}

func (c *ChapterResponse) GetStatistics() (value *ChapterStatisticsResponse) {
	if c == nil {
		return
	}
	return c.Statistics
}

func (c *ChapterResponse) GetLastConversionError() (value *string) {
	if c == nil {
		return
	}
	return c.LastConversionError
}

func (c *ChapterResponse) SetChapterId(chapterId string) {
	c.ChapterId = chapterId
}

func (c *ChapterResponse) SetName(name string) {
	c.Name = name
}

func (c *ChapterResponse) SetLastConversionDateUnix(lastConversionDateUnix *int) {
	c.LastConversionDateUnix = lastConversionDateUnix
}

func (c *ChapterResponse) SetConversionProgress(conversionProgress *float64) {
	c.ConversionProgress = conversionProgress
}

func (c *ChapterResponse) SetCanBeDownloaded(canBeDownloaded bool) {
	c.CanBeDownloaded = canBeDownloaded
}

func (c *ChapterResponse) SetState(state ChapterState) {
	c.State = state
}

func (c *ChapterResponse) SetHasVideo(hasVideo *bool) {
	c.HasVideo = hasVideo
}

func (c *ChapterResponse) SetHasVisualContent(hasVisualContent *bool) {
	c.HasVisualContent = hasVisualContent
}

func (c *ChapterResponse) SetVoiceIds(voiceIds []string) {
	c.VoiceIds = voiceIds
}

func (c *ChapterResponse) SetStatistics(statistics *ChapterStatisticsResponse) {
	c.Statistics = statistics
}

func (c *ChapterResponse) SetLastConversionError(lastConversionError *string) {
	c.LastConversionError = lastConversionError
}

func (c *ChapterResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterResponse to string"
	}
	return string(jsonData)
}

func (c *ChapterResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ChapterResponse) MarshalJSON() ([]byte, error) {
	type embed ChapterResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
