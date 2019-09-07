package codegen

import (
	"bytes"
	"dzc/pkg/pkg/pkg"
	"text/template"

	"github.com/pkg/errors"
)

const (
	tmp = `
//{{ .Pkg.Name }}

#include <stdint.h>
#include <stddef.h>

#define true 1
#define false 0

#define i8_t int8_t
#define u8_t uint8_t
#define i16_t int16_t
#define u16_t uint16_t
#define i32_t int32_t
#define u32_t uint32_t
#define i64_t int64_t
#define u64_t uint64_t

{{ range .Consts -}}
#define {{ .Name }} {{ .Value }}
{{ end }}

//stub
int main() {
	return 0;
}
`
)

func GenPkgCode(pkg *pkg.Info) (string, error) {
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
