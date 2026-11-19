package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ChapterContentBlockResponseModel struct {
	BlockId         string                                       `json:"block_id" xml:"block_id" required:"true"`
	Nodes           []*ChapterContentBlockResponseModelNodesItem `json:"nodes" xml:"nodes" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ChapterContentBlockResponseModel) GetBlockId() (value string) {
	if c == nil {
		return
	}
	return c.BlockId
}

func (c *ChapterContentBlockResponseModel) GetNodes() (value []*ChapterContentBlockResponseModelNodesItem) {
	if c == nil {
		return
	}
	return c.Nodes
}

func (c *ChapterContentBlockResponseModel) SetBlockId(blockId string) {
	c.BlockId = blockId
}

func (c *ChapterContentBlockResponseModel) SetNodes(nodes []*ChapterContentBlockResponseModelNodesItem) {
	c.Nodes = nodes
}

func (c *ChapterContentBlockResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterContentBlockResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterContentBlockResponseModel to string"
	}
	return string(jsonData)
}

func (c *ChapterContentBlockResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ChapterContentBlockResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ChapterContentBlockResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ChapterContentBlockResponseModel) MarshalJSON() ([]byte, error) {
	type embed ChapterContentBlockResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
