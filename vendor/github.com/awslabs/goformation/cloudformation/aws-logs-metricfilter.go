package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSLogsMetricFilter AWS CloudFormation Resource (AWS::Logs::MetricFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html
type AWSLogsMetricFilter struct {

	// FilterPattern AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-filterpattern
	FilterPattern *Value `json:"FilterPattern,omitempty"`

	// LogGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-loggroupname
	LogGroupName *Value `json:"LogGroupName,omitempty"`

	// MetricTransformations AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-metrictransformations
	MetricTransformations []AWSLogsMetricFilter_MetricTransformation `json:"MetricTransformations,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsMetricFilter) AWSCloudFormationType() string {
	return "AWS::Logs::MetricFilter"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSLogsMetricFilter) MarshalJSON() ([]byte, error) {
	type Properties AWSLogsMetricFilter
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
func (r *AWSLogsMetricFilter) UnmarshalJSON(b []byte) error {
	type Properties AWSLogsMetricFilter
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
		*r = AWSLogsMetricFilter(*res.Properties)
	}

	return nil
}

// GetAllAWSLogsMetricFilterResources retrieves all AWSLogsMetricFilter items from an AWS CloudFormation template
func (t *Template) GetAllAWSLogsMetricFilterResources() map[string]AWSLogsMetricFilter {
	results := map[string]AWSLogsMetricFilter{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSLogsMetricFilter:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Logs::MetricFilter" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSLogsMetricFilter{}
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

// GetAWSLogsMetricFilterWithName retrieves all AWSLogsMetricFilter items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsMetricFilterWithName(name string) (AWSLogsMetricFilter, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSLogsMetricFilter:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Logs::MetricFilter" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSLogsMetricFilter{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSLogsMetricFilter{}, errors.New("resource not found")
}
