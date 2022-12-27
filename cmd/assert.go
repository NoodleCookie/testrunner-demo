package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"testrunner/common"
	"testrunner/test_report"
	"testrunner/test_runner"
)

var assertCmd = &cobra.Command{
	Use:  "assert",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_ = os.Setenv(common.PhaseEnv, string(common.Asserting))
		testrunner := test_runner.Testrunner{}
		err := testrunner.Run(args[0])
		if err != nil {
			panic(err)
		}
		_, err = test_report.GetReport().Gen(".")
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(assertCmd)
}
