package cmds_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/tpm-morphia-cli/cmds"
	_ "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/tpm-morphia-cli/cmds/genAllEntities"
	_ "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/tpm-morphia-cli/cmds/genEntity"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const (
	TPMMorphiaCliVersion = "v0.0.1-SNAPSHOT"
)

var genEntityCmdArgs = []string{
	"gen", "entity", "--out-dir", "../..", "--schema-file", "../../examples/example7/schema/schema.yml", "--name", "author", "--with-format",
}

func TestGenEntityCmd(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	cmds.RootCmd.SetArgs(genEntityCmdArgs)
	cmds.Version = TPMMorphiaCliVersion
	err := cmds.RootCmd.Execute()
	require.NoError(t, err)
}

var genAllEntitiesCmdArgs = []string{
	"gen", "all-entities", "--out-dir", "../..", "--schema-file", "../../examples/example7/schema/schema.yml", "--with-format",
}

func TestGenAllEntitiesCmd(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	cmds.RootCmd.SetArgs(genAllEntitiesCmdArgs)
	cmds.Version = TPMMorphiaCliVersion
	err := cmds.RootCmd.Execute()
	require.NoError(t, err)
}
