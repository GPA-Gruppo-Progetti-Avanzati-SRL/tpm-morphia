package main

import (
	"flag"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system/gen/version"
	"log"
	"os"
)

var resourceDir = flag.String("res-dir", "resources", "Resource Directory")
var versionTxtTemplate = flag.String("txt-tpl", "version.txt", "Nome file di template versione testuale")
var outFile = flag.String("out", "vinfo.go", "Output File")

const VersionFileName = "version.txt"

func main() {

	log.Println("-----------------------------")
	log.Println("TPM-Morphia Command: Version")
	flag.Parse()

	if dir, err1 := os.Getwd(); err1 != nil {
		log.Fatal(err1)
	} else {
		log.Println("Current Directory...........: " + dir)
	}

	log.Println("Resource Directory: " + *resourceDir)
	log.Println("Version Info Go Source FileName: " + *outFile)
	log.Println("Version Text Template FileName: " + *versionTxtTemplate)
	log.Println("Version Text FileName: " + VersionFileName)

	version.UpdateVersionInfo("VERSION", *resourceDir, *outFile, *versionTxtTemplate, VersionFileName)
}
