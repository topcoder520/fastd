package resume

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/topcoder520/fastd/internal/tool"

	"github.com/pkg/errors"
)

func TaskPrint() error {
	downloading, err := ioutil.ReadDir(filepath.Join(os.Getenv(tool.HomeName), tool.SaveFolder))
	if err != nil {
		return errors.WithStack(err)
	}

	folders := []string{}
	for _, d := range downloading {
		if d.IsDir() {
			folders = append(folders, d.Name())
		}
	}

	folderString := strings.Join(folders, "\n")
	fmt.Printf("Currently on going download: \n")
	fmt.Println(folderString)

	return nil
}

func Resume(task string) (*tool.State, error) {
	return tool.Read(task)
}
