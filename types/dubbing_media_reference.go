package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingMediaReference struct {
	Src             string  `json:"src" xml:"src" required:"true"`
	ContentType     string  `json:"content_type" xml:"content_type" required:"true"`
	BucketName      string  `json:"bucket_name" xml:"bucket_name" required:"true"`
	RandomPathSlug  string  `json:"random_path_slug" xml:"random_path_slug" required:"true"`
	DurationSecs    float64 `json:"duration_secs" xml:"duration_secs" required:"true"`
	IsAudio         bool    `json:"is_audio" xml:"is_audio" required:"true"`
	Url             string  `json:"url" xml:"url" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingMediaReference) GetSrc() (value string) {
	if d == nil {
		return
	}
	return d.Src
}

func (d *DubbingMediaReference) GetContentType() (value string) {
	if d == nil {
		return
	}
	return d.ContentType
}

func (d *DubbingMediaReference) GetBucketName() (value string) {
	if d == nil {
		return
	}
	return d.BucketName
}

func (d *DubbingMediaReference) GetRandomPathSlug() (value string) {
	if d == nil {
		return
	}
	return d.RandomPathSlug
}

func (d *DubbingMediaReference) GetDurationSecs() (value float64) {
	if d == nil {
		return
	}
	return d.DurationSecs
}

func (d *DubbingMediaReference) GetIsAudio() (value bool) {
	if d == nil {
		return
	}
	return d.IsAudio
}

func (d *DubbingMediaReference) GetUrl() (value string) {
	if d == nil {
		return
	}
	return d.Url
}

func (d *DubbingMediaReference) SetSrc(src string) {
	d.Src = src
}

func (d *DubbingMediaReference) SetContentType(contentType string) {
	d.ContentType = contentType
}

func (d *DubbingMediaReference) SetBucketName(bucketName string) {
	d.BucketName = bucketName
}

func (d *DubbingMediaReference) SetRandomPathSlug(randomPathSlug string) {
	d.RandomPathSlug = randomPathSlug
}

func (d *DubbingMediaReference) SetDurationSecs(durationSecs float64) {
	d.DurationSecs = durationSecs
}

func (d *DubbingMediaReference) SetIsAudio(isAudio bool) {
	d.IsAudio = isAudio
}

func (d *DubbingMediaReference) SetUrl(url string) {
	d.Url = url
}

func (d *DubbingMediaReference) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingMediaReference) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingMediaReference to string"
	}
	return string(jsonData)
}

func (d *DubbingMediaReference) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingMediaReference
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingMediaReference(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingMediaReference) MarshalJSON() ([]byte, error) {
	type embed DubbingMediaReference
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
