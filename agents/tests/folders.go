package tests

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

type BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost struct {
	// The name of the folder to create
	Name string `json:"name" xml:"name" required:"true"`
	// The ID of the parent folder. If not provided, the folder will be created at the root level.
	ParentFolderId *string `json:"parent_folder_id,omitempty" xml:"parent_folder_id,omitempty"`
}

func (b *BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost) GetParentFolderId() (value *string) {
	if b == nil {
		return
	}
	return b.ParentFolderId
}

func (b *BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost) SetName(name string) {
	b.Name = name
}

func (b *BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost) SetParentFolderId(parentFolderId *string) {
	b.ParentFolderId = parentFolderId
}

func (b *BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateAgentTestFolderV1ConvaiAgentTestingFoldersPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch struct {
	// The folder ID.
	FolderId string `json:"-" url:"-"`
	// The new name for the folder
	Name string `json:"name" xml:"name" required:"true"`
}

func (b *BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch) GetFolderId() (value string) {
	if b == nil {
		return
	}
	return b.FolderId
}

func (b *BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch) SetFolderId(folderId string) {
	b.FolderId = folderId
}

func (b *BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch) SetName(name string) {
	b.Name = name
}

func (b *BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch(tmp)
	return nil
}

func (b *BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateAgentTestFolderV1ConvaiAgentTestingFoldersFolderIdPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	deleteFoldersRequestFieldForce = big.NewInt(1 << 0)
)

type DeleteFoldersRequest struct {
	Force *bool `json:"-" url:"force,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteFoldersRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetForce sets the Force field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteFoldersRequest) SetForce(force *bool) {
	d.Force = force
	d.require(deleteFoldersRequestFieldForce)
}
