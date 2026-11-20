package agentbuilder

import (
	"fmt"
	"sort"
	"strings"

	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
)

// IssueCode identifies a problem found in an agent. The codes are shared by every ElevenLabs SDK
// (see the agent spec).
type IssueCode string

const (
	// MissingName: the agent's name is empty.
	MissingName IssueCode = "missing_name"
	// MissingPrompt (warning): the assembled config has no prompt.
	MissingPrompt IssueCode = "missing_prompt"
	// DuplicateToolName: two different tools share a name. The same tool used twice is fine.
	DuplicateToolName IssueCode = "duplicate_tool_name"
	// InvalidTool: a tool's own validation fails. Path is "tools.NAME".
	InvalidTool IssueCode = "invalid_tool"
	// InvalidWorkflow: the workflow's own validation fails.
	InvalidWorkflow IssueCode = "invalid_workflow"
	// InvalidConfig: the raw ConversationConfig or PlatformSettings can't be converted to JSON.
	InvalidConfig IssueCode = "invalid_config"
	// OptionConflict: the raw config sets a field that a typed option, Tools or a system tool
	// also sets.
	OptionConflict IssueCode = "option_conflict"
	// UnresolvedTool: ToJSON or Build has no id for a tool the agent manages. Never reported by
	// Validate.
	UnresolvedTool IssueCode = "unresolved_tool"
	// SystemToolInWorkflow: a workflow tool node or subagent uses a system tool.
	SystemToolInWorkflow IssueCode = "system_tool_in_workflow"
	// InvalidModel is Go-specific: the validated config still doesn't decode into the generated
	// types.BodyCreateAgentV1ConvaiAgentsCreatePost, for example because a loaded field has the
	// wrong shape. ToJSON, which Push sends, still writes such fields back as loaded.
	InvalidModel IssueCode = "invalid_model"
)

// Severity says whether an Issue stops Validate, ToJSON and Build (an error) or is only reported
// (a warning). It is the workflow builder's type.
type Severity = agentworkflows.Severity

const (
	SeverityError   = agentworkflows.SeverityError
	SeverityWarning = agentworkflows.SeverityWarning
)

// Issue is a single problem found in an agent.
type Issue struct {
	Code     IssueCode
	Severity Severity
	Message  string
	// Path is where the problem is, such as "tools.get_order" or "workflow", or "" when it is about
	// the whole agent.
	Path string
}

// ValidationError is returned by Validate, ToJSON, Build and Push when an agent has errors. It
// lists every error found, not just the first.
type ValidationError struct {
	// AgentName is the agent's name.
	AgentName string
	Issues    []Issue
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
	fmt.Fprintf(&b, "invalid agent %q:", e.AgentName)
	for _, issue := range e.Issues {
		b.WriteString("\n  - ")
		b.WriteString(issue.Message)
	}
	return b.String()
}

// AgentAction is what Push did to the agent.
type AgentAction string

const (
	AgentCreated AgentAction = "created"
	AgentUpdated AgentAction = "updated"
)

// ToolAction is what Push did to one tool.
type ToolAction string

const (
	ToolCreated   ToolAction = "created"
	ToolUpdated   ToolAction = "updated"
	ToolUnchanged ToolAction = "unchanged"
	// ToolDetached: the agent used the tool before and no longer does. It isn't deleted.
	ToolDetached ToolAction = "detached"
)

// ToolResult is what Push did to one tool, and the tool's id.
type ToolResult struct {
	ID     string
	Action ToolAction
}

// PushResult is what Push did.
type PushResult struct {
	// AgentID is the agent's id. Empty in a *PushError's partial result if the agent wasn't
	// written.
	AgentID string
	// Action is what happened to the agent. Empty in a *PushError's partial result if the agent
	// wasn't written.
	Action AgentAction
	// Tools is each tool Push touched, by name.
	Tools map[string]ToolResult
}

// PushError is returned by Push when an API call fails after it started writing. Result lists
// what was written before the failure, so nothing created goes unreported; with a project, those
// ids are saved to it too. Unwrap returns the failed call's error.
type PushError struct {
	// Message names the step that failed, such as `Couldn't create tool "get_order"`.
	Message string
	Result  PushResult
	Cause   error
}

func (e *PushError) Error() string { return e.Message + ": " + e.Cause.Error() }

func (e *PushError) Unwrap() error { return e.Cause }

// ProjectError is returned when a project's files can't be read or written: a missing or
// malformed file, a path outside the project, or a format the SDK can't read yet.
type ProjectError struct {
	Message string
}

func (e *ProjectError) Error() string { return e.Message }

func projectErrorf(format string, args ...any) error {
	return &ProjectError{Message: fmt.Sprintf(format, args...)}
}
