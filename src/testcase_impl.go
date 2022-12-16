package src

import (
	"net/http"
)

type SimpleTestcase struct {
	Req http.Request  `yaml:"request" json:"request"`
	Res http.Response `yaml:"response" json:"response"`
}

func (s SimpleTestcase) Request() http.Request {
	return s.Req
}

func (s SimpleTestcase) Response() http.Response {
	return s.Res
}
