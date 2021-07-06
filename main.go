package main

import (
	"fmt"
	"os"

	"github.com/topcoder520/fastd/cmd"
)

//https://github.com/Imputes/fdlr的练习项目
func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
