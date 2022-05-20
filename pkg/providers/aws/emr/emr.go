package emr

import (
	"github.com/aquasecurity/defsec/internal/types"
)

// resource "aws_emr_security_configuration" "foo" {
// 	name = "emrsc_other"

// 	configuration = <<EOF
//   {
// 	"EncryptionConfiguration": {
// 	  "AtRestEncryptionConfiguration": {
// 		"S3EncryptionConfiguration": {
// 		  "EncryptionMode": "SSE-S3"
// 		},
// 		"LocalDiskEncryptionConfiguration": {
// 		  "EncryptionKeyProviderType": "AwsKms",
// 		  "AwsKmsKey": "arn:aws:kms:us-west-2:187416307283:alias/tf_emr_test_key"
// 		}
// 	  },
// 	  "EnableInTransitEncryption": false,
// 	  "EnableAtRestEncryption": true
// 	}
//   }
//   EOF
//   }

// type conf struct {
// 	EncryptionConfiguration struct {
// 		AtRestEncryptionConfiguration struct {
// 			S3EncryptionConfiguration struct {
// 				EncryptionMode string `json:"EncryptionMode"`
// 			} `json:"S3EncryptionConfiguration"`
// 			LocalDiskEncryptionConfiguration struct {
// 				EncryptionKeyProviderType string `json:"EncryptionKeyProviderType"`
// 				AwsKmsKey                 string `json:"AwsKmsKey"`
// 			} `json:"LocalDiskEncryptionConfiguration"`
// 		} `json:"AtRestEncryptionConfiguration"`
// 		EnableInTransitEncryption bool `json:"EnableInTransitEncryption"`
// 		EnableAtRestEncryption    bool `json:"EnableAtRestEncryption"`
// 	} `json:"EncryptionConfiguration"`
// }

type EMR struct {
	Clusters              []Cluster
	SecurityConfiguration []SecurityConfiguration
}

type Cluster struct {
	types.Metadata
	Settings ClusterSettings
}

type ClusterSettings struct {
	types.Metadata
	Name         types.StringValue
	ReleaseLabel types.StringValue
	ServiceRole  types.StringValue
}

type SecurityConfiguration struct {
	types.Metadata
	Name types.StringValue
	// TODO: unmarshal the value?
	Configuration types.StringValue
}
