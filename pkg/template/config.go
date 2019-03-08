package template

import (
	"text/template"
)

type ConfigSetter = func(*config)

type config struct {
	options        []string
	leftDelim      string
	rightDelim     string
	extraFunctions template.FuncMap
}

func NewConfig(s ...ConfigSetter) *config {
	defConfig := &config{
		options:        []string{"missingkey=default"},
		leftDelim:      "{{",
		rightDelim:     "}}",
		extraFunctions: template.FuncMap{},
	}

	for _, setter := range s {
		setter(defConfig)
	}

	return defConfig
}

// Add more validation to each setter

func SetOptions(v ...string) ConfigSetter {
	return func(c *config) {
		if len(v) > 0 {
			c.options = v
		}
	}
}

func SetLeftDelimiter(d string) ConfigSetter {
	return func(c *config) {
		if d != "" {
			c.leftDelim = d
		}
	}
}

func SetRightDelimiter(d string) ConfigSetter {
	return func(c *config) {
		if d != "" {
			c.rightDelim = d
		}
	}
}

func SetExtraFunctions(e template.FuncMap) ConfigSetter {
	return func(c *config) {
		if e != nil {
			c.extraFunctions = e
		}
	}
}
