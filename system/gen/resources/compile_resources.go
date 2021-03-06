package resources

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var packageTemplate = template.Must(template.New("").Funcs(map[string]interface{}{"conv": FormatByteSlice}).Parse(`// Code generated by go generate; DO NOT EDIT.
// generated using files from resources directory
// DO NOT COMMIT this file
package resources

func init(){
	{{- range $name, $file := . }}
    	resources.Add("{{ $name }}", []byte{ {{ conv $file }} })
	{{- end }}
}
`))

func FormatByteSlice(sl []byte) string {
	builder := strings.Builder{}
	for _, v := range sl {
		builder.WriteString(fmt.Sprintf("%d,", int(v)))
	}
	return builder.String()
}

func CompileResources(resourceDirectory string, outputFile string, mountPoint string) {
	resources := make(map[string][]byte)
	resDirUnix := filepath.ToSlash(resourceDirectory)
	err := filepath.Walk(resourceDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println("Error :", err)
			return err
		}

		pathUnix := filepath.ToSlash(path)
		relativePath := strings.TrimPrefix(pathUnix, resDirUnix)
		if info.IsDir() {
			log.Println(pathUnix, "is a directory, skipping...")
			return nil
		} else {
			log.Println(pathUnix, "is a file, baking in...")
			b, err := ioutil.ReadFile(path)
			if err != nil {
				log.Printf("Error reading %s: %s", path, err)
				return err
			}

			boxPath := mountPoint + relativePath
			log.Printf("Box path is %s", boxPath)

			resources[boxPath] = b
		}
		return nil
	})

	if err != nil {
		log.Fatal("Error walking through resources directory:", err)
	}

	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Error creating outputFile file:", err)
	}
	defer f.Close()

	builder := &bytes.Buffer{}

	err = packageTemplate.Execute(builder, resources)
	if err != nil {
		log.Fatal("Error executing template", err)
	}

	data, err := format.Source(builder.Bytes())
	if err != nil {
		log.Fatal("Error formatting generated code", err)
	}
	err = ioutil.WriteFile(outputFile, data, os.ModePerm)
	if err != nil {
		log.Fatal("Error writing outputFile file", err)
	}

	log.Println("Resources done...")
	log.Println("DO NOT COMMIT " + outputFile)
}
