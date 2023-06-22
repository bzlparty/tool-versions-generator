package output

import (
	"os"
	"text/template"

	"github.com/bzlparty/tool-versions-generator/pkg/github"
)

const OUTPUT_TEMPLATE = `# This file was generated from https://github.com/{{.Repo}}/releases
TOOL_VERSIONS = { {{range $key, $val := .ResultMap}}
    "{{$key}}": { {{range $val}}
        "{{.Platform}}": "{{.Integrity}}", {{end}}
    }, {{end}}
}`

type Output struct {
	data     OutputData
	template *template.Template
}

type OutputData struct {
	ResultMap map[string][]github.PlatformAsset
	Repo      string
}

func (o *Output) Write(filepath string) (err error) {
	temp, err := o.template.Parse(OUTPUT_TEMPLATE)
	out := os.Stdout

	if err != nil {
		return
	}

	if filepath != "" {
		out, err = os.Create(filepath)

		if err != nil {
			return
		}
	}

	if err = temp.Execute(out, o.data); err != nil {
		return
	}

	return nil
}

func NewOutput(data OutputData) *Output {
	t := template.New("output")
	return &Output{
		data:     data,
		template: t,
	}
}
