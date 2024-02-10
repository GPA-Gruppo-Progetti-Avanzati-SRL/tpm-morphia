package pogs

import (
	"embed"
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/tpm-morphia-cli/cmds"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"text/template"
)

//go:embed templates/*
var templates embed.FS

var (
	targetFolder           string
	definitionFileName     string
	withBackupConflictMode bool
)

const semLogContextCmd = "gen-cmd-pogs::"

var genPogsCmd = &cobra.Command{
	Use:   "pogs",
	Short: "Generates Plain Old Golang Struct",
	Long:  `The command generates golang structs`,
	Run: func(cmd *cobra.Command, args []string) {
		const semLogContext = semLogContextCmd + "run"
		if err := validateArgs(); err != nil {
			log.Error().Err(err).Msg(semLogContext)
			return
		}

		log.Info().Str("target-folder", targetFolder).Msg(semLogContext)
		err := doWork()
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			return
		}
	},
}

func doWork() error {

	const semLogContext = semLogContextCmd + "do-work"

	/*
		for _, e := range entities {
			metadata := map[string]interface{}{
				"Name":    e.Name,
				"Version": cmds.Version,
				"UUID":    util.NewUUID(),
			}

			src, err := schematics.GetSource(
				templates, "templates",
				schematics.SourceWithFuncMap(getFuncMap()),
				schematics.SourceWithMetadata(metadata),
				schematics.SourceWithModel(&e),
			)
			if err != nil {
				log.Error().Err(err).Msg(semLogContext)
				return err
			}

			cm := schematics.ConflictModeOverwrite
			if withBackupConflictMode {
				cm = schematics.ConflictModeBackup
			}
			err = schematics.Apply(
				targetFolder,
				src,
				schematics.WithApplyDefaultConflictMode(cm), schematics.WithApplyProduceDiff())
			if err != nil {
				log.Error().Err(err).Msg(semLogContext)
				return err
			}
		}
	*/

	return nil
}

func validateArgs() error {
	const semLogContext = semLogContextCmd + "validate-args"

	var err error

	if definitionFileName == "" {
		err = errors.New("error: definition file name parameter must be provided")
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	if util.FileExists(definitionFileName) {
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			return err
		}
	} else {
		err = fmt.Errorf("error: the definition file %s cannot be found", definitionFileName)
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	if targetFolder == "" {
		err = errors.New("error: the output folder name parameter must be provided")
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	if !util.FileExists(targetFolder) {
		err = errors.New("error: out-dir parameter must point to a valid existing folder")
		log.Error().Err(err).Str("out-dir", targetFolder).Msg(semLogContext)
		return err
	}

	return nil
}

func init() {
	cmds.GenCmd.AddCommand(genPogsCmd)
	genPogsCmd.Flags().StringVarP(&definitionFileName, "def-file", "d", "", "the file containing structs definitions")
	genPogsCmd.Flags().StringVarP(&targetFolder, "out-dir", "o", "", "the name of target folder the structs will be created in")
	genPogsCmd.Flags().BoolVarP(&withBackupConflictMode, "with-bak", "k", false, "the current artifacts if ppresent will be backed-up and a diff file produced")
}

func getFuncMap() template.FuncMap {
	fMap := template.FuncMap{
		"camelize": func(s string) string {
			return util.Camelize(s)
		},
		"dasherize": func(s string) string {
			return util.Dasherize(s)
		},
		"classify": func(s string) string {
			return util.Classify(s)
		},
		"decamelize": func(s string) string {
			return util.Decamelize(s)
		},
		"underscore": func(s string) string {
			return util.Underscore(s)
		},
		"uuid": func() string {
			return util.NewUUID()
		},
	}

	return fMap
}
