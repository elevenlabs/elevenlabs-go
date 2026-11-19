package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ContentSchema struct {
	Type        string
	Array       *ArraySchema
	Audio       *AudioSchema
	Boolean     *BooleanSchema
	Image       *ImageSchema
	Integer     *IntegerSchema
	Number      *NumberSchema
	Object      *ObjectSchema
	FieldString *StringSchema
	Video       *VideoSchema
	Voice       *VoiceSchema
}

func (c *ContentSchema) GetArray() *ArraySchema {
	if c == nil {
		return nil
	}
	if c.Array == nil {
		return nil
	}
	return c.Array
}

func (c *ContentSchema) GetAudio() *AudioSchema {
	if c == nil {
		return nil
	}
	if c.Audio == nil {
		return nil
	}
	return c.Audio
}

func (c *ContentSchema) GetBoolean() *BooleanSchema {
	if c == nil {
		return nil
	}
	if c.Boolean == nil {
		return nil
	}
	return c.Boolean
}

func (c *ContentSchema) GetImage() *ImageSchema {
	if c == nil {
		return nil
	}
	if c.Image == nil {
		return nil
	}
	return c.Image
}

func (c *ContentSchema) GetInteger() *IntegerSchema {
	if c == nil {
		return nil
	}
	if c.Integer == nil {
		return nil
	}
	return c.Integer
}

func (c *ContentSchema) GetNumber() *NumberSchema {
	if c == nil {
		return nil
	}
	if c.Number == nil {
		return nil
	}
	return c.Number
}

func (c *ContentSchema) GetObject() *ObjectSchema {
	if c == nil {
		return nil
	}
	if c.Object == nil {
		return nil
	}
	return c.Object
}

func (c *ContentSchema) GetFieldString() *StringSchema {
	if c == nil {
		return nil
	}
	if c.FieldString == nil {
		return nil
	}
	return c.FieldString
}

func (c *ContentSchema) GetVideo() *VideoSchema {
	if c == nil {
		return nil
	}
	if c.Video == nil {
		return nil
	}
	return c.Video
}

func (c *ContentSchema) GetVoice() *VoiceSchema {
	if c == nil {
		return nil
	}
	if c.Voice == nil {
		return nil
	}
	return c.Voice
}

func (c *ContentSchema) GetType() string {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *ContentSchema) UnmarshalJSON(data []byte) error {
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
	case "array":
		value := new(ArraySchema)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Array = value
	case "audio":
		value := new(AudioSchema)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Audio = value
	case "boolean":
		value := new(BooleanSchema)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Boolean = value
	case "image":
		value := new(ImageSchema)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Image = value
	case "integer":
		value := new(IntegerSchema)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Integer = value
	case "number":
		value := new(NumberSchema)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Number = value
	case "object":
		value := new(ObjectSchema)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Object = value
	case "string":
		value := new(StringSchema)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.FieldString = value
	case "video":
		value := new(VideoSchema)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Video = value
	case "voice":
		value := new(VoiceSchema)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Voice = value
	}
	return nil
}

func (c ContentSchema) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.Array != nil {
		return internal.MarshalJSONWithDiscriminant(c.Array, "type", "array")
	}
	if c.Audio != nil {
		return internal.MarshalJSONWithDiscriminant(c.Audio, "type", "audio")
	}
	if c.Boolean != nil {
		return internal.MarshalJSONWithDiscriminant(c.Boolean, "type", "boolean")
	}
	if c.Image != nil {
		return internal.MarshalJSONWithDiscriminant(c.Image, "type", "image")
	}
	if c.Integer != nil {
		return internal.MarshalJSONWithDiscriminant(c.Integer, "type", "integer")
	}
	if c.Number != nil {
		return internal.MarshalJSONWithDiscriminant(c.Number, "type", "number")
	}
	if c.Object != nil {
		return internal.MarshalJSONWithDiscriminant(c.Object, "type", "object")
	}
	if c.FieldString != nil {
		return internal.MarshalJSONWithDiscriminant(c.FieldString, "type", "string")
	}
	if c.Video != nil {
		return internal.MarshalJSONWithDiscriminant(c.Video, "type", "video")
	}
	if c.Voice != nil {
		return internal.MarshalJSONWithDiscriminant(c.Voice, "type", "voice")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type ContentSchemaVisitor interface {
	VisitArray(*ArraySchema) error
	VisitAudio(*AudioSchema) error
	VisitBoolean(*BooleanSchema) error
	VisitImage(*ImageSchema) error
	VisitInteger(*IntegerSchema) error
	VisitNumber(*NumberSchema) error
	VisitObject(*ObjectSchema) error
	VisitFieldString(*StringSchema) error
	VisitVideo(*VideoSchema) error
	VisitVoice(*VoiceSchema) error
}

func (c *ContentSchema) Accept(visitor ContentSchemaVisitor) error {
	if c.Array != nil {
		return visitor.VisitArray(c.Array)
	}
	if c.Audio != nil {
		return visitor.VisitAudio(c.Audio)
	}
	if c.Boolean != nil {
		return visitor.VisitBoolean(c.Boolean)
	}
	if c.Image != nil {
		return visitor.VisitImage(c.Image)
	}
	if c.Integer != nil {
		return visitor.VisitInteger(c.Integer)
	}
	if c.Number != nil {
		return visitor.VisitNumber(c.Number)
	}
	if c.Object != nil {
		return visitor.VisitObject(c.Object)
	}
	if c.FieldString != nil {
		return visitor.VisitFieldString(c.FieldString)
	}
	if c.Video != nil {
		return visitor.VisitVideo(c.Video)
	}
	if c.Voice != nil {
		return visitor.VisitVoice(c.Voice)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *ContentSchema) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.Array != nil {
		fields = append(fields, "array")
	}
	if c.Audio != nil {
		fields = append(fields, "audio")
	}
	if c.Boolean != nil {
		fields = append(fields, "boolean")
	}
	if c.Image != nil {
		fields = append(fields, "image")
	}
	if c.Integer != nil {
		fields = append(fields, "integer")
	}
	if c.Number != nil {
		fields = append(fields, "number")
	}
	if c.Object != nil {
		fields = append(fields, "object")
	}
	if c.FieldString != nil {
		fields = append(fields, "string")
	}
	if c.Video != nil {
		fields = append(fields, "video")
	}
	if c.Voice != nil {
		fields = append(fields, "voice")
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
