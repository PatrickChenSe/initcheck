package main

import (
	"fmt"
	"github.com/patrickchense/initcheck/internal/initchecker"

	"golang.org/x/tools/go/analysis"
)

func New(conf any) ([]*analysis.Analyzer, error) {

	fmt.Printf("My configuration (%[1]T): %#[1]v\n", conf)

	return []*analysis.Analyzer{initchecker.InitChecker}, nil
}
