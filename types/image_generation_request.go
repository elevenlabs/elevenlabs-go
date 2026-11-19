package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ImageGenerationRequest struct {
	ModelId                string
	BytedanceSeedream5Lite *BytedanceSeedream5LiteRequest
	BytedanceSeedream5Pro  *BytedanceSeedream5ProRequest
	Gemini25FlashImage     *Gemini25FlashImageRequest
	Gemini3ProImage        *Gemini3ProImageRequest
	Gemini31FlashImage     *Gemini31FlashImageRequest
	Gemini31FlashLiteImage *Gemini31FlashLiteImageRequest
	GptImage1              *GptImage1Request
	GptImage15             *GptImage15Request
	GptImage2              *GptImage2Request
	GptImage25Flare        *GptImage25FlareRequest
	GptImage25Sunburst     *GptImage25SunburstRequest
}

func (i *ImageGenerationRequest) GetBytedanceSeedream5Lite() *BytedanceSeedream5LiteRequest {
	if i == nil {
		return nil
	}
	if i.BytedanceSeedream5Lite == nil {
		return nil
	}
	return i.BytedanceSeedream5Lite
}

func (i *ImageGenerationRequest) GetBytedanceSeedream5Pro() *BytedanceSeedream5ProRequest {
	if i == nil {
		return nil
	}
	if i.BytedanceSeedream5Pro == nil {
		return nil
	}
	return i.BytedanceSeedream5Pro
}

func (i *ImageGenerationRequest) GetGemini25FlashImage() *Gemini25FlashImageRequest {
	if i == nil {
		return nil
	}
	if i.Gemini25FlashImage == nil {
		return nil
	}
	return i.Gemini25FlashImage
}

func (i *ImageGenerationRequest) GetGemini3ProImage() *Gemini3ProImageRequest {
	if i == nil {
		return nil
	}
	if i.Gemini3ProImage == nil {
		return nil
	}
	return i.Gemini3ProImage
}

func (i *ImageGenerationRequest) GetGemini31FlashImage() *Gemini31FlashImageRequest {
	if i == nil {
		return nil
	}
	if i.Gemini31FlashImage == nil {
		return nil
	}
	return i.Gemini31FlashImage
}

func (i *ImageGenerationRequest) GetGemini31FlashLiteImage() *Gemini31FlashLiteImageRequest {
	if i == nil {
		return nil
	}
	if i.Gemini31FlashLiteImage == nil {
		return nil
	}
	return i.Gemini31FlashLiteImage
}

func (i *ImageGenerationRequest) GetGptImage1() *GptImage1Request {
	if i == nil {
		return nil
	}
	if i.GptImage1 == nil {
		return nil
	}
	return i.GptImage1
}

func (i *ImageGenerationRequest) GetGptImage15() *GptImage15Request {
	if i == nil {
		return nil
	}
	if i.GptImage15 == nil {
		return nil
	}
	return i.GptImage15
}

func (i *ImageGenerationRequest) GetGptImage2() *GptImage2Request {
	if i == nil {
		return nil
	}
	if i.GptImage2 == nil {
		return nil
	}
	return i.GptImage2
}

func (i *ImageGenerationRequest) GetGptImage25Flare() *GptImage25FlareRequest {
	if i == nil {
		return nil
	}
	if i.GptImage25Flare == nil {
		return nil
	}
	return i.GptImage25Flare
}

func (i *ImageGenerationRequest) GetGptImage25Sunburst() *GptImage25SunburstRequest {
	if i == nil {
		return nil
	}
	if i.GptImage25Sunburst == nil {
		return nil
	}
	return i.GptImage25Sunburst
}

func (i *ImageGenerationRequest) GetModelId() string {
	if i == nil {
		return ""
	}
	return i.ModelId
}

func (i *ImageGenerationRequest) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		ModelId string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	i.ModelId = unmarshaler.ModelId
	if unmarshaler.ModelId == "" {
		return fmt.Errorf("%T did not include discriminant model_id", i)
	}
	switch unmarshaler.ModelId {
	case "bytedance-seedream-5-lite":
		value := new(BytedanceSeedream5LiteRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.BytedanceSeedream5Lite = value
	case "bytedance-seedream-5-pro":
		value := new(BytedanceSeedream5ProRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.BytedanceSeedream5Pro = value
	case "gemini-2.5-flash-image":
		value := new(Gemini25FlashImageRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.Gemini25FlashImage = value
	case "gemini-3-pro-image":
		value := new(Gemini3ProImageRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.Gemini3ProImage = value
	case "gemini-3.1-flash-image":
		value := new(Gemini31FlashImageRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.Gemini31FlashImage = value
	case "gemini-3.1-flash-lite-image":
		value := new(Gemini31FlashLiteImageRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.Gemini31FlashLiteImage = value
	case "gpt-image-1":
		value := new(GptImage1Request)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.GptImage1 = value
	case "gpt-image-1.5":
		value := new(GptImage15Request)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.GptImage15 = value
	case "gpt-image-2":
		value := new(GptImage2Request)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.GptImage2 = value
	case "gpt-image-2.5-flare":
		value := new(GptImage25FlareRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.GptImage25Flare = value
	case "gpt-image-2.5-sunburst":
		value := new(GptImage25SunburstRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.GptImage25Sunburst = value
	}
	return nil
}

func (i ImageGenerationRequest) MarshalJSON() ([]byte, error) {
	if err := i.validate(); err != nil {
		return nil, err
	}
	if i.BytedanceSeedream5Lite != nil {
		return internal.MarshalJSONWithDiscriminant(i.BytedanceSeedream5Lite, "model_id", "bytedance-seedream-5-lite")
	}
	if i.BytedanceSeedream5Pro != nil {
		return internal.MarshalJSONWithDiscriminant(i.BytedanceSeedream5Pro, "model_id", "bytedance-seedream-5-pro")
	}
	if i.Gemini25FlashImage != nil {
		return internal.MarshalJSONWithDiscriminant(i.Gemini25FlashImage, "model_id", "gemini-2.5-flash-image")
	}
	if i.Gemini3ProImage != nil {
		return internal.MarshalJSONWithDiscriminant(i.Gemini3ProImage, "model_id", "gemini-3-pro-image")
	}
	if i.Gemini31FlashImage != nil {
		return internal.MarshalJSONWithDiscriminant(i.Gemini31FlashImage, "model_id", "gemini-3.1-flash-image")
	}
	if i.Gemini31FlashLiteImage != nil {
		return internal.MarshalJSONWithDiscriminant(i.Gemini31FlashLiteImage, "model_id", "gemini-3.1-flash-lite-image")
	}
	if i.GptImage1 != nil {
		return internal.MarshalJSONWithDiscriminant(i.GptImage1, "model_id", "gpt-image-1")
	}
	if i.GptImage15 != nil {
		return internal.MarshalJSONWithDiscriminant(i.GptImage15, "model_id", "gpt-image-1.5")
	}
	if i.GptImage2 != nil {
		return internal.MarshalJSONWithDiscriminant(i.GptImage2, "model_id", "gpt-image-2")
	}
	if i.GptImage25Flare != nil {
		return internal.MarshalJSONWithDiscriminant(i.GptImage25Flare, "model_id", "gpt-image-2.5-flare")
	}
	if i.GptImage25Sunburst != nil {
		return internal.MarshalJSONWithDiscriminant(i.GptImage25Sunburst, "model_id", "gpt-image-2.5-sunburst")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", i)
}

type ImageGenerationRequestVisitor interface {
	VisitBytedanceSeedream5Lite(*BytedanceSeedream5LiteRequest) error
	VisitBytedanceSeedream5Pro(*BytedanceSeedream5ProRequest) error
	VisitGemini25FlashImage(*Gemini25FlashImageRequest) error
	VisitGemini3ProImage(*Gemini3ProImageRequest) error
	VisitGemini31FlashImage(*Gemini31FlashImageRequest) error
	VisitGemini31FlashLiteImage(*Gemini31FlashLiteImageRequest) error
	VisitGptImage1(*GptImage1Request) error
	VisitGptImage15(*GptImage15Request) error
	VisitGptImage2(*GptImage2Request) error
	VisitGptImage25Flare(*GptImage25FlareRequest) error
	VisitGptImage25Sunburst(*GptImage25SunburstRequest) error
}

func (i *ImageGenerationRequest) Accept(visitor ImageGenerationRequestVisitor) error {
	if i.BytedanceSeedream5Lite != nil {
		return visitor.VisitBytedanceSeedream5Lite(i.BytedanceSeedream5Lite)
	}
	if i.BytedanceSeedream5Pro != nil {
		return visitor.VisitBytedanceSeedream5Pro(i.BytedanceSeedream5Pro)
	}
	if i.Gemini25FlashImage != nil {
		return visitor.VisitGemini25FlashImage(i.Gemini25FlashImage)
	}
	if i.Gemini3ProImage != nil {
		return visitor.VisitGemini3ProImage(i.Gemini3ProImage)
	}
	if i.Gemini31FlashImage != nil {
		return visitor.VisitGemini31FlashImage(i.Gemini31FlashImage)
	}
	if i.Gemini31FlashLiteImage != nil {
		return visitor.VisitGemini31FlashLiteImage(i.Gemini31FlashLiteImage)
	}
	if i.GptImage1 != nil {
		return visitor.VisitGptImage1(i.GptImage1)
	}
	if i.GptImage15 != nil {
		return visitor.VisitGptImage15(i.GptImage15)
	}
	if i.GptImage2 != nil {
		return visitor.VisitGptImage2(i.GptImage2)
	}
	if i.GptImage25Flare != nil {
		return visitor.VisitGptImage25Flare(i.GptImage25Flare)
	}
	if i.GptImage25Sunburst != nil {
		return visitor.VisitGptImage25Sunburst(i.GptImage25Sunburst)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", i)
}

func (i *ImageGenerationRequest) validate() error {
	if i == nil {
		return fmt.Errorf("type %T is nil", i)
	}
	var fields []string
	if i.BytedanceSeedream5Lite != nil {
		fields = append(fields, "bytedance-seedream-5-lite")
	}
	if i.BytedanceSeedream5Pro != nil {
		fields = append(fields, "bytedance-seedream-5-pro")
	}
	if i.Gemini25FlashImage != nil {
		fields = append(fields, "gemini-2.5-flash-image")
	}
	if i.Gemini3ProImage != nil {
		fields = append(fields, "gemini-3-pro-image")
	}
	if i.Gemini31FlashImage != nil {
		fields = append(fields, "gemini-3.1-flash-image")
	}
	if i.Gemini31FlashLiteImage != nil {
		fields = append(fields, "gemini-3.1-flash-lite-image")
	}
	if i.GptImage1 != nil {
		fields = append(fields, "gpt-image-1")
	}
	if i.GptImage15 != nil {
		fields = append(fields, "gpt-image-1.5")
	}
	if i.GptImage2 != nil {
		fields = append(fields, "gpt-image-2")
	}
	if i.GptImage25Flare != nil {
		fields = append(fields, "gpt-image-2.5-flare")
	}
	if i.GptImage25Sunburst != nil {
		fields = append(fields, "gpt-image-2.5-sunburst")
	}
	if len(fields) == 0 {
		if i.ModelId != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", i, i.ModelId)
		}
		return fmt.Errorf("type %T is empty", i)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", i, fields)
	}
	if i.ModelId != "" {
		field := fields[0]
		if i.ModelId != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				i,
				i.ModelId,
				i,
			)
		}
	}
	return nil
}
