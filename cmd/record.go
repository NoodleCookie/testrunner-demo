package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"testrunner/common"
	"testrunner/test_runner"
)

var recordCmd = &cobra.Command{
	Use:  "record",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_ = os.Setenv(common.PhaseEnv, string(common.Recording))
		testrunner := test_runner.Testrunner{}
		err := testrunner.Run(args[0])
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(recordCmd)
}
