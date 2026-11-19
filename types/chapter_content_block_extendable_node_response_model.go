package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Not used. Make sure you anticipate new types in the future.
type ChapterContentBlockExtendableNodeResponseModel struct {
	type_           string
	extraProperties map[string]interface{}
}

func (c *ChapterContentBlockExtendableNodeResponseModel) Type() string {
	return c.type_
}

func (c *ChapterContentBlockExtendableNodeResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ChapterContentBlockExtendableNodeResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ChapterContentBlockExtendableNodeResponseModel to string"
	}
	return string(jsonData)
}

func (c *ChapterContentBlockExtendableNodeResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ChapterContentBlockExtendableNodeResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ChapterContentBlockExtendableNodeResponseModel(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "_other" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "_other", *literals.Type)
	}
	c.type_ = "_other"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ChapterContentBlockExtendableNodeResponseModel) MarshalJSON() ([]byte, error) {
	type embed ChapterContentBlockExtendableNodeResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "_other",
	}
	return json.Marshal(marshaler)
}
