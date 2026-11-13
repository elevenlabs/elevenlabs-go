package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ChapterContentBlockResponseModelNodesItem struct {
	Type    string
	Other   *ChapterContentBlockExtendableNodeResponseModel
	TtsNode *ChapterContentBlockTtsNodeResponseModel
}

func (c *ChapterContentBlockResponseModelNodesItem) GetOther() *ChapterContentBlockExtendableNodeResponseModel {
	if c == nil {
		return nil
	}
	if c.Other == nil {
		return nil
	}
	return c.Other
}

func (c *ChapterContentBlockResponseModelNodesItem) GetTtsNode() *ChapterContentBlockTtsNodeResponseModel {
	if c == nil {
		return nil
	}
	if c.TtsNode == nil {
		return nil
	}
	return c.TtsNode
}

func (c *ChapterContentBlockResponseModelNodesItem) GetType() string {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *ChapterContentBlockResponseModelNodesItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", c)
	}
	switch unmarshaler.Type {
	case "_other":
		value := new(ChapterContentBlockExtendableNodeResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Other = value
	case "tts_node":
		value := new(ChapterContentBlockTtsNodeResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.TtsNode = value
	}
	return nil
}

func (c ChapterContentBlockResponseModelNodesItem) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.Other != nil {
		return internal.MarshalJSONWithDiscriminant(c.Other, "type", "_other")
	}
	if c.TtsNode != nil {
		return internal.MarshalJSONWithDiscriminant(c.TtsNode, "type", "tts_node")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type ChapterContentBlockResponseModelNodesItemVisitor interface {
	VisitOther(*ChapterContentBlockExtendableNodeResponseModel) error
	VisitTtsNode(*ChapterContentBlockTtsNodeResponseModel) error
}

func (c *ChapterContentBlockResponseModelNodesItem) Accept(visitor ChapterContentBlockResponseModelNodesItemVisitor) error {
	if c.Other != nil {
		return visitor.VisitOther(c.Other)
	}
	if c.TtsNode != nil {
		return visitor.VisitTtsNode(c.TtsNode)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *ChapterContentBlockResponseModelNodesItem) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.Other != nil {
		fields = append(fields, "_other")
	}
	if c.TtsNode != nil {
		fields = append(fields, "tts_node")
	}
	if len(fields) == 0 {
		if c.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.Type)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.Type != "" {
		field := fields[0]
		if c.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.Type,
				c,
			)
		}
	}
	return nil
}
