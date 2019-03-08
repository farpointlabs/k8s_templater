package template

import (
	"bytes"
	"fmt"
	"text/template"
)

type Template struct {
	config *config
	name   string
}

func New(n string, s ...ConfigSetter) *Template {
	return &Template{
		config: NewConfig(s...),
		name:   n,
	}
}

func (t *Template) buildTemplater(data string) (*template.Template, error) {
	return template.New(t.name).
		Delims(t.config.leftDelim, t.config.rightDelim).
		Funcs(t.config.extraFunctions).
		Option(t.config.options...).
		Parse(data)
}

func (t *Template) Execute(data string, values interface{}) (string, error) {
	var buffer bytes.Buffer
	templater, _ := t.buildTemplater(data)
	err := templater.Execute(&buffer, values)
	if err != nil {
		retErr := err
		if e, ok := err.(template.ExecError); ok {
			retErr = fmt.Errorf("error (ExecError) evaluating the template named '%s': %s", e.Name, err)
		}
		return "", retErr
	}
	return buffer.String(), nil
}
