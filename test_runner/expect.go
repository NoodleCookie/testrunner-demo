package test_runner

type Expect struct {
	status   int    `yaml:"status,omitempty"`
	response string `yaml:"response,omitempty"`
}
