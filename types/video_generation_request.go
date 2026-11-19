package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type VideoGenerationRequest struct {
	ModelId                 string
	BytedanceSeedanceV2     *BytedanceSeedance2Request
	BytedanceSeedanceV2Fast *BytedanceSeedance2FastRequest
	BytedanceSeedanceV2Mini *BytedanceSeedance2MiniRequest
	BytedanceSeedanceV25    *BytedanceSeedance25Request
	CreatifyAurora          *CreatifyAuroraRequest
	Veo31FastGenerate001    *Veo31FastRequest
	Veo31Generate001        *Veo31Request
}

func (v *VideoGenerationRequest) GetBytedanceSeedanceV2() *BytedanceSeedance2Request {
	if v == nil {
		return nil
	}
	if v.BytedanceSeedanceV2 == nil {
		return nil
	}
	return v.BytedanceSeedanceV2
}

func (v *VideoGenerationRequest) GetBytedanceSeedanceV2Fast() *BytedanceSeedance2FastRequest {
	if v == nil {
		return nil
	}
	if v.BytedanceSeedanceV2Fast == nil {
		return nil
	}
	return v.BytedanceSeedanceV2Fast
}

func (v *VideoGenerationRequest) GetBytedanceSeedanceV2Mini() *BytedanceSeedance2MiniRequest {
	if v == nil {
		return nil
	}
	if v.BytedanceSeedanceV2Mini == nil {
		return nil
	}
	return v.BytedanceSeedanceV2Mini
}

func (v *VideoGenerationRequest) GetBytedanceSeedanceV25() *BytedanceSeedance25Request {
	if v == nil {
		return nil
	}
	if v.BytedanceSeedanceV25 == nil {
		return nil
	}
	return v.BytedanceSeedanceV25
}

func (v *VideoGenerationRequest) GetCreatifyAurora() *CreatifyAuroraRequest {
	if v == nil {
		return nil
	}
	if v.CreatifyAurora == nil {
		return nil
	}
	return v.CreatifyAurora
}

func (v *VideoGenerationRequest) GetVeo31FastGenerate001() *Veo31FastRequest {
	if v == nil {
		return nil
	}
	if v.Veo31FastGenerate001 == nil {
		return nil
	}
	return v.Veo31FastGenerate001
}

func (v *VideoGenerationRequest) GetVeo31Generate001() *Veo31Request {
	if v == nil {
		return nil
	}
	if v.Veo31Generate001 == nil {
		return nil
	}
	return v.Veo31Generate001
}

func (v *VideoGenerationRequest) GetModelId() string {
	if v == nil {
		return ""
	}
	return v.ModelId
}

func (v *VideoGenerationRequest) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		ModelId string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	v.ModelId = unmarshaler.ModelId
	if unmarshaler.ModelId == "" {
		return fmt.Errorf("%T did not include discriminant model_id", v)
	}
	switch unmarshaler.ModelId {
	case "bytedance-seedance-v2":
		value := new(BytedanceSeedance2Request)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		v.BytedanceSeedanceV2 = value
	case "bytedance-seedance-v2-fast":
		value := new(BytedanceSeedance2FastRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		v.BytedanceSeedanceV2Fast = value
	case "bytedance-seedance-v2-mini":
		value := new(BytedanceSeedance2MiniRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		v.BytedanceSeedanceV2Mini = value
	case "bytedance-seedance-v2.5":
		value := new(BytedanceSeedance25Request)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		v.BytedanceSeedanceV25 = value
	case "creatify-aurora":
		value := new(CreatifyAuroraRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		v.CreatifyAurora = value
	case "veo-3.1-fast-generate-001":
		value := new(Veo31FastRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		v.Veo31FastGenerate001 = value
	case "veo-3.1-generate-001":
		value := new(Veo31Request)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		v.Veo31Generate001 = value
	}
	return nil
}

func (v VideoGenerationRequest) MarshalJSON() ([]byte, error) {
	if err := v.validate(); err != nil {
		return nil, err
	}
	if v.BytedanceSeedanceV2 != nil {
		return internal.MarshalJSONWithDiscriminant(v.BytedanceSeedanceV2, "model_id", "bytedance-seedance-v2")
	}
	if v.BytedanceSeedanceV2Fast != nil {
		return internal.MarshalJSONWithDiscriminant(v.BytedanceSeedanceV2Fast, "model_id", "bytedance-seedance-v2-fast")
	}
	if v.BytedanceSeedanceV2Mini != nil {
		return internal.MarshalJSONWithDiscriminant(v.BytedanceSeedanceV2Mini, "model_id", "bytedance-seedance-v2-mini")
	}
	if v.BytedanceSeedanceV25 != nil {
		return internal.MarshalJSONWithDiscriminant(v.BytedanceSeedanceV25, "model_id", "bytedance-seedance-v2.5")
	}
	if v.CreatifyAurora != nil {
		return internal.MarshalJSONWithDiscriminant(v.CreatifyAurora, "model_id", "creatify-aurora")
	}
	if v.Veo31FastGenerate001 != nil {
		return internal.MarshalJSONWithDiscriminant(v.Veo31FastGenerate001, "model_id", "veo-3.1-fast-generate-001")
	}
	if v.Veo31Generate001 != nil {
		return internal.MarshalJSONWithDiscriminant(v.Veo31Generate001, "model_id", "veo-3.1-generate-001")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", v)
}

type VideoGenerationRequestVisitor interface {
	VisitBytedanceSeedanceV2(*BytedanceSeedance2Request) error
	VisitBytedanceSeedanceV2Fast(*BytedanceSeedance2FastRequest) error
	VisitBytedanceSeedanceV2Mini(*BytedanceSeedance2MiniRequest) error
	VisitBytedanceSeedanceV25(*BytedanceSeedance25Request) error
	VisitCreatifyAurora(*CreatifyAuroraRequest) error
	VisitVeo31FastGenerate001(*Veo31FastRequest) error
	VisitVeo31Generate001(*Veo31Request) error
}

func (v *VideoGenerationRequest) Accept(visitor VideoGenerationRequestVisitor) error {
	if v.BytedanceSeedanceV2 != nil {
		return visitor.VisitBytedanceSeedanceV2(v.BytedanceSeedanceV2)
	}
	if v.BytedanceSeedanceV2Fast != nil {
		return visitor.VisitBytedanceSeedanceV2Fast(v.BytedanceSeedanceV2Fast)
	}
	if v.BytedanceSeedanceV2Mini != nil {
		return visitor.VisitBytedanceSeedanceV2Mini(v.BytedanceSeedanceV2Mini)
	}
	if v.BytedanceSeedanceV25 != nil {
		return visitor.VisitBytedanceSeedanceV25(v.BytedanceSeedanceV25)
	}
	if v.CreatifyAurora != nil {
		return visitor.VisitCreatifyAurora(v.CreatifyAurora)
	}
	if v.Veo31FastGenerate001 != nil {
		return visitor.VisitVeo31FastGenerate001(v.Veo31FastGenerate001)
	}
	if v.Veo31Generate001 != nil {
		return visitor.VisitVeo31Generate001(v.Veo31Generate001)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", v)
}

func (v *VideoGenerationRequest) validate() error {
	if v == nil {
		return fmt.Errorf("type %T is nil", v)
	}
	var fields []string
	if v.BytedanceSeedanceV2 != nil {
		fields = append(fields, "bytedance-seedance-v2")
	}
	if v.BytedanceSeedanceV2Fast != nil {
		fields = append(fields, "bytedance-seedance-v2-fast")
	}
	if v.BytedanceSeedanceV2Mini != nil {
		fields = append(fields, "bytedance-seedance-v2-mini")
	}
	if v.BytedanceSeedanceV25 != nil {
		fields = append(fields, "bytedance-seedance-v2.5")
	}
	if v.CreatifyAurora != nil {
		fields = append(fields, "creatify-aurora")
	}
	if v.Veo31FastGenerate001 != nil {
		fields = append(fields, "veo-3.1-fast-generate-001")
	}
	if v.Veo31Generate001 != nil {
		fields = append(fields, "veo-3.1-generate-001")
	}
	if len(fields) == 0 {
		if v.ModelId != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", v, v.ModelId)
		}
		return fmt.Errorf("type %T is empty", v)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", v, fields)
	}
	if v.ModelId != "" {
		field := fields[0]
		if v.ModelId != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				v,
				v.ModelId,
				v,
			)
		}
	}
	return nil
}
