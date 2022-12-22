package test_runner

import (
	"io/fs"
	"path/filepath"
	"testrunner/common"
	"testrunner/test_report"
	"testrunner/test_suite"
)

type Testrunner struct {
}

func (tr *Testrunner) Run(root string) error {
	if common.CurrentPhase() == common.Asserting {
		test_report.BuildReport()
	}
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() || !test_suite.IsTestSuite(path) {
			return nil
		}

		suite, err := test_suite.BuildTestSuite(path)
		if err != nil {
			return err
		}
		return suite.Execute()
	})
}
