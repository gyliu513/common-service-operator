// +build !ignore_autogenerated

//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceConfig":       schema_pkg_apis_operator_v1alpha1_CommonServiceConfig(ref),
		"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceConfigSpec":   schema_pkg_apis_operator_v1alpha1_CommonServiceConfigSpec(ref),
		"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceConfigStatus": schema_pkg_apis_operator_v1alpha1_CommonServiceConfigStatus(ref),
		"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceSet":          schema_pkg_apis_operator_v1alpha1_CommonServiceSet(ref),
		"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceSetSpec":      schema_pkg_apis_operator_v1alpha1_CommonServiceSetSpec(ref),
		"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceSetStatus":    schema_pkg_apis_operator_v1alpha1_CommonServiceSetStatus(ref),
		"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.MetaOperator":              schema_pkg_apis_operator_v1alpha1_MetaOperator(ref),
		"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.MetaOperatorSpec":          schema_pkg_apis_operator_v1alpha1_MetaOperatorSpec(ref),
		"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.MetaOperatorStatus":        schema_pkg_apis_operator_v1alpha1_MetaOperatorStatus(ref),
	}
}

func schema_pkg_apis_operator_v1alpha1_CommonServiceConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CommonServiceConfig is the Schema for the commonserviceconfigs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceConfigStatus"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceConfigSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceConfigSpec", "github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceConfigStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_operator_v1alpha1_CommonServiceConfigSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CommonServiceConfigSpec defines the desired state of CommonServiceConfig",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"services": {
						SchemaProps: spec.SchemaProps{
							Description: "Services is a list of configuration of service",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.ConfigService"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.ConfigService"},
	}
}

func schema_pkg_apis_operator_v1alpha1_CommonServiceConfigStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CommonServiceConfigStatus defines the observed state of CommonServiceConfig",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"serviceStatus": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html ServiceStatus defines all the status of a operator",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CrStatus"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CrStatus"},
	}
}

func schema_pkg_apis_operator_v1alpha1_CommonServiceSet(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CommonServiceSet is the Schema for the commonservicesets API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
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
							Ref: ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceSetSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceSetStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceSetSpec", "github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.CommonServiceSetStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_operator_v1alpha1_CommonServiceSetSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CommonServiceSetSpec defines the desired state of CommonServiceSet",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"services": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.SetService"),
									},
								},
							},
						},
					},
				},
				Required: []string{"services"},
			},
		},
		Dependencies: []string{
			"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.SetService"},
	}
}

func schema_pkg_apis_operator_v1alpha1_CommonServiceSetStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CommonServiceSetStatus defines the observed state of CommonServiceSet",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html Conditions represents the current state of the Set Service",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.Condition"),
									},
								},
							},
						},
					},
					"members": {
						SchemaProps: spec.SchemaProps{
							Description: "Members represnets the current operators of the set",
							Ref:         ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.MembersStatus"),
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is the cluster running phase",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"phase"},
			},
		},
		Dependencies: []string{
			"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.Condition", "github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.MembersStatus"},
	}
}

func schema_pkg_apis_operator_v1alpha1_MetaOperator(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MetaOperator is the Schema for the metaoperators API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.MetaOperatorStatus"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.MetaOperatorSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.MetaOperatorSpec", "github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.MetaOperatorStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_operator_v1alpha1_MetaOperatorSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MetaOperatorSpec defines the desired state of MetaOperator",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"operators": {
						SchemaProps: spec.SchemaProps{
							Description: "Operators is a list of operator definition",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.Operator"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/IBM/common-service-operator/pkg/apis/operator/v1alpha1.Operator"},
	}
}

func schema_pkg_apis_operator_v1alpha1_MetaOperatorStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MetaOperatorStatus defines the observed state of MetaOperator",
				Type:        []string{"object"},
			},
		},
	}
}
