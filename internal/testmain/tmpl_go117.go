//go:build !go1.18
// +build !go1.18

package testmain

var testmainData = `
package main

import (
{{if not .TestMain}}
	"os"
{{end}}
	"testing"
	"testing/internal/testdeps"

{{if .ImportTest}}
	{{if .ExecutesTest}}_test{{else}}_{{end}} {{.Package.ImportPath | printf "%q"}}
{{end -}}
{{- if .ImportXTest -}}
	{{if .ExecutesXTest}}_xtest{{else}}_{{end}} {{.Package.ImportPath | printf "%s_test" | printf "%q"}}
{{end}}
)

var tests = []testing.InternalTest{
{{- range .Tests}}
	{"{{.Name}}", {{.Location}}.{{.Name}}},
{{- end}}
}

var benchmarks = []testing.InternalBenchmark{
{{- range .Benchmarks}}
	{"{{.Name}}", {{.Location}}.{{.Name}}},
{{- end}}
}

var examples = []testing.InternalExample{
{{- range .Examples }}
{{- if .Executable }}
	{"{{.Name}}", {{.Location}}.{{.Name}}, {{.Output | printf "%q"}}, {{.Unordered}}},
{{- end }}
{{- end }}
}

func main() {
	m := testing.MainStart(testdeps.TestDeps{}, tests, benchmarks, examples)
{{with .TestMain}}
	{{.Location}}.{{.Name}}(m)
{{else}}
	os.Exit(m.Run())
{{end -}}
}

`
