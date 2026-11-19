package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RecordingResponse struct {
	// The ID of the recording.
	RecordingId string `json:"recording_id" xml:"recording_id" required:"true"`
	// The MIME type of the recording.
	MimeType string `json:"mime_type" xml:"mime_type" required:"true"`
	// The size of the recording in bytes.
	SizeBytes int `json:"size_bytes" xml:"size_bytes" required:"true"`
	// The date of the recording in Unix time.
	UploadDateUnix int `json:"upload_date_unix" xml:"upload_date_unix" required:"true"`
	// The transcription of the recording.
	Transcription   string `json:"transcription" xml:"transcription" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RecordingResponse) GetRecordingId() (value string) {
	if r == nil {
		return
	}
	return r.RecordingId
}

func (r *RecordingResponse) GetMimeType() (value string) {
	if r == nil {
		return
	}
	return r.MimeType
}

func (r *RecordingResponse) GetSizeBytes() (value int) {
	if r == nil {
		return
	}
	return r.SizeBytes
}

func (r *RecordingResponse) GetUploadDateUnix() (value int) {
	if r == nil {
		return
	}
	return r.UploadDateUnix
}

func (r *RecordingResponse) GetTranscription() (value string) {
	if r == nil {
		return
	}
	return r.Transcription
}

func (r *RecordingResponse) SetRecordingId(recordingId string) {
	r.RecordingId = recordingId
}

func (r *RecordingResponse) SetMimeType(mimeType string) {
	r.MimeType = mimeType
}

func (r *RecordingResponse) SetSizeBytes(sizeBytes int) {
	r.SizeBytes = sizeBytes
}

func (r *RecordingResponse) SetUploadDateUnix(uploadDateUnix int) {
	r.UploadDateUnix = uploadDateUnix
}

func (r *RecordingResponse) SetTranscription(transcription string) {
	r.Transcription = transcription
}

func (r *RecordingResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RecordingResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RecordingResponse to string"
	}
	return string(jsonData)
}

func (r *RecordingResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RecordingResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RecordingResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RecordingResponse) MarshalJSON() ([]byte, error) {
	type embed RecordingResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
