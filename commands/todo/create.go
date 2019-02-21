package todo

import (
	"github.com/opsbot/cli-go/api"
	"github.com/opsbot/cli-go/utils"
	"github.com/spf13/cobra"
)

// CreateCommand returns a cobra command
func CreateCommand() *cobra.Command {
	var todoTitle string
	var todoBody string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "create todo",
		Run: func(cmd *cobra.Command, args []string) {
			data := todo.Create(todo.Todo{
				Title: todoTitle,
				Body:  todoBody,
			})
			utils.PrettyPrint(data)
		},
	}

	cmd.Flags().StringVarP(&todoTitle, "title", "t", "", "the todo title")
	cmd.Flags().StringVarP(&todoBody, "body", "b", "", "the todo body")
	cmd.MarkFlagRequired("slug")
	cmd.MarkFlagRequired("title")

	return cmd
}
