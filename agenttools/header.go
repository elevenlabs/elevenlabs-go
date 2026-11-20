package agenttools

import (
	"github.com/elevenlabs/elevenlabs-go/agentworkflows"
	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

// Header is the value of a webhook request header. Create one with HeaderValue, HeaderSecret,
// HeaderVariable or HeaderEnvVar. The zero value is an empty fixed value.
type Header struct {
	wire any
}

// HeaderValue is a fixed value.
func HeaderValue(text string) Header { return Header{wire: text} }

// HeaderSecret is a secret from the workspace's secret store, so the value never appears in the
// config.
func HeaderSecret(secretID string) Header {
	return Header{wire: map[string]any{"secret_id": secretID}}
}

// HeaderVariable is the value of a dynamic variable when the tool runs.
func HeaderVariable(name string) Header {
	return Header{wire: map[string]any{"variable_name": name}}
}

// HeaderEnvVar is a workspace environment variable, by label.
func HeaderEnvVar(label string) Header {
	return Header{wire: map[string]any{"env_var_label": label}}
}

// Wire returns the header value in API wire format (a fresh copy).
func (h Header) Wire() any {
	if h.wire == nil {
		return ""
	}
	return wirejson.Clone(h.wire)
}

// AuthConnection is the auth connection a webhook tool authenticates with. Create one with
// AuthConnectionID or AuthConnectionEnvVar.
type AuthConnection struct {
	wire map[string]any
}

// AuthConnectionID is an auth connection from the workspace's store, by id.
func AuthConnectionID(authConnectionID string) AuthConnection {
	return AuthConnection{wire: map[string]any{"auth_connection_id": authConnectionID}}
}

// AuthConnectionEnvVar is an environment variable of type auth connection, by label, resolved for
// the current environment.
func AuthConnectionEnvVar(label string) AuthConnection {
	return AuthConnection{wire: map[string]any{"env_var_label": label}}
}

// Wire returns the auth connection in API wire format (a fresh copy).
func (a AuthConnection) Wire() map[string]any {
	return cloneMap(a.wire)
}

// Destination is where a transfer system tool hands the conversation off to. It is the workflow
// builder's type, shared by both builders; create one with DestinationPhone,
// DestinationPhoneVariable, DestinationSIP, DestinationSIPVariable or DestinationAgent (here or in
// agentworkflows).
type Destination = agentworkflows.Destination

// AgentDestinationOptions are the optional settings of DestinationAgent.
type AgentDestinationOptions = agentworkflows.AgentDestinationOptions

// DestinationPhone is a phone number in E.164 format.
func DestinationPhone(phoneNumber string) Destination {
	return agentworkflows.DestinationPhone(phoneNumber)
}

// DestinationPhoneVariable is a phone number read from a dynamic variable at transfer time.
func DestinationPhoneVariable(variableName string) Destination {
	return agentworkflows.DestinationPhoneVariable(variableName)
}

// DestinationSIP is a SIP URI. A transfer to one needs TransferType "sip_refer".
func DestinationSIP(sipURI string) Destination {
	return agentworkflows.DestinationSIP(sipURI)
}

// DestinationSIPVariable is a SIP URI read from a dynamic variable at transfer time. A transfer to
// one needs TransferType "sip_refer".
func DestinationSIPVariable(variableName string) Destination {
	return agentworkflows.DestinationSIPVariable(variableName)
}

// DestinationAgent is another ElevenLabs agent. opts may be nil.
func DestinationAgent(agentID string, opts *AgentDestinationOptions) Destination {
	return agentworkflows.DestinationAgent(agentID, opts)
}
