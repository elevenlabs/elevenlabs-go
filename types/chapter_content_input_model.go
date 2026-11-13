package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ChapterContentInputModel struct {
	Blocks          []*ChapterContentBlockInputModel `json:"blocks" xml:"blocks" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ChapterContentInputModel) GetBlocks() (value []*ChapterContentBlockInputModel) {
	if c == nil {
		return
	}
	return c.Blocks
}

func (c *ChapterContentInputModel) SetBlocks(blocks []*ChapterContentBlockInputModel) {
	c.Blocks = blocks
}

func (c *ChapterContentInputModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterContentInputModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterContentInputModel to string"
	}
	return string(jsonData)
}

func (c *ChapterContentInputModel) UnmarshalJSON(data []byte) error {
	type unmarshaler ChapterContentInputModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ChapterContentInputModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ChapterContentInputModel) MarshalJSON() ([]byte, error) {
	type embed ChapterContentInputModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
