package genEntity

import (
	"errors"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-common/util/fileutil"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/gen/gomongodb"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/schema"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/tpm-morphia-cli/cmds"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-schematics/schematics"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	targetFolder           string
	schemaFileName         string
	entityName             string
	withBackupConflictMode bool
	withNewConflictMode    bool
	withFormatCode         bool
)

const semLogContextCmd = "cmd-gen-entity::"

var genEntityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Generates entity",
	Long:  `The command generates golang entity`,
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

	sch, err := schema.ReadSchemaDefinitionFromFile(schemaFileName)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	cm := schematics.ConflictModeOverwrite
	if withNewConflictMode {
		cm = schematics.ConflictModeNew
	} else if withBackupConflictMode {
		cm = schematics.ConflictModeBackup
	}

	cfg := gomongodb.GeneratorConfig{
		Schema:               sch,
		TargetFolder:         targetFolder,
		EntityName:           entityName,
		Version:              cmds.Version,
		FormatCode:           withFormatCode,
		ConflictModeHandling: cm,
	}

	err = gomongodb.GenerateEntity(&cfg)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	return nil
}

func validateArgs() error {
	const semLogContext = semLogContextCmd + "validate-args"

	var err error

	if schemaFileName == "" {
		err = errors.New("error: definition file name parameter must be provided")
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	if fileutil.FileExists(schemaFileName) {
		if err != nil {
			log.Error().Err(err).Msg(semLogContext)
			return err
		}
	} else {
		err = fmt.Errorf("error: the definition file %s cannot be found", schemaFileName)
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	if targetFolder == "" {
		err = errors.New("error: the output folder name parameter must be provided")
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	if !fileutil.FileExists(targetFolder) {
		err = errors.New("error: out-dir parameter must point to a valid existing folder")
		log.Error().Err(err).Str("out-dir", targetFolder).Msg(semLogContext)
		return err
	}

	return nil
}

func init() {
	cmds.GenCmd.AddCommand(genEntityCmd)
	genEntityCmd.Flags().StringVarP(&schemaFileName, "schema-file", "s", "", "the file containing structs definitions")
	genEntityCmd.Flags().StringVarP(&targetFolder, "out-dir", "o", "", "the name of target folder the structs will be created in")
	genEntityCmd.Flags().StringVarP(&entityName, "name", "n", "", "the name of the structs that will be processed")
	genEntityCmd.Flags().BoolVarP(&withBackupConflictMode, "with-bak", "k", false, "the current artifacts if present will be backed-up and a diff file produced")
	genEntityCmd.Flags().BoolVarP(&withNewConflictMode, "with-new", "w", false, "if the current artifacts is present the new one will be written to a .new file and a diff file produced")

	genEntityCmd.Flags().BoolVarP(&withFormatCode, "with-format", "f", false, "the go code will be formatted")
}
