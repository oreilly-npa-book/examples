package main

import (
	"fmt"

	// It's conventional to place third-party dependencies in a separate section,
	// to help distinguish them from standard-library and intra-module imports.
	//
	// You can also optionally specify an alias prior to the package path,
	// which may or may not be different from the original package name. In this
	// case, the `logrus` package is being aliased to `log`.
	log "github.com/sirupsen/logrus"
)

func main() {
	vlanIDs := []int{
		100, 200, 300,
	}

	log.Infof("Hello from logrus! There are %d VLANs in the vlanIDs slice.", len(vlanIDs))

	fmt.Println("End of program.")
}
