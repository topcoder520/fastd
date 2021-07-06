package cmd

import (
	"github.com/topcoder520/fastd/internal/errHandle"
	"github.com/topcoder520/fastd/internal/executioner"
	"github.com/topcoder520/fastd/internal/resume"
	"github.com/topcoder520/fastd/internal/tool"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(resumeCmd)
}

var resumeCmd = &cobra.Command{
	Use:     "resume",
	Short:   "resume downloading task",
	Example: `fdlr resume URL or file name`,
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		errHandle.ExitWithError(resumeTask(args))
	},
}

func resumeTask(args []string) error {
	task := ""
	if tool.IsVaildURL(args[0]) {
		task = tool.GetFilenameFrom(args[0])
	} else {
		task = args[0]
	}

	state, err := resume.Resume(task)
	if err != nil {
		return errors.WithStack(err)
	}
	return executioner.Do(state.URL, state, conc)
}
