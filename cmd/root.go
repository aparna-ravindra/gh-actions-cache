package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const VERSION = "1.0.0"

var rootCmd = &cobra.Command{
	Use:   "gh-actions-cache",
	Short: "Works with GitHub Actions Cache. ",
}

func Execute() {
	addCommandsToRoot()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetHelpTemplate(getRootHelp())
}

func addCommandsToRoot() {
	rootCmd.AddCommand(NewCmdList())
	rootCmd.AddCommand(NewCmdDelete())
}

func getRootHelp() string {
	return `
gh-actions-cache: Works with GitHub Actions Cache. 

USAGE:
	gh actions-cache <command> [flags]

CORE COMMANDS:
	list:		list caches with result length cap of 100
	delete:		delete caches with a key

INHERITED FLAGS
	--help		Show help for command

EXAMPLES:
	$ gh actions-cache list
	$ gh actions-cache list --limit 100
	$ gh actions-cache list --order desc
	$ gh actions-cache delete Linux-node-f5dbf39c9d11eba80242ac13
`
}
