package sa1032

import (
	"testing"

	"honnef.co/go/tools/analysis/lint/testutil"
)

func TestTestdata(t *testing.T) {
	testutil.Run(t, SCAnalyzer)
}
