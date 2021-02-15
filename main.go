//go:generate go run system/gen/version/cmd/gen.go -res-dir resources -txt-tpl system/gen/version/cmd/version.txt -out system/vinfo.go
//go:generate go run system/gen/resources/cmd/gen.go -res-dir resources -out system/resources/blob.go
package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/config"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/mongodb"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system/resources"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {

	if !resources.Has("/resources/version.txt") {
		_ = level.Info(system.GetLogger()).Log(system.DefaultLogMessageField, "Welcome To TPM-Morphia!", "goos", runtime.GOOS, "goarch", runtime.GOARCH)
		_ = level.Error(system.GetLogger()).Log(system.DefaultLogMessageField, "go generate not invoked during the build!")
		// os.Exit(-1)
	} else {
		versionInfo, _ := resources.Get("/resources/version.txt")
		_, _ = os.Stderr.WriteString(fmt.Sprintf("%s\n", versionInfo))
	}

	cfg, err := config.NewBuilder().Build(context.Background())
	if err != nil {
		_ = level.Error(system.GetLogger()).Log(system.DefaultLogMessageField, err.Error())
		os.Exit(-1)
	}

	cDefs, err := cfg.FindCollectionToProcess(system.GetLogger())
	if err != nil {
		_ = level.Error(system.GetLogger()).Log(system.DefaultLogMessageField, err.Error())
		os.Exit(-1)
	}

	for _, collDefFile := range cDefs {
		if schemaFile, err := ioutil.ReadFile(collDefFile); err != nil {
			_ = level.Error(system.GetLogger()).Log(system.DefaultLogMessageField, err.Error())
			return
		} else {
			r := bytes.NewReader(schemaFile)
			cDef, err := schema.ReadCollectionDefinition(system.GetLogger(), r)
			if err != nil {
				_ = level.Error(system.GetLogger()).Log(system.DefaultLogMessageField, err.Error())
				return
			}

			if genDriver, err := mongodb.NewCodeGenCollection(cDef); err != nil {
				_ = level.Error(system.GetLogger()).Log(system.DefaultLogMessageField, err.Error())
				return
			} else {
				if err := mongodb.Generate(system.GetLogger(), cfg, genDriver); err != nil {
					_ = level.Error(system.GetLogger()).Log(system.DefaultLogMessageField, err.Error())
					return
				}
			}
		}
	}
}
