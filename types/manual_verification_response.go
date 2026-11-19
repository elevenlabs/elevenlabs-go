package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ManualVerificationResponse struct {
	// The extra text of the manual verification.
	ExtraText string `json:"extra_text" xml:"extra_text" required:"true"`
	// The date of the manual verification in Unix time.
	RequestTimeUnix int `json:"request_time_unix" xml:"request_time_unix" required:"true"`
	// The files of the manual verification.
	Files           []*ManualVerificationFileResponse `json:"files" xml:"files" required:"true"`
	extraProperties map[string]interface{}
}

func (m *ManualVerificationResponse) GetExtraText() (value string) {
	if m == nil {
		return
	}
	return m.ExtraText
}

func (m *ManualVerificationResponse) GetRequestTimeUnix() (value int) {
	if m == nil {
		return
	}
	return m.RequestTimeUnix
}

func (m *ManualVerificationResponse) GetFiles() (value []*ManualVerificationFileResponse) {
	if m == nil {
		return
	}
	return m.Files
}

func (m *ManualVerificationResponse) SetExtraText(extraText string) {
	m.ExtraText = extraText
}

func (m *ManualVerificationResponse) SetRequestTimeUnix(requestTimeUnix int) {
	m.RequestTimeUnix = requestTimeUnix
}

func (m *ManualVerificationResponse) SetFiles(files []*ManualVerificationFileResponse) {
	m.Files = files
}

func (m *ManualVerificationResponse) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m ManualVerificationResponse) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: ManualVerificationResponse to string"
	}
	return string(jsonData)
}

func (m *ManualVerificationResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias ManualVerificationResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = ManualVerificationResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *ManualVerificationResponse) MarshalJSON() ([]byte, error) {
	type embed ManualVerificationResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
