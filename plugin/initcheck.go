package main

import (
	"fmt"
	"github.com/patrickchense/initcheck/pkg/initchecker"

	"golang.org/x/tools/go/analysis"
)

// New plugin, follow doc https://golangci-lint.run/contributing/new-linters/#configure-a-plugin
func New(conf any) ([]*analysis.Analyzer, error) {

	fmt.Printf("My configuration (%[1]T): %#[1]v\n", conf)

	return []*analysis.Analyzer{initchecker.InitChecker}, nil
}
