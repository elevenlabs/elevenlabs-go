package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ChapterContentResponseModel struct {
	Blocks          []*ChapterContentBlockResponseModel `json:"blocks" xml:"blocks" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ChapterContentResponseModel) GetBlocks() (value []*ChapterContentBlockResponseModel) {
	if c == nil {
		return
	}
	return c.Blocks
}

func (c *ChapterContentResponseModel) SetBlocks(blocks []*ChapterContentBlockResponseModel) {
	c.Blocks = blocks
}

func (c *ChapterContentResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterContentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterContentResponseModel to string"
	}
	return string(jsonData)
}

func (c *ChapterContentResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler ChapterContentResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ChapterContentResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ChapterContentResponseModel) MarshalJSON() ([]byte, error) {
	type embed ChapterContentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
