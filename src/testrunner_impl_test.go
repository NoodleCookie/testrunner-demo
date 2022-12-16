package src

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestSimpleTestrunner_Resolve(t *testing.T) {
	// given
	testcaseFilePath := filepath.Join(".", "test", "testcase.yaml")

	// when
	testrunner := SimpleTestrunner{}
	testcase, err := testrunner.Resolve(testcaseFilePath)

	// then
	if err != nil {
		t.Fatalf("err %s", err.Error())
	}

	if testcase.Request().URL == nil {
		t.Fatalf("err: request url cannot be nil")
	}

	fmt.Println(testcase.Request().URL)
}

func TestSimpleTestrunner_Send(t *testing.T) {
	// given
	testcaseFilePath := filepath.Join(".", "test", "testcase.yaml")

	// when
	testrunner := SimpleTestrunner{}
	testcase, _ := testrunner.Resolve(testcaseFilePath)
	result, err := testrunner.Send(testcase)

	// then
	if err != nil {
		t.Fatalf("err %s", err.Error())
	}

	fmt.Println(result.Response())
}

func TestSimpleTestrunner_Verify(t *testing.T) {
	// given
	testcaseFilePath := filepath.Join(".", "test", "testcase.yaml")

	// when
	testrunner := SimpleTestrunner{}
	testcase, _ := testrunner.Resolve(testcaseFilePath)
	result, _ := testrunner.Send(testcase)
	verify, _, err := testrunner.Verify(result.Response(), testcase)

	// then
	if !verify {
		t.Fatalf("verify failed %s", err.Error())
	}
}
