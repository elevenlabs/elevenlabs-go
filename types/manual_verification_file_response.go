package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ManualVerificationFileResponse struct {
	// The ID of the file.
	FileId string `json:"file_id" xml:"file_id" required:"true"`
	// The name of the file.
	FileName string `json:"file_name" xml:"file_name" required:"true"`
	// The MIME type of the file.
	MimeType string `json:"mime_type" xml:"mime_type" required:"true"`
	// The size of the file in bytes.
	SizeBytes int `json:"size_bytes" xml:"size_bytes" required:"true"`
	// The date of the file in Unix time.
	UploadDateUnix  int `json:"upload_date_unix" xml:"upload_date_unix" required:"true"`
	extraProperties map[string]interface{}
}

func (m *ManualVerificationFileResponse) GetFileId() (value string) {
	if m == nil {
		return
	}
	return m.FileId
}

func (m *ManualVerificationFileResponse) GetFileName() (value string) {
	if m == nil {
		return
	}
	return m.FileName
}

func (m *ManualVerificationFileResponse) GetMimeType() (value string) {
	if m == nil {
		return
	}
	return m.MimeType
}

func (m *ManualVerificationFileResponse) GetSizeBytes() (value int) {
	if m == nil {
		return
	}
	return m.SizeBytes
}

func (m *ManualVerificationFileResponse) GetUploadDateUnix() (value int) {
	if m == nil {
		return
	}
	return m.UploadDateUnix
}

func (m *ManualVerificationFileResponse) SetFileId(fileId string) {
	m.FileId = fileId
}

func (m *ManualVerificationFileResponse) SetFileName(fileName string) {
	m.FileName = fileName
}

func (m *ManualVerificationFileResponse) SetMimeType(mimeType string) {
	m.MimeType = mimeType
}

func (m *ManualVerificationFileResponse) SetSizeBytes(sizeBytes int) {
	m.SizeBytes = sizeBytes
}

func (m *ManualVerificationFileResponse) SetUploadDateUnix(uploadDateUnix int) {
	m.UploadDateUnix = uploadDateUnix
}

func (m *ManualVerificationFileResponse) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m ManualVerificationFileResponse) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: ManualVerificationFileResponse to string"
	}
	return string(jsonData)
}

func (m *ManualVerificationFileResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias ManualVerificationFileResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = ManualVerificationFileResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *ManualVerificationFileResponse) MarshalJSON() ([]byte, error) {
	type embed ManualVerificationFileResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
