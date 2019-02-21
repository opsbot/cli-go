package todo

import (
	"github.com/opsbot/cli-go/api"
	"github.com/opsbot/cli-go/utils"
	"github.com/spf13/cobra"
)

// DeleteCommand returns a cobra command
func DeleteCommand() *cobra.Command {
	var todoSlug string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete todo",
		Run: func(cmd *cobra.Command, args []string) {
			data := todo.Delete(todoSlug)
			utils.PrettyPrint(data)
		},
	}

	cmd.Flags().StringVarP(&todoSlug, "slug", "s", "", "the todo slug to fetch")
	cmd.MarkFlagRequired("slug")

	return cmd
}
