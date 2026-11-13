package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PlainPlatformMetadata struct {
	ThreadId        string `json:"thread_id" xml:"thread_id" required:"true"`
	WorkspaceId     string `json:"workspace_id" xml:"workspace_id" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PlainPlatformMetadata) GetThreadId() (value string) {
	if p == nil {
		return
	}
	return p.ThreadId
}

func (p *PlainPlatformMetadata) GetWorkspaceId() (value string) {
	if p == nil {
		return
	}
	return p.WorkspaceId
}

func (p *PlainPlatformMetadata) SetThreadId(threadId string) {
	p.ThreadId = threadId
}

func (p *PlainPlatformMetadata) SetWorkspaceId(workspaceId string) {
	p.WorkspaceId = workspaceId
}

func (p *PlainPlatformMetadata) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PlainPlatformMetadata) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PlainPlatformMetadata to string"
	}
	return string(jsonData)
}

func (p *PlainPlatformMetadata) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PlainPlatformMetadata
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PlainPlatformMetadata(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "conversation_initiation_source")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PlainPlatformMetadata) MarshalJSON() ([]byte, error) {
	type embed PlainPlatformMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
