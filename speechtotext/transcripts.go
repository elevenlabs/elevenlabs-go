package speechtotext

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type GetTranscriptsResponse struct {
	SpeechToTextChunkResponseModel        *types.SpeechToTextChunkResponseModel
	MultichannelSpeechToTextResponseModel *types.MultichannelSpeechToTextResponseModel
	typ                                   string
}

func (g *GetTranscriptsResponse) GetSpeechToTextChunkResponseModel() *types.SpeechToTextChunkResponseModel {
	if g == nil {
		return nil
	}
	if g.SpeechToTextChunkResponseModel == nil {
		return nil
	}
	return g.SpeechToTextChunkResponseModel
}

func (g *GetTranscriptsResponse) GetMultichannelSpeechToTextResponseModel() *types.MultichannelSpeechToTextResponseModel {
	if g == nil {
		return nil
	}
	if g.MultichannelSpeechToTextResponseModel == nil {
		return nil
	}
	return g.MultichannelSpeechToTextResponseModel
}

func (g *GetTranscriptsResponse) UnmarshalJSON(data []byte) error {
	var valueSpeechToTextChunkResponseModel *types.SpeechToTextChunkResponseModel
	if err := json.Unmarshal(data, &valueSpeechToTextChunkResponseModel); err == nil {
		g.typ = "SpeechToTextChunkResponseModel"
		g.SpeechToTextChunkResponseModel = valueSpeechToTextChunkResponseModel
		return nil
	}
	var valueMultichannelSpeechToTextResponseModel *types.MultichannelSpeechToTextResponseModel
	if err := json.Unmarshal(data, &valueMultichannelSpeechToTextResponseModel); err == nil {
		g.typ = "MultichannelSpeechToTextResponseModel"
		g.MultichannelSpeechToTextResponseModel = valueMultichannelSpeechToTextResponseModel
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, g)
}

func (g GetTranscriptsResponse) MarshalJSON() ([]byte, error) {
	if g.typ == "SpeechToTextChunkResponseModel" || g.SpeechToTextChunkResponseModel != nil {
		return json.Marshal(g.SpeechToTextChunkResponseModel)
	}
	if g.typ == "MultichannelSpeechToTextResponseModel" || g.MultichannelSpeechToTextResponseModel != nil {
		return json.Marshal(g.MultichannelSpeechToTextResponseModel)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", g)
}

type GetTranscriptsResponseVisitor interface {
	VisitSpeechToTextChunkResponseModel(*types.SpeechToTextChunkResponseModel) error
	VisitMultichannelSpeechToTextResponseModel(*types.MultichannelSpeechToTextResponseModel) error
}

func (g *GetTranscriptsResponse) Accept(visitor GetTranscriptsResponseVisitor) error {
	if g.typ == "SpeechToTextChunkResponseModel" || g.SpeechToTextChunkResponseModel != nil {
		return visitor.VisitSpeechToTextChunkResponseModel(g.SpeechToTextChunkResponseModel)
	}
	if g.typ == "MultichannelSpeechToTextResponseModel" || g.MultichannelSpeechToTextResponseModel != nil {
		return visitor.VisitMultichannelSpeechToTextResponseModel(g.MultichannelSpeechToTextResponseModel)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", g)
}
