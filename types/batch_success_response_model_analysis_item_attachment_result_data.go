package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BatchSuccessResponseModelAnalysisItemAttachmentResultData struct {
	status          string
	Data            *AnalysisItemAttachmentResultData `json:"data" xml:"data" required:"true"`
	extraProperties map[string]interface{}
}

func (b *BatchSuccessResponseModelAnalysisItemAttachmentResultData) GetData() (value *AnalysisItemAttachmentResultData) {
	if b == nil {
		return
	}
	return b.Data
}

func (b *BatchSuccessResponseModelAnalysisItemAttachmentResultData) SetData(data *AnalysisItemAttachmentResultData) {
	b.Data = data
}

func (b *BatchSuccessResponseModelAnalysisItemAttachmentResultData) Status() string {
	return b.status
}

func (b *BatchSuccessResponseModelAnalysisItemAttachmentResultData) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BatchSuccessResponseModelAnalysisItemAttachmentResultData) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BatchSuccessResponseModelAnalysisItemAttachmentResultData to string"
	}
	return string(jsonData)
}

func (b *BatchSuccessResponseModelAnalysisItemAttachmentResultData) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BatchSuccessResponseModelAnalysisItemAttachmentResultData
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BatchSuccessResponseModelAnalysisItemAttachmentResultData(tmp)
	var literals struct {
		Status *string `json:"status"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Status != nil && *literals.Status != "success" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", b, "success", *literals.Status)
	}
	b.status = "success"
	extraProperties, err := internal.ExtractExtraProperties(data, *b, "status")
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BatchSuccessResponseModelAnalysisItemAttachmentResultData) MarshalJSON() ([]byte, error) {
	type embed BatchSuccessResponseModelAnalysisItemAttachmentResultData
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*b),
		Status: "success",
	}
	return json.Marshal(marshaler)
}
