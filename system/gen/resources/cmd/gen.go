package main

import (
	"flag"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system/gen/resources"
	"log"
	"os"
)

const blob = "blob.go"
const resourcesPath = "resources"

var outfile = flag.String("out", "blob.go", "Output File")
var resourcedir = flag.String("res-dir", "resources", "Resource Directory")

func main() {
	log.Println("-------------------------------")
	log.Println("TPM-Morphia Command: Resources")
	flag.Parse()

	dir, err1 := os.Getwd()
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Println("Current Directory Is: " + dir)

	log.Println("Resource Directory: " + *resourcedir)
	log.Println("Output File: " + *outfile)

	if _, err := os.Stat(*resourcedir); os.IsNotExist(err) {
		log.Fatal("Resources directory does not exists")
	}

	resources.CompileResources(*resourcedir, *outfile, "/resources")
}
