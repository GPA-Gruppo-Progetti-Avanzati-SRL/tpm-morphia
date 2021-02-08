//go:generate go run system/gen/version/cmd/gen.go -res-dir resources -txt-tpl system/gen/version/cmd/version.txt -out system/vinfo.go
//go:generate go run system/gen/resources/cmd/gen.go -res-dir resources -out system/resources/blob.go
package main

import (
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system/resources"
	"github.com/go-kit/kit/log/level"
	"os"
	"runtime"
)

func main() {

	if !resources.Has("/resources/version.txt") {
		_ = level.Info(system.GetLogger()).Log(system.DefaultLogMessageField, "Welcome To TPM-Exporter!", "goos", runtime.GOOS, "goarch", runtime.GOARCH)
		_ = level.Error(system.GetLogger()).Log(system.DefaultLogMessageField, "go generate not invoked during the build!")
		// os.Exit(-1)
	} else {
		versionInfo, _ := resources.Get("/resources/version.txt")
		os.Stderr.WriteString(fmt.Sprintf("%s\n", versionInfo))
	}

}
