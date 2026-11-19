package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PendingBlocksMetadataModel struct {
	TargetGlobalOffsetMs *int     `json:"target_global_offset_ms" xml:"target_global_offset_ms" required:"true"`
	BlockIds             []string `json:"block_ids" xml:"block_ids" required:"true"`
	extraProperties      map[string]interface{}
}

func (p *PendingBlocksMetadataModel) GetTargetGlobalOffsetMs() (value *int) {
	if p == nil {
		return
	}
	return p.TargetGlobalOffsetMs
}

func (p *PendingBlocksMetadataModel) GetBlockIds() (value []string) {
	if p == nil {
		return
	}
	return p.BlockIds
}

func (p *PendingBlocksMetadataModel) SetTargetGlobalOffsetMs(targetGlobalOffsetMs *int) {
	p.TargetGlobalOffsetMs = targetGlobalOffsetMs
}

func (p *PendingBlocksMetadataModel) SetBlockIds(blockIds []string) {
	p.BlockIds = blockIds
}

func (p *PendingBlocksMetadataModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PendingBlocksMetadataModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PendingBlocksMetadataModel to string"
	}
	return string(jsonData)
}

func (p *PendingBlocksMetadataModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PendingBlocksMetadataModel) MarshalJSON() ([]byte, error) {
	type embed PendingBlocksMetadataModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
