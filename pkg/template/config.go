package template

import (
	"text/template"
)

var validOptions = []string{"default", "error", "zero", "error"}

type Config struct {
	options        []string
	leftDelim      string
	rightDelim     string
	extraFunctions template.FuncMap
}

func newConfig() *Config {
	return &Config{
		options:        []string{"missingkey=default"},
		leftDelim:      "{{",
		rightDelim:     "}}",
		extraFunctions: template.FuncMap{},
	}
}
