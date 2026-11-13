package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BranchSyncStatus struct {
	// Latest automatic attempt to pull the synced branch's new versions in, run after each of them. Null when none has run since sync was turned on.
	LastAttempt     *AutoRebaseAttempt `json:"last_attempt,omitempty" xml:"last_attempt,omitempty"`
	extraProperties map[string]interface{}
}

func (b *BranchSyncStatus) GetLastAttempt() (value *AutoRebaseAttempt) {
	if b == nil {
		return
	}
	return b.LastAttempt
}

func (b *BranchSyncStatus) SetLastAttempt(lastAttempt *AutoRebaseAttempt) {
	b.LastAttempt = lastAttempt
}

func (b *BranchSyncStatus) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BranchSyncStatus) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BranchSyncStatus to string"
	}
	return string(jsonData)
}

func (b *BranchSyncStatus) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BranchSyncStatus) MarshalJSON() ([]byte, error) {
	type embed BranchSyncStatus
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
