package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type Contributor struct {
	Name            string  `json:"name" xml:"name" required:"true"`
	Role            string  `json:"role" xml:"role" required:"true"`
	Bio             *string `json:"bio,omitempty" xml:"bio,omitempty"`
	ProfileId       *string `json:"profile_id,omitempty" xml:"profile_id,omitempty"`
	extraProperties map[string]interface{}
}

func (c *Contributor) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *Contributor) GetRole() (value string) {
	if c == nil {
		return
	}
	return c.Role
}

func (c *Contributor) GetBio() (value *string) {
	if c == nil {
		return
	}
	return c.Bio
}

func (c *Contributor) GetProfileId() (value *string) {
	if c == nil {
		return
	}
	return c.ProfileId
}

func (c *Contributor) SetName(name string) {
	c.Name = name
}

func (c *Contributor) SetRole(role string) {
	c.Role = role
}

func (c *Contributor) SetBio(bio *string) {
	c.Bio = bio
}

func (c *Contributor) SetProfileId(profileId *string) {
	c.ProfileId = profileId
}

func (c *Contributor) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c Contributor) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: Contributor to string"
	}
	return string(jsonData)
}

func (c *Contributor) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *Contributor) MarshalJSON() ([]byte, error) {
	type embed Contributor
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
