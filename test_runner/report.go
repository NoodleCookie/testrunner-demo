package test_runner

import "time"

type Report struct {
	testCase  string
	results   []Result
	totalTime time.Duration
}
