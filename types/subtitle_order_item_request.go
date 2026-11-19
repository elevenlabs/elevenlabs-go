package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitleOrderItemRequest struct {
	// The type of order item.
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The IDs of the uploaded media files to generate subtitles for.
	MediaIds []MediaId `json:"media_ids" xml:"media_ids" required:"true" maxItems:"50"`
	// The language code of the source media (e.g. 'en', 'es').
	SourceLanguage string `json:"source_language" xml:"source_language" required:"true"`
	// List of target language codes. Subtitles will be generated for each media file in each destination language.
	DestinationLanguages []string           `json:"destination_languages" xml:"destination_languages" required:"true"`
	CueOptions           *CueOptionsRequest `json:"cue_options,omitempty" xml:"cue_options,omitempty"`
	// Whether subtitles should use SDH format, which includes descriptions for deaf and hard-of-hearing viewers.
	Sdh *bool `json:"sdh,omitempty" xml:"sdh,omitempty"`
	// Optional free-text instructions for the subtitling team.
	Instructions    *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SubtitleOrderItemRequest) GetKind() (value *string) {
	if s == nil {
		return
	}
	return s.Kind
}

func (s *SubtitleOrderItemRequest) GetMediaIds() (value []MediaId) {
	if s == nil {
		return
	}
	return s.MediaIds
}

func (s *SubtitleOrderItemRequest) GetSourceLanguage() (value string) {
	if s == nil {
		return
	}
	return s.SourceLanguage
}

func (s *SubtitleOrderItemRequest) GetDestinationLanguages() (value []string) {
	if s == nil {
		return
	}
	return s.DestinationLanguages
}

func (s *SubtitleOrderItemRequest) GetCueOptions() (value *CueOptionsRequest) {
	if s == nil {
		return
	}
	return s.CueOptions
}

func (s *SubtitleOrderItemRequest) GetSdh() (value *bool) {
	if s == nil {
		return
	}
	return s.Sdh
}

func (s *SubtitleOrderItemRequest) GetInstructions() (value *string) {
	if s == nil {
		return
	}
	return s.Instructions
}

func (s *SubtitleOrderItemRequest) SetKind(kind *string) {
	s.Kind = kind
}

func (s *SubtitleOrderItemRequest) SetMediaIds(mediaIds []MediaId) {
	s.MediaIds = mediaIds
}

func (s *SubtitleOrderItemRequest) SetSourceLanguage(sourceLanguage string) {
	s.SourceLanguage = sourceLanguage
}

func (s *SubtitleOrderItemRequest) SetDestinationLanguages(destinationLanguages []string) {
	s.DestinationLanguages = destinationLanguages
}

func (s *SubtitleOrderItemRequest) SetCueOptions(cueOptions *CueOptionsRequest) {
	s.CueOptions = cueOptions
}

func (s *SubtitleOrderItemRequest) SetSdh(sdh *bool) {
	s.Sdh = sdh
}

func (s *SubtitleOrderItemRequest) SetInstructions(instructions *string) {
	s.Instructions = instructions
}

func (s *SubtitleOrderItemRequest) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitleOrderItemRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitleOrderItemRequest to string"
	}
	return string(jsonData)
}

func (s *SubtitleOrderItemRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitleOrderItemRequest) MarshalJSON() ([]byte, error) {
	type embed SubtitleOrderItemRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
