package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type FollowUpRequestCounts struct {
	// How many times contacts asked.
	Requested int `json:"requested" xml:"requested" required:"true"`
	// Requests by what the campaign did about them; adds up to `requested`. Missing keys mean none.
	RequestedByDisposition map[string]int `json:"requested_by_disposition" required:"true"`
	// Leads that asked at least once.
	DistinctLeads   int `json:"distinct_leads" xml:"distinct_leads" required:"true"`
	extraProperties map[string]interface{}
}

func (f *FollowUpRequestCounts) GetRequested() (value int) {
	if f == nil {
		return
	}
	return f.Requested
}

func (f *FollowUpRequestCounts) GetRequestedByDisposition() (value map[string]int) {
	if f == nil {
		return
	}
	return f.RequestedByDisposition
}

func (f *FollowUpRequestCounts) GetDistinctLeads() (value int) {
	if f == nil {
		return
	}
	return f.DistinctLeads
}

func (f *FollowUpRequestCounts) SetRequested(requested int) {
	f.Requested = requested
}

func (f *FollowUpRequestCounts) SetRequestedByDisposition(requestedByDisposition map[string]int) {
	f.RequestedByDisposition = requestedByDisposition
}

func (f *FollowUpRequestCounts) SetDistinctLeads(distinctLeads int) {
	f.DistinctLeads = distinctLeads
}

func (f *FollowUpRequestCounts) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FollowUpRequestCounts) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FollowUpRequestCounts to string"
	}
	return string(jsonData)
}

func (f *FollowUpRequestCounts) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias FollowUpRequestCounts
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = FollowUpRequestCounts(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FollowUpRequestCounts) MarshalJSON() ([]byte, error) {
	type embed FollowUpRequestCounts
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
