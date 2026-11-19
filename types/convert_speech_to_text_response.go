package types

import (
	"encoding/json"
	"fmt"
)

type ConvertSpeechToTextResponse struct {
	SpeechToTextChunkResponseModel        *SpeechToTextChunkResponseModel
	MultichannelSpeechToTextResponseModel *MultichannelSpeechToTextResponseModel
	SpeechToTextWebhookResponseModel      *SpeechToTextWebhookResponseModel
	typ                                   string
}

func (c *ConvertSpeechToTextResponse) GetSpeechToTextChunkResponseModel() *SpeechToTextChunkResponseModel {
	if c == nil {
		return nil
	}
	if c.SpeechToTextChunkResponseModel == nil {
		return nil
	}
	return c.SpeechToTextChunkResponseModel
}

func (c *ConvertSpeechToTextResponse) GetMultichannelSpeechToTextResponseModel() *MultichannelSpeechToTextResponseModel {
	if c == nil {
		return nil
	}
	if c.MultichannelSpeechToTextResponseModel == nil {
		return nil
	}
	return c.MultichannelSpeechToTextResponseModel
}

func (c *ConvertSpeechToTextResponse) GetSpeechToTextWebhookResponseModel() *SpeechToTextWebhookResponseModel {
	if c == nil {
		return nil
	}
	if c.SpeechToTextWebhookResponseModel == nil {
		return nil
	}
	return c.SpeechToTextWebhookResponseModel
}

func (c *ConvertSpeechToTextResponse) UnmarshalJSON(data []byte) error {
	var valueSpeechToTextChunkResponseModel *SpeechToTextChunkResponseModel
	if err := json.Unmarshal(data, &valueSpeechToTextChunkResponseModel); err == nil {
		c.typ = "SpeechToTextChunkResponseModel"
		c.SpeechToTextChunkResponseModel = valueSpeechToTextChunkResponseModel
		return nil
	}
	var valueMultichannelSpeechToTextResponseModel *MultichannelSpeechToTextResponseModel
	if err := json.Unmarshal(data, &valueMultichannelSpeechToTextResponseModel); err == nil {
		c.typ = "MultichannelSpeechToTextResponseModel"
		c.MultichannelSpeechToTextResponseModel = valueMultichannelSpeechToTextResponseModel
		return nil
	}
	var valueSpeechToTextWebhookResponseModel *SpeechToTextWebhookResponseModel
	if err := json.Unmarshal(data, &valueSpeechToTextWebhookResponseModel); err == nil {
		c.typ = "SpeechToTextWebhookResponseModel"
		c.SpeechToTextWebhookResponseModel = valueSpeechToTextWebhookResponseModel
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, c)
}

func (c ConvertSpeechToTextResponse) MarshalJSON() ([]byte, error) {
	if c.typ == "SpeechToTextChunkResponseModel" || c.SpeechToTextChunkResponseModel != nil {
		return json.Marshal(c.SpeechToTextChunkResponseModel)
	}
	if c.typ == "MultichannelSpeechToTextResponseModel" || c.MultichannelSpeechToTextResponseModel != nil {
		return json.Marshal(c.MultichannelSpeechToTextResponseModel)
	}
	if c.typ == "SpeechToTextWebhookResponseModel" || c.SpeechToTextWebhookResponseModel != nil {
		return json.Marshal(c.SpeechToTextWebhookResponseModel)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", c)
}

type ConvertSpeechToTextResponseVisitor interface {
	VisitSpeechToTextChunkResponseModel(*SpeechToTextChunkResponseModel) error
	VisitMultichannelSpeechToTextResponseModel(*MultichannelSpeechToTextResponseModel) error
	VisitSpeechToTextWebhookResponseModel(*SpeechToTextWebhookResponseModel) error
}

func (c *ConvertSpeechToTextResponse) Accept(visitor ConvertSpeechToTextResponseVisitor) error {
	if c.typ == "SpeechToTextChunkResponseModel" || c.SpeechToTextChunkResponseModel != nil {
		return visitor.VisitSpeechToTextChunkResponseModel(c.SpeechToTextChunkResponseModel)
	}
	if c.typ == "MultichannelSpeechToTextResponseModel" || c.MultichannelSpeechToTextResponseModel != nil {
		return visitor.VisitMultichannelSpeechToTextResponseModel(c.MultichannelSpeechToTextResponseModel)
	}
	if c.typ == "SpeechToTextWebhookResponseModel" || c.SpeechToTextWebhookResponseModel != nil {
		return visitor.VisitSpeechToTextWebhookResponseModel(c.SpeechToTextWebhookResponseModel)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", c)
}
