package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// One output of a template run, discriminated on `type`: the `type` of the
// port's `content_schema`. Check `status` for where it is in its lifecycle.
type TemplateOutput struct {
	Type        string
	Array       *TemplateArrayOutput
	Audio       *TemplateAudioOutput
	Boolean     *TemplateBooleanOutput
	Image       *TemplateImageOutput
	Integer     *TemplateIntegerOutput
	Number      *TemplateNumberOutput
	Object      *TemplateObjectOutput
	FieldString *TemplateStringOutput
	Video       *TemplateVideoOutput
}

func (t *TemplateOutput) GetArray() *TemplateArrayOutput {
	if t == nil {
		return nil
	}
	if t.Array == nil {
		return nil
	}
	return t.Array
}

func (t *TemplateOutput) GetAudio() *TemplateAudioOutput {
	if t == nil {
		return nil
	}
	if t.Audio == nil {
		return nil
	}
	return t.Audio
}

func (t *TemplateOutput) GetBoolean() *TemplateBooleanOutput {
	if t == nil {
		return nil
	}
	if t.Boolean == nil {
		return nil
	}
	return t.Boolean
}

func (t *TemplateOutput) GetImage() *TemplateImageOutput {
	if t == nil {
		return nil
	}
	if t.Image == nil {
		return nil
	}
	return t.Image
}

func (t *TemplateOutput) GetInteger() *TemplateIntegerOutput {
	if t == nil {
		return nil
	}
	if t.Integer == nil {
		return nil
	}
	return t.Integer
}

func (t *TemplateOutput) GetNumber() *TemplateNumberOutput {
	if t == nil {
		return nil
	}
	if t.Number == nil {
		return nil
	}
	return t.Number
}

func (t *TemplateOutput) GetObject() *TemplateObjectOutput {
	if t == nil {
		return nil
	}
	if t.Object == nil {
		return nil
	}
	return t.Object
}

func (t *TemplateOutput) GetFieldString() *TemplateStringOutput {
	if t == nil {
		return nil
	}
	if t.FieldString == nil {
		return nil
	}
	return t.FieldString
}

func (t *TemplateOutput) GetVideo() *TemplateVideoOutput {
	if t == nil {
		return nil
	}
	if t.Video == nil {
		return nil
	}
	return t.Video
}

func (t *TemplateOutput) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *TemplateOutput) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	t.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", t)
	}
	switch unmarshaler.Type {
	case "array":
		value := new(TemplateArrayOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Array = value
	case "audio":
		value := new(TemplateAudioOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Audio = value
	case "boolean":
		value := new(TemplateBooleanOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Boolean = value
	case "image":
		value := new(TemplateImageOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Image = value
	case "integer":
		value := new(TemplateIntegerOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Integer = value
	case "number":
		value := new(TemplateNumberOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Number = value
	case "object":
		value := new(TemplateObjectOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Object = value
	case "string":
		value := new(TemplateStringOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.FieldString = value
	case "video":
		value := new(TemplateVideoOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Video = value
	}
	return nil
}

func (t TemplateOutput) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.Array != nil {
		return internal.MarshalJSONWithDiscriminant(t.Array, "type", "array")
	}
	if t.Audio != nil {
		return internal.MarshalJSONWithDiscriminant(t.Audio, "type", "audio")
	}
	if t.Boolean != nil {
		return internal.MarshalJSONWithDiscriminant(t.Boolean, "type", "boolean")
	}
	if t.Image != nil {
		return internal.MarshalJSONWithDiscriminant(t.Image, "type", "image")
	}
	if t.Integer != nil {
		return internal.MarshalJSONWithDiscriminant(t.Integer, "type", "integer")
	}
	if t.Number != nil {
		return internal.MarshalJSONWithDiscriminant(t.Number, "type", "number")
	}
	if t.Object != nil {
		return internal.MarshalJSONWithDiscriminant(t.Object, "type", "object")
	}
	if t.FieldString != nil {
		return internal.MarshalJSONWithDiscriminant(t.FieldString, "type", "string")
	}
	if t.Video != nil {
		return internal.MarshalJSONWithDiscriminant(t.Video, "type", "video")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TemplateOutputVisitor interface {
	VisitArray(*TemplateArrayOutput) error
	VisitAudio(*TemplateAudioOutput) error
	VisitBoolean(*TemplateBooleanOutput) error
	VisitImage(*TemplateImageOutput) error
	VisitInteger(*TemplateIntegerOutput) error
	VisitNumber(*TemplateNumberOutput) error
	VisitObject(*TemplateObjectOutput) error
	VisitFieldString(*TemplateStringOutput) error
	VisitVideo(*TemplateVideoOutput) error
}

func (t *TemplateOutput) Accept(visitor TemplateOutputVisitor) error {
	if t.Array != nil {
		return visitor.VisitArray(t.Array)
	}
	if t.Audio != nil {
		return visitor.VisitAudio(t.Audio)
	}
	if t.Boolean != nil {
		return visitor.VisitBoolean(t.Boolean)
	}
	if t.Image != nil {
		return visitor.VisitImage(t.Image)
	}
	if t.Integer != nil {
		return visitor.VisitInteger(t.Integer)
	}
	if t.Number != nil {
		return visitor.VisitNumber(t.Number)
	}
	if t.Object != nil {
		return visitor.VisitObject(t.Object)
	}
	if t.FieldString != nil {
		return visitor.VisitFieldString(t.FieldString)
	}
	if t.Video != nil {
		return visitor.VisitVideo(t.Video)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *TemplateOutput) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.Array != nil {
		fields = append(fields, "array")
	}
	if t.Audio != nil {
		fields = append(fields, "audio")
	}
	if t.Boolean != nil {
		fields = append(fields, "boolean")
	}
	if t.Image != nil {
		fields = append(fields, "image")
	}
	if t.Integer != nil {
		fields = append(fields, "integer")
	}
	if t.Number != nil {
		fields = append(fields, "number")
	}
	if t.Object != nil {
		fields = append(fields, "object")
	}
	if t.FieldString != nil {
		fields = append(fields, "string")
	}
	if t.Video != nil {
		fields = append(fields, "video")
	}
	if len(fields) == 0 {
		if t.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", t, t.Type)
		}
		return fmt.Errorf("type %T is empty", t)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", t, fields)
	}
	if t.Type != "" {
		field := fields[0]
		if t.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				t,
				t.Type,
				t,
			)
		}
	}
	return nil
}
