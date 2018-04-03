package main

import (
	"fmt"
	"github.com/stelligent/config-lint/assertion"
)

// FIXME can the filenames be taken out of the Linter interface, and provided to the constructor or Linters that need them?

// Linter provides the interface for all supported linters
type Linter interface {
	Validate(ruleSet assertion.RuleSet, tags []string, ruleIDs []string) (assertion.ValidationReport, error)
	Search(ruleSet assertion.RuleSet, searchExpression string)
}

// ResourceLoader provides the interface that a Linter needs to load a collection of Resource objects
type ResourceLoader interface {
	Load(filename string) ([]assertion.Resource, error)
}

func makeLinter(linterType string, args []string, log assertion.LoggingFunction) Linter {
	switch linterType {
	case "Kubernetes":
		return KubernetesLinter{Filenames: args, Log: log}
	case "Terraform":
		return TerraformLinter{Filenames: args, Log: log}
	case "SecurityGroup":
		return AWSResourceLinter{Loader: SecurityGroupLoader{}, Log: log}
	case "IAMUser":
		return AWSResourceLinter{Loader: IAMUserLoader{}, Log: log}
	case "LintRules":
		return RulesLinter{Filenames: args, Log: log}
	case "YAML":
		return YAMLLinter{Filenames: args, Log: log}
	default:
		fmt.Printf("Type not supported: %s\n", linterType)
		return nil
	}
}
