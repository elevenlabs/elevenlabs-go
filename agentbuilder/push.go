package agentbuilder

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/elevenlabs/elevenlabs-go/agenttools"
	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
	"github.com/elevenlabs/elevenlabs-go/option"
)

// PushOptions are the optional settings of Push.
type PushOptions struct {
	// AgentID is the agent to update. Without it (and without a project that knows the agent),
	// Push creates a new agent, so keep the returned id. Not allowed with Project, which records
	// the id itself.
	AgentID string
	// Project is an agents-as-code project in the CLI's layout. Its agents.json and tools.json say
	// which agent and tools on the server this definition is; Push records new ids there and writes
	// the configs it sent. A project opened with OpenProject is saved, also when a request fails
	// partway.
	Project *Project
	// BranchID updates this branch of the agent instead of its main branch.
	BranchID string
	// VersionDescription describes the new version.
	VersionDescription string
}

// PullOptions are the optional settings of Pull.
type PullOptions struct {
	// Project receives the agent and its tools, with their ids. A project opened with OpenProject
	// is saved.
	Project *Project
	// BranchID pulls this branch instead of the main one.
	BranchID string
}

// remoteTool is a tool as the API returned it.
type remoteTool struct {
	id   string
	name string
	// config is the tool_config as JSON text, and tree the same parsed.
	config json.RawMessage
	tree   any
}

func parseRemoteTool(data json.RawMessage) (remoteTool, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return remoteTool{}, fmt.Errorf("agentbuilder: unexpected tool response: %w", err)
	}
	var remote remoteTool
	if err := json.Unmarshal(fields["id"], &remote.id); err != nil || remote.id == "" {
		return remoteTool{}, errors.New("agentbuilder: the tool response has no id")
	}
	remote.config = fields["tool_config"]
	if len(remote.config) > 0 {
		tree, err := wirejson.Decode(remote.config)
		if err != nil {
			return remoteTool{}, err
		}
		remote.tree = tree
	}
	remote.name = remote.id
	if config, ok := remote.tree.(map[string]any); ok {
		if name, ok := config["name"].(string); ok {
			remote.name = name
		}
	}
	return remote, nil
}

// Push creates or updates an agent and every webhook and client tool it uses, in one call.
//
//  1. It validates the agent, so an invalid one writes nothing (a *ValidationError).
//  2. It creates each tool that's new and updates each one that changed. Which tools already
//     exist comes from the project's tools.json, or, without a project, from the tools the agent
//     (opts.AgentID) already uses, by name: on its main branch, or on opts.BranchID. Tools only
//     other agents use are never touched.
//  3. It creates the agent, or updates it, with the tools' ids filled in.
//
// Settings the SDK doesn't model, kept from a loaded config, are sent too. Tools the agent stops
// using are reported as detached, not deleted. When a request fails after Push started writing,
// it returns a *PushError whose Result lists what was written. reqOpts apply to every request.
// opts may be nil.
//
//	result, err := client.Agents.Push(ctx, agent, nil)                                       // create
//	result, err = client.Agents.Push(ctx, agent, &agentbuilder.PushOptions{AgentID: id})      // update
//	result, err = client.Agents.Push(ctx, agent, &agentbuilder.PushOptions{Project: project}) // ids in a project
func (c *AgentsClient) Push(ctx context.Context, agent *Agent, opts *PushOptions, reqOpts ...option.RequestOption) (*PushResult, error) {
	if agent == nil {
		return nil, errors.New("agentbuilder: Push needs an agent")
	}
	if opts == nil {
		opts = &PushOptions{}
	}
	project, branchID := opts.Project, opts.BranchID
	if project != nil && opts.AgentID != "" {
		return nil, errors.New("agentbuilder: pass AgentID or Project to Push, not both: the project records the agent's id")
	}
	if err := agent.Validate(nil); err != nil {
		return nil, err
	}
	tools := agent.Tools()
	agentID := opts.AgentID
	// With a project, every id it records is known before any request.
	projectIDs := make([]string, len(tools))
	if project != nil {
		if err := project.checkBranch(agent, branchID); err != nil {
			return nil, err
		}
		var err error
		if agentID, err = project.AgentID(agent); err != nil {
			return nil, err
		}
		for i, tool := range tools {
			if projectIDs[i], err = project.ToolID(tool); err != nil {
				return nil, err
			}
		}
	}

	result := PushResult{Tools: map[string]ToolResult{}}
	ids := map[string]string{}
	step := "Couldn't read the agent's tools"
	fail := func(err error) (*PushResult, error) {
		if project != nil && project.dir != "" {
			// The push error is the one to report; the ids it carries are in its result.
			_, _ = project.Save()
		}
		return nil, &PushError{Message: step, Result: result, Cause: err}
	}

	// The tools the agent already uses. With a project they're matched by recorded id; without
	// one, by name among the tools this agent uses, so a same-named tool another agent uses is
	// never touched.
	var inUse []remoteTool
	if agentID != "" {
		var err error
		if branchID != "" {
			inUse, err = c.toolsOnBranch(ctx, agentID, branchID, reqOpts)
		} else {
			inUse, err = c.toolsUsedBy(ctx, agentID, reqOpts)
		}
		if err != nil {
			return fail(err)
		}
	}
	usedByName := map[string]remoteTool{}
	for _, remote := range inUse {
		usedByName[remote.name] = remote
	}

	for i, tool := range tools {
		name := tool.Name()
		knownID := projectIDs[i]
		if project == nil {
			knownID = usedByName[name].id
		}
		var remote *remoteTool
		if knownID != "" {
			step = fmt.Sprintf("Couldn't read tool %q (%s)", name, knownID)
			for j := range inUse {
				if inUse[j].id == knownID {
					remote = &inUse[j]
					break
				}
			}
			if remote == nil {
				fetched, err := c.getTool(ctx, knownID, reqOpts)
				if err != nil {
					return fail(err)
				}
				remote = &fetched
			}
		}
		var config any
		if project != nil {
			config = project.unchangedRaw(tool)
		}
		if config == nil {
			wire, err := tool.ToJSON()
			if err != nil {
				return fail(err)
			}
			config = wire
		}
		if remote != nil && sameRemoteTool(*remote, tool, config) {
			ids[name] = remote.id
			result.Tools[name] = ToolResult{ID: remote.id, Action: ToolUnchanged}
			// The server already matches, but the project may hold an older definition: bring it
			// up to date.
			if project != nil {
				if err := project.recordTool(tool, remote.id, config, false); err != nil {
					return fail(err)
				}
			}
			continue
		}
		body := map[string]any{"tool_config": config}
		if remote != nil {
			step = fmt.Sprintf("Couldn't update tool %q (%s)", name, remote.id)
			if _, err := c.call(ctx, http.MethodPatch, toolPath(remote.id), nil, body, reqOpts); err != nil {
				return fail(err)
			}
			ids[name] = remote.id
			result.Tools[name] = ToolResult{ID: remote.id, Action: ToolUpdated}
		} else {
			step = fmt.Sprintf("Couldn't create tool %q", name)
			response, err := c.call(ctx, http.MethodPost, "/v1/convai/tools", nil, body, reqOpts)
			if err != nil {
				return fail(err)
			}
			created, err := parseRemoteTool(response)
			if err != nil {
				return fail(err)
			}
			ids[name] = created.id
			result.Tools[name] = ToolResult{ID: created.id, Action: ToolCreated}
		}
		if project != nil {
			if err := project.recordTool(tool, ids[name], config, false); err != nil {
				return fail(err)
			}
		}
	}

	step = fmt.Sprintf("Couldn't assemble agent %q", agent.Name())
	config, err := agent.ToJSON(&BuildOptions{ToolIDs: ids})
	if err != nil {
		return fail(err)
	}

	// Every id the new config still uses, wherever it is: the agent's tool_ids and the workflow's
	// nodes.
	kept := map[string]bool{}
	promptIDs, _ := getPath(config, toolIDsPath)
	for _, id := range append(stringList(promptIDs), workflowToolIDs(config["workflow"], nil)...) {
		kept[id] = true
	}
	for _, remote := range inUse {
		if _, touched := result.Tools[remote.name]; !kept[remote.id] && !touched {
			result.Tools[remote.name] = ToolResult{ID: remote.id, Action: ToolDetached}
		}
	}

	versionID := ""
	if agentID == "" {
		step = fmt.Sprintf("Couldn't create agent %q", agent.Name())
		response, err := c.call(ctx, http.MethodPost, "/v1/convai/agents/create", nil, config, reqOpts)
		if err != nil {
			return fail(err)
		}
		var created struct {
			AgentID string `json:"agent_id"`
		}
		if err := json.Unmarshal(response, &created); err != nil || created.AgentID == "" {
			return fail(errors.New("agentbuilder: the create-agent response has no agent_id"))
		}
		result.AgentID, result.Action = created.AgentID, AgentCreated
	} else {
		step = fmt.Sprintf("Couldn't update agent %q (%s)", agent.Name(), agentID)
		body := config
		if opts.VersionDescription != "" {
			body = wirejson.Clone(config).(map[string]any)
			body["version_description"] = opts.VersionDescription
		}
		response, err := c.call(ctx, http.MethodPatch, agentPath(agentID), branchQuery(branchID), body, reqOpts)
		if err != nil {
			return fail(err)
		}
		var updated struct {
			VersionID *string `json:"version_id"`
		}
		if json.Unmarshal(response, &updated) == nil && updated.VersionID != nil {
			versionID = *updated.VersionID
		}
		result.AgentID, result.Action = agentID, AgentUpdated
	}
	if project != nil {
		step = fmt.Sprintf("Couldn't record agent %q in the project", agent.Name())
		if err := project.recordAgent(agent, result.AgentID, config, branchID, versionID, false); err != nil {
			return fail(err)
		}
		if project.dir != "" {
			if _, err := project.Save(); err != nil {
				return nil, &PushError{Message: "Couldn't save the project", Result: result, Cause: err}
			}
		}
	}
	return &result, nil
}

// Pull loads an existing agent as an *Agent, with its webhook and client tools as agenttools
// tools; other kinds stay attached by id. With opts.Project, it also writes the agent and its
// tools there, in the CLI's layout, and records their ids, matching entries by id. reqOpts apply
// to every request. opts may be nil.
//
//	agent, err := client.Agents.Pull(ctx, agentID, &agentbuilder.PullOptions{Project: project})
//	agent.Update(&agentbuilder.Options{FirstMessage: ptr("Hello!")})
//	_, err = client.Agents.Push(ctx, agent, &agentbuilder.PushOptions{Project: project})
func (c *AgentsClient) Pull(ctx context.Context, agentID string, opts *PullOptions, reqOpts ...option.RequestOption) (*Agent, error) {
	if opts == nil {
		opts = &PullOptions{}
	}
	project, branchID := opts.Project, opts.BranchID
	if project != nil {
		if err := project.checkBranchByID(agentID, branchID); err != nil {
			return nil, err
		}
	}
	response, err := c.get(ctx, agentPath(agentID), branchQuery(branchID), reqOpts)
	if err != nil {
		return nil, err
	}
	text, config, err := pulledConfig(response)
	if err != nil {
		return nil, err
	}
	var versionID string
	var live struct {
		VersionID *string `json:"version_id"`
	}
	if json.Unmarshal(response, &live) == nil && live.VersionID != nil {
		versionID = *live.VersionID
	}

	toolsByID := map[string]agenttools.Tool{}
	toolConfigs := map[string]any{}
	var order []string
	for _, id := range toolIDsIn(config, nodeOrder(field(text, "workflow"))) {
		remote, err := c.getTool(ctx, id, reqOpts)
		if err != nil {
			return nil, err
		}
		tool, err := agenttools.FromJSON(remote.config)
		if err != nil {
			return nil, fmt.Errorf("agentbuilder: tool %s: %w", id, err)
		}
		// tools.json lists webhook and client tools; others (MCP...) stay attached by id.
		if kind := tool.Kind(); kind == agenttools.ToolKindWebhook || kind == agenttools.ToolKindClient {
			toolsByID[id] = tool
			toolConfigs[id] = remote.tree
			order = append(order, id)
		}
	}

	agent, err := FromJSON(text, &LoadOptions{ToolsByID: toolsByID})
	if err != nil {
		return nil, err
	}
	if project != nil {
		for _, id := range order {
			if err := project.recordTool(toolsByID[id], id, toolConfigs[id], true); err != nil {
				return nil, err
			}
		}
		if err := project.recordAgent(agent, agentID, config, branchID, versionID, true); err != nil {
			return nil, err
		}
		if project.dir != "" {
			if _, err := project.Save(); err != nil {
				return nil, err
			}
		}
	}
	return agent, nil
}

// pulledConfig builds the CLI's pulled-config shape from a get-agent response: name,
// conversation_config, platform_settings, tags, and workflow only when present, so ids and server
// metadata stay out of the file. It returns the config as JSON text (in the response's order) and
// parsed.
func pulledConfig(response json.RawMessage) ([]byte, map[string]any, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(response, &fields); err != nil {
		return nil, nil, fmt.Errorf("agentbuilder: unexpected get-agent response: %w", err)
	}
	or := func(key, fallback string) json.RawMessage {
		value := bytes.TrimSpace(fields[key])
		if len(value) == 0 || string(value) == "null" {
			return json.RawMessage(fallback)
		}
		return value
	}
	var buf bytes.Buffer
	buf.WriteString(`{"name":`)
	buf.Write(or("name", "null"))
	buf.WriteString(`,"conversation_config":`)
	buf.Write(or("conversation_config", "{}"))
	buf.WriteString(`,"platform_settings":`)
	buf.Write(or("platform_settings", "{}"))
	buf.WriteString(`,"tags":`)
	buf.Write(or("tags", "[]"))
	if workflow := bytes.TrimSpace(fields["workflow"]); len(workflow) > 0 && workflow[0] == '{' {
		buf.WriteString(`,"workflow":`)
		buf.Write(workflow)
	}
	buf.WriteByte('}')
	config, err := decodeObject(buf.Bytes(), "the agent config")
	if err != nil {
		return nil, nil, err
	}
	return buf.Bytes(), config, nil
}

// toolIDsIn lists every tool id an agent config uses, in the agent's tool_ids and the workflow's
// nodes, once each, in order.
func toolIDsIn(config map[string]any, workflowOrder []string) []string {
	promptIDs, _ := getPath(config, toolIDsPath)
	return unique(append(stringList(promptIDs), workflowToolIDs(config["workflow"], workflowOrder)...))
}

func (c *AgentsClient) getTool(ctx context.Context, id string, opts []option.RequestOption) (remoteTool, error) {
	response, err := c.get(ctx, toolPath(id), nil, opts)
	if err != nil {
		return remoteTool{}, err
	}
	return parseRemoteTool(response)
}

// toolsOnBranch lists the tools a branch of the agent uses. The tools API's used_by_agent_id
// covers the main branch only, so they come from the branch's config, by id.
func (c *AgentsClient) toolsOnBranch(ctx context.Context, agentID, branchID string, opts []option.RequestOption) ([]remoteTool, error) {
	response, err := c.get(ctx, agentPath(agentID), branchQuery(branchID), opts)
	if err != nil {
		return nil, err
	}
	branch, err := decodeObject(response, "the get-agent response")
	if err != nil {
		return nil, err
	}
	var tools []remoteTool
	for _, id := range toolIDsIn(branch, nodeOrder(field(response, "workflow"))) {
		tool, err := c.getTool(ctx, id, opts)
		if err != nil {
			return nil, err
		}
		tools = append(tools, tool)
	}
	return tools, nil
}

// toolsUsedBy lists the tools the agent's main branch uses, following the cursor.
func (c *AgentsClient) toolsUsedBy(ctx context.Context, agentID string, opts []option.RequestOption) ([]remoteTool, error) {
	var tools []remoteTool
	cursor := ""
	for {
		query := url.Values{"used_by_agent_id": {agentID}, "page_size": {strconv.Itoa(100)}}
		if cursor != "" {
			query.Set("cursor", cursor)
		}
		response, err := c.get(ctx, "/v1/convai/tools", query, opts)
		if err != nil {
			return nil, err
		}
		var page struct {
			Tools      []json.RawMessage `json:"tools"`
			HasMore    bool              `json:"has_more"`
			NextCursor *string           `json:"next_cursor"`
		}
		if err := json.Unmarshal(response, &page); err != nil {
			return nil, fmt.Errorf("agentbuilder: unexpected list-tools response: %w", err)
		}
		for _, item := range page.Tools {
			tool, err := parseRemoteTool(item)
			if err != nil {
				return nil, err
			}
			tools = append(tools, tool)
		}
		if !page.HasMore || page.NextCursor == nil || *page.NextCursor == "" {
			return tools, nil
		}
		cursor = *page.NextCursor
	}
}

// sameRemoteTool reports whether the server already has this tool as defined. Modeled fields are
// compared through the builder, so both sides normalize alike; the config to send must also be
// contained in the server's, so a change to a field the builder doesn't model still counts.
func sameRemoteTool(remote remoteTool, tool agenttools.Tool, config any) bool {
	if len(remote.config) == 0 {
		return false
	}
	loaded, err := agenttools.FromJSON(remote.config)
	if err != nil {
		return false
	}
	local, ok := toolDefinition(tool)
	if !ok {
		return false
	}
	theirs, ok := toolDefinition(loaded)
	return ok && local == theirs && isSubset(config, remote.tree)
}
