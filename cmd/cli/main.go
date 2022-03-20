/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"cobra_starter/cmd/cli/cmd"
	"cobra_starter/cmd/cli/shared"
)

func main() {
	shared.Version = version
	cmd.Execute()
}

var version = "Development"
