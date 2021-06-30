package errHandle

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

func ExitWithError(err error) {
	if err != nil {
		log.Printf("%v\n", errors.Cause(err))
		os.Exit(1)
	}
}
