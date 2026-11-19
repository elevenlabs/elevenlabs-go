package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OrderItemRequestOutput struct {
	Kind          string
	Dub           *DubOrderItemRequest
	Subtitles     *SubtitleOrderItemRequest
	Transcription *TranscriptionOrderItemRequest
}

func (o *OrderItemRequestOutput) GetDub() *DubOrderItemRequest {
	if o == nil {
		return nil
	}
	if o.Dub == nil {
		return nil
	}
	return o.Dub
}

func (o *OrderItemRequestOutput) GetSubtitles() *SubtitleOrderItemRequest {
	if o == nil {
		return nil
	}
	if o.Subtitles == nil {
		return nil
	}
	return o.Subtitles
}

func (o *OrderItemRequestOutput) GetTranscription() *TranscriptionOrderItemRequest {
	if o == nil {
		return nil
	}
	if o.Transcription == nil {
		return nil
	}
	return o.Transcription
}

func (o *OrderItemRequestOutput) GetKind() string {
	if o == nil {
		return ""
	}
	return o.Kind
}

func (o *OrderItemRequestOutput) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	o.Kind = unmarshaler.Kind
	if unmarshaler.Kind == "" {
		return fmt.Errorf("%T did not include discriminant kind", o)
	}
	switch unmarshaler.Kind {
	case "dub":
		value := new(DubOrderItemRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		o.Dub = value
	case "subtitles":
		value := new(SubtitleOrderItemRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		o.Subtitles = value
	case "transcription":
		value := new(TranscriptionOrderItemRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		o.Transcription = value
	}
	return nil
}

func (o OrderItemRequestOutput) MarshalJSON() ([]byte, error) {
	if err := o.validate(); err != nil {
		return nil, err
	}
	if o.Dub != nil {
		return internal.MarshalJSONWithDiscriminant(o.Dub, "kind", "dub")
	}
	if o.Subtitles != nil {
		return internal.MarshalJSONWithDiscriminant(o.Subtitles, "kind", "subtitles")
	}
	if o.Transcription != nil {
		return internal.MarshalJSONWithDiscriminant(o.Transcription, "kind", "transcription")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", o)
}

type OrderItemRequestOutputVisitor interface {
	VisitDub(*DubOrderItemRequest) error
	VisitSubtitles(*SubtitleOrderItemRequest) error
	VisitTranscription(*TranscriptionOrderItemRequest) error
}

func (o *OrderItemRequestOutput) Accept(visitor OrderItemRequestOutputVisitor) error {
	if o.Dub != nil {
		return visitor.VisitDub(o.Dub)
	}
	if o.Subtitles != nil {
		return visitor.VisitSubtitles(o.Subtitles)
	}
	if o.Transcription != nil {
		return visitor.VisitTranscription(o.Transcription)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", o)
}

func (o *OrderItemRequestOutput) validate() error {
	if o == nil {
		return fmt.Errorf("type %T is nil", o)
	}
	var fields []string
	if o.Dub != nil {
		fields = append(fields, "dub")
	}
	if o.Subtitles != nil {
		fields = append(fields, "subtitles")
	}
	if o.Transcription != nil {
		fields = append(fields, "transcription")
	}
	if len(fields) == 0 {
		if o.Kind != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", o, o.Kind)
		}
		return fmt.Errorf("type %T is empty", o)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", o, fields)
	}
	if o.Kind != "" {
		field := fields[0]
		if o.Kind != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				o,
				o.Kind,
				o,
			)
		}
	}
	return nil
}
