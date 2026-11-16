package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type HistoryAlignmentsResponseModel struct {
	Alignment           *HistoryAlignmentResponseModel `json:"alignment" xml:"alignment" required:"true"`
	NormalizedAlignment *HistoryAlignmentResponseModel `json:"normalized_alignment" xml:"normalized_alignment" required:"true"`
	extraProperties     map[string]interface{}
}

func (h *HistoryAlignmentsResponseModel) GetAlignment() (value *HistoryAlignmentResponseModel) {
	if h == nil {
		return
	}
	return h.Alignment
}

func (h *HistoryAlignmentsResponseModel) GetNormalizedAlignment() (value *HistoryAlignmentResponseModel) {
	if h == nil {
		return
	}
	return h.NormalizedAlignment
}

func (h *HistoryAlignmentsResponseModel) SetAlignment(alignment *HistoryAlignmentResponseModel) {
	h.Alignment = alignment
}

func (h *HistoryAlignmentsResponseModel) SetNormalizedAlignment(normalizedAlignment *HistoryAlignmentResponseModel) {
	h.NormalizedAlignment = normalizedAlignment
}

func (h *HistoryAlignmentsResponseModel) GetExtraProperties() map[string]interface{} {
	if h == nil {
		return nil
	}
	return h.extraProperties
}

func (h HistoryAlignmentsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: HistoryAlignmentsResponseModel to string"
	}
	return string(jsonData)
}

func (h *HistoryAlignmentsResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, h); err != nil {
		return err
	}
	type alias HistoryAlignmentsResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*h = HistoryAlignmentsResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *h)
	if err != nil {
		return err
	}
	h.extraProperties = extraProperties
	return nil
}

func (h *HistoryAlignmentsResponseModel) MarshalJSON() ([]byte, error) {
	type embed HistoryAlignmentsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*h),
	}
	return json.Marshal(marshaler)
}
