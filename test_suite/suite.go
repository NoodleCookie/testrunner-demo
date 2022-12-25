package test_suite

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
	"testrunner/common"
	"testrunner/test_report"
	"testrunner/util"
)

type Suite struct {
	name  string
	cases []Case
}

func (s *Suite) SetVar(key string, value any) {
	//TODO implement me
	panic("implement me")
}

func (s *Suite) Var() map[string]any {
	//TODO implement me
	panic("implement me")
}

func BuildTestSuite(dir string) (*Suite, error) {
	filename, ok := util.FileExistWithExtensionName(filepath.Join(dir, common.DescriptionFileName), common.SupportYamlExt...)
	if !ok {
		return nil, errors.New("not found description file, please create file with name description.yaml or description.yml")
	}

	base := filepath.Base(dir)
	suite := &Suite{name: base, cases: make([]Case, 0)}

	description := Description{}
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("check your %s has correct privilege", filename))
	}

	err = yaml.Unmarshal(file, &description)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("check your %s is legal", filename))
	}

	if description.Import == nil || len(description.Import) == 0 {
		return nil, errors.New("you must import your testcase into description")
	}

	for _, caseName := range description.Import {

		if strings.HasPrefix(caseName, "../") {
			return nil, errors.New(fmt.Sprintf("your imported testcase must be created in current dir %s", dir))
		}

		caseName := filepath.Join(dir, caseName)
		completedCaseFileName, ok := util.FileExistWithExtensionName(caseName, common.SupportYamlExt...)
		if !ok {
			return nil, errors.New(fmt.Sprintf("miss the testcase %s(.yaml/.yml)", caseName))
		}

		testcase := &Case{}
		readFile, err := os.ReadFile(completedCaseFileName)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to read testcase file %s", readFile))
		}

		if err := yaml.Unmarshal(readFile, testcase); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("failed to unmarsahl testcase file %s", readFile))
		}
		testcase.name = caseName
		suite.cases = append(suite.cases, *testcase)
	}
	return suite, nil
}

func IsTestSuite(dir string) bool {
	file, ok := util.FileExistWithExtensionName(filepath.Join(dir, common.DescriptionFileName), common.SupportYamlExt...)
	if !ok {
		return false
	}

	readFile, err := os.ReadFile(file)
	if err != nil {
		return false
	}

	description := &Description{}
	if err := yaml.Unmarshal(readFile, description); err != nil {
		return false
	}
	return description.Import != nil
}

func (s *Suite) Execute() error {
	s.report()
	for _, testcase := range s.cases {
		if err := testcase.Execute(); err != nil {
			return err
		}
	}
	return nil
}

func (s *Suite) report() {
	if common.CurrentPhase() == common.Asserting {
		report := test_report.GetReport()
		report.AppendSuite(s.name)
	}
}
