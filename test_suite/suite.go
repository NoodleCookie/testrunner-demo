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
	name   string
	cases  []Case
	option struct {
		variables map[string]any
	}
}

type description struct {
	Import []string          `yaml:"import,omitempty"`
	Vars   map[string]string `yaml:"vars,omitempty"`
}

func BuildTestSuite(dir string) (*Suite, error) {
	filename, ok := util.FileExistWithExtensionName(filepath.Join(dir, common.DescriptionFileName), common.SupportYamlExt...)
	if !ok {
		return nil, errors.New("not found description file, please create file with name description.yaml or description.yml")
	}

	base := filepath.Base(dir)
	suite := &Suite{name: base, cases: make([]Case, 0)}

	description := description{}
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

		//todo: modify current dir condition
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
		testcase.SetName(completedCaseFileName)
		suite.cases = append(suite.cases, *testcase)
	}

	suite.option.variables = make(map[string]any, 0)
	if description.Vars != nil {
		for key, value := range description.Vars {
			suite.option.variables[key] = value
		}
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

	description := &description{}
	if err := yaml.Unmarshal(readFile, description); err != nil {
		return false
	}
	return description.Import != nil
}

func (s *Suite) Execute() error {
	s.report()
	for _, testcase := range s.cases {

		s.deliverSuiteVariables(&testcase)

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

func (s *Suite) deliverSuiteVariables(c *Case) {
	if s.option.variables != nil {
		for key, value := range s.option.variables {
			c.SetVar(key, value)
		}
	}
}
