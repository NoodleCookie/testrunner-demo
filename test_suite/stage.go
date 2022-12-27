package test_suite

import (
	"bytes"
	"encoding/json"
	"html/template"
	"strings"
)

type Stage struct {
	Type      StageType  `yaml:"type,omitempty"`
	Name      string     `yaml:"name,omitempty"`
	Request   Request    `yaml:"request,omitempty"`
	Actual    *Actual    `yaml:"actual,omitempty"`
	Assertion *Assertion `yaml:"assert,omitempty"`
	option    struct {
		renderStage *Stage
		variables   map[string]any
	}
}

func (s *Stage) SetVar(key string, value any) {
	if s.option.variables == nil {
		s.option.variables = make(map[string]any, 0)
	}
	s.option.variables[key] = value
}

func (s *Stage) Var() map[string]any {
	return s.option.variables
}

func (s *Stage) GetRequest() Request {
	if s.option.renderStage != nil {
		return s.getRenderRequest()
	}
	return s.Request
}

func (s *Stage) getRenderRequest() Request {
	return s.option.renderStage.Request
}

func (s *Stage) Execute() error {
	s.render()
	if s.Type == API {
		return s.executeApi()
	}
	return nil
}

func (s *Stage) render() {
	marshal, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	temp, err := template.New(s.Name).Parse(string(marshal))
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer([]byte{})
	if err := temp.Execute(buffer, s.option.variables); err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(escapeHTMLChar(buffer.String())), &s.option.renderStage); err != nil {
		panic(err)
	}
}

func escapeHTMLChar(o string) string {
	return strings.Replace(o, "&#34;", "\\\"", -1)
}
