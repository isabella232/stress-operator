// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.Stress":              schema_pkg_apis_thelastpickle_v1alpha1_Stress(ref),
		"github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressContext":       schema_pkg_apis_thelastpickle_v1alpha1_StressContext(ref),
		"github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressContextSpec":   schema_pkg_apis_thelastpickle_v1alpha1_StressContextSpec(ref),
		"github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressContextStatus": schema_pkg_apis_thelastpickle_v1alpha1_StressContextStatus(ref),
		"github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressSpec":          schema_pkg_apis_thelastpickle_v1alpha1_StressSpec(ref),
		"github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressStatus":        schema_pkg_apis_thelastpickle_v1alpha1_StressStatus(ref),
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_Stress(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Stress is the Schema for the Stresses API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressSpec", "github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_StressContext(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StressContext is the Schema for the stresscontexts API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressContextSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressContextStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressContextSpec", "github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressContextStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_StressContextSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StressContextSpec defines the desired state of StressContext",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"installPrometheus": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"installGrafana": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_StressContextStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StressContextStatus defines the observed state of StressContext",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_StressSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StressSpec defines the desired state of Stress",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"cassandraConfig": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.CassandraConfig"),
						},
					},
					"image": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"imagePullPolicy": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"stressConfig": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressConfig"),
						},
					},
					"jobConfig": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.JobConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.CassandraConfig", "github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.JobConfig", "github.com/jsanda/stress-operator/pkg/apis/thelastpickle/v1alpha1.StressConfig"},
	}
}

func schema_pkg_apis_thelastpickle_v1alpha1_StressStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StressStatus defines the observed state of Stress",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"jobStatus": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/api/batch/v1.JobStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/batch/v1.JobStatus"},
	}
}
