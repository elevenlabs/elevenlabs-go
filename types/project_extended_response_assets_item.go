package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ProjectExtendedResponseAssetsItem struct {
	Type  string
	Audio *ProjectExternalAudioResponseModel
	Image *ProjectImageResponseModel
	Video *ProjectVideoResponseModel
}

func (p *ProjectExtendedResponseAssetsItem) GetAudio() *ProjectExternalAudioResponseModel {
	if p == nil {
		return nil
	}
	if p.Audio == nil {
		return nil
	}
	return p.Audio
}

func (p *ProjectExtendedResponseAssetsItem) GetImage() *ProjectImageResponseModel {
	if p == nil {
		return nil
	}
	if p.Image == nil {
		return nil
	}
	return p.Image
}

func (p *ProjectExtendedResponseAssetsItem) GetVideo() *ProjectVideoResponseModel {
	if p == nil {
		return nil
	}
	if p.Video == nil {
		return nil
	}
	return p.Video
}

func (p *ProjectExtendedResponseAssetsItem) GetType() string {
	if p == nil {
		return ""
	}
	return p.Type
}

func (p *ProjectExtendedResponseAssetsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	p.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", p)
	}
	switch unmarshaler.Type {
	case "audio":
		value := new(ProjectExternalAudioResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Audio = value
	case "image":
		value := new(ProjectImageResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Image = value
	case "video":
		value := new(ProjectVideoResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Video = value
	}
	return nil
}

func (p ProjectExtendedResponseAssetsItem) MarshalJSON() ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	if p.Audio != nil {
		return internal.MarshalJSONWithDiscriminant(p.Audio, "type", "audio")
	}
	if p.Image != nil {
		return internal.MarshalJSONWithDiscriminant(p.Image, "type", "image")
	}
	if p.Video != nil {
		return internal.MarshalJSONWithDiscriminant(p.Video, "type", "video")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", p)
}

type ProjectExtendedResponseAssetsItemVisitor interface {
	VisitAudio(*ProjectExternalAudioResponseModel) error
	VisitImage(*ProjectImageResponseModel) error
	VisitVideo(*ProjectVideoResponseModel) error
}

func (p *ProjectExtendedResponseAssetsItem) Accept(visitor ProjectExtendedResponseAssetsItemVisitor) error {
	if p.Audio != nil {
		return visitor.VisitAudio(p.Audio)
	}
	if p.Image != nil {
		return visitor.VisitImage(p.Image)
	}
	if p.Video != nil {
		return visitor.VisitVideo(p.Video)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", p)
}

func (p *ProjectExtendedResponseAssetsItem) validate() error {
	if p == nil {
		return fmt.Errorf("type %T is nil", p)
	}
	var fields []string
	if p.Audio != nil {
		fields = append(fields, "audio")
	}
	if p.Image != nil {
		fields = append(fields, "image")
	}
	if p.Video != nil {
		fields = append(fields, "video")
	}
	if len(fields) == 0 {
		if p.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", p, p.Type)
		}
		return fmt.Errorf("type %T is empty", p)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", p, fields)
	}
	if p.Type != "" {
		field := fields[0]
		if p.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				p,
				p.Type,
				p,
			)
		}
	}
	return nil
}
