package genAllEntities

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
	withBackupConflictMode bool
	withNewConflictMode    bool
	withFormatCode         bool
)

const semLogContextCmd = "cmd-gen-all-entities::"

var genAllEntitiesCmd = &cobra.Command{
	Use:   "all-entities",
	Short: "Generates all entities in the schema",
	Long:  `The command generates all the golang entities in the schema`,
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

	cm := schematics.ConflictModeOverwrite
	if withNewConflictMode {
		cm = schematics.ConflictModeNew
	} else if withBackupConflictMode {
		cm = schematics.ConflictModeBackup
	}

	sch, err := schema.ReadSchemaDefinitionFromFile(schemaFileName)
	if err != nil {
		log.Error().Err(err).Msg(semLogContext)
		return err
	}

	cfg := gomongodb.GeneratorConfig{
		Schema:               sch,
		TargetFolder:         targetFolder,
		Version:              cmds.Version,
		FormatCode:           withFormatCode,
		ConflictModeHandling: cm,
	}

	err = gomongodb.GenerateAllEntities(&cfg)
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
	cmds.GenCmd.AddCommand(genAllEntitiesCmd)
	genAllEntitiesCmd.Flags().StringVarP(&schemaFileName, "schema-file", "s", "", "the file containing structs definitions")
	genAllEntitiesCmd.Flags().StringVarP(&targetFolder, "out-dir", "o", "", "the name of target folder the structs will be created in")
	genAllEntitiesCmd.Flags().BoolVarP(&withBackupConflictMode, "with-bak", "k", false, "the current artifacts if present will be backed-up and a diff file produced")
	genAllEntitiesCmd.Flags().BoolVarP(&withNewConflictMode, "with-new", "w", false, "if the current artifacts is present the new one will be written to a .new file and a diff file produced")
	genAllEntitiesCmd.Flags().BoolVarP(&withFormatCode, "with-format", "f", false, "the go code will be formatted")
}
