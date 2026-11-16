package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DeliverableInfo struct {
	// A time-limited URL to download the delivered file.
	SignedUrl string `json:"signed_url" xml:"signed_url" required:"true"`
	// The MIME type of the delivered file (e.g. 'video/mp4').
	ContentType string `json:"content_type" xml:"content_type" required:"true"`
	// The name of the delivered file.
	Name string `json:"name" xml:"name" required:"true"`
	// The version number of the deliverable.
	Version         *int `json:"version,omitempty" xml:"version,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DeliverableInfo) GetSignedUrl() (value string) {
	if d == nil {
		return
	}
	return d.SignedUrl
}

func (d *DeliverableInfo) GetContentType() (value string) {
	if d == nil {
		return
	}
	return d.ContentType
}

func (d *DeliverableInfo) GetName() (value string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DeliverableInfo) GetVersion() (value *int) {
	if d == nil {
		return
	}
	return d.Version
}

func (d *DeliverableInfo) SetSignedUrl(signedUrl string) {
	d.SignedUrl = signedUrl
}

func (d *DeliverableInfo) SetContentType(contentType string) {
	d.ContentType = contentType
}

func (d *DeliverableInfo) SetName(name string) {
	d.Name = name
}

func (d *DeliverableInfo) SetVersion(version *int) {
	d.Version = version
}

func (d *DeliverableInfo) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeliverableInfo) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeliverableInfo to string"
	}
	return string(jsonData)
}

func (d *DeliverableInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DeliverableInfo
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DeliverableInfo(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeliverableInfo) MarshalJSON() ([]byte, error) {
	type embed DeliverableInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
