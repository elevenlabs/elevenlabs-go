package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RagIndexBatchSuccessfulResponseModel struct {
	status          string
	Data            *RagDocumentIndexResponseModel `json:"data" xml:"data" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RagIndexBatchSuccessfulResponseModel) GetData() (value *RagDocumentIndexResponseModel) {
	if r == nil {
		return
	}
	return r.Data
}

func (r *RagIndexBatchSuccessfulResponseModel) SetData(data *RagDocumentIndexResponseModel) {
	r.Data = data
}

func (r *RagIndexBatchSuccessfulResponseModel) Status() string {
	return r.status
}

func (r *RagIndexBatchSuccessfulResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RagIndexBatchSuccessfulResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RagIndexBatchSuccessfulResponseModel to string"
	}
	return string(jsonData)
}

func (r *RagIndexBatchSuccessfulResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RagIndexBatchSuccessfulResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RagIndexBatchSuccessfulResponseModel(tmp)
	var literals struct {
		Status *string `json:"status"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Status != nil && *literals.Status != "success" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", r, "success", *literals.Status)
	}
	r.status = "success"
	extraProperties, err := internal.ExtractExtraProperties(data, *r, "status")
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RagIndexBatchSuccessfulResponseModel) MarshalJSON() ([]byte, error) {
	type embed RagIndexBatchSuccessfulResponseModel
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*r),
		Status: "success",
	}
	return json.Marshal(marshaler)
}
