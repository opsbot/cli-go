package cli

import (
	"github.com/opsbot/cli-go/cli/todo"
	"github.com/spf13/cobra"
)

// AddCommands add the commands from subcommand groups to the root command
func AddCommands(cmd *cobra.Command) {
	cmd.AddCommand(
		todo.CreateCommand(),
		todo.DeleteCommand(),
		todo.GetCommand(),
		todo.ListCommand(),
		todo.PutCommand(),
		VersionCommand(),
	)
}
