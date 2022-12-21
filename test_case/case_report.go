package test_case

import "time"

type CaseReport struct {
	TestCase  string
	Results   []Result
	totalTime time.Duration
}
