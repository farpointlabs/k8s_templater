package command

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/bmatcuk/doublestar"
	"gopkg.in/yaml.v2"
)

type templateData struct {
	values    map[string]interface{}
	templates []string
}

func findFiles(dir string) ([]string, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}
	g := filepath.Clean(filepath.Join(abs, "**/*.{yml,yaml}"))
	matches, err := doublestar.Glob(g)
	if err != nil {
		return nil, err
	}
	sort.Strings(matches)
	return matches, nil
}

func getAllFiles(dir string) (*templateData, error) {
	resolvedData := &templateData{
		values: map[string]interface{}{},
	}

	matches, err := findFiles(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range matches {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}

		if regexp.MustCompile(`[v,V]alues\.(yml|yaml)`).Match([]byte(f)) {
			yaml.Unmarshal(data, resolvedData.values)
		} else {
			resolvedData.templates = append(resolvedData.templates, string(data))
		}
	}

	return resolvedData, nil
}
