package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceNow struct {
	// Represents a calendar date as individual numeric components.
	Today *DateModel `json:"today" xml:"today" required:"true"`
	// Current hour (0-23)
	Hour int `json:"hour" xml:"hour" required:"true" min:"0"`
	// Current minute (0-59)
	Minute int `json:"minute" xml:"minute" required:"true" min:"0"`
	// Unix timestamp in milliseconds when this was generated
	FetchedAtMs     int `json:"fetched_at_ms" xml:"fetched_at_ms" required:"true"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceNow) GetToday() (value *DateModel) {
	if w == nil {
		return
	}
	return w.Today
}

func (w *WorkspaceNow) GetHour() (value int) {
	if w == nil {
		return
	}
	return w.Hour
}

func (w *WorkspaceNow) GetMinute() (value int) {
	if w == nil {
		return
	}
	return w.Minute
}

func (w *WorkspaceNow) GetFetchedAtMs() (value int) {
	if w == nil {
		return
	}
	return w.FetchedAtMs
}

func (w *WorkspaceNow) SetToday(today *DateModel) {
	w.Today = today
}

func (w *WorkspaceNow) SetHour(hour int) {
	w.Hour = hour
}

func (w *WorkspaceNow) SetMinute(minute int) {
	w.Minute = minute
}

func (w *WorkspaceNow) SetFetchedAtMs(fetchedAtMs int) {
	w.FetchedAtMs = fetchedAtMs
}

func (w *WorkspaceNow) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceNow) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceNow to string"
	}
	return string(jsonData)
}

func (w *WorkspaceNow) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, w); err != nil {
		return err
	}
	type alias WorkspaceNow
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*w = WorkspaceNow(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceNow) MarshalJSON() ([]byte, error) {
	type embed WorkspaceNow
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
