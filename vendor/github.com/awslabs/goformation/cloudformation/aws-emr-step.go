package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEMRStep AWS CloudFormation Resource (AWS::EMR::Step)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html
type AWSEMRStep struct {

	// ActionOnFailure AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-elasticmapreduce-step-actiononfailure
	ActionOnFailure *Value `json:"ActionOnFailure,omitempty"`

	// HadoopJarStep AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-elasticmapreduce-step-hadoopjarstep
	HadoopJarStep *AWSEMRStep_HadoopJarStepConfig `json:"HadoopJarStep,omitempty"`

	// JobFlowId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-elasticmapreduce-step-jobflowid
	JobFlowId *Value `json:"JobFlowId,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-step.html#cfn-elasticmapreduce-step-name
	Name *Value `json:"Name,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRStep) AWSCloudFormationType() string {
	return "AWS::EMR::Step"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEMRStep) MarshalJSON() ([]byte, error) {
	type Properties AWSEMRStep
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSEMRStep) UnmarshalJSON(b []byte) error {
	type Properties AWSEMRStep
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSEMRStep(*res.Properties)
	}

	return nil
}

// GetAllAWSEMRStepResources retrieves all AWSEMRStep items from an AWS CloudFormation template
func (t *Template) GetAllAWSEMRStepResources() map[string]AWSEMRStep {
	results := map[string]AWSEMRStep{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEMRStep:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EMR::Step" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSEMRStep{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSEMRStepWithName retrieves all AWSEMRStep items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRStepWithName(name string) (AWSEMRStep, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEMRStep:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EMR::Step" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSEMRStep{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSEMRStep{}, errors.New("resource not found")
}
