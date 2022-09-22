//go:generate go run system/gen/version/cmd/gen.go -res-dir resources -txt-tpl system/gen/version/cmd/version.txt -out system/vinfo.go
//go:generate go run system/gen/resources/cmd/gen.go -res-dir resources -out system/resources/blob.go
package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/config"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/mongodb"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/rs/zerolog/log"

	"io/ioutil"
	"os"
)

//go:embed VERSION
var version string

// appLogo contains the ASCII splash screen
//
//go:embed app-logo.txt
var appLogo []byte

func main() {

	fmt.Println(string(appLogo))
	fmt.Printf("Version: %s\n", version)

	cfg, err := config.NewBuilder().Build(context.Background())
	if err != nil {
		log.Error().Err(err).Send()
		os.Exit(-1)
	}

	cDefs, err := cfg.FindCollectionToProcess()
	if err != nil {
		log.Error().Err(err).Send()
		os.Exit(-1)
	}

	for _, collDefFile := range cDefs {
		if schemaFile, err := ioutil.ReadFile(collDefFile); err != nil {
			log.Error().Err(err).Send()
			return
		} else {
			r := bytes.NewReader(schemaFile)
			cDef, err := schema.ReadCollectionDefinition(r)
			if err != nil {
				log.Error().Err(err).Send()
				return
			}

			if genDriver, err := mongodb.NewCodeGenCollection(cDef); err != nil {
				log.Error().Err(err).Send()
				return
			} else {
				if err := mongodb.Generate(cfg, genDriver); err != nil {
					log.Error().Err(err).Send()
					return
				}
			}
		}
	}
}
