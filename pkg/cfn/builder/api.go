package builder

import (
	"fmt"
	"reflect"

	cfn "github.com/aws/aws-sdk-go/service/cloudformation"
	gfn "github.com/awslabs/goformation/cloudformation"
)

const (
	clusterTemplateDescription                = "EKS cluster"
	clusterTemplateDescriptionDefaultFeatures = " (with dedicated VPC & IAM role) "

	nodeGroupTemplateDescription                = "EKS nodes"
	nodeGroupTemplateDescriptionDefaultFeatures = " (Amazon Linux 2 with SSH) "

	templateDescriptionSuffix = " [created and managed by eksctl]"
)

// ResourceSet is an interface which cluster and nodegroup builders
// must implement
type ResourceSet interface {
	AddAllResources() error
	WithIAM() bool
	RenderJSON() ([]byte, error)
	GetAllOutputs(cfn.Stack) error
}

type resourceSet struct {
	template *gfn.Template
	outputs  []string
	withIAM  bool
}

func newResourceSet() *resourceSet {
	return &resourceSet{
		template: gfn.NewTemplate(),
	}
}

// makeName is syntactic sugar for {"Fn::Sub": "${AWS::Stack}-<name>"}
func makeName(suffix string) *gfn.Value {
	return gfn.MakeFnSubString(fmt.Sprintf("${%s}-%s", gfn.StackName, suffix))
}

// makeSlice makes a slice from a list of arguments
func makeSlice(i ...*gfn.Value) []*gfn.Value {
	return i
}

// makeSlice makes a slice from a list of string arguments
func makeStringSlice(s ...string) []*gfn.Value {
	slice := []*gfn.Value{}
	for _, i := range s {
		slice = append(slice, gfn.NewString(i))
	}
	return slice
}

func (r *resourceSet) newParameter(name, valueType, defaultValue string) *gfn.Value {
	p := map[string]string{"Type": valueType}
	if defaultValue != "" {
		p["Default"] = defaultValue
	}
	r.template.Parameters[name] = p
	return gfn.MakeRef(name)
}

func (r *resourceSet) newStringParameter(name, defaultValue string) *gfn.Value {
	return r.newParameter(name, "String", defaultValue)
}

func (r *resourceSet) newNumberParameter(name, defaultValue string) *gfn.Value {
	return r.newParameter(name, "Number", defaultValue)
}

// makeAutoNameTag create a new Name tag in the following format:
// {Key: "Name", Value: !Sub "${AWS::StackName}/<logicalResourceName>"}
func makeAutoNameTag(suffix string) gfn.Tag {
	return gfn.Tag{
		Key:   gfn.NewString("Name"),
		Value: gfn.MakeFnSubString(fmt.Sprintf("${%s}/%s", gfn.StackName, suffix)),
	}
}

// maybeSetNameTag adds a Name tag to any resource that supports tags
// it calls makeAutoNameTag to format the tag value
func maybeSetNameTag(name string, resource interface{}) {
	e := reflect.ValueOf(resource).Elem()
	if e.Kind() == reflect.Struct {
		f := e.FieldByName("Tags")
		if f.IsValid() && f.CanSet() {
			tag := reflect.ValueOf(makeAutoNameTag(name))
			if f.Type() == reflect.ValueOf([]gfn.Tag{}).Type() {
				f.Set(reflect.Append(f, tag))
			}
		}
	}
}

// newResource adds a resource, and adds Name tag if possible, it returns a reference
func (r *resourceSet) newResource(name string, resource interface{}) *gfn.Value {
	maybeSetNameTag(name, resource)
	r.template.Resources[name] = resource
	return gfn.MakeRef(name)
}

// renderJSON renders template as JSON
func (r *resourceSet) renderJSON() ([]byte, error) {
	return r.template.JSON()
}
