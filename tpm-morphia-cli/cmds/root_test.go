package cmds_test

import (
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/tpm-morphia-cli/cmds"
	_ "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/tpm-morphia-cli/cmds/pogs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const (
	TPMMorphiaCliVersion = "v0.0.1-SNAPSHOT"
)

var pogsCmdArgs = []string{
	"gen", "pogs", "--out-dir", "../../examples/ex001", "--with-bak",
}

func TestGenPOGSCmd(t *testing.T) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	cmds.RootCmd.SetArgs(pogsCmdArgs)
	cmds.Version = TPMMorphiaCliVersion
	err := cmds.RootCmd.Execute()
	require.NoError(t, err)
}
