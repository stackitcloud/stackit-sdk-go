package pkgstutter

import (
	"strings"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "pkgstutter",
	Doc:  "Prevents stuttering package names, e.g. github.com/foo/bar/wait/wait",
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	pkgPath := pass.Pkg.Path()
	parts := strings.Split(pkgPath, "/")

	// Check for adjacent identical parts in the path
	for i := 1; i < len(parts); i++ {
		if parts[i] == parts[i-1] {
			// If a stutter is found, report it at the package declaration
			// of the first file in the package to pinpoint the error.
			if len(pass.Files) > 0 {
				pass.Reportf(
					pass.Files[0].Package,
					"package path %q contains stuttering (%s/%s)",
					pkgPath, parts[i-1], parts[i],
				)
			}
			// Break after the first finding to avoid spamming multiple errors for the same package
			break
		}
	}

	return nil, nil
}

func init() {
	register.Plugin("pkgstutter", New)
}

func New(settings any) (register.LinterPlugin, error) {
	return &plugin{}, nil
}

type plugin struct{}

func (p *plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{Analyzer}, nil
}

func (p *plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
