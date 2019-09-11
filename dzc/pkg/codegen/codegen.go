package codegen

import (
	"bytes"
	"text/template"

	"dzc/pkg/ast/pkginfo"

	"github.com/pkg/errors"
)

const (
	tmp =
`
//{{ .Pkg.Name }}

#include <stdint.h>
#include <stddef.h>

#define true 1
#define false 0

{{ range .Consts -}}
#define {{ .Name }} {{ .Value }}
{{ end -}}

`
)

func GenPkgCode(pkg *pkginfo.PkgInfo) (string, error) {
	t, err := template.New("").Parse(tmp)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse template")
	}

	buf := bytes.Buffer{}
	if err := t.Execute(&buf, pkg); err != nil {
		return "", errors.Wrap(err, "failed to execute template")
	}

	return buf.String(), nil
}
