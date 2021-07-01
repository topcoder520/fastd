package cmd

import (
	"fastd/internal/errHandle"
	"fastd/internal/executioner"
	"fastd/internal/tool"
	"fmt"
	"os"
	"runtime"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var conc int

var downloadCmd = &cobra.Command{
	Use:     "download",
	Short:   "downloads a file from URL or file name",
	Example: "fastd download [-c=goroutines_count] URL",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		errHandle.ExitWithError(download(args))
	},
}

func init() {
	RootCmd.AddCommand(downloadCmd)
	downloadCmd.Flags().IntVarP(&conc, "goroutines count", "c", runtime.NumCPU(), "default is your CPU threads count")
}

func download(args []string) error {
	//获取下载文件保存路径
	folder, err := tool.GetFolderFrom(args[0])
	if err != nil {
		return errors.WithStack(err)
	}
	if tool.IsFolderExisted(folder) {
		fmt.Printf("Task already exist, remove it first \n")
		folder, err = tool.GetFolderFrom(args[0])
		if err != nil {
			return errors.WithStack(err)
		}
		//如果文件路径已存在则删除
		if err := os.RemoveAll(folder); err != nil {
			return errors.WithStack(err)
		}
	}

	return executioner.Do(args[0], nil, conc)
}
