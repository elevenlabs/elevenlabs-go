package agentworkflows

import (
	"sort"
	"strings"
)

// IssueCode identifies a problem found in a workflow. The codes are shared by every ElevenLabs SDK
// (see the workflow spec).
type IssueCode string

const (
	// DuplicateNodeID: two different nodes use the same id.
	DuplicateNodeID IssueCode = "duplicate_node_id"
	// UnknownNode: an edge endpoint, Node(id) or Update refers to a node that doesn't exist.
	UnknownNode IssueCode = "unknown_node"
	// UnknownEdge: InsertBetween is given a pair of nodes with no edge between them.
	UnknownEdge IssueCode = "unknown_edge"
	// InvalidOption: an option doesn't apply to the node's kind, a schema override names a tool
	// the node doesn't run, an option value can't be converted to JSON, or (Go-specific) When or
	// a Back option is given the zero Condition.
	InvalidOption IssueCode = "invalid_option"
	// DuplicateEdge: a second edge connects the same two nodes and can't be merged into the first
	// as its backward condition.
	DuplicateEdge IssueCode = "duplicate_edge"
	// SelfLoop: an edge's source and target are the same node.
	SelfLoop IssueCode = "self_loop"
	// EdgeIntoStart: an edge targets the start node.
	EdgeIntoStart IssueCode = "edge_into_start"
	// EndHasOutgoingEdges: an end node has an outgoing edge.
	EndHasOutgoingEdges IssueCode = "end_has_outgoing_edges"
	// ToolEdgeNeedsResult: an edge leaving a tool node has a condition other than a result.
	ToolEdgeNeedsResult IssueCode = "tool_edge_needs_result"
	// ResultEdgeNeedsTool: a result condition is used on an edge that doesn't leave a tool node.
	ResultEdgeNeedsTool IssueCode = "result_edge_needs_tool"
	// CannotRemoveStart: Remove is called on the start node.
	CannotRemoveStart IssueCode = "cannot_remove_start"
	// CannotInsertBidirectional: InsertBetween targets an edge that also has a backward condition.
	CannotInsertBidirectional IssueCode = "cannot_insert_bidirectional"
	// MissingInsertResult: an InsertBetween callback doesn't return a node.
	MissingInsertResult IssueCode = "missing_insert_result"
	// UnreachableNode (warning): a node can't be reached from the start node.
	UnreachableNode IssueCode = "unreachable_node"
	// UnknownNodeType (warning): a node's type isn't in this SDK version's generated types, so
	// agents.Create / agents.Update can't send Build's output. Send ToJSON's output directly instead.
	UnknownNodeType IssueCode = "unknown_node_type"
	// InvalidModel is Go-specific: the validated workflow still doesn't decode into the generated
	// types.AgentWorkflowRequestModel, for example because a loaded field has the wrong shape.
	InvalidModel IssueCode = "invalid_model"
)

// Severity says whether an Issue stops Build (an error) or is only reported (a warning).
type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
)

// Issue is a single problem found in a workflow.
type Issue struct {
	Code     IssueCode
	Severity Severity
	Message  string
	// NodeID is the node the problem is about, or "" when it isn't about one node.
	NodeID string
}

// ValidationError is returned by Build, ToJSON and ToJSONString when a workflow has errors. It
// lists every error found, not just the first.
type ValidationError struct {
	Issues []Issue
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
	b.WriteString("invalid workflow:")
	for _, issue := range e.Issues {
		b.WriteString("\n  - ")
		b.WriteString(issue.Message)
	}
	return b.String()
}
