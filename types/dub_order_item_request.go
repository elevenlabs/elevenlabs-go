package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubOrderItemRequest struct {
	// The type of order item.
	Kind    *string `json:"kind,omitempty" xml:"kind,omitempty"`
	MediaId MediaId `json:"media_id" xml:"media_id" required:"true" pattern:"^prodmedia_[a-z0-9]{26}$"`
	// The language code of the source media (e.g. 'en', 'es').
	SourceLanguage string `json:"source_language" xml:"source_language" required:"true"`
	// List of target language codes to dub the media into.
	DestinationLanguages []string `json:"destination_languages" xml:"destination_languages" required:"true"`
	// Whether to generate captions for the dubbed outputs.
	IncludeCaptions bool `json:"include_captions" xml:"include_captions" required:"true"`
	// Whether to generate captions for the source language.
	IncludeSourceCaptions bool `json:"include_source_captions" xml:"include_source_captions" required:"true"`
	// Optional free-text instructions for the dubbing team.
	Instructions *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	// Whether captions should use SDH format, which includes descriptions for deaf and hard-of-hearing viewers.
	CaptionsSdh     *bool `json:"captions_sdh,omitempty" xml:"captions_sdh,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubOrderItemRequest) GetKind() (value *string) {
	if d == nil {
		return
	}
	return d.Kind
}

func (d *DubOrderItemRequest) GetMediaId() (value MediaId) {
	if d == nil {
		return
	}
	return d.MediaId
}

func (d *DubOrderItemRequest) GetSourceLanguage() (value string) {
	if d == nil {
		return
	}
	return d.SourceLanguage
}

func (d *DubOrderItemRequest) GetDestinationLanguages() (value []string) {
	if d == nil {
		return
	}
	return d.DestinationLanguages
}

func (d *DubOrderItemRequest) GetIncludeCaptions() (value bool) {
	if d == nil {
		return
	}
	return d.IncludeCaptions
}

func (d *DubOrderItemRequest) GetIncludeSourceCaptions() (value bool) {
	if d == nil {
		return
	}
	return d.IncludeSourceCaptions
}

func (d *DubOrderItemRequest) GetInstructions() (value *string) {
	if d == nil {
		return
	}
	return d.Instructions
}

func (d *DubOrderItemRequest) GetCaptionsSdh() (value *bool) {
	if d == nil {
		return
	}
	return d.CaptionsSdh
}

func (d *DubOrderItemRequest) SetKind(kind *string) {
	d.Kind = kind
}

func (d *DubOrderItemRequest) SetMediaId(mediaId MediaId) {
	d.MediaId = mediaId
}

func (d *DubOrderItemRequest) SetSourceLanguage(sourceLanguage string) {
	d.SourceLanguage = sourceLanguage
}

func (d *DubOrderItemRequest) SetDestinationLanguages(destinationLanguages []string) {
	d.DestinationLanguages = destinationLanguages
}

func (d *DubOrderItemRequest) SetIncludeCaptions(includeCaptions bool) {
	d.IncludeCaptions = includeCaptions
}

func (d *DubOrderItemRequest) SetIncludeSourceCaptions(includeSourceCaptions bool) {
	d.IncludeSourceCaptions = includeSourceCaptions
}

func (d *DubOrderItemRequest) SetInstructions(instructions *string) {
	d.Instructions = instructions
}

func (d *DubOrderItemRequest) SetCaptionsSdh(captionsSdh *bool) {
	d.CaptionsSdh = captionsSdh
}

func (d *DubOrderItemRequest) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubOrderItemRequest) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubOrderItemRequest to string"
	}
	return string(jsonData)
}

func (d *DubOrderItemRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubOrderItemRequest) MarshalJSON() ([]byte, error) {
	type embed DubOrderItemRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
