package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingMetadataPageResponseModel struct {
	Dubs            []*DubbingMetadataResponse `json:"dubs" xml:"dubs" required:"true"`
	NextCursor      *string                    `json:"next_cursor" xml:"next_cursor" required:"true"`
	HasMore         bool                       `json:"has_more" xml:"has_more" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingMetadataPageResponseModel) GetDubs() (value []*DubbingMetadataResponse) {
	if d == nil {
		return
	}
	return d.Dubs
}

func (d *DubbingMetadataPageResponseModel) GetNextCursor() (value *string) {
	if d == nil {
		return
	}
	return d.NextCursor
}

func (d *DubbingMetadataPageResponseModel) GetHasMore() (value bool) {
	if d == nil {
		return
	}
	return d.HasMore
}

func (d *DubbingMetadataPageResponseModel) SetDubs(dubs []*DubbingMetadataResponse) {
	d.Dubs = dubs
}

func (d *DubbingMetadataPageResponseModel) SetNextCursor(nextCursor *string) {
	d.NextCursor = nextCursor
}

func (d *DubbingMetadataPageResponseModel) SetHasMore(hasMore bool) {
	d.HasMore = hasMore
}

func (d *DubbingMetadataPageResponseModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingMetadataPageResponseModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingMetadataPageResponseModel to string"
	}
	return string(jsonData)
}

func (d *DubbingMetadataPageResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingMetadataPageResponseModel) MarshalJSON() ([]byte, error) {
	type embed DubbingMetadataPageResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
