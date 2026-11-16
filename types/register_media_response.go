package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RegisterMediaResponse struct {
	MediaId         MediaId `json:"media_id" xml:"media_id" required:"true" pattern:"^prodmedia_[a-z0-9]{26}$"`
	extraProperties map[string]interface{}
}

func (r *RegisterMediaResponse) GetMediaId() (value MediaId) {
	if r == nil {
		return
	}
	return r.MediaId
}

func (r *RegisterMediaResponse) SetMediaId(mediaId MediaId) {
	r.MediaId = mediaId
}

func (r *RegisterMediaResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RegisterMediaResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegisterMediaResponse to string"
	}
	return string(jsonData)
}

func (r *RegisterMediaResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RegisterMediaResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RegisterMediaResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RegisterMediaResponse) MarshalJSON() ([]byte, error) {
	type embed RegisterMediaResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
