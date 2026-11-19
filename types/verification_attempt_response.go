package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VerificationAttemptResponse struct {
	// The text of the verification attempt.
	Text string `json:"text" xml:"text" required:"true"`
	// The date of the verification attempt in Unix time.
	DateUnix int `json:"date_unix" xml:"date_unix" required:"true"`
	// Whether the verification attempt was accepted.
	Accepted bool `json:"accepted" xml:"accepted" required:"true"`
	// The similarity of the verification attempt.
	Similarity float64 `json:"similarity" xml:"similarity" required:"true"`
	// The Levenshtein distance of the verification attempt.
	LevenshteinDistance float64 `json:"levenshtein_distance" xml:"levenshtein_distance" required:"true"`
	// The recording of the verification attempt.
	Recording       *RecordingResponse `json:"recording,omitempty" xml:"recording,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VerificationAttemptResponse) GetText() (value string) {
	if v == nil {
		return
	}
	return v.Text
}

func (v *VerificationAttemptResponse) GetDateUnix() (value int) {
	if v == nil {
		return
	}
	return v.DateUnix
}

func (v *VerificationAttemptResponse) GetAccepted() (value bool) {
	if v == nil {
		return
	}
	return v.Accepted
}

func (v *VerificationAttemptResponse) GetSimilarity() (value float64) {
	if v == nil {
		return
	}
	return v.Similarity
}

func (v *VerificationAttemptResponse) GetLevenshteinDistance() (value float64) {
	if v == nil {
		return
	}
	return v.LevenshteinDistance
}

func (v *VerificationAttemptResponse) GetRecording() (value *RecordingResponse) {
	if v == nil {
		return
	}
	return v.Recording
}

func (v *VerificationAttemptResponse) SetText(text string) {
	v.Text = text
}

func (v *VerificationAttemptResponse) SetDateUnix(dateUnix int) {
	v.DateUnix = dateUnix
}

func (v *VerificationAttemptResponse) SetAccepted(accepted bool) {
	v.Accepted = accepted
}

func (v *VerificationAttemptResponse) SetSimilarity(similarity float64) {
	v.Similarity = similarity
}

func (v *VerificationAttemptResponse) SetLevenshteinDistance(levenshteinDistance float64) {
	v.LevenshteinDistance = levenshteinDistance
}

func (v *VerificationAttemptResponse) SetRecording(recording *RecordingResponse) {
	v.Recording = recording
}

func (v *VerificationAttemptResponse) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VerificationAttemptResponse) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VerificationAttemptResponse to string"
	}
	return string(jsonData)
}

func (v *VerificationAttemptResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VerificationAttemptResponse) MarshalJSON() ([]byte, error) {
	type embed VerificationAttemptResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
