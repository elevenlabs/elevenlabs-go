package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UserAudioChunk struct {
	// Base64-encoded audio data from the user's microphone.
	UserAudioChunk  string `json:"user_audio_chunk" xml:"user_audio_chunk" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UserAudioChunk) GetUserAudioChunk() (value string) {
	if u == nil {
		return
	}
	return u.UserAudioChunk
}

func (u *UserAudioChunk) SetUserAudioChunk(userAudioChunk string) {
	u.UserAudioChunk = userAudioChunk
}

func (u *UserAudioChunk) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserAudioChunk) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserAudioChunk to string"
	}
	return string(jsonData)
}

func (u *UserAudioChunk) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UserAudioChunk
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UserAudioChunk(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UserAudioChunk) MarshalJSON() ([]byte, error) {
	type embed UserAudioChunk
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
