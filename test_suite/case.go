package test_suite

import (
	"testrunner/common"
	"testrunner/test_report"
	"time"
)

type Case struct {
	Name   string  `yaml:"name,omitempty"`
	Stages []Stage `yaml:"stages,omitempty"`
	option struct {
		filters []func(stage Stage) bool
	}
}

var report CaseReport

func (c *Case) AddFilter(filter func(stage Stage) bool) *Case {
	if c.option.filters == nil {
		c.option.filters = make([]func(stage Stage) bool, 0)
	}
	c.option.filters = append(c.option.filters, filter)
	return c
}

func (c *Case) Execute() error {
	c.report()
	if c.Stages != nil {
		start := time.Now()
		for _, stage := range c.filter() {
			err := stage.Execute()
			if err != nil {
				return err
			}
		}
		duration := time.Since(start)
		report.totalTime = duration
	}
	return nil
}

func (c *Case) report() {
	if common.CurrentPhase() == common.Asserting {
		report := test_report.GetReport()
		report.AppendCase(c.Name)
	}
}

func (c *Case) filter() []Stage {

	if c.option.filters == nil {
		return c.Stages
	}

	if c.Stages == nil {
		return []Stage{}
	}

	var result []Stage
	for _, stage := range c.Stages {
		for _, filter := range c.option.filters {
			if filter(stage) {
				result = append(result, stage)
				break
			}
		}
	}
	return result
}
