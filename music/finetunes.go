package music

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	"io"
	big "math/big"
)

type BodyCreateMusicFinetuneV1MusicFinetunesPost struct {
	// Name for the finetune (5-200 characters).
	Name string `json:"name" xml:"name" required:"true" maxLength:"200" minLength:"5"`
	// Primary musical genre of the finetune.
	PrimaryGenre string `json:"primary_genre" xml:"primary_genre" required:"true"`
	// Audio files to train on.
	Files []io.Reader `json:"files,omitempty" xml:"files,omitempty" maxItems:"50"`
	// Tags to associate with the finetune.
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty"`
	// Finetune visibility. Only 'private' and 'workspace' can be set.
	Visibility *CreateFinetunesRequestVisibility `json:"visibility,omitempty" xml:"visibility,omitempty"`
	ModelId    *types.MusicModelId               `json:"model_id,omitempty" xml:"model_id,omitempty"`
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) GetPrimaryGenre() (value string) {
	if b == nil {
		return
	}
	return b.PrimaryGenre
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) GetFiles() (value []io.Reader) {
	if b == nil {
		return
	}
	return b.Files
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) GetTags() (value []string) {
	if b == nil {
		return
	}
	return b.Tags
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) GetVisibility() (value *CreateFinetunesRequestVisibility) {
	if b == nil {
		return
	}
	return b.Visibility
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) GetModelId() (value *types.MusicModelId) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) SetName(name string) {
	b.Name = name
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) SetPrimaryGenre(primaryGenre string) {
	b.PrimaryGenre = primaryGenre
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) SetFiles(files []io.Reader) {
	b.Files = files
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) SetTags(tags []string) {
	b.Tags = tags
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) SetVisibility(visibility *CreateFinetunesRequestVisibility) {
	b.Visibility = visibility
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) SetModelId(modelId *types.MusicModelId) {
	b.ModelId = modelId
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateMusicFinetuneV1MusicFinetunesPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateMusicFinetuneV1MusicFinetunesPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateMusicFinetuneV1MusicFinetunesPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateMusicFinetuneV1MusicFinetunesPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type CreateFinetunesRequestVisibility string

const (
	CreateFinetunesRequestVisibilityPrivate   CreateFinetunesRequestVisibility = "private"
	CreateFinetunesRequestVisibilityWorkspace CreateFinetunesRequestVisibility = "workspace"
)

func NewCreateFinetunesRequestVisibilityFromString(s string) (CreateFinetunesRequestVisibility, error) {
	switch s {
	case "private":
		return CreateFinetunesRequestVisibilityPrivate, nil
	case "workspace":
		return CreateFinetunesRequestVisibilityWorkspace, nil
	}
	var t CreateFinetunesRequestVisibility
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreateFinetunesRequestVisibility) Ptr() *CreateFinetunesRequestVisibility {
	return &e
}

type UpdateMusicFinetuneRequestModel struct {
	FinetuneId string `json:"-" url:"-"`
	// Updated name for the finetune.
	Name *string `json:"name,omitempty" xml:"name,omitempty" maxLength:"200" minLength:"5"`
	// Replacement set of tags.
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty" maxItems:"5"`
	// Updated primary musical genre.
	PrimaryGenre *string `json:"primary_genre,omitempty" xml:"primary_genre,omitempty"`
	// Finetune visibility. Only 'private' and 'workspace' can be set.
	Visibility *UpdateMusicFinetuneRequestModelVisibility `json:"visibility,omitempty" xml:"visibility,omitempty"`
}

func (u *UpdateMusicFinetuneRequestModel) GetFinetuneId() (value string) {
	if u == nil {
		return
	}
	return u.FinetuneId
}

func (u *UpdateMusicFinetuneRequestModel) GetName() (value *string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UpdateMusicFinetuneRequestModel) GetTags() (value []string) {
	if u == nil {
		return
	}
	return u.Tags
}

func (u *UpdateMusicFinetuneRequestModel) GetPrimaryGenre() (value *string) {
	if u == nil {
		return
	}
	return u.PrimaryGenre
}

func (u *UpdateMusicFinetuneRequestModel) GetVisibility() (value *UpdateMusicFinetuneRequestModelVisibility) {
	if u == nil {
		return
	}
	return u.Visibility
}

func (u *UpdateMusicFinetuneRequestModel) SetFinetuneId(finetuneId string) {
	u.FinetuneId = finetuneId
}

func (u *UpdateMusicFinetuneRequestModel) SetName(name *string) {
	u.Name = name
}

func (u *UpdateMusicFinetuneRequestModel) SetTags(tags []string) {
	u.Tags = tags
}

func (u *UpdateMusicFinetuneRequestModel) SetPrimaryGenre(primaryGenre *string) {
	u.PrimaryGenre = primaryGenre
}

func (u *UpdateMusicFinetuneRequestModel) SetVisibility(visibility *UpdateMusicFinetuneRequestModelVisibility) {
	u.Visibility = visibility
}

func (u *UpdateMusicFinetuneRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdateMusicFinetuneRequestModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateMusicFinetuneRequestModel to string"
	}
	return string(jsonData)
}

func (u *UpdateMusicFinetuneRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, u)
}

func (u *UpdateMusicFinetuneRequestModel) MarshalJSON() ([]byte, error) {
	type embed UpdateMusicFinetuneRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}

type UpdateMusicFinetuneRequestModelVisibility string

const (
	UpdateMusicFinetuneRequestModelVisibilityPrivate   UpdateMusicFinetuneRequestModelVisibility = "private"
	UpdateMusicFinetuneRequestModelVisibilityWorkspace UpdateMusicFinetuneRequestModelVisibility = "workspace"
)

func NewUpdateMusicFinetuneRequestModelVisibilityFromString(s string) (UpdateMusicFinetuneRequestModelVisibility, error) {
	switch s {
	case "private":
		return UpdateMusicFinetuneRequestModelVisibilityPrivate, nil
	case "workspace":
		return UpdateMusicFinetuneRequestModelVisibilityWorkspace, nil
	}
	var t UpdateMusicFinetuneRequestModelVisibility
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e UpdateMusicFinetuneRequestModelVisibility) Ptr() *UpdateMusicFinetuneRequestModelVisibility {
	return &e
}

type ListFinetunesRequestSort string

const (
	ListFinetunesRequestSortCreatedAt ListFinetunesRequestSort = "created_at"
	ListFinetunesRequestSortName      ListFinetunesRequestSort = "name"
)

func NewListFinetunesRequestSortFromString(s string) (ListFinetunesRequestSort, error) {
	switch s {
	case "created_at":
		return ListFinetunesRequestSortCreatedAt, nil
	case "name":
		return ListFinetunesRequestSortName, nil
	}
	var t ListFinetunesRequestSort
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListFinetunesRequestSort) Ptr() *ListFinetunesRequestSort {
	return &e
}

type ListFinetunesRequestSortDirection string

const (
	ListFinetunesRequestSortDirectionAsc  ListFinetunesRequestSortDirection = "asc"
	ListFinetunesRequestSortDirectionDesc ListFinetunesRequestSortDirection = "desc"
)

func NewListFinetunesRequestSortDirectionFromString(s string) (ListFinetunesRequestSortDirection, error) {
	switch s {
	case "asc":
		return ListFinetunesRequestSortDirectionAsc, nil
	case "desc":
		return ListFinetunesRequestSortDirectionDesc, nil
	}
	var t ListFinetunesRequestSortDirection
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListFinetunesRequestSortDirection) Ptr() *ListFinetunesRequestSortDirection {
	return &e
}

var (
	listFinetunesRequestFieldCursor        = big.NewInt(1 << 0)
	listFinetunesRequestFieldPageSize      = big.NewInt(1 << 1)
	listFinetunesRequestFieldVisibility    = big.NewInt(1 << 2)
	listFinetunesRequestFieldCreatedBy     = big.NewInt(1 << 3)
	listFinetunesRequestFieldSort          = big.NewInt(1 << 4)
	listFinetunesRequestFieldSortDirection = big.NewInt(1 << 5)
)

type ListFinetunesRequest struct {
	Cursor        *string                            `json:"-" url:"cursor,omitempty"`
	PageSize      *int                               `json:"-" url:"page_size,omitempty"`
	Visibility    *types.FinetuneVisibility          `json:"-" url:"visibility,omitempty"`
	CreatedBy     *types.FinetuneCreatedBy           `json:"-" url:"created_by,omitempty"`
	Sort          *ListFinetunesRequestSort          `json:"-" url:"sort,omitempty"`
	SortDirection *ListFinetunesRequestSortDirection `json:"-" url:"sort_direction,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListFinetunesRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFinetunesRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listFinetunesRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFinetunesRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listFinetunesRequestFieldPageSize)
}

// SetVisibility sets the Visibility field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFinetunesRequest) SetVisibility(visibility *types.FinetuneVisibility) {
	l.Visibility = visibility
	l.require(listFinetunesRequestFieldVisibility)
}

// SetCreatedBy sets the CreatedBy field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFinetunesRequest) SetCreatedBy(createdBy *types.FinetuneCreatedBy) {
	l.CreatedBy = createdBy
	l.require(listFinetunesRequestFieldCreatedBy)
}

// SetSort sets the Sort field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFinetunesRequest) SetSort(sort *ListFinetunesRequestSort) {
	l.Sort = sort
	l.require(listFinetunesRequestFieldSort)
}

// SetSortDirection sets the SortDirection field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListFinetunesRequest) SetSortDirection(sortDirection *ListFinetunesRequestSortDirection) {
	l.SortDirection = sortDirection
	l.require(listFinetunesRequestFieldSortDirection)
}
