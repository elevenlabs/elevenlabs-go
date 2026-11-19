package agenttools

import (
	"sort"
	"strings"

	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
)

// IssueCode identifies a problem found in a tool. The codes are shared by every ElevenLabs SDK (see
// the tool spec).
type IssueCode string

const (
	// MissingURL: a webhook tool has no method and URL, or an empty URL.
	MissingURL IssueCode = "missing_url"
	// MissingPathParam: the URL has a {name} placeholder with no path param.
	MissingPathParam IssueCode = "missing_path_param"
	// UnknownPathParam: a path param has no {name} placeholder in the URL, or its placeholder is
	// filled by the auth connection.
	UnknownPathParam IssueCode = "unknown_path_param"
	// ParamNotLiteral: a path or query param is an object or array.
	ParamNotLiteral IssueCode = "param_not_literal"
	// BodyNotObject: a body, response, query or client params schema isn't an object.
	BodyNotObject IssueCode = "body_not_object"
	// BodyNotAllowed: a body is set on a GET.
	BodyNotAllowed IssueCode = "body_not_allowed"
	// MissingBody: a POST, PUT or PATCH has no body.
	MissingBody IssueCode = "missing_body"
	// DuplicateParam: a name is used in more than one of the path params, query params and body
	// properties.
	DuplicateParam IssueCode = "duplicate_param"
	// ConflictingValueSource: a schema sets more than one of Describe, FromVariable, Constant,
	// SystemProvided and Omit.
	ConflictingValueSource IssueCode = "conflicting_value_source"
	// OmittedRequired: Omit is used on a schema that isn't optional, or on a path param.
	OmittedRequired IssueCode = "omitted_required"
	// InvalidChoices: Choices is used on a schema that isn't a string.
	InvalidChoices IssueCode = "invalid_choices"
	// AllowedValuesNeedsDescription: AllowedFrom is used without Describe.
	AllowedValuesNeedsDescription IssueCode = "allowed_values_needs_description"
	// ConstantTypeMismatch: a constant doesn't match the schema type, is "" on a literal (which the
	// API reads as unset), or can't be sent exactly (see Schema).
	ConstantTypeMismatch IssueCode = "constant_type_mismatch"
	// UnknownRequiredProperty: a name in RequireAnyOf / RequireAllOf isn't a property of the object.
	UnknownRequiredProperty IssueCode = "unknown_required_property"
	// MissingArrayItems: an array schema has no items (only possible when loaded).
	MissingArrayItems IssueCode = "missing_array_items"
	// DuplicateHeader: two header names differ only in case.
	DuplicateHeader IssueCode = "duplicate_header"
	// InvalidDestination: a transfer destination of the wrong kind, or a SIP URI without
	// transfer_type "sip_refer".
	InvalidDestination IssueCode = "invalid_destination"
	// NoTransfers: a transfer system tool has no Transfer entries.
	NoTransfers IssueCode = "no_transfers"
	// InvalidOption: an option doesn't apply to the tool's kind, is outside its documented range, or
	// can't be converted to JSON.
	InvalidOption IssueCode = "invalid_option"
	// MissingValueSource: a literal has none of Describe, FromVariable, Constant, SystemProvided and
	// Omit.
	MissingValueSource IssueCode = "missing_value_source"
	// UnknownToolType (warning): a loaded tool's type has no constructor in this SDK version. It
	// round-trips through ToJSON but can't be edited.
	UnknownToolType IssueCode = "unknown_tool_type"
	// InvalidModel is reported by Build and BuildConfig only: the tool doesn't decode into the
	// generated types.ToolRequestModelToolConfig, for example because a loaded field has the wrong
	// shape. ToJSON still writes such fields back as loaded.
	InvalidModel IssueCode = "invalid_model"
)

// Severity says whether an Issue stops Build (an error) or is only reported (a warning). It is the
// workflow builder's type.
type Severity = agentworkflows.Severity

const (
	SeverityError   = agentworkflows.SeverityError
	SeverityWarning = agentworkflows.SeverityWarning
)

// Issue is a single problem found in a tool.
type Issue struct {
	Code     IssueCode
	Severity Severity
	Message  string
	// Path is where the problem is, such as "body.reason" or "path_params.order_id", or "" when it
	// is about the whole tool.
	Path string
}

// ValidationError is returned by Build, BuildConfig, ToJSON and ToJSONString when a tool has
// errors. It lists every error found, not just the first.
type ValidationError struct {
	// ToolName is the tool's name.
	ToolName string
	Issues   []Issue
}

// Codes returns the issue codes as a sorted list.
func (e *ValidationError) Codes() []string {
	codes := make([]string, len(e.Issues))
	for i, issue := range e.Issues {
		codes[i] = string(issue.Code)
	}
	sort.Strings(codes)
	return codes
}

func (e *ValidationError) Error() string {
	var b strings.Builder
	b.WriteString("invalid tool \"")
	b.WriteString(e.ToolName)
	b.WriteString("\":")
	for _, issue := range e.Issues {
		b.WriteString("\n  - ")
		b.WriteString(issue.Message)
	}
	return b.String()
}

func errorIssue(code IssueCode, path, message string) Issue {
	return Issue{Code: code, Severity: SeverityError, Message: message, Path: path}
}
