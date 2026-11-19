package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Cost of running post-call analysis on this conversation.
// Present once an analysis pass has run, billed or not.
type AnalysisCharging struct {
	// Cumulative LLM cost of running post-call analysis on this conversation.
	Total *AnalysisRunningTotal `json:"total" xml:"total" required:"true"`
	// LLM cost of the most recent post-call analysis pass on this conversation.
	LastRun         *AnalysisRunSnapshot `json:"last_run" xml:"last_run" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AnalysisCharging) GetTotal() (value *AnalysisRunningTotal) {
	if a == nil {
		return
	}
	return a.Total
}

func (a *AnalysisCharging) GetLastRun() (value *AnalysisRunSnapshot) {
	if a == nil {
		return
	}
	return a.LastRun
}

func (a *AnalysisCharging) SetTotal(total *AnalysisRunningTotal) {
	a.Total = total
}

func (a *AnalysisCharging) SetLastRun(lastRun *AnalysisRunSnapshot) {
	a.LastRun = lastRun
}

func (a *AnalysisCharging) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AnalysisCharging) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AnalysisCharging to string"
	}
	return string(jsonData)
}

func (a *AnalysisCharging) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AnalysisCharging
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AnalysisCharging(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AnalysisCharging) MarshalJSON() ([]byte, error) {
	type embed AnalysisCharging
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
