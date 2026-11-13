package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type FineTuningResponse struct {
	// Whether the user is allowed to fine-tune the voice.
	IsAllowedToFineTune *bool `json:"is_allowed_to_fine_tune,omitempty" xml:"is_allowed_to_fine_tune,omitempty"`
	// The state of the fine-tuning process for each model.
	State map[string]FineTuningResponseStateValue `json:"state,omitempty"`
	// List of verification failures in the fine-tuning process.
	VerificationFailures []string `json:"verification_failures,omitempty" xml:"verification_failures,omitempty"`
	// The number of verification attempts in the fine-tuning process.
	VerificationAttemptsCount *int `json:"verification_attempts_count,omitempty" xml:"verification_attempts_count,omitempty"`
	// Whether a manual verification was requested for the fine-tuning process.
	ManualVerificationRequested *bool `json:"manual_verification_requested,omitempty" xml:"manual_verification_requested,omitempty"`
	// The language of the fine-tuning process.
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The progress of the fine-tuning process.
	Progress map[string]*float64 `json:"progress,omitempty"`
	// The message of the fine-tuning process.
	Message map[string]*string `json:"message,omitempty"`
	// The duration of the dataset in seconds.
	DatasetDurationSeconds *float64 `json:"dataset_duration_seconds,omitempty" xml:"dataset_duration_seconds,omitempty"`
	// The number of verification attempts.
	VerificationAttempts []*VerificationAttemptResponse `json:"verification_attempts,omitempty" xml:"verification_attempts,omitempty"`
	// List of slice IDs.
	SliceIds []string `json:"slice_ids,omitempty" xml:"slice_ids,omitempty"`
	// The manual verification of the fine-tuning process.
	ManualVerification *ManualVerificationResponse `json:"manual_verification,omitempty" xml:"manual_verification,omitempty"`
	// The maximum number of verification attempts.
	MaxVerificationAttempts *int `json:"max_verification_attempts,omitempty" xml:"max_verification_attempts,omitempty"`
	// The next maximum verification attempts reset time in Unix milliseconds.
	NextMaxVerificationAttemptsResetUnixMs *int `json:"next_max_verification_attempts_reset_unix_ms,omitempty" xml:"next_max_verification_attempts_reset_unix_ms,omitempty"`
	FinetuningState                        any  `json:"finetuning_state,omitempty" xml:"finetuning_state,omitempty"`
	extraProperties                        map[string]interface{}
}

func (f *FineTuningResponse) GetIsAllowedToFineTune() (value *bool) {
	if f == nil {
		return
	}
	return f.IsAllowedToFineTune
}

func (f *FineTuningResponse) GetState() (value map[string]FineTuningResponseStateValue) {
	if f == nil {
		return
	}
	return f.State
}

func (f *FineTuningResponse) GetVerificationFailures() (value []string) {
	if f == nil {
		return
	}
	return f.VerificationFailures
}

func (f *FineTuningResponse) GetVerificationAttemptsCount() (value *int) {
	if f == nil {
		return
	}
	return f.VerificationAttemptsCount
}

func (f *FineTuningResponse) GetManualVerificationRequested() (value *bool) {
	if f == nil {
		return
	}
	return f.ManualVerificationRequested
}

func (f *FineTuningResponse) GetLanguage() (value *string) {
	if f == nil {
		return
	}
	return f.Language
}

func (f *FineTuningResponse) GetProgress() (value map[string]*float64) {
	if f == nil {
		return
	}
	return f.Progress
}

func (f *FineTuningResponse) GetMessage() (value map[string]*string) {
	if f == nil {
		return
	}
	return f.Message
}

func (f *FineTuningResponse) GetDatasetDurationSeconds() (value *float64) {
	if f == nil {
		return
	}
	return f.DatasetDurationSeconds
}

func (f *FineTuningResponse) GetVerificationAttempts() (value []*VerificationAttemptResponse) {
	if f == nil {
		return
	}
	return f.VerificationAttempts
}

func (f *FineTuningResponse) GetSliceIds() (value []string) {
	if f == nil {
		return
	}
	return f.SliceIds
}

func (f *FineTuningResponse) GetManualVerification() (value *ManualVerificationResponse) {
	if f == nil {
		return
	}
	return f.ManualVerification
}

func (f *FineTuningResponse) GetMaxVerificationAttempts() (value *int) {
	if f == nil {
		return
	}
	return f.MaxVerificationAttempts
}

func (f *FineTuningResponse) GetNextMaxVerificationAttemptsResetUnixMs() (value *int) {
	if f == nil {
		return
	}
	return f.NextMaxVerificationAttemptsResetUnixMs
}

func (f *FineTuningResponse) GetFinetuningState() (value any) {
	if f == nil {
		return
	}
	return f.FinetuningState
}

func (f *FineTuningResponse) SetIsAllowedToFineTune(isAllowedToFineTune *bool) {
	f.IsAllowedToFineTune = isAllowedToFineTune
}

func (f *FineTuningResponse) SetState(state map[string]FineTuningResponseStateValue) {
	f.State = state
}

func (f *FineTuningResponse) SetVerificationFailures(verificationFailures []string) {
	f.VerificationFailures = verificationFailures
}

func (f *FineTuningResponse) SetVerificationAttemptsCount(verificationAttemptsCount *int) {
	f.VerificationAttemptsCount = verificationAttemptsCount
}

func (f *FineTuningResponse) SetManualVerificationRequested(manualVerificationRequested *bool) {
	f.ManualVerificationRequested = manualVerificationRequested
}

func (f *FineTuningResponse) SetLanguage(language *string) {
	f.Language = language
}

func (f *FineTuningResponse) SetProgress(progress map[string]*float64) {
	f.Progress = progress
}

func (f *FineTuningResponse) SetMessage(message map[string]*string) {
	f.Message = message
}

func (f *FineTuningResponse) SetDatasetDurationSeconds(datasetDurationSeconds *float64) {
	f.DatasetDurationSeconds = datasetDurationSeconds
}

func (f *FineTuningResponse) SetVerificationAttempts(verificationAttempts []*VerificationAttemptResponse) {
	f.VerificationAttempts = verificationAttempts
}

func (f *FineTuningResponse) SetSliceIds(sliceIds []string) {
	f.SliceIds = sliceIds
}

func (f *FineTuningResponse) SetManualVerification(manualVerification *ManualVerificationResponse) {
	f.ManualVerification = manualVerification
}

func (f *FineTuningResponse) SetMaxVerificationAttempts(maxVerificationAttempts *int) {
	f.MaxVerificationAttempts = maxVerificationAttempts
}

func (f *FineTuningResponse) SetNextMaxVerificationAttemptsResetUnixMs(nextMaxVerificationAttemptsResetUnixMs *int) {
	f.NextMaxVerificationAttemptsResetUnixMs = nextMaxVerificationAttemptsResetUnixMs
}

func (f *FineTuningResponse) SetFinetuningState(finetuningState any) {
	f.FinetuningState = finetuningState
}

func (f *FineTuningResponse) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FineTuningResponse) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FineTuningResponse to string"
	}
	return string(jsonData)
}

func (f *FineTuningResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, f); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FineTuningResponse) MarshalJSON() ([]byte, error) {
	type embed FineTuningResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
