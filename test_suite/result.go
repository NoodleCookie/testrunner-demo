package test_suite

type Result struct {
	request  Request
	response Response
	compare  Compare
	Pass     bool
}
