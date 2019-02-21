package todo

import (
	"github.com/opsbot/cli-go/api"
	"github.com/opsbot/cli-go/utils"
	"github.com/spf13/cobra"
)

// ListCommand returns a cobra command
func ListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list todos",
		Run: func(cmd *cobra.Command, args []string) {
			data := todo.List()
			utils.PrettyPrint(data)
		},
	}
	return cmd
}
