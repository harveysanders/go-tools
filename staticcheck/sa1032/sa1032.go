package sa1032

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/analysis/callcheck"
	"honnef.co/go/tools/analysis/lint"
	"honnef.co/go/tools/internal/passes/buildir"
	"honnef.co/go/tools/knowledge"
)

var SCAnalyzer = lint.InitializeAnalyzer(&lint.Analyzer{
	Analyzer: &analysis.Analyzer{
		Name:     "SA1032",
		Requires: []*analysis.Analyzer{buildir.Analyzer},
		Run:      callcheck.Analyzer(checkErrorsRules),
	},
	Doc: &lint.Documentation{
		Title:    `\'err\' and \'target\' arguments passed to \'errors.Is\' may be flipped`,
		Text:     `The first argument to \'errors.Is\' is the wrapped error to check. The second argument is the target error to check against. If these arguments are flipped, the check will always return false. This check ensures that the arguments are passed in the correct order.`,
		Since:    "TODO",
		Severity: lint.SeverityWarning,
		MergeIf:  lint.MergeIfAny,
	},
})

var Analyzer = SCAnalyzer.Analyzer

var checkErrorsRules = map[string]callcheck.Check{
	"errors.Is": CheckFlippedErrTarget(
		knowledge.Arg("errors.Is.err"),
		knowledge.Arg("errors.Is.target")),
}

func CheckFlippedErrTarget(errArgPos, targetArgPos int) callcheck.Check {
	return func(call *callcheck.Call) {
		errArg := call.Args[errArgPos]
		targetArg := call.Args[targetArgPos]

		errExported := ast.IsExported(errArg.Value.Value.Name())
		targetExported := ast.IsExported(targetArg.Value.Value.Name())

		// If neither are exported, we can't tell which is which
		// If both are exported. it's probably a style issue
		// If the target is exported and the error is not, it likely correct usage
		// If the error is exported and the target is not, it's likely a mistake
		if errExported && !targetExported {
			call.Invalid("flipped err and target arguments")
		}
	}
}
