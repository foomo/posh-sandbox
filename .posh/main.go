package main

import (
	"github.com/foomo/posh-sandbox/posh/pkg"
	"github.com/foomo/posh/cmd"
)

func init() {
	cmd.Init(pkg.New)
}

func main() {
	cmd.Execute()
}
