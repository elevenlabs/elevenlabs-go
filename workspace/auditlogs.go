package workspace

import (
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

var (
	listAuditLogsRequestFieldLimit          = big.NewInt(1 << 0)
	listAuditLogsRequestFieldCursor         = big.NewInt(1 << 1)
	listAuditLogsRequestFieldTimeFromUnixMs = big.NewInt(1 << 2)
	listAuditLogsRequestFieldTimeToUnixMs   = big.NewInt(1 << 3)
	listAuditLogsRequestFieldActorUid       = big.NewInt(1 << 4)
	listAuditLogsRequestFieldClassName      = big.NewInt(1 << 5)
	listAuditLogsRequestFieldActivityName   = big.NewInt(1 << 6)
)

type ListAuditLogsRequest struct {
	Limit          *int                     `json:"-" url:"limit,omitempty"`
	Cursor         *string                  `json:"-" url:"cursor,omitempty"`
	TimeFromUnixMs *int                     `json:"-" url:"time_from_unix_ms,omitempty"`
	TimeToUnixMs   *int                     `json:"-" url:"time_to_unix_ms,omitempty"`
	ActorUid       *string                  `json:"-" url:"actor_uid,omitempty"`
	ClassName      *types.AuditClassName    `json:"-" url:"class_name,omitempty"`
	ActivityName   *types.AuditActivityName `json:"-" url:"activity_name,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListAuditLogsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetLimit sets the Limit field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuditLogsRequest) SetLimit(limit *int) {
	l.Limit = limit
	l.require(listAuditLogsRequestFieldLimit)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuditLogsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listAuditLogsRequestFieldCursor)
}

// SetTimeFromUnixMs sets the TimeFromUnixMs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuditLogsRequest) SetTimeFromUnixMs(timeFromUnixMs *int) {
	l.TimeFromUnixMs = timeFromUnixMs
	l.require(listAuditLogsRequestFieldTimeFromUnixMs)
}

// SetTimeToUnixMs sets the TimeToUnixMs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuditLogsRequest) SetTimeToUnixMs(timeToUnixMs *int) {
	l.TimeToUnixMs = timeToUnixMs
	l.require(listAuditLogsRequestFieldTimeToUnixMs)
}

// SetActorUid sets the ActorUid field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuditLogsRequest) SetActorUid(actorUid *string) {
	l.ActorUid = actorUid
	l.require(listAuditLogsRequestFieldActorUid)
}

// SetClassName sets the ClassName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuditLogsRequest) SetClassName(className *types.AuditClassName) {
	l.ClassName = className
	l.require(listAuditLogsRequestFieldClassName)
}

// SetActivityName sets the ActivityName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAuditLogsRequest) SetActivityName(activityName *types.AuditActivityName) {
	l.ActivityName = activityName
	l.require(listAuditLogsRequestFieldActivityName)
}
