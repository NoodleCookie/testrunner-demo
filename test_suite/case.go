package test_suite

import (
	"testrunner/common"
	"testrunner/test_report"
	"time"
)

type Case struct {
	name   string
	Stages []Stage           `yaml:"stages,omitempty"`
	Vars   map[string]string `yaml:"vars,omitempty"`

	option struct {
		filters   []func(stage Stage) bool
		variables map[string]any
	}
}

func (c *Case) SetVar(key string, value any) {
	if c.option.variables == nil {
		c.option.variables = make(map[string]any, 0)
	}
	c.option.variables[key] = value
}

func (c *Case) Var() map[string]any {
	return c.option.variables
}

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

			c.classifyStageVariables(&stage)

			if stage.Execute() != nil {
				return stage.Execute()
			}

			c.classifyStageResponse(&stage)

		}
		duration := time.Since(start)
		test_report.GetReport().SetCaseRunTime(duration)
	}
	return nil
}

func (c *Case) report() {
	if common.CurrentPhase() == common.Asserting {
		report := test_report.GetReport()
		report.AppendCase(c.name)
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

func (c *Case) classifyStageResponse(stage *Stage) {
	if stage.Type != API {
		return
	}

	//c.setVar(fmt.Sprintf("%s.actual.status", c.name), stage.Actual.Status)
	//
	//for key, value := range stage.Actual.Headers {
	//	c.setVar(fmt.Sprintf("%s.actual.headers.%s", c.name, key), value)
	//}
	//
	//c.setVar(fmt.Sprintf("%s.actual.body", c.name), stage.Actual.Body)
}

func (c *Case) classifyStageVariables(s *Stage) {
	if c.option.variables != nil {
		for key, value := range c.option.variables {
			s.SetVar(key, value)
		}
	}

	if c.Vars != nil {
		for key, value := range c.Vars {
			s.SetVar(key, value)
		}
	}
}
