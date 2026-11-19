package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LeadsProgress struct {
	Total int `json:"total" xml:"total" required:"true"`
	// Unfinished leads whose first call hasn't been placed yet.
	Uncalled int `json:"uncalled" xml:"uncalled" required:"true"`
	// Unfinished leads whose next call an attention item holds back.
	LeadsNeedingAttention int `json:"leads_needing_attention" xml:"leads_needing_attention" required:"true"`
	// Other unfinished leads: a call under way, awaiting analysis or a follow-up.
	Active int `json:"active" xml:"active" required:"true"`
	// Leads that won't be called again.
	Finished int `json:"finished" xml:"finished" required:"true"`
	// Finished leads by why they finished. Missing keys mean none.
	FinishedByDisposition map[string]int `json:"finished_by_disposition" required:"true"`
	extraProperties       map[string]interface{}
}

func (l *LeadsProgress) GetTotal() (value int) {
	if l == nil {
		return
	}
	return l.Total
}

func (l *LeadsProgress) GetUncalled() (value int) {
	if l == nil {
		return
	}
	return l.Uncalled
}

func (l *LeadsProgress) GetLeadsNeedingAttention() (value int) {
	if l == nil {
		return
	}
	return l.LeadsNeedingAttention
}

func (l *LeadsProgress) GetActive() (value int) {
	if l == nil {
		return
	}
	return l.Active
}

func (l *LeadsProgress) GetFinished() (value int) {
	if l == nil {
		return
	}
	return l.Finished
}

func (l *LeadsProgress) GetFinishedByDisposition() (value map[string]int) {
	if l == nil {
		return
	}
	return l.FinishedByDisposition
}

func (l *LeadsProgress) SetTotal(total int) {
	l.Total = total
}

func (l *LeadsProgress) SetUncalled(uncalled int) {
	l.Uncalled = uncalled
}

func (l *LeadsProgress) SetLeadsNeedingAttention(leadsNeedingAttention int) {
	l.LeadsNeedingAttention = leadsNeedingAttention
}

func (l *LeadsProgress) SetActive(active int) {
	l.Active = active
}

func (l *LeadsProgress) SetFinished(finished int) {
	l.Finished = finished
}

func (l *LeadsProgress) SetFinishedByDisposition(finishedByDisposition map[string]int) {
	l.FinishedByDisposition = finishedByDisposition
}

func (l *LeadsProgress) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LeadsProgress) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LeadsProgress to string"
	}
	return string(jsonData)
}

func (l *LeadsProgress) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LeadsProgress
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LeadsProgress(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LeadsProgress) MarshalJSON() ([]byte, error) {
	type embed LeadsProgress
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
