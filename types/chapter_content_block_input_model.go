package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ChapterContentBlockInputModel struct {
	SubType         *ChapterContentBlockInputModelSubType       `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
	Nodes           []*ChapterContentParagraphTtsNodeInputModel `json:"nodes" xml:"nodes" required:"true"`
	BlockId         *string                                     `json:"block_id,omitempty" xml:"block_id,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ChapterContentBlockInputModel) GetSubType() (value *ChapterContentBlockInputModelSubType) {
	if c == nil {
		return
	}
	return c.SubType
}

func (c *ChapterContentBlockInputModel) GetNodes() (value []*ChapterContentParagraphTtsNodeInputModel) {
	if c == nil {
		return
	}
	return c.Nodes
}

func (c *ChapterContentBlockInputModel) GetBlockId() (value *string) {
	if c == nil {
		return
	}
	return c.BlockId
}

func (c *ChapterContentBlockInputModel) SetSubType(subType *ChapterContentBlockInputModelSubType) {
	c.SubType = subType
}

func (c *ChapterContentBlockInputModel) SetNodes(nodes []*ChapterContentParagraphTtsNodeInputModel) {
	c.Nodes = nodes
}

func (c *ChapterContentBlockInputModel) SetBlockId(blockId *string) {
	c.BlockId = blockId
}

func (c *ChapterContentBlockInputModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterContentBlockInputModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterContentBlockInputModel to string"
	}
	return string(jsonData)
}

func (c *ChapterContentBlockInputModel) UnmarshalJSON(data []byte) error {
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

func (c *ChapterContentBlockInputModel) MarshalJSON() ([]byte, error) {
	type embed ChapterContentBlockInputModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
