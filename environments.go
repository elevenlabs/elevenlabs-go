package elevenlabs

// Environments defines all of the API environments.
// These values can be used with the WithBaseURL
// RequestOption to override the client's default environment,
// if any.
var Environments = struct {
	Production          string
	ProductionUs        string
	ProductionEu        string
	ProductionIndia     string
	ProductionSingapore string
}{
	Production:          "https://api.elevenlabs.io",
	ProductionUs:        "https://api.us.elevenlabs.io",
	ProductionEu:        "https://api.eu.residency.elevenlabs.io",
	ProductionIndia:     "https://api.in.residency.elevenlabs.io",
	ProductionSingapore: "https://api.sg.residency.elevenlabs.io",
}
