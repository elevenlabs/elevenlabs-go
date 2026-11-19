package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Unified container for all platform limits.
//
// Can be used by:
// - Billing groups (WorkspaceGroupDBModel.platform_limits)
// - Child workspaces (SubscriptionDBModel.platform_limits)
// - API keys (XiApiKeyMetadataDBModel.platform_limits)
//
// All fields are required when platform_limits exists. Use limit=None for unlimited.
type PlatformLimits struct {
	// Platform limit with usage stored in Firestore. Example: credit usage tracked in Firestore.
	Credits *StoredUsagePlatformLimit `json:"credits,omitempty" xml:"credits,omitempty"`
	// Platform limit with usage computed externally. Example: PVCs use the count_owned_pro_voices_in_billing_group function to compute the usage.
	Pvc *ComputedUsagePlatformLimit `json:"pvc,omitempty" xml:"pvc,omitempty"`
	// Platform limit with usage computed externally. Example: PVCs use the count_owned_pro_voices_in_billing_group function to compute the usage.
	Concurrency *ComputedUsagePlatformLimit `json:"concurrency,omitempty" xml:"concurrency,omitempty"`
	// Platform limit with usage computed externally. Example: PVCs use the count_owned_pro_voices_in_billing_group function to compute the usage.
	DubbingConcurrency *ComputedUsagePlatformLimit `json:"dubbing_concurrency,omitempty" xml:"dubbing_concurrency,omitempty"`
	// Platform limit with usage computed externally. Example: PVCs use the count_owned_pro_voices_in_billing_group function to compute the usage.
	MusicConcurrency *ComputedUsagePlatformLimit `json:"music_concurrency,omitempty" xml:"music_concurrency,omitempty"`
	extraProperties  map[string]interface{}
}

func (p *PlatformLimits) GetCredits() (value *StoredUsagePlatformLimit) {
	if p == nil {
		return
	}
	return p.Credits
}

func (p *PlatformLimits) GetPvc() (value *ComputedUsagePlatformLimit) {
	if p == nil {
		return
	}
	return p.Pvc
}

func (p *PlatformLimits) GetConcurrency() (value *ComputedUsagePlatformLimit) {
	if p == nil {
		return
	}
	return p.Concurrency
}

func (p *PlatformLimits) GetDubbingConcurrency() (value *ComputedUsagePlatformLimit) {
	if p == nil {
		return
	}
	return p.DubbingConcurrency
}

func (p *PlatformLimits) GetMusicConcurrency() (value *ComputedUsagePlatformLimit) {
	if p == nil {
		return
	}
	return p.MusicConcurrency
}

func (p *PlatformLimits) SetCredits(credits *StoredUsagePlatformLimit) {
	p.Credits = credits
}

func (p *PlatformLimits) SetPvc(pvc *ComputedUsagePlatformLimit) {
	p.Pvc = pvc
}

func (p *PlatformLimits) SetConcurrency(concurrency *ComputedUsagePlatformLimit) {
	p.Concurrency = concurrency
}

func (p *PlatformLimits) SetDubbingConcurrency(dubbingConcurrency *ComputedUsagePlatformLimit) {
	p.DubbingConcurrency = dubbingConcurrency
}

func (p *PlatformLimits) SetMusicConcurrency(musicConcurrency *ComputedUsagePlatformLimit) {
	p.MusicConcurrency = musicConcurrency
}

func (p *PlatformLimits) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PlatformLimits) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PlatformLimits to string"
	}
	return string(jsonData)
}

func (p *PlatformLimits) UnmarshalJSON(data []byte) error {
	type unmarshaler PlatformLimits
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PlatformLimits(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PlatformLimits) MarshalJSON() ([]byte, error) {
	type embed PlatformLimits
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
