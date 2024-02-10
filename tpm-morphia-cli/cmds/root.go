package cmds

import (
	"github.com/spf13/cobra"
	"os"
)

var Version = ""

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tpm-morphia-cli",
	Short: "command line tool for mongodb persistencce golang code generatiion",
	Long:  `The command line supports the generation of some artifacts.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var GenCmd = &cobra.Command{
	Use:   "gen",
	Short: "command for generating MongoDb artifacts for the golang language",
	Long:  `The command line supports the generation of some artifacts targeted at the MongoDb persistence in Golang language.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(GenCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tpm-iso20022-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
