package common

import "os"

type TestrunnerPhase string

const (
	PhaseEnv            = "PHASE"
	DescriptionFileName = "description"
	Recording           = TestrunnerPhase("record")
	Asserting           = TestrunnerPhase("assert")
	ReportFile          = "testrunner-report"
	ReportFileExt       = "json"
)

var (
	SupportYamlExt = []string{"yaml", "yml"}
)

func CurrentPhase() TestrunnerPhase {
	phase := os.Getenv(PhaseEnv)
	if phase == "" || TestrunnerPhase(phase) == Asserting {
		return Asserting
	}
	if TestrunnerPhase(phase) == Recording {
		return Recording
	}
	return ""
}
