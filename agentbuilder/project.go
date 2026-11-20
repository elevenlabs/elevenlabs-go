package agentbuilder

import (
	"strings"

	"github.com/elevenlabs/elevenlabs-go/agenttools"
	"github.com/elevenlabs/elevenlabs-go/internal/wirejson"
)

// The CLI's project layout.
const (
	agentsFile      = "agents.json"
	toolsFile       = "tools.json"
	agentConfigsDir = "agent_configs"
	toolConfigsDir  = "tool_configs"
)

// agentEntry is one entry in agents.json. Empty strings are unset.
type agentEntry struct {
	config    string
	id        string
	branchID  string
	versionID string
	// branches holds the entry's branches as loaded, by name: *wirejson.Object values with
	// config, branch_id and version_id.
	branches *wirejson.Object
}

// toolEntry is one entry in tools.json. Empty strings are unset.
type toolEntry struct {
	toolType string
	config   string
	id       string
}

type agentSlot struct {
	entry agentEntry
	agent *Agent
}

type toolSlot struct {
	entry toolEntry
	tool  agenttools.Tool
	// raw is the config as loaded (or last recorded), in API format, sent as it is while the tool
	// is unchanged so fields this SDK doesn't model survive.
	raw any
	// baseline is the tool's definition as loaded (see toolDefinition), to tell whether it was
	// edited since.
	baseline string
}

// Project is an agents-as-code project in the ElevenLabs CLI's layout: agents.json and
// tools.json list the agents and tools with their ids, and each one's config sits in
// agent_configs/ or tool_configs/ in API format. A project made by `elevenlabs agents init` or
// `pull` works here, and one written here works with the CLI.
//
// Pass it to client.Agents.Push and client.Agents.Pull: the project records which agent and tools
// on the server each definition is, so pushing again updates them instead of making copies. Files
// are written only by Push and Pull, after they know every id, and a file nothing changed keeps
// its exact text.
//
//	project, err := agentbuilder.OpenProject(".")
//	agent, err := project.Agent("Support line")
//	agent.Update(&agentbuilder.Options{Prompt: ptr("...")})
//	result, err := client.Agents.Push(ctx, agent, &agentbuilder.PushOptions{Project: project})
//
// YAML and TOML configs, tests.json and creating branches aren't supported yet. A Project is not
// safe for concurrent use.
type Project struct {
	dir    string
	files  map[string]string
	agents []*agentSlot
	tools  []*toolSlot
	// changed holds the paths written since loading or the last save, with their new text, in the
	// order they were first written.
	changed      map[string]string
	changedOrder []string
	// taken holds the paths in use, on disk too, so a new file never overwrites one the project
	// doesn't list.
	taken map[string]bool
}

// ProjectFromFiles reads a project from file contents, keyed by path relative to the project
// root. It does no I/O. Only agents.json, tools.json and the config files they list are read;
// missing index files mean an empty project.
func ProjectFromFiles(files map[string]string) (*Project, error) {
	return newProject(files, "", nil)
}

func newProject(files map[string]string, dir string, taken []string) (*Project, error) {
	p := &Project{dir: dir, files: map[string]string{}, changed: map[string]string{}, taken: map[string]bool{}}
	for path, text := range files {
		p.files[normalizePath(path)] = text
		p.taken[normalizePath(path)] = true
	}
	for _, path := range taken {
		p.taken[normalizePath(path)] = true
	}
	if err := p.load(); err != nil {
		return nil, err
	}
	return p, nil
}

// Dir is the directory the project was opened from, or "" for ProjectFromFiles. Save writes
// there.
func (p *Project) Dir() string { return p.dir }

// Agents lists every agent in the project, in agents.json order.
func (p *Project) Agents() []*Agent {
	out := make([]*Agent, len(p.agents))
	for i, slot := range p.agents {
		out[i] = slot.agent
	}
	return out
}

// Agent returns the agent named name: the same *Agent every time, so edits to it are what Push
// sends. It returns a *ProjectError if no agent, or more than one, has that name.
func (p *Project) Agent(name string) (*Agent, error) {
	var matches []*agentSlot
	for _, slot := range p.agents {
		if slot.agent.Name() == name {
			matches = append(matches, slot)
		}
	}
	switch len(matches) {
	case 0:
		return nil, projectErrorf("No agent named %q in %s", name, agentsFile)
	case 1:
		return matches[0].agent, nil
	}
	return nil, projectErrorf("%d agents in %s are named %q", len(matches), agentsFile, name)
}

// Tools lists every tool in tools.json, in order.
func (p *Project) Tools() []agenttools.Tool {
	out := make([]agenttools.Tool, len(p.tools))
	for i, slot := range p.tools {
		out[i] = slot.tool
	}
	return out
}

// AgentID is the id agents.json records for agent, or "" if it has none yet. The entry is the one
// the agent was loaded from, else the only one with its name; more than one is a *ProjectError.
func (p *Project) AgentID(agent *Agent) (string, error) {
	slot, err := p.agentSlot(agent)
	if slot == nil || err != nil {
		return "", err
	}
	return slot.entry.id, nil
}

// ToolID is the id tools.json records for tool, or "" if it has none yet. The entry is the one
// the tool was loaded from, else the only one with its name; more than one is a *ProjectError.
func (p *Project) ToolID(tool agenttools.Tool) (string, error) {
	slot, err := p.toolSlot(tool)
	if slot == nil || err != nil {
		return "", err
	}
	return slot.entry.id, nil
}

// ToFiles returns the project's files, keyed by path: the index files and every config file they
// list, branch configs included, with what Push and Pull wrote since loading. Files nothing
// changed keep their exact text.
func (p *Project) ToFiles() (map[string]string, error) {
	files := map[string]string{}
	for _, index := range []struct {
		path  string
		build func() (string, error)
	}{{agentsFile, p.agentsIndex}, {toolsFile, p.toolsIndex}} {
		if text, ok := p.text(index.path); ok {
			files[index.path] = text
			continue
		}
		text, err := index.build()
		if err != nil {
			return nil, err
		}
		files[index.path] = text
	}
	var paths []string
	for _, slot := range p.agents {
		paths = append(paths, slot.entry.config)
		for _, branch := range branchesOf(slot.entry) {
			paths = append(paths, branch.config)
		}
	}
	for _, slot := range p.tools {
		paths = append(paths, slot.entry.config)
	}
	for _, path := range paths {
		if text, ok := p.text(path); ok {
			files[path] = text
		}
	}
	return files, nil
}

// text is a file's current text: as written since loading, or as loaded.
func (p *Project) text(path string) (string, bool) {
	key := normalizePath(path)
	if text, ok := p.changed[key]; ok {
		return text, true
	}
	text, ok := p.files[key]
	return text, ok
}

// ChangedFiles lists the paths Push and Pull changed since the project was loaded or last saved.
func (p *Project) ChangedFiles() []string { return append([]string{}, p.changedOrder...) }

// Save writes the files Push and Pull changed to the project's directory, and returns their
// paths. Push and Pull call it themselves for a project opened with OpenProject. A project from
// ProjectFromFiles has no directory; use ToFiles.
func (p *Project) Save() ([]string, error) {
	if p.dir == "" {
		return nil, projectErrorf("This project wasn't opened from a directory; use ToFiles instead")
	}
	if len(p.changedOrder) == 0 {
		return nil, nil
	}
	written := make([][2]string, len(p.changedOrder))
	for i, path := range p.changedOrder {
		written[i] = [2]string{path, p.changed[path]}
	}
	if err := writeProjectFiles(p.dir, written); err != nil {
		return nil, err
	}
	for _, file := range written {
		p.files[file[0]] = file[1]
	}
	paths := p.changedOrder
	p.changed, p.changedOrder = map[string]string{}, nil
	return paths, nil
}

// agentSlot is the entry for an agent: the one it was loaded from, else the only one with its
// name.
func (p *Project) agentSlot(agent *Agent) (*agentSlot, error) {
	var named []*agentSlot
	for _, slot := range p.agents {
		if slot.agent == agent {
			return slot, nil
		}
		if slot.agent.Name() == agent.Name() {
			named = append(named, slot)
		}
	}
	if len(named) > 1 {
		return nil, projectErrorf("%d entries in %s are named %q; push the loaded one", len(named), agentsFile, agent.Name())
	}
	if len(named) == 1 {
		return named[0], nil
	}
	return nil, nil
}

// toolSlot is the entry for a tool: the one it was loaded from, else the only one with its name.
func (p *Project) toolSlot(tool agenttools.Tool) (*toolSlot, error) {
	var named []*toolSlot
	for _, slot := range p.tools {
		if sameTool(slot.tool, tool) {
			return slot, nil
		}
		if slot.tool.Name() == tool.Name() {
			named = append(named, slot)
		}
	}
	if len(named) > 1 {
		return nil, projectErrorf("%d entries in %s are named %q; push the loaded one", len(named), toolsFile, tool.Name())
	}
	if len(named) == 1 {
		return named[0], nil
	}
	return nil, nil
}

// unchangedRaw is the loaded config to send for a tool whose definition matches its project
// entry (the loaded tool, or one rebuilt in code with the same name), so fields this SDK doesn't
// model survive. It is nil otherwise.
func (p *Project) unchangedRaw(tool agenttools.Tool) any {
	slot, err := p.toolSlot(tool)
	if err != nil || slot == nil || slot.raw == nil || slot.baseline == "" {
		return nil
	}
	if definition, ok := toolDefinition(tool); ok && definition == slot.baseline {
		return wirejson.Clone(slot.raw)
	}
	return nil
}

// recordTool records a tool's id and the config now on the server, adding an entry if it's new.
// Push finds the entry by definition; Pull (byID), by id.
func (p *Project) recordTool(tool agenttools.Tool, id string, config any, byID bool) error {
	var slot *toolSlot
	if byID {
		for _, s := range p.tools {
			if s.entry.id == id {
				slot = s
				break
			}
		}
	} else {
		var err error
		if slot, err = p.toolSlot(tool); err != nil {
			return err
		}
	}
	if slot == nil {
		path := uniquePath(toolConfigsDir, tool.Name(), p.taken)
		p.taken[path] = true
		slot = &toolSlot{entry: toolEntry{config: path}}
		p.tools = append(p.tools, slot)
	}
	slot.tool = tool
	slot.entry.id = id
	slot.entry.toolType = string(tool.Kind())
	slot.raw = wirejson.Clone(config)
	slot.baseline, _ = toolDefinition(tool)
	text, err := formatConfig(config)
	if err != nil {
		return err
	}
	p.write(slot.entry.config, text)
	index, err := p.toolsIndex()
	if err != nil {
		return err
	}
	p.write(toolsFile, index)
	return nil
}

// recordAgent records an agent's id and its config, adding an entry if it's new. Push matches like
// recordTool; Pull (byID) by id. A branch's config goes to the branch's own file.
func (p *Project) recordAgent(agent *Agent, id string, config any, branchID, versionID string, byID bool) error {
	var slot *agentSlot
	if byID {
		for _, s := range p.agents {
			if s.entry.id == id {
				slot = s
				break
			}
		}
	} else {
		var err error
		if slot, err = p.agentSlot(agent); err != nil {
			return err
		}
	}
	if slot == nil {
		path := uniquePath(agentConfigsDir, agent.Name(), p.taken)
		p.taken[path] = true
		slot = &agentSlot{entry: agentEntry{config: path}, agent: agent}
		p.agents = append(p.agents, slot)
	}
	slot.entry.id = id
	var branch *projectBranch
	if branchID != "" {
		for _, b := range branchesOf(slot.entry) {
			if b.branchID == branchID {
				branch = &b
				break
			}
		}
	}
	// A branch has its own file; the slot's agent is the main branch's, which a branch must not
	// replace.
	path := slot.entry.config
	if branch == nil {
		slot.agent = agent
	} else {
		// checkBranch made sure it has one; never fall back to the main branch's file.
		if branch.config == "" {
			return checkBranchOf(nil, agent.Name(), branchID)
		}
		path = branch.config
	}
	// Like the CLI: the version the server reports, on the entry it belongs to.
	if versionID != "" {
		if branch != nil {
			set(branch.obj, "version_id", versionID)
		} else {
			slot.entry.versionID = versionID
		}
	}
	text, err := formatConfig(config)
	if err != nil {
		return err
	}
	p.write(path, text)
	index, err := p.agentsIndex()
	if err != nil {
		return err
	}
	p.write(agentsFile, index)
	return nil
}

// checkBranch rejects a branch the project has no file for, before anything is written: its
// settings would otherwise land in the main branch's file, and a later CLI push would send them
// to the main branch.
func (p *Project) checkBranch(agent *Agent, branchID string) error {
	slot, err := p.agentSlot(agent)
	if err != nil {
		return err
	}
	return checkBranchOf(slot, agent.Name(), branchID)
}

// checkBranchByID is checkBranch for Pull, which finds the agent by id.
func (p *Project) checkBranchByID(agentID, branchID string) error {
	var slot *agentSlot
	for _, s := range p.agents {
		if s.entry.id == agentID {
			slot = s
			break
		}
	}
	return checkBranchOf(slot, agentID, branchID)
}

func checkBranchOf(slot *agentSlot, label, branchID string) error {
	if branchID == "" || (slot != nil && slot.entry.branchID == branchID) {
		return nil
	}
	if slot != nil {
		// The first branch entry with that id is the one recorded; it needs its own file.
		for _, branch := range branchesOf(slot.entry) {
			if branch.branchID == branchID {
				if branch.config != "" {
					return nil
				}
				break
			}
		}
	}
	return projectErrorf("Branch %s of %q has no config file in %s. Add it under \"branches\" (the CLI's `agents pull --all-branches` does), or push the main branch", branchID, label, agentsFile)
}

// projectBranch is one branch of an agents.json entry.
type projectBranch struct {
	config   string
	branchID string
	obj      *wirejson.Object
}

func branchesOf(entry agentEntry) []projectBranch {
	if entry.branches == nil {
		return nil
	}
	var out []projectBranch
	for _, name := range entry.branches.Keys {
		obj, ok := entry.branches.Values[name].(*wirejson.Object)
		if !ok {
			continue
		}
		config, _ := get(obj, "config").(string)
		branchID, _ := get(obj, "branch_id").(string)
		out = append(out, projectBranch{config: config, branchID: branchID, obj: obj})
	}
	return out
}

// write records a file's new text. A file back to what was loaded is no longer changed.
func (p *Project) write(path, text string) {
	key := normalizePath(path)
	if loaded, ok := p.files[key]; ok && loaded == text {
		if _, pending := p.changed[key]; pending {
			delete(p.changed, key)
			order := p.changedOrder[:0:0]
			for _, k := range p.changedOrder {
				if k != key {
					order = append(order, k)
				}
			}
			p.changedOrder = order
		}
		return
	}
	if _, pending := p.changed[key]; !pending {
		p.changedOrder = append(p.changedOrder, key)
	}
	p.changed[key] = text
}

func (p *Project) load() error {
	agentsIndex, err := p.readIndex(agentsFile, "agents")
	if err != nil {
		return err
	}
	toolsIndex, err := p.readIndex(toolsFile, "tools")
	if err != nil {
		return err
	}

	byID := map[string]agenttools.Tool{}
	for _, item := range toolsIndex {
		obj, _ := item.(*wirejson.Object)
		entry := toolEntry{}
		entry.toolType, _ = get(obj, "type").(string)
		entry.config, _ = get(obj, "config").(string)
		entry.id, _ = get(obj, "id").(string)
		text, err := p.readConfig(get(obj, "config"), toolsFile)
		if err != nil {
			return err
		}
		tool, err := agenttools.FromJSON(text)
		if err != nil {
			return projectErrorf("%s isn't a valid tool config: %v", entry.config, err)
		}
		raw, _ := wirejson.Decode([]byte(text))
		baseline, _ := toolDefinition(tool)
		p.tools = append(p.tools, &toolSlot{entry: entry, tool: tool, raw: raw, baseline: baseline})
		if entry.id != "" {
			byID[entry.id] = tool
		}
	}
	for _, item := range agentsIndex {
		obj, _ := item.(*wirejson.Object)
		entry := agentEntry{}
		entry.config, _ = get(obj, "config").(string)
		entry.id, _ = get(obj, "id").(string)
		entry.branchID, _ = get(obj, "branch_id").(string)
		entry.versionID, _ = get(obj, "version_id").(string)
		entry.branches, _ = get(obj, "branches").(*wirejson.Object)
		text, err := p.readConfig(get(obj, "config"), agentsFile)
		if err != nil {
			return err
		}
		agent, err := FromJSON(text, &LoadOptions{ToolsByID: byID})
		if err != nil {
			return projectErrorf("%s isn't a valid agent config: %v", entry.config, err)
		}
		p.agents = append(p.agents, &agentSlot{entry: entry, agent: agent})
	}
	return nil
}

// readIndex reads the list under key in an index file; a missing file is an empty list.
func (p *Project) readIndex(path, key string) ([]any, error) {
	text, ok := p.files[path]
	if !ok {
		return nil, nil
	}
	tree, err := decodeOrdered([]byte(text))
	if err != nil {
		return nil, projectErrorf("%s isn't valid JSON: %v", path, err)
	}
	obj, isObj := tree.(*wirejson.Object)
	if !isObj {
		return nil, projectErrorf("%s must be a JSON object", path)
	}
	value := get(obj, key)
	if value == nil {
		return nil, nil
	}
	list, ok := value.([]any)
	if !ok {
		return nil, projectErrorf("%s: %q must be a list", path, key)
	}
	for _, item := range list {
		if _, ok := item.(*wirejson.Object); !ok {
			return nil, projectErrorf("%s: every %q entry must be an object", path, key)
		}
	}
	return list, nil
}

// readConfig returns the text of a config file an index points at.
func (p *Project) readConfig(value any, index string) (string, error) {
	path, ok := value.(string)
	if !ok {
		return "", projectErrorf("%s has an entry without a \"config\" path", index)
	}
	if !strings.HasSuffix(strings.ToLower(path), ".json") {
		return "", projectErrorf("%s points at %s: only JSON config files are supported yet (YAML and TOML aren't)", index, path)
	}
	text, ok := p.files[normalizePath(path)]
	if !ok {
		return "", projectErrorf("%s points at %s, which doesn't exist", index, path)
	}
	if _, err := wirejson.Decode([]byte(text)); err != nil {
		return "", projectErrorf("%s isn't valid JSON: %v", path, err)
	}
	return text, nil
}

func (p *Project) agentsIndex() (string, error) {
	agents := make([]any, len(p.agents))
	for i, slot := range p.agents {
		// The CLI writes fields in this order and leaves unset ones out.
		var branches any
		if slot.entry.branches != nil {
			branches = slot.entry.branches
		}
		agents[i] = orderedFields(
			"config", slot.entry.config,
			"id", slot.entry.id,
			"branch_id", slot.entry.branchID,
			"version_id", slot.entry.versionID,
			"branches", branches,
		)
	}
	return formatIndex(&wirejson.Object{Keys: []string{"agents"}, Values: map[string]any{"agents": agents}})
}

func (p *Project) toolsIndex() (string, error) {
	tools := make([]any, len(p.tools))
	for i, slot := range p.tools {
		tools[i] = orderedFields("type", slot.entry.toolType, "config", slot.entry.config, "id", slot.entry.id)
	}
	return formatIndex(&wirejson.Object{Keys: []string{"tools"}, Values: map[string]any{"tools": tools}})
}
