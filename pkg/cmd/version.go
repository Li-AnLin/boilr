package cmd

import (
	"fmt"

	cli "github.com/spf13/cobra"

	"github.com/Li-AnLin/boilr/pkg/boilr"
	"github.com/Li-AnLin/boilr/pkg/util/tlog"
	"github.com/Li-AnLin/boilr/pkg/util/validate"
)

// Version contains the cli-command for printing the current version of the tool.
var Version = &cli.Command{
	Use:   "version",
	Short: "Show the boilr version information",
	Run: func(c *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{})

		shouldntPrettify := GetBoolFlag(c, "dont-prettify")
		if shouldntPrettify {
			fmt.Println(boilr.Version)
		} else {
			tlog.Info(fmt.Sprint("This is a forked version of boilr from github.com/Li-AnLin/boilr!"))
			tlog.Info(fmt.Sprint("You are running Git Commit Hash: ", boilr.Version))
			tlog.Info(fmt.Sprint("See https://github.com/Li-AnLin/boilr/commit/", boilr.Version))
		}
	},
}
