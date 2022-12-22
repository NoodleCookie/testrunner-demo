package test_suite

import "time"

type CaseReport struct {
	TestCase  string
	Results   []Result
	totalTime time.Duration
}
