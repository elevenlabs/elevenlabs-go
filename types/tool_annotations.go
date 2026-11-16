package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Additional properties describing a Tool to clients.
//
// NOTE: all properties in ToolAnnotations are **hints**.
// They are not guaranteed to provide a faithful description of
// tool behavior (including descriptive properties like `title`).
//
// Clients should never make tool use decisions based on ToolAnnotations
// received from untrusted servers.
type ToolAnnotations struct {
	Title           *string `json:"title,omitempty" xml:"title,omitempty"`
	ReadOnlyHint    *bool   `json:"readOnlyHint,omitempty" xml:"readOnlyHint,omitempty"`
	DestructiveHint *bool   `json:"destructiveHint,omitempty" xml:"destructiveHint,omitempty"`
	IdempotentHint  *bool   `json:"idempotentHint,omitempty" xml:"idempotentHint,omitempty"`
	OpenWorldHint   *bool   `json:"openWorldHint,omitempty" xml:"openWorldHint,omitempty"`
	extraProperties map[string]interface{}
}

func (t *ToolAnnotations) GetTitle() (value *string) {
	if t == nil {
		return
	}
	return t.Title
}

func (t *ToolAnnotations) GetReadOnlyHint() (value *bool) {
	if t == nil {
		return
	}
	return t.ReadOnlyHint
}

func (t *ToolAnnotations) GetDestructiveHint() (value *bool) {
	if t == nil {
		return
	}
	return t.DestructiveHint
}

func (t *ToolAnnotations) GetIdempotentHint() (value *bool) {
	if t == nil {
		return
	}
	return t.IdempotentHint
}

func (t *ToolAnnotations) GetOpenWorldHint() (value *bool) {
	if t == nil {
		return
	}
	return t.OpenWorldHint
}

func (t *ToolAnnotations) SetTitle(title *string) {
	t.Title = title
}

func (t *ToolAnnotations) SetReadOnlyHint(readOnlyHint *bool) {
	t.ReadOnlyHint = readOnlyHint
}

func (t *ToolAnnotations) SetDestructiveHint(destructiveHint *bool) {
	t.DestructiveHint = destructiveHint
}

func (t *ToolAnnotations) SetIdempotentHint(idempotentHint *bool) {
	t.IdempotentHint = idempotentHint
}

func (t *ToolAnnotations) SetOpenWorldHint(openWorldHint *bool) {
	t.OpenWorldHint = openWorldHint
}

func (t *ToolAnnotations) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t ToolAnnotations) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ToolAnnotations to string"
	}
	return string(jsonData)
}

func (t *ToolAnnotations) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *ToolAnnotations) MarshalJSON() ([]byte, error) {
	type embed ToolAnnotations
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
