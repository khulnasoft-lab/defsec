package rules

import (
	"github.com/khulnasoft-lab/defsec/internal/rules"
	"github.com/khulnasoft-lab/defsec/pkg/framework"
	"github.com/khulnasoft-lab/defsec/pkg/scan"
	ruleTypes "github.com/khulnasoft-lab/defsec/pkg/types/rules"
)

func Register(rule scan.Rule) ruleTypes.RegisteredRule {
	return rules.Register(rule)
}

func Deregister(rule ruleTypes.RegisteredRule) {
	rules.Deregister(rule)
}

func GetRegistered(fw ...framework.Framework) []ruleTypes.RegisteredRule {
	return rules.GetFrameworkRules(fw...)
}

func GetSpecRules(spec string) []ruleTypes.RegisteredRule {
	return rules.GetSpecRules(spec)
}
