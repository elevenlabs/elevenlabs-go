package agentworkflows

import "github.com/elevenlabs/elevenlabs-go/internal/wirejson"

// DestinationKind says what a Destination points at.
type DestinationKind string

const (
	DestinationKindPhone         DestinationKind = "phone"
	DestinationKindPhoneVariable DestinationKind = "phone_dynamic_variable"
	DestinationKindSIP           DestinationKind = "sip_uri"
	DestinationKindSIPVariable   DestinationKind = "sip_uri_dynamic_variable"
	DestinationKindAgent         DestinationKind = "agent"
)

const (
	destinationNodeTypePhoneNumber = "phone_number"
	destinationNodeTypeAgent       = "standalone_agent"
)

// Destination is where a transfer hands the conversation off to. Create one with
// DestinationPhone, DestinationPhoneVariable, DestinationSIP, DestinationSIPVariable or
// DestinationAgent. The zero value is not a valid destination.
//
// The type is shared with the agent tools builder, whose transfer system tools take the same
// destinations; its accessors give that builder the parts it needs.
type Destination struct {
	kind    DestinationKind
	value   string // phone number, SIP URI, variable name or agent id
	nodeID  string
	hasNode bool
}

// AgentDestinationOptions are the optional settings of DestinationAgent.
type AgentDestinationOptions struct {
	// NodeID is the node in the other agent's workflow to start from. Empty means its default
	// entry point.
	NodeID string
}

// DestinationPhone is a phone number in E.164 format.
func DestinationPhone(phoneNumber string) Destination {
	return Destination{kind: DestinationKindPhone, value: phoneNumber}
}

// DestinationPhoneVariable is a phone number read from a dynamic variable at transfer time.
func DestinationPhoneVariable(variableName string) Destination {
	return Destination{kind: DestinationKindPhoneVariable, value: variableName}
}

// DestinationSIP is a SIP URI.
func DestinationSIP(sipURI string) Destination {
	return Destination{kind: DestinationKindSIP, value: sipURI}
}

// DestinationSIPVariable is a SIP URI read from a dynamic variable at transfer time.
func DestinationSIPVariable(variableName string) Destination {
	return Destination{kind: DestinationKindSIPVariable, value: variableName}
}

// DestinationAgent is another ElevenLabs agent.
func DestinationAgent(agentID string, opts *AgentDestinationOptions) Destination {
	d := Destination{kind: DestinationKindAgent, value: agentID}
	if opts != nil && opts.NodeID != "" {
		d.nodeID = opts.NodeID
		d.hasNode = true
	}
	return d
}

// Kind says what the destination points at.
func (d Destination) Kind() DestinationKind { return d.kind }

// IsAgent reports whether the destination is another agent (as opposed to a phone number or SIP
// URI).
func (d Destination) IsAgent() bool { return d.kind == DestinationKindAgent }

// AgentID is the agent of an agent destination, or "".
func (d Destination) AgentID() string {
	if d.kind != DestinationKindAgent {
		return ""
	}
	return d.value
}

// NodeID is the entry node of an agent destination, or "" for its default entry point.
func (d Destination) NodeID() string { return d.nodeID }

// NodeType is the workflow node type a transfer to this destination creates: "phone_number" for
// phone and SIP destinations, "standalone_agent" for agent destinations.
func (d Destination) NodeType() string {
	if d.kind == DestinationKindAgent {
		return destinationNodeTypeAgent
	}
	return destinationNodeTypePhoneNumber
}

// TransferDestination is the API's transfer_destination object for a phone or SIP destination
// (a fresh copy), or nil for an agent destination.
func (d Destination) TransferDestination() map[string]any {
	switch d.kind {
	case DestinationKindPhone, DestinationKindPhoneVariable:
		return map[string]any{"type": string(d.kind), "phone_number": d.value}
	case DestinationKindSIP, DestinationKindSIPVariable:
		return map[string]any{"type": string(d.kind), "sip_uri": d.value}
	}
	return nil
}

// Wire returns the node fields this destination sets, in API wire format (a fresh copy):
// {"type": "phone_number", "transfer_destination": {...}} or
// {"type": "standalone_agent", "agent_id": ..., "node_id"?: ...}.
func (d Destination) Wire() map[string]any {
	if d.kind == DestinationKindAgent {
		wire := map[string]any{"type": destinationNodeTypeAgent, "agent_id": d.value}
		if d.hasNode {
			wire["node_id"] = d.nodeID
		}
		return wire
	}
	return map[string]any{"type": destinationNodeTypePhoneNumber, "transfer_destination": d.TransferDestination()}
}

func cloneMap(m map[string]any) map[string]any {
	if m == nil {
		return nil
	}
	return wirejson.Clone(m).(map[string]any)
}
