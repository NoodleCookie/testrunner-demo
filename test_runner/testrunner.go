package test_runner

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
	"testrunner/common"
	"testrunner/test_case"
	"testrunner/util"
)

type Testrunner struct {
}

func (tr *Testrunner) Run(dir string) error {
	description, err := tr.CheckDescription(dir)
	if err != nil {
		return err
	}
	for _, caseName := range description.Import {
		caseFile := filepath.Join(dir, caseName)
		testcase, _ := util.FileExistWithExtensionName(caseFile, common.SupportYamlFileType...)
		err := tr.ExecuteCase(testcase)
		if err != nil {
			return err
		}
	}
	return nil
}

func (tr *Testrunner) CheckDescription(dir string) (*test_case.TestDescription, error) {
	filename, ok := util.FileExistWithExtensionName(filepath.Join(dir, common.DescriptionFileName), common.SupportYamlFileType...)
	if !ok {
		return nil, errors.New("not found description file, please create file with name description.yaml or description.yml")
	}

	description := test_case.TestDescription{}
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
		caseFile := filepath.Join(dir, caseName)
		if _, ok := util.FileExistWithExtensionName(caseFile, common.SupportYamlFileType...); !ok {
			return nil, errors.New(fmt.Sprintf("miss the testcase %s(.yaml/.yml)", caseFile))
		}
	}
	return &description, nil
}

func (tr *Testrunner) ExecuteCase(caseFile string) error {
	file, err := os.ReadFile(caseFile)
	if err != nil {
		return err
	}

	testcase := &test_case.Case{}
	if err := yaml.Unmarshal(file, testcase); err != nil {
		return err
	}

	return testcase.Execute()
}

func (tr *Testrunner) Exec(caseFile string) (report *test_case.SuiteReport, err error) {

	return nil, nil
}
