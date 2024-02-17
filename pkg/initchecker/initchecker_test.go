package initchecker

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	tests := []struct {
		name     string
		pkgPaths []string
	}{
		{
			name:     "Use var for slice init",
			pkgPaths: []string{"a"}, // "a" testdata/a
		},
		{
			name:     "Use make for map init",
			pkgPaths: []string{"b"},
		},
		{
			name:     "Use var for zero value struct init",
			pkgPaths: []string{"c"},
		},
		{
			name:     "Use := for simple type variable init",
			pkgPaths: []string{"d"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			analysistest.Run(t, testdata, InitChecker, tt.pkgPaths...)
		})
	}
}
