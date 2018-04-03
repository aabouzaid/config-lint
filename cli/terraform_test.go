package main

import (
	"testing"
)

func TestTerraformLinter(t *testing.T) {
	emptyTags := []string{}
	emptyIds := []string{}
	filenames := []string{"./testdata/resources/terraform_instance.tf"}
	linter := TerraformLinter{Filenames: filenames, Log: testLogger}
	ruleSet := loadRulesForTest("./testdata/rules/terraform_instance.yml", t)
	report, err := linter.Validate(ruleSet, emptyTags, emptyIds)
	if err != nil {
		t.Error("Expecting TestTerraformLinter to not return an error")
	}
	if len(report.ResourcesScanned) != 1 {
		t.Errorf("TestTerraformLinter scanned %d resources, expecting 1", len(report.ResourcesScanned))
	}
	if len(report.FilesScanned) != 1 {
		t.Errorf("TestTerraformLinter scanned %d files, expecting 1", len(report.FilesScanned))
	}
	if len(report.Violations) != 0 {
		t.Errorf("TestTerraformLinter returned %d violations, expecting 0", len(report.Violations))
	}
}
