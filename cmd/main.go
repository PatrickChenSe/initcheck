package main

import (
	"github.com/patrickchense/initcheck/pkg/initchecker"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(initchecker.InitChecker)
}
