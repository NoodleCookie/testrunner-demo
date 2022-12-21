package test_case

type Case struct {
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
	if c.Stages != nil {
		for _, stage := range c.filter() {
			err := stage.Execute()
			if err != nil {
				return err
			}
		}
	}
	return nil
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
