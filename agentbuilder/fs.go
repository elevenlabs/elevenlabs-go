package agentbuilder

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// OpenProject reads the project in dir: agents.json, tools.json and every config file they list.
// A directory without index files opens as an empty project, and Save creates them. A config path
// that leads outside the project, including through a symlinked directory, is a *ProjectError.
func OpenProject(dir string) (*Project, error) {
	files, taken, err := readProject(dir)
	if err != nil {
		return nil, err
	}
	return newProject(files, dir, taken)
}

// readProject reads the index files and every config they list. taken lists the files already in
// the config directories.
func readProject(dir string) (map[string]string, []string, error) {
	root, err := projectRoot(dir)
	if err != nil {
		return nil, nil, err
	}
	files := map[string]string{}
	for _, index := range []string{agentsFile, toolsFile} {
		text, ok, err := readInProject(root, index)
		if err != nil {
			return nil, nil, err
		}
		if !ok {
			continue
		}
		files[index] = text
		var parsed struct {
			Agents []struct {
				Config   any                        `json:"config"`
				Branches map[string]json.RawMessage `json:"branches"`
			} `json:"agents"`
			Tools []struct {
				Config any `json:"config"`
			} `json:"tools"`
		}
		if json.Unmarshal([]byte(text), &parsed) != nil {
			continue // the project reports it, with the parse error
		}
		var configs []any
		for _, entry := range parsed.Agents {
			configs = append(configs, entry.Config)
			for _, raw := range entry.Branches {
				var branch struct {
					Config any `json:"config"`
				}
				if json.Unmarshal(raw, &branch) == nil {
					configs = append(configs, branch.Config)
				}
			}
		}
		for _, entry := range parsed.Tools {
			configs = append(configs, entry.Config)
		}
		for _, config := range configs {
			path, isString := config.(string)
			if !isString {
				return nil, nil, projectErrorf("%s has an entry without a \"config\" path", index)
			}
			content, ok, err := readInProject(root, path)
			if err != nil {
				return nil, nil, err
			}
			if ok {
				files[normalizePath(path)] = content
			}
		}
	}
	var taken []string
	for _, sub := range []string{agentConfigsDir, toolConfigsDir} {
		entries, err := os.ReadDir(filepath.Join(root, sub))
		if err != nil {
			continue // no directory yet
		}
		for _, entry := range entries {
			taken = append(taken, sub+"/"+entry.Name())
		}
	}
	return files, taken, nil
}

func projectRoot(dir string) (string, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(abs)
}

func readInProject(root, path string) (string, bool, error) {
	target, err := resolveInProject(root, path)
	if err != nil {
		return "", false, err
	}
	data, err := os.ReadFile(target)
	if errors.Is(err, fs.ErrNotExist) {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return string(data), true, nil
}

// writeProjectFiles writes each file owner-only, creating directories as needed, never through a
// symlink and never outside the project.
func writeProjectFiles(dir string, files [][2]string) error {
	root, err := projectRoot(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		path, text := file[0], file[1]
		target, err := resolveInProject(root, path)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		// The directories may have been created through a symlink since the check.
		if _, err := resolveInProject(root, path); err != nil {
			return err
		}
		if info, err := os.Lstat(target); err == nil && info.Mode()&os.ModeSymlink != 0 {
			return projectErrorf("Refusing to write %s: it's a symlink", path)
		}
		f, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|oNoFollow, 0o600)
		if err != nil {
			if isSymlinkError(err) {
				return projectErrorf("Refusing to write %s: it's a symlink", path)
			}
			return err
		}
		_ = f.Chmod(0o600) // a file that already existed keeps its mode otherwise
		_, err = f.WriteString(text)
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// resolveInProject is the CLI's resolve_in_project: the absolute path for a project-relative one,
// refusing anything that escapes the project, including through a symlinked directory.
func resolveInProject(root, path string) (string, error) {
	if strings.TrimSpace(path) == "" || strings.IndexFunc(path, unicode.IsControl) >= 0 {
		return "", projectErrorf("Invalid config path %q", path)
	}
	candidate := filepath.FromSlash(path)
	if !filepath.IsAbs(candidate) {
		candidate = filepath.Join(root, candidate)
	}
	candidate = filepath.Clean(candidate)
	inside := func(p string) bool {
		return p == root || strings.HasPrefix(p, strings.TrimSuffix(root, string(filepath.Separator))+string(filepath.Separator))
	}
	if !inside(candidate) {
		return "", projectErrorf("Config path %s is outside the project", path)
	}
	// The deepest part of the path that exists, with symlinks resolved, must still be inside.
	for ancestor := candidate; ; {
		if real, err := filepath.EvalSymlinks(ancestor); err == nil {
			if !inside(real) {
				return "", projectErrorf("Config path %s leads outside the project through a symlink", path)
			}
			return candidate, nil
		}
		parent := filepath.Dir(ancestor)
		if parent == ancestor {
			return "", projectErrorf("Config path %s is outside the project", path)
		}
		ancestor = parent
	}
}
