package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// OCSF User object.
//
// Spec: https://schema.ocsf.io/1.6.0/objects/user
type UserModel struct {
	// Username
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Unique user identifier
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// OCSF User type IDs.
	//
	// Spec: https://schema.ocsf.io/1.6.0/objects/user
	TypeId *int `json:"type_id,omitempty" xml:"type_id,omitempty"`
	// Account type description
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// User email address
	EmailAddr *string `json:"email_addr,omitempty" xml:"email_addr,omitempty"`
	// Full name of the user
	FullName *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// User's domain
	Domain          *string `json:"domain,omitempty" xml:"domain,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UserModel) GetName() (value *string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UserModel) GetUid() (value *string) {
	if u == nil {
		return
	}
	return u.Uid
}

func (u *UserModel) GetTypeId() (value *int) {
	if u == nil {
		return
	}
	return u.TypeId
}

func (u *UserModel) GetType() (value *string) {
	if u == nil {
		return
	}
	return u.Type
}

func (u *UserModel) GetEmailAddr() (value *string) {
	if u == nil {
		return
	}
	return u.EmailAddr
}

func (u *UserModel) GetFullName() (value *string) {
	if u == nil {
		return
	}
	return u.FullName
}

func (u *UserModel) GetDomain() (value *string) {
	if u == nil {
		return
	}
	return u.Domain
}

func (u *UserModel) SetName(name *string) {
	u.Name = name
}

func (u *UserModel) SetUid(uid *string) {
	u.Uid = uid
}

func (u *UserModel) SetTypeId(typeId *int) {
	u.TypeId = typeId
}

func (u *UserModel) SetType(type_ *string) {
	u.Type = type_
}

func (u *UserModel) SetEmailAddr(emailAddr *string) {
	u.EmailAddr = emailAddr
}

func (u *UserModel) SetFullName(fullName *string) {
	u.FullName = fullName
}

func (u *UserModel) SetDomain(domain *string) {
	u.Domain = domain
}

func (u *UserModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserModel to string"
	}
	return string(jsonData)
}

func (u *UserModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UserModel) MarshalJSON() ([]byte, error) {
	type embed UserModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
