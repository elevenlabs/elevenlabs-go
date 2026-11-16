package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ZendeskPlatformMetadata struct {
	TicketId        string `json:"ticket_id" xml:"ticket_id" required:"true"`
	extraProperties map[string]interface{}
}

func (z *ZendeskPlatformMetadata) GetTicketId() (value string) {
	if z == nil {
		return
	}
	return z.TicketId
}

func (z *ZendeskPlatformMetadata) SetTicketId(ticketId string) {
	z.TicketId = ticketId
}

func (z *ZendeskPlatformMetadata) GetExtraProperties() map[string]interface{} {
	if z == nil {
		return nil
	}
	return z.extraProperties
}

func (z ZendeskPlatformMetadata) String() string {
	jsonData, err := json.MarshalIndent(z, "", "  ")
	if err != nil {
		return "error converting struct: ZendeskPlatformMetadata to string"
	}
	return string(jsonData)
}

func (z *ZendeskPlatformMetadata) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, z); err != nil {
		return err
	}
	type alias ZendeskPlatformMetadata
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*z = ZendeskPlatformMetadata(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *z, "conversation_initiation_source")
	if err != nil {
		return err
	}
	z.extraProperties = extraProperties
	return nil
}

func (z *ZendeskPlatformMetadata) MarshalJSON() ([]byte, error) {
	type embed ZendeskPlatformMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*z),
	}
	return json.Marshal(marshaler)
}
