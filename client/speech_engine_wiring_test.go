// Hand-written (listed in .fernignore). client/client.go is frozen so client.SpeechEngine can be
// the hand-written *speechengine.SpeechEngineClient. These tests check that wiring, and guard
// against the cost of freezing the file: a resource the generator adds later would otherwise be
// missing from the root client without any error.

package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/elevenlabs/elevenlabs-go/option"
	"github.com/elevenlabs/elevenlabs-go/speechengine"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// The shared cross-SDK golden vector (see speechengine/jwt_test.go).
const (
	wiringAPIKey = "sk_0123456789abcdef0123456789abcdef"
	wiringToken  = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJodHRwczovL2FwaS5lbGV2ZW5sYWJzLmlvL2NvbnZhaS9zcGVlY2gtZW5naW5lIiwic3ViIjoiY29udmFpX3NwZWVjaF9lbmdpbmVfdXBzdHJlYW0iLCJpYXQiOjEyMzQ1Njc4OTAsImV4cCI6NDEwMjQ0NDgwMH0.UObM4hx1W8NvGViU3aqPv1-IEXhmXgWDgoS9RZZWILk"
)

func TestSpeechEngineCreateReturnsAResource(t *testing.T) {
	// A complete response for engine "seng_1", taken from the generated wire test fixture.
	body, err := os.ReadFile("testdata/speech_engine_response.json")
	if err != nil {
		t.Fatal(err)
	}
	var gotKey, gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotKey, gotPath = r.Header.Get("xi-api-key"), r.Method+" "+r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}))
	defer server.Close()

	key, name := wiringAPIKey, "My engine"
	c := NewClient(option.WithBaseURL(server.URL), option.WithApiKey(&key))
	var _ *speechengine.SpeechEngineClient = c.SpeechEngine

	engine, err := c.SpeechEngine.Create(context.Background(), &types.CreateSpeechEngineRequest{
		Name:         &name,
		SpeechEngine: &types.SpeechEngineConfig{WsUrl: "wss://example.com/ws"},
	})
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	if gotPath != "POST /v1/speech-engine" || gotKey != wiringAPIKey {
		t.Fatalf("unexpected request: %s with key %q", gotPath, gotKey)
	}
	if engine.EngineID != "seng_1" || engine.Config == nil || engine.Config.Name != "My engine" {
		t.Fatalf("unexpected resource: %+v", engine)
	}

	// The resource carries the client's API key, so it verifies the API's inbound connections.
	req := httptest.NewRequest(http.MethodGet, "/ws", nil)
	req.Header.Set(speechengine.AuthorizationHeader, wiringToken)
	if err := engine.VerifyRequest(req); err != nil {
		t.Fatalf("resource must verify with the client's key: %v", err)
	}

	got, err := c.SpeechEngine.Get(context.Background(), "seng_1")
	if err != nil || got.EngineID != "seng_1" || gotPath != "GET /v1/speech-engine/seng_1" {
		t.Fatalf("get: %+v %v (%s)", got, err, gotPath)
	}

	// A per-call key is the one the resource verifies with.
	other := "sk_other"
	got, err = c.SpeechEngine.Get(context.Background(), "seng_1", option.WithApiKey(&other))
	if err != nil {
		t.Fatal(err)
	}
	if err := got.VerifyRequest(req); err == nil {
		t.Fatal("a resource fetched with another key must not verify the golden token")
	}

	// The generated method is still reachable through the embedded client.
	raw, err := c.SpeechEngine.Client.Get(context.Background(), "seng_1")
	if err != nil || raw.SpeechEngineId != "seng_1" {
		t.Fatalf("generated get: %+v %v", raw, err)
	}

	if err := c.SpeechEngine.Resource("seng_2").VerifyRequest(req); err != nil {
		t.Fatalf("Resource must carry the client's key: %v", err)
	}
}

// wrappedResources maps a generated resource package to the hand-written package whose wrapper
// the frozen client/client.go constructs instead of the generated client: client.Agents is an
// *agentbuilder.AgentsClient, which embeds the generated agents client.
var wrappedResources = map[string]string{
	"agents/client": "agentbuilder",
}

// TestRootClientWiresEveryResource fails when the generator adds a top-level resource that the
// frozen client/client.go does not expose. Fix it by merging the regenerated client.go in by
// hand, keeping the SpeechEngine field as *speechengine.SpeechEngineClient and the Agents field as
// *agentbuilder.AgentsClient.
func TestRootClientWiresEveryResource(t *testing.T) {
	source, err := os.ReadFile("client.go")
	if err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir("..")
	if err != nil {
		t.Fatal(err)
	}
	found := 0
	for _, entry := range entries {
		if !entry.IsDir() || entry.Name() == "client" {
			continue
		}
		for _, pkg := range []string{entry.Name(), entry.Name() + "/client"} {
			data, err := os.ReadFile(filepath.Join("..", pkg, "client.go"))
			if err != nil || !strings.Contains(string(data), "func NewClient(options *core.RequestOptions)") {
				continue
			}
			found++
			wired := pkg
			if wrapper, ok := wrappedResources[pkg]; ok {
				wired = wrapper
			}
			if !strings.Contains(string(source), `"github.com/elevenlabs/elevenlabs-go/`+wired+`"`) {
				t.Errorf("generated resource package %q is not wired into the frozen client/client.go", pkg)
			}
		}
	}
	if found == 0 {
		t.Fatal("found no generated resource packages; the scan is broken")
	}
}
