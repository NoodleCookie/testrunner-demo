package cmd

import (
	"github.com/spf13/cobra"
	"testrunner/test_report"
	"testrunner/test_runner"
)

var execCmd = &cobra.Command{
	Use: "exec",
	Run: func(cmd *cobra.Command, args []string) {
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
	rootCmd.AddCommand(execCmd)
}
