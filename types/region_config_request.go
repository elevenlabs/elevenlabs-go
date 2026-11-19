package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RegionConfigRequest struct {
	// Valid Twilio region IDs.
	RegionId TwilioRegionId `json:"region_id" xml:"region_id" required:"true"`
	// Auth Token for this region
	Token string `json:"token" xml:"token" required:"true"`
	// Valid Twilio edge locations.
	EdgeLocation    TwilioEdgeLocation `json:"edge_location" xml:"edge_location" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RegionConfigRequest) GetRegionId() (value TwilioRegionId) {
	if r == nil {
		return
	}
	return r.RegionId
}

func (r *RegionConfigRequest) GetToken() (value string) {
	if r == nil {
		return
	}
	return r.Token
}

func (r *RegionConfigRequest) GetEdgeLocation() (value TwilioEdgeLocation) {
	if r == nil {
		return
	}
	return r.EdgeLocation
}

func (r *RegionConfigRequest) SetRegionId(regionId TwilioRegionId) {
	r.RegionId = regionId
}

func (r *RegionConfigRequest) SetToken(token string) {
	r.Token = token
}

func (r *RegionConfigRequest) SetEdgeLocation(edgeLocation TwilioEdgeLocation) {
	r.EdgeLocation = edgeLocation
}

func (r *RegionConfigRequest) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RegionConfigRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RegionConfigRequest to string"
	}
	return string(jsonData)
}

func (r *RegionConfigRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RegionConfigRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RegionConfigRequest(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RegionConfigRequest) MarshalJSON() ([]byte, error) {
	type embed RegionConfigRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
