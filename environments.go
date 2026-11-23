package elevenlabs

// Environment defines the environment with multiple base URLs.
type Environment struct {
	Base                string
	Production          string
	ProductionUs        string
	ProductionEu        string
	ProductionIndia     string
	ProductionSingapore string
}

// Environments defines all of the API environments.
// These values can be used with the WithEnvironment
// RequestOption to override the client's default environment,
// if any.
var Environments = struct {
	Production          Environment
	ProductionUs        Environment
	ProductionEu        Environment
	ProductionIndia     Environment
	ProductionSingapore Environment
}{
	Production: Environment{
		Base:                "https://api.elevenlabs.io",
		Production:          "wss://api.elevenlabs.io",
		ProductionUs:        "wss://api.us.elevenlabs.io",
		ProductionEu:        "wss://api.eu.residency.elevenlabs.io",
		ProductionIndia:     "wss://api.in.residency.elevenlabs.io",
		ProductionSingapore: "wss://api.sg.residency.elevenlabs.io",
	},
	ProductionUs: Environment{
		Base:                "https://api.us.elevenlabs.io",
		Production:          "wss://api.elevenlabs.io",
		ProductionUs:        "wss://api.us.elevenlabs.io",
		ProductionEu:        "wss://api.eu.residency.elevenlabs.io",
		ProductionIndia:     "wss://api.in.residency.elevenlabs.io",
		ProductionSingapore: "wss://api.sg.residency.elevenlabs.io",
	},
	ProductionEu: Environment{
		Base:                "https://api.eu.residency.elevenlabs.io",
		Production:          "wss://api.elevenlabs.io",
		ProductionUs:        "wss://api.us.elevenlabs.io",
		ProductionEu:        "wss://api.eu.residency.elevenlabs.io",
		ProductionIndia:     "wss://api.in.residency.elevenlabs.io",
		ProductionSingapore: "wss://api.sg.residency.elevenlabs.io",
	},
	ProductionIndia: Environment{
		Base:                "https://api.in.residency.elevenlabs.io",
		Production:          "wss://api.elevenlabs.io",
		ProductionUs:        "wss://api.us.elevenlabs.io",
		ProductionEu:        "wss://api.eu.residency.elevenlabs.io",
		ProductionIndia:     "wss://api.in.residency.elevenlabs.io",
		ProductionSingapore: "wss://api.sg.residency.elevenlabs.io",
	},
	ProductionSingapore: Environment{
		Base:                "https://api.sg.residency.elevenlabs.io",
		Production:          "wss://api.elevenlabs.io",
		ProductionUs:        "wss://api.us.elevenlabs.io",
		ProductionEu:        "wss://api.eu.residency.elevenlabs.io",
		ProductionIndia:     "wss://api.in.residency.elevenlabs.io",
		ProductionSingapore: "wss://api.sg.residency.elevenlabs.io",
	},
}
