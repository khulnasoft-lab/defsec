package executor

import (
	"context"
	"testing"

	"github.com/aquasecurity/defsec/parsers/terraform"
	"github.com/aquasecurity/defsec/providers"
	"github.com/aquasecurity/defsec/severity"

	"github.com/aquasecurity/defsec/parsers/terraform/parser"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/test/testutil/filesystem"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var panicRule = rules.Rule{
	Provider:  providers.AWSProvider,
	Service:   "service",
	ShortCode: "abc",
	Severity:  severity.High,
	CustomChecks: rules.CustomChecks{
		Terraform: &rules.TerraformCustomCheck{
			RequiredTypes:  []string{"resource"},
			RequiredLabels: []string{"problem"},
			Check: func(resourceBlock *terraform.Block, _ *terraform.Module) (results rules.Results) {
				if resourceBlock.GetAttribute("panic").IsTrue() {
					panic("This is fine")
				}
				return
			},
		},
	},
}

func Test_PanicInCheckNotAllowed(t *testing.T) {

	reg := rules.Register(panicRule, nil)
	defer rules.Deregister(reg)

	fs, err := filesystem.New()
	require.NoError(t, err)
	defer func() { _ = fs.Close() }()

	require.NoError(t, fs.WriteTextFile("project/main.tf", `
resource "problem" "this" {
	panic = true
}
`))
	p := parser.New(parser.OptionStopOnHCLError(true))
	err = p.ParseDirectory(fs.RealPath("/project"))
	require.NoError(t, err)
	modules, _, err := p.EvaluateAll(context.TODO())
	require.NoError(t, err)
	results, _, _ := New().Execute(modules)
	assert.Equal(t, len(results.GetFailed()), 0)
}

func Test_PanicInCheckAllowed(t *testing.T) {

	reg := rules.Register(panicRule, nil)
	defer rules.Deregister(reg)

	fs, err := filesystem.New()
	require.NoError(t, err)
	defer func() { _ = fs.Close() }()

	require.NoError(t, fs.WriteTextFile("project/main.tf", `
resource "problem" "this" {
	panic = true
}
`))

	p := parser.New(parser.OptionStopOnHCLError(true))
	err = p.ParseDirectory(fs.RealPath("/project"))
	require.NoError(t, err)
	modules, _, err := p.EvaluateAll(context.TODO())
	require.NoError(t, err)
	_, _, err = New(OptionStopOnErrors(false)).Execute(modules)
	assert.Error(t, err)
}

func Test_PanicNotInCheckNotIncludePassed(t *testing.T) {

	reg := rules.Register(panicRule, nil)
	defer rules.Deregister(reg)

	fs, err := filesystem.New()
	require.NoError(t, err)
	defer func() { _ = fs.Close() }()

	require.NoError(t, fs.WriteTextFile("project/main.tf", `
resource "problem" "this" {
	panic = true
}
`))

	p := parser.New(parser.OptionStopOnHCLError(true))
	err = p.ParseDirectory(fs.RealPath("/project"))
	require.NoError(t, err)
	modules, _, err := p.EvaluateAll(context.TODO())
	require.NoError(t, err)
	results, _, _ := New().Execute(modules)
	assert.Equal(t, len(results.GetFailed()), 0)
}

func Test_PanicNotInCheckNotIncludePassedStopOnError(t *testing.T) {

	reg := rules.Register(panicRule, nil)
	defer rules.Deregister(reg)

	fs, err := filesystem.New()
	require.NoError(t, err)
	defer func() { _ = fs.Close() }()

	require.NoError(t, fs.WriteTextFile("project/main.tf", `
resource "problem" "this" {
	panic = true
}
`))

	p := parser.New(parser.OptionStopOnHCLError(true))
	err = p.ParseDirectory(fs.RealPath("/project"))
	require.NoError(t, err)
	modules, _, err := p.EvaluateAll(context.TODO())
	require.NoError(t, err)

	_, _, err = New(OptionStopOnErrors(false)).Execute(modules)
	assert.Error(t, err)
}
