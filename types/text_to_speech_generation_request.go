package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TextToSpeechGenerationRequest struct {
	ModelId              string
	ElevenFlashV25       *ElevenFlashV25Request
	ElevenMultilingualV2 *ElevenMultilingualV2Request
	ElevenV3             *ElevenV3Request
}

func (t *TextToSpeechGenerationRequest) GetElevenFlashV25() *ElevenFlashV25Request {
	if t == nil {
		return nil
	}
	if t.ElevenFlashV25 == nil {
		return nil
	}
	return t.ElevenFlashV25
}

func (t *TextToSpeechGenerationRequest) GetElevenMultilingualV2() *ElevenMultilingualV2Request {
	if t == nil {
		return nil
	}
	if t.ElevenMultilingualV2 == nil {
		return nil
	}
	return t.ElevenMultilingualV2
}

func (t *TextToSpeechGenerationRequest) GetElevenV3() *ElevenV3Request {
	if t == nil {
		return nil
	}
	if t.ElevenV3 == nil {
		return nil
	}
	return t.ElevenV3
}

func (t *TextToSpeechGenerationRequest) GetModelId() string {
	if t == nil {
		return ""
	}
	return t.ModelId
}

func (t *TextToSpeechGenerationRequest) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		ModelId string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	t.ModelId = unmarshaler.ModelId
	if unmarshaler.ModelId == "" {
		return fmt.Errorf("%T did not include discriminant model_id", t)
	}
	switch unmarshaler.ModelId {
	case "eleven_flash_v2_5":
		value := new(ElevenFlashV25Request)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.ElevenFlashV25 = value
	case "eleven_multilingual_v2":
		value := new(ElevenMultilingualV2Request)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.ElevenMultilingualV2 = value
	case "eleven_v3":
		value := new(ElevenV3Request)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.ElevenV3 = value
	}
	return nil
}

func (t TextToSpeechGenerationRequest) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.ElevenFlashV25 != nil {
		return internal.MarshalJSONWithDiscriminant(t.ElevenFlashV25, "model_id", "eleven_flash_v2_5")
	}
	if t.ElevenMultilingualV2 != nil {
		return internal.MarshalJSONWithDiscriminant(t.ElevenMultilingualV2, "model_id", "eleven_multilingual_v2")
	}
	if t.ElevenV3 != nil {
		return internal.MarshalJSONWithDiscriminant(t.ElevenV3, "model_id", "eleven_v3")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TextToSpeechGenerationRequestVisitor interface {
	VisitElevenFlashV25(*ElevenFlashV25Request) error
	VisitElevenMultilingualV2(*ElevenMultilingualV2Request) error
	VisitElevenV3(*ElevenV3Request) error
}

func (t *TextToSpeechGenerationRequest) Accept(visitor TextToSpeechGenerationRequestVisitor) error {
	if t.ElevenFlashV25 != nil {
		return visitor.VisitElevenFlashV25(t.ElevenFlashV25)
	}
	if t.ElevenMultilingualV2 != nil {
		return visitor.VisitElevenMultilingualV2(t.ElevenMultilingualV2)
	}
	if t.ElevenV3 != nil {
		return visitor.VisitElevenV3(t.ElevenV3)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *TextToSpeechGenerationRequest) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.ElevenFlashV25 != nil {
		fields = append(fields, "eleven_flash_v2_5")
	}
	if t.ElevenMultilingualV2 != nil {
		fields = append(fields, "eleven_multilingual_v2")
	}
	if t.ElevenV3 != nil {
		fields = append(fields, "eleven_v3")
	}
	if len(fields) == 0 {
		if t.ModelId != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", t, t.ModelId)
		}
		return fmt.Errorf("type %T is empty", t)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", t, fields)
	}
	if t.ModelId != "" {
		field := fields[0]
		if t.ModelId != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				t,
				t.ModelId,
				t,
			)
		}
	}
	return nil
}
