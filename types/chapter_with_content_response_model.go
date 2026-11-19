package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ChapterWithContentResponseModel struct {
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
	State ChapterWithContentResponseModelState `json:"state" xml:"state" required:"true"`
	// Whether the chapter has a video.
	HasVideo *bool `json:"has_video,omitempty" xml:"has_video,omitempty"`
	// Whether the chapter has any visual content (video, image, or text clips).
	HasVisualContent *bool `json:"has_visual_content,omitempty" xml:"has_visual_content,omitempty"`
	// List of voice ids used by the chapter
	VoiceIds []string `json:"voice_ids,omitempty" xml:"voice_ids,omitempty"`
	// The statistics of the chapter.
	Statistics *ChapterStatisticsResponse `json:"statistics,omitempty" xml:"statistics,omitempty"`
	// The last conversion error of the chapter.
	LastConversionError *string                      `json:"last_conversion_error,omitempty" xml:"last_conversion_error,omitempty"`
	Content             *ChapterContentResponseModel `json:"content" xml:"content" required:"true"`
	extraProperties     map[string]interface{}
}

func (c *ChapterWithContentResponseModel) GetChapterId() (value string) {
	if c == nil {
		return
	}
	return c.ChapterId
}

func (c *ChapterWithContentResponseModel) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *ChapterWithContentResponseModel) GetLastConversionDateUnix() (value *int) {
	if c == nil {
		return
	}
	return c.LastConversionDateUnix
}

func (c *ChapterWithContentResponseModel) GetConversionProgress() (value *float64) {
	if c == nil {
		return
	}
	return c.ConversionProgress
}

func (c *ChapterWithContentResponseModel) GetCanBeDownloaded() (value bool) {
	if c == nil {
		return
	}
	return c.CanBeDownloaded
}

func (c *ChapterWithContentResponseModel) GetState() (value ChapterWithContentResponseModelState) {
	if c == nil {
		return
	}
	return c.State
}

func (c *ChapterWithContentResponseModel) GetHasVideo() (value *bool) {
	if c == nil {
		return
	}
	return c.HasVideo
}

func (c *ChapterWithContentResponseModel) GetHasVisualContent() (value *bool) {
	if c == nil {
		return
	}
	return c.HasVisualContent
}

func (c *ChapterWithContentResponseModel) GetVoiceIds() (value []string) {
	if c == nil {
		return
	}
	return c.VoiceIds
}

func (c *ChapterWithContentResponseModel) GetStatistics() (value *ChapterStatisticsResponse) {
	if c == nil {
		return
	}
	return c.Statistics
}

func (c *ChapterWithContentResponseModel) GetLastConversionError() (value *string) {
	if c == nil {
		return
	}
	return c.LastConversionError
}

func (c *ChapterWithContentResponseModel) GetContent() (value *ChapterContentResponseModel) {
	if c == nil {
		return
	}
	return c.Content
}

func (c *ChapterWithContentResponseModel) SetChapterId(chapterId string) {
	c.ChapterId = chapterId
}

func (c *ChapterWithContentResponseModel) SetName(name string) {
	c.Name = name
}

func (c *ChapterWithContentResponseModel) SetLastConversionDateUnix(lastConversionDateUnix *int) {
	c.LastConversionDateUnix = lastConversionDateUnix
}

func (c *ChapterWithContentResponseModel) SetConversionProgress(conversionProgress *float64) {
	c.ConversionProgress = conversionProgress
}

func (c *ChapterWithContentResponseModel) SetCanBeDownloaded(canBeDownloaded bool) {
	c.CanBeDownloaded = canBeDownloaded
}

func (c *ChapterWithContentResponseModel) SetState(state ChapterWithContentResponseModelState) {
	c.State = state
}

func (c *ChapterWithContentResponseModel) SetHasVideo(hasVideo *bool) {
	c.HasVideo = hasVideo
}

func (c *ChapterWithContentResponseModel) SetHasVisualContent(hasVisualContent *bool) {
	c.HasVisualContent = hasVisualContent
}

func (c *ChapterWithContentResponseModel) SetVoiceIds(voiceIds []string) {
	c.VoiceIds = voiceIds
}

func (c *ChapterWithContentResponseModel) SetStatistics(statistics *ChapterStatisticsResponse) {
	c.Statistics = statistics
}

func (c *ChapterWithContentResponseModel) SetLastConversionError(lastConversionError *string) {
	c.LastConversionError = lastConversionError
}

func (c *ChapterWithContentResponseModel) SetContent(content *ChapterContentResponseModel) {
	c.Content = content
}

func (c *ChapterWithContentResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterWithContentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterWithContentResponseModel to string"
	}
	return string(jsonData)
}

func (c *ChapterWithContentResponseModel) UnmarshalJSON(data []byte) error {
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

func (c *ChapterWithContentResponseModel) MarshalJSON() ([]byte, error) {
	type embed ChapterWithContentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
