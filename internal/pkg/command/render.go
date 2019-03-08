package command

import (
	"fmt"
	"strings"

	"github.com/farpointlabs/k8s_templater/pkg/template"
	"github.com/pkg/errors"
)

type RenderHandler struct {
	RootPath    string
	Output      string
	Overrides   interface{}
	Environment string
}

func (r *RenderHandler) List(path string) error {
	d, err := getAllFiles(path)
	if err != nil {
		return errors.Wrap(err, "failed to get files")
	}

	t := template.New("")
	res := []string{}
	for _, td := range d.templates {
		s, err := t.Execute(td, d.values)

		if err != nil {
			return errors.Wrap(err, "failed to execute template")
		}

		res = append(res, s)
	}

	fmt.Println(strings.Join(res, "---\n"))

	return nil
}
