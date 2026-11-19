package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Audit log entry with Firestore document ID for API responses.
type WorkspaceAuditLogEntryResponse struct {
	// Event metadata
	Metadata map[string]any `json:"metadata,omitempty"`
	// Event time in milliseconds since epoch
	Time *int `json:"time,omitempty" xml:"time,omitempty"`
	// Activity ID
	ActivityId *WorkspaceAuditLogEntryResponseActivityId `json:"activity_id" xml:"activity_id" required:"true"`
	// Activity name
	ActivityName string `json:"activity_name" xml:"activity_name" required:"true"`
	// Event category
	CategoryName *string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// Category UID for IAM
	CategoryUid *int `json:"category_uid,omitempty" xml:"category_uid,omitempty"`
	// Event class name
	ClassName *string `json:"class_name,omitempty" xml:"class_name,omitempty"`
	// Event class UID
	ClassUid *int `json:"class_uid,omitempty" xml:"class_uid,omitempty"`
	// OCSF Severity levels.
	//
	// Spec: https://schema.ocsf.io/1.6.0/objects/severity_id
	SeverityId *int `json:"severity_id,omitempty" xml:"severity_id,omitempty"`
	// OCSF Status levels.
	//
	// Spec: https://schema.ocsf.io/1.6.0/objects/status_id
	StatusId int `json:"status_id" xml:"status_id" required:"true"`
	// OCSF Actor object - describes the entity that performed the action.
	//
	// Spec: https://schema.ocsf.io/1.6.0/objects/actor
	Actor *ActorModel `json:"actor" xml:"actor" required:"true"`
	// Device information
	Device *DeviceModel `json:"device,omitempty" xml:"device,omitempty"`
	// HTTP request details
	HttpRequest *HttpRequestModel `json:"http_request,omitempty" xml:"http_request,omitempty"`
	// Human-readable event description
	Message string `json:"message" xml:"message" required:"true"`
	// Attributes not mapped to OCSF
	Unmapped map[string]any `json:"unmapped,omitempty"`
	// Firestore document ID
	Id string `json:"id" xml:"id" required:"true"`
	// Event time in human-readable RFC 3339 format, derived from 'time'.
	TimeDt *string `json:"time_dt,omitempty" xml:"time_dt,omitempty"`
	// OCSF type_uid is class_uid * 100 + activity_id.
	TypeUid *int `json:"type_uid,omitempty" xml:"type_uid,omitempty"`
	// OCSF type_name combines class_name and activity_name.
	TypeName        *string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkspaceAuditLogEntryResponse) GetMetadata() (value map[string]any) {
	if w == nil {
		return
	}
	return w.Metadata
}

func (w *WorkspaceAuditLogEntryResponse) GetTime() (value *int) {
	if w == nil {
		return
	}
	return w.Time
}

func (w *WorkspaceAuditLogEntryResponse) GetActivityId() (value *WorkspaceAuditLogEntryResponseActivityId) {
	if w == nil {
		return
	}
	return w.ActivityId
}

func (w *WorkspaceAuditLogEntryResponse) GetActivityName() (value string) {
	if w == nil {
		return
	}
	return w.ActivityName
}

func (w *WorkspaceAuditLogEntryResponse) GetCategoryName() (value *string) {
	if w == nil {
		return
	}
	return w.CategoryName
}

func (w *WorkspaceAuditLogEntryResponse) GetCategoryUid() (value *int) {
	if w == nil {
		return
	}
	return w.CategoryUid
}

func (w *WorkspaceAuditLogEntryResponse) GetClassName() (value *string) {
	if w == nil {
		return
	}
	return w.ClassName
}

func (w *WorkspaceAuditLogEntryResponse) GetClassUid() (value *int) {
	if w == nil {
		return
	}
	return w.ClassUid
}

func (w *WorkspaceAuditLogEntryResponse) GetSeverityId() (value *int) {
	if w == nil {
		return
	}
	return w.SeverityId
}

func (w *WorkspaceAuditLogEntryResponse) GetStatusId() (value int) {
	if w == nil {
		return
	}
	return w.StatusId
}

func (w *WorkspaceAuditLogEntryResponse) GetActor() (value *ActorModel) {
	if w == nil {
		return
	}
	return w.Actor
}

func (w *WorkspaceAuditLogEntryResponse) GetDevice() (value *DeviceModel) {
	if w == nil {
		return
	}
	return w.Device
}

func (w *WorkspaceAuditLogEntryResponse) GetHttpRequest() (value *HttpRequestModel) {
	if w == nil {
		return
	}
	return w.HttpRequest
}

func (w *WorkspaceAuditLogEntryResponse) GetMessage() (value string) {
	if w == nil {
		return
	}
	return w.Message
}

func (w *WorkspaceAuditLogEntryResponse) GetUnmapped() (value map[string]any) {
	if w == nil {
		return
	}
	return w.Unmapped
}

func (w *WorkspaceAuditLogEntryResponse) GetId() (value string) {
	if w == nil {
		return
	}
	return w.Id
}

func (w *WorkspaceAuditLogEntryResponse) GetTimeDt() (value *string) {
	if w == nil {
		return
	}
	return w.TimeDt
}

func (w *WorkspaceAuditLogEntryResponse) GetTypeUid() (value *int) {
	if w == nil {
		return
	}
	return w.TypeUid
}

func (w *WorkspaceAuditLogEntryResponse) GetTypeName() (value *string) {
	if w == nil {
		return
	}
	return w.TypeName
}

func (w *WorkspaceAuditLogEntryResponse) SetMetadata(metadata map[string]any) {
	w.Metadata = metadata
}

func (w *WorkspaceAuditLogEntryResponse) SetTime(time *int) {
	w.Time = time
}

func (w *WorkspaceAuditLogEntryResponse) SetActivityId(activityId *WorkspaceAuditLogEntryResponseActivityId) {
	w.ActivityId = activityId
}

func (w *WorkspaceAuditLogEntryResponse) SetActivityName(activityName string) {
	w.ActivityName = activityName
}

func (w *WorkspaceAuditLogEntryResponse) SetCategoryName(categoryName *string) {
	w.CategoryName = categoryName
}

func (w *WorkspaceAuditLogEntryResponse) SetCategoryUid(categoryUid *int) {
	w.CategoryUid = categoryUid
}

func (w *WorkspaceAuditLogEntryResponse) SetClassName(className *string) {
	w.ClassName = className
}

func (w *WorkspaceAuditLogEntryResponse) SetClassUid(classUid *int) {
	w.ClassUid = classUid
}

func (w *WorkspaceAuditLogEntryResponse) SetSeverityId(severityId *int) {
	w.SeverityId = severityId
}

func (w *WorkspaceAuditLogEntryResponse) SetStatusId(statusId int) {
	w.StatusId = statusId
}

func (w *WorkspaceAuditLogEntryResponse) SetActor(actor *ActorModel) {
	w.Actor = actor
}

func (w *WorkspaceAuditLogEntryResponse) SetDevice(device *DeviceModel) {
	w.Device = device
}

func (w *WorkspaceAuditLogEntryResponse) SetHttpRequest(httpRequest *HttpRequestModel) {
	w.HttpRequest = httpRequest
}

func (w *WorkspaceAuditLogEntryResponse) SetMessage(message string) {
	w.Message = message
}

func (w *WorkspaceAuditLogEntryResponse) SetUnmapped(unmapped map[string]any) {
	w.Unmapped = unmapped
}

func (w *WorkspaceAuditLogEntryResponse) SetId(id string) {
	w.Id = id
}

func (w *WorkspaceAuditLogEntryResponse) SetTimeDt(timeDt *string) {
	w.TimeDt = timeDt
}

func (w *WorkspaceAuditLogEntryResponse) SetTypeUid(typeUid *int) {
	w.TypeUid = typeUid
}

func (w *WorkspaceAuditLogEntryResponse) SetTypeName(typeName *string) {
	w.TypeName = typeName
}

func (w *WorkspaceAuditLogEntryResponse) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceAuditLogEntryResponse) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceAuditLogEntryResponse to string"
	}
	return string(jsonData)
}

func (w *WorkspaceAuditLogEntryResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceAuditLogEntryResponse) MarshalJSON() ([]byte, error) {
	type embed WorkspaceAuditLogEntryResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
