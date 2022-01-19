package monitor

import "github.com/aquasecurity/defsec/types"

type Monitor struct {
	types.Metadata
	LogProfiles []LogProfile
}

type LogProfile struct {
	types.Metadata
	RetentionPolicy RetentionPolicy
	Categories      []types.StringValue
	Locations       []types.StringValue
}

func (p LogProfile) GetMetadata() *types.Metadata {
	return &p.Metadata
}

func (p LogProfile) GetRawValue() interface{} {
	return nil
}

type RetentionPolicy struct {
	types.Metadata
	Enabled types.BoolValue
	Days    types.IntValue
}
