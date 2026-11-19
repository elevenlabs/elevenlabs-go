package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A Voice Library voice the uploaded audio matched, offered in place of the clone.
type LibraryVoiceMatchResponseModel struct {
	// The ID of the matched Voice Library voice.
	LibraryVoiceId string `json:"library_voice_id" xml:"library_voice_id" required:"true"`
	// The public owner ID of the Voice Library voice. Pass it with library_voice_id to POST /v1/voices/add/{public_user_id}/{voice_id} to add the voice.
	PublicOwnerId   string `json:"public_owner_id" xml:"public_owner_id" required:"true"`
	extraProperties map[string]interface{}
}

func (l *LibraryVoiceMatchResponseModel) GetLibraryVoiceId() (value string) {
	if l == nil {
		return
	}
	return l.LibraryVoiceId
}

func (l *LibraryVoiceMatchResponseModel) GetPublicOwnerId() (value string) {
	if l == nil {
		return
	}
	return l.PublicOwnerId
}

func (l *LibraryVoiceMatchResponseModel) SetLibraryVoiceId(libraryVoiceId string) {
	l.LibraryVoiceId = libraryVoiceId
}

func (l *LibraryVoiceMatchResponseModel) SetPublicOwnerId(publicOwnerId string) {
	l.PublicOwnerId = publicOwnerId
}

func (l *LibraryVoiceMatchResponseModel) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LibraryVoiceMatchResponseModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LibraryVoiceMatchResponseModel to string"
	}
	return string(jsonData)
}

func (l *LibraryVoiceMatchResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LibraryVoiceMatchResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LibraryVoiceMatchResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LibraryVoiceMatchResponseModel) MarshalJSON() ([]byte, error) {
	type embed LibraryVoiceMatchResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
