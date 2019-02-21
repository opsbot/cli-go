package todo

import (
	"github.com/opsbot/cli-go/api"
	"github.com/opsbot/cli-go/utils"
	"github.com/spf13/cobra"
)

// GetCommand returns a cobra command
func GetCommand() *cobra.Command {
	var todoSlug string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "get todo",
		Run: func(cmd *cobra.Command, args []string) {
			data := todo.Get(todoSlug)
			utils.PrettyPrint(data)
		},
	}

	cmd.Flags().StringVarP(&todoSlug, "slug", "s", "", "the todo slug to fetch")
	cmd.MarkFlagRequired("slug")

	return cmd
}
