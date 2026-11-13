package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BatchFailureResponseModel struct {
	status          string
	ErrorCode       int    `json:"error_code" xml:"error_code" required:"true"`
	ErrorStatus     string `json:"error_status" xml:"error_status" required:"true"`
	ErrorMessage    string `json:"error_message" xml:"error_message" required:"true"`
	extraProperties map[string]interface{}
}

func (b *BatchFailureResponseModel) GetErrorCode() (value int) {
	if b == nil {
		return
	}
	return b.ErrorCode
}

func (b *BatchFailureResponseModel) GetErrorStatus() (value string) {
	if b == nil {
		return
	}
	return b.ErrorStatus
}

func (b *BatchFailureResponseModel) GetErrorMessage() (value string) {
	if b == nil {
		return
	}
	return b.ErrorMessage
}

func (b *BatchFailureResponseModel) SetErrorCode(errorCode int) {
	b.ErrorCode = errorCode
}

func (b *BatchFailureResponseModel) SetErrorStatus(errorStatus string) {
	b.ErrorStatus = errorStatus
}

func (b *BatchFailureResponseModel) SetErrorMessage(errorMessage string) {
	b.ErrorMessage = errorMessage
}

func (b *BatchFailureResponseModel) Status() string {
	return b.status
}

func (b *BatchFailureResponseModel) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BatchFailureResponseModel) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BatchFailureResponseModel to string"
	}
	return string(jsonData)
}

func (b *BatchFailureResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BatchFailureResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BatchFailureResponseModel(tmp)
	var literals struct {
		Status *string `json:"status"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Status != nil && *literals.Status != "failure" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", b, "failure", *literals.Status)
	}
	b.status = "failure"
	extraProperties, err := internal.ExtractExtraProperties(data, *b, "status")
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BatchFailureResponseModel) MarshalJSON() ([]byte, error) {
	type embed BatchFailureResponseModel
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*b),
		Status: "failure",
	}
	return json.Marshal(marshaler)
}
