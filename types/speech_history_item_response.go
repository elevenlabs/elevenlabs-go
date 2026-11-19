package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeechHistoryItemResponse struct {
	// The ID of the history item.
	HistoryItemId string `json:"history_item_id" xml:"history_item_id" required:"true"`
	// The ID of the request.
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// The ID of the voice used.
	VoiceId *string `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	// The ID of the model.
	ModelId *string `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// The name of the voice.
	VoiceName *string `json:"voice_name,omitempty" xml:"voice_name,omitempty"`
	// The category of the voice. Either 'premade', 'cloned', 'generated' or 'professional'.
	VoiceCategory *SpeechHistoryItemResponseVoiceCategory `json:"voice_category,omitempty" xml:"voice_category,omitempty"`
	// The text used to generate the audio item.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// Unix timestamp of when the item was created.
	DateUnix int `json:"date_unix" xml:"date_unix" required:"true"`
	// The character count change from.
	CharacterCountChangeFrom int `json:"character_count_change_from" xml:"character_count_change_from" required:"true"`
	// The character count change to.
	CharacterCountChangeTo int `json:"character_count_change_to" xml:"character_count_change_to" required:"true"`
	// The content type of the generated item.
	ContentType string `json:"content_type" xml:"content_type" required:"true"`
	State       any    `json:"state" xml:"state" required:"true"`
	// The settings of the history item.
	Settings map[string]any `json:"settings,omitempty"`
	// The ID of the share link.
	ShareLinkId *string `json:"share_link_id,omitempty" xml:"share_link_id,omitempty"`
	// The source of the history item. Either TTS (text to speech), STS (speech to text), AN (audio native), Projects, Dubbing, PlayAPI, PD (pronunciation dictionary) or ConvAI (Agents Platform).
	Source *SpeechHistoryItemResponseSource `json:"source,omitempty" xml:"source,omitempty"`
	// The alignments of the history item.
	Alignments *HistoryAlignmentsResponseModel `json:"alignments,omitempty" xml:"alignments,omitempty"`
	// The dialogue (voice and text pairs) used to generate the audio item. If this is set then the top level `text` and `voice_id` fields will be empty.
	Dialogue []*DialogueInputResponseModel `json:"dialogue,omitempty" xml:"dialogue,omitempty"`
	// The output format the audio was originally generated in.
	OutputFormat    *string `json:"output_format,omitempty" xml:"output_format,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SpeechHistoryItemResponse) GetHistoryItemId() (value string) {
	if s == nil {
		return
	}
	return s.HistoryItemId
}

func (s *SpeechHistoryItemResponse) GetRequestId() (value *string) {
	if s == nil {
		return
	}
	return s.RequestId
}

func (s *SpeechHistoryItemResponse) GetVoiceId() (value *string) {
	if s == nil {
		return
	}
	return s.VoiceId
}

func (s *SpeechHistoryItemResponse) GetModelId() (value *string) {
	if s == nil {
		return
	}
	return s.ModelId
}

func (s *SpeechHistoryItemResponse) GetVoiceName() (value *string) {
	if s == nil {
		return
	}
	return s.VoiceName
}

func (s *SpeechHistoryItemResponse) GetVoiceCategory() (value *SpeechHistoryItemResponseVoiceCategory) {
	if s == nil {
		return
	}
	return s.VoiceCategory
}

func (s *SpeechHistoryItemResponse) GetText() (value *string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SpeechHistoryItemResponse) GetDateUnix() (value int) {
	if s == nil {
		return
	}
	return s.DateUnix
}

func (s *SpeechHistoryItemResponse) GetCharacterCountChangeFrom() (value int) {
	if s == nil {
		return
	}
	return s.CharacterCountChangeFrom
}

func (s *SpeechHistoryItemResponse) GetCharacterCountChangeTo() (value int) {
	if s == nil {
		return
	}
	return s.CharacterCountChangeTo
}

func (s *SpeechHistoryItemResponse) GetContentType() (value string) {
	if s == nil {
		return
	}
	return s.ContentType
}

func (s *SpeechHistoryItemResponse) GetState() (value any) {
	if s == nil {
		return
	}
	return s.State
}

func (s *SpeechHistoryItemResponse) GetSettings() (value map[string]any) {
	if s == nil {
		return
	}
	return s.Settings
}

func (s *SpeechHistoryItemResponse) GetShareLinkId() (value *string) {
	if s == nil {
		return
	}
	return s.ShareLinkId
}

func (s *SpeechHistoryItemResponse) GetSource() (value *SpeechHistoryItemResponseSource) {
	if s == nil {
		return
	}
	return s.Source
}

func (s *SpeechHistoryItemResponse) GetAlignments() (value *HistoryAlignmentsResponseModel) {
	if s == nil {
		return
	}
	return s.Alignments
}

func (s *SpeechHistoryItemResponse) GetDialogue() (value []*DialogueInputResponseModel) {
	if s == nil {
		return
	}
	return s.Dialogue
}

func (s *SpeechHistoryItemResponse) GetOutputFormat() (value *string) {
	if s == nil {
		return
	}
	return s.OutputFormat
}

func (s *SpeechHistoryItemResponse) SetHistoryItemId(historyItemId string) {
	s.HistoryItemId = historyItemId
}

func (s *SpeechHistoryItemResponse) SetRequestId(requestId *string) {
	s.RequestId = requestId
}

func (s *SpeechHistoryItemResponse) SetVoiceId(voiceId *string) {
	s.VoiceId = voiceId
}

func (s *SpeechHistoryItemResponse) SetModelId(modelId *string) {
	s.ModelId = modelId
}

func (s *SpeechHistoryItemResponse) SetVoiceName(voiceName *string) {
	s.VoiceName = voiceName
}

func (s *SpeechHistoryItemResponse) SetVoiceCategory(voiceCategory *SpeechHistoryItemResponseVoiceCategory) {
	s.VoiceCategory = voiceCategory
}

func (s *SpeechHistoryItemResponse) SetText(text *string) {
	s.Text = text
}

func (s *SpeechHistoryItemResponse) SetDateUnix(dateUnix int) {
	s.DateUnix = dateUnix
}

func (s *SpeechHistoryItemResponse) SetCharacterCountChangeFrom(characterCountChangeFrom int) {
	s.CharacterCountChangeFrom = characterCountChangeFrom
}

func (s *SpeechHistoryItemResponse) SetCharacterCountChangeTo(characterCountChangeTo int) {
	s.CharacterCountChangeTo = characterCountChangeTo
}

func (s *SpeechHistoryItemResponse) SetContentType(contentType string) {
	s.ContentType = contentType
}

func (s *SpeechHistoryItemResponse) SetState(state any) {
	s.State = state
}

func (s *SpeechHistoryItemResponse) SetSettings(settings map[string]any) {
	s.Settings = settings
}

func (s *SpeechHistoryItemResponse) SetShareLinkId(shareLinkId *string) {
	s.ShareLinkId = shareLinkId
}

func (s *SpeechHistoryItemResponse) SetSource(source *SpeechHistoryItemResponseSource) {
	s.Source = source
}

func (s *SpeechHistoryItemResponse) SetAlignments(alignments *HistoryAlignmentsResponseModel) {
	s.Alignments = alignments
}

func (s *SpeechHistoryItemResponse) SetDialogue(dialogue []*DialogueInputResponseModel) {
	s.Dialogue = dialogue
}

func (s *SpeechHistoryItemResponse) SetOutputFormat(outputFormat *string) {
	s.OutputFormat = outputFormat
}

func (s *SpeechHistoryItemResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeechHistoryItemResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeechHistoryItemResponse to string"
	}
	return string(jsonData)
}

func (s *SpeechHistoryItemResponse) UnmarshalJSON(data []byte) error {
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

func (s *SpeechHistoryItemResponse) MarshalJSON() ([]byte, error) {
	type embed SpeechHistoryItemResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
