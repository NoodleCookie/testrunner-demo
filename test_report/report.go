package test_report

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"testrunner/common"
	"testrunner/util"
)

var report Report

func GetReport() *Report {
	return &report
}

func BuildReport() *Report {
	report = Report{}
	return &report
}

type stageReport struct {
	Pass   bool        `json:"pass"`
	Name   string      `json:"name"`
	Detail interface{} `json:"detail,omitempty"`
}

type caseReport struct {
	Pass   bool           `json:"pass"`
	Name   string         `json:"name"`
	Stages []*stageReport `json:"stages,omitempty"`
}

type suiteReport struct {
	Pass  bool          `json:"pass"`
	Name  string        `json:"name"`
	Cases []*caseReport `json:"cases,omitempty"`
}

type Report struct {
	Pass   bool           `json:"pass"`
	Suites []*suiteReport `json:"suites,omitempty"`
}

func (r *Report) AppendSuite(name string) {
	if r.Suites == nil {
		r.Suites = make([]*suiteReport, 0)
	}
	r.Suites = append(r.Suites, &suiteReport{Name: name})
}

func (r *Report) getLastSuite() *suiteReport {
	return r.Suites[len(r.Suites)-1]
}

func (r *Report) AppendCase(name string) {
	suite := r.getLastSuite()
	if r.getLastSuite().Cases == nil {
		suite.Cases = make([]*caseReport, 0)
	}
	suite.Cases = append(suite.Cases, &caseReport{Name: name})
}

func (r *Report) getLastCase() *caseReport {
	suite := r.getLastSuite()
	return suite.Cases[len(suite.Cases)-1]
}

func (r *Report) AppendStage(name string, pass bool, detail interface{}) {
	lastCase := r.getLastCase()
	if lastCase.Stages == nil {
		lastCase.Stages = make([]*stageReport, 0)
	}
	lastCase.Stages = append(lastCase.Stages, &stageReport{Pass: pass, Name: name, Detail: detail})
}

func (cr *caseReport) integrate() {
	cr.Pass = true
	for _, stage := range cr.Stages {
		if !stage.Pass {
			cr.Pass = false
			break
		}
	}
}

func (sr *suiteReport) integrate() {
	sr.Pass = true
	for _, testcase := range sr.Cases {
		testcase.integrate()
		if !testcase.Pass {
			sr.Pass = false
			break
		}
	}
}

func (r *Report) integrate() {
	r.Pass = true
	for _, suite := range r.Suites {
		suite.integrate()
		if !suite.Pass {
			r.Pass = false
			break
		}
	}
}

func (r *Report) Gen(path string) (string, error) {
	if !util.FileOrDirExist(path) {
		if err := os.MkdirAll(path, 0766); err != nil {
			return "", errors.Wrap(err, "failed to generate report")
		}
	}

	r.integrate()

	marshal, err := json.Marshal(r)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal report")
	}

	reportPath := filepath.Join(path, fmt.Sprintf("%s.%s", common.ReportFile, common.ReportFileExt))
	if err := os.WriteFile(reportPath, marshal, 0766); err != nil {
		return "", errors.Wrap(err, "failed to write report")
	}

	return reportPath, nil
}
