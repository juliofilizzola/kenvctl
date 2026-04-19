package command

import "github.com/spf13/cobra"

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new environment variable",
	RunE:  Add,
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all environment variables",
	RunE:  List,
}

var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an environment variable",
	RunE:  Remove,
}

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an environment variable",
	RunE:  Set,
}
