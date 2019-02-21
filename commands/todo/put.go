package todo

import (
	"github.com/opsbot/cli-go/api"
	"github.com/opsbot/cli-go/utils"
	"github.com/spf13/cobra"
)

// PutCommand returns a cobra command
func PutCommand() *cobra.Command {
	var todoSlug string
	var todoTitle string
	var todoBody string

	cmd := &cobra.Command{
		Use:   "put",
		Short: "update todo",
		Run: func(cmd *cobra.Command, args []string) {
			todoDoc := todo.Todo{
				Title: todoTitle,
				Body:  todoBody,
			}
			data := todo.Put(todoSlug, todoDoc)
			utils.PrettyPrint(data)
		},
	}

	cmd.Flags().StringVarP(&todoSlug, "slug", "s", "", "the todo slug to udpate")
	cmd.Flags().StringVarP(&todoTitle, "title", "t", "", "the todo title")
	cmd.Flags().StringVarP(&todoBody, "body", "b", "", "the todo body")
	cmd.MarkFlagRequired("slug")
	cmd.MarkFlagRequired("title")

	return cmd
}
