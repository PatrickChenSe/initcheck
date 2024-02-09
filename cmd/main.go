package main

import (
	"github.com/patrickchense/initcheck/internal/initchecker"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(initchecker.InitChecker)
}
