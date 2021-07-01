package cmd

import (
	"fastd/internal/resume"

	"fastd/internal/errHandle"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(taskCmd)
}

var taskCmd = &cobra.Command{
	Use:     "task",
	Short:   "show current downloading task",
	Example: `fdlr task`,
	Run: func(cmd *cobra.Command, args []string) {
		errHandle.ExitWithError(task())
	},
}

func task() error {
	err := resume.TaskPrint()
	return errors.WithStack(err)
}
