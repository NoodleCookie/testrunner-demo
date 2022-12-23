package main

import "testrunner/cmd"

func main() {
	//command := cobra.Command{
	//	Args: cobra.ExactArgs(0),
	//	Run: func(cmd *cobra.Command, args []string) {
	//		fmt.Println("yes")
	//		testrunner := test_runner.Testrunner{}
	//		testrunner.Run("/Users/yandi.lin/Documents/test_mesh/testrunner-demo/test_runner/unit_test/data/testsuite_correct_baidu_request")
	//		report := test_report.GetReport()
	//		report.Gen("./gen")
	//	},
	//}
	cmd.Execute()
}
