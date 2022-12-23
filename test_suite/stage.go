package test_suite

type Stage struct {
	Type      StageType `yaml:"type,omitempty"`
	Name      string    `yaml:"name,omitempty"`
	Request   Request   `yaml:"request,omitempty"`
	Assertion Assertion `yaml:"assert,omitempty"`
}

func (s *Stage) Execute() error {
	if s.Type == API {
		return s.executeApi()
	}
	return nil
}
