package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   os.Args[0],
	Short: "file downloader written in Go",
}
